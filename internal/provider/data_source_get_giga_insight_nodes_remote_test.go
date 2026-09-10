package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetGigaInsightNodesDataSource_Read_Happy exercises GetGigaInsightNodesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetGigaInsightNodesDataSource_Read_Happy(t *testing.T) {
	r := &GetGigaInsightNodesDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetGigaInsightNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetGigaInsightNodesDataSource_Read_NilClient exercises GetGigaInsightNodesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetGigaInsightNodesDataSource_Read_NilClient(t *testing.T) {
	r := &GetGigaInsightNodesDataSource{}
	m := GetGigaInsightNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetGigaInsightNodesDataSource_Read_BuildError exercises GetGigaInsightNodesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetGigaInsightNodesDataSource_Read_BuildError(t *testing.T) {
	r := &GetGigaInsightNodesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetGigaInsightNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetGigaInsightNodesDataSource_Read_SendError exercises GetGigaInsightNodesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetGigaInsightNodesDataSource_Read_SendError(t *testing.T) {
	r := &GetGigaInsightNodesDataSource{client: newTransportErrorClient(t)}
	m := GetGigaInsightNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetGigaInsightNodesDataSource_Read_InvalidJSON exercises GetGigaInsightNodesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetGigaInsightNodesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetGigaInsightNodesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetGigaInsightNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
