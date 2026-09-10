package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllIcapServersDataSource_Read_Happy exercises LoadAllIcapServersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllIcapServersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllIcapServersDataSource{client: newMockClientStatus(t, 200, "{\"icapServers\":[]}")}
	m := LoadAllIcapServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllIcapServersDataSource_Read_NilClient exercises LoadAllIcapServersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllIcapServersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllIcapServersDataSource{}
	m := LoadAllIcapServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllIcapServersDataSource_Read_BuildError exercises LoadAllIcapServersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllIcapServersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllIcapServersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllIcapServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllIcapServersDataSource_Read_SendError exercises LoadAllIcapServersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllIcapServersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllIcapServersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllIcapServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllIcapServersDataSource_Read_InvalidJSON exercises LoadAllIcapServersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllIcapServersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllIcapServersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllIcapServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
