package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodesDataSource_Read_Happy exercises GetNodesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetNodesDataSource_Read_Happy(t *testing.T) {
	r := &GetNodesDataSource{client: newMockClientStatus(t, 200, "{\"nodes\":[]}")}
	m := GetNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetNodesDataSource_Read_NilClient exercises GetNodesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetNodesDataSource_Read_NilClient(t *testing.T) {
	r := &GetNodesDataSource{}
	m := GetNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetNodesDataSource_Read_BuildError exercises GetNodesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetNodesDataSource_Read_BuildError(t *testing.T) {
	r := &GetNodesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNodesDataSource_Read_SendError exercises GetNodesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetNodesDataSource_Read_SendError(t *testing.T) {
	r := &GetNodesDataSource{client: newTransportErrorClient(t)}
	m := GetNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNodesDataSource_Read_InvalidJSON exercises GetNodesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetNodesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetNodesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
