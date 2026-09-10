package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeStatusDataSource_Read_Happy exercises GetNodeStatusDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetNodeStatusDataSource_Read_Happy(t *testing.T) {
	r := &GetNodeStatusDataSource{client: newMockClientStatus(t, 200, "{\"nodes\":[]}")}
	m := GetNodeStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetNodeStatusDataSource_Read_NilClient exercises GetNodeStatusDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetNodeStatusDataSource_Read_NilClient(t *testing.T) {
	r := &GetNodeStatusDataSource{}
	m := GetNodeStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetNodeStatusDataSource_Read_BuildError exercises GetNodeStatusDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetNodeStatusDataSource_Read_BuildError(t *testing.T) {
	r := &GetNodeStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := GetNodeStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNodeStatusDataSource_Read_SendError exercises GetNodeStatusDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetNodeStatusDataSource_Read_SendError(t *testing.T) {
	r := &GetNodeStatusDataSource{client: newTransportErrorClient(t)}
	m := GetNodeStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNodeStatusDataSource_Read_InvalidJSON exercises GetNodeStatusDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetNodeStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetNodeStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetNodeStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
