package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllImageServersDataSource_Read_Happy exercises LoadAllImageServersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllImageServersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllImageServersDataSource{client: newMockClientStatus(t, 200, "{\"imageServers\":[]}")}
	m := LoadAllImageServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllImageServersDataSource_Read_NilClient exercises LoadAllImageServersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllImageServersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllImageServersDataSource{}
	m := LoadAllImageServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllImageServersDataSource_Read_BuildError exercises LoadAllImageServersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllImageServersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllImageServersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllImageServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllImageServersDataSource_Read_SendError exercises LoadAllImageServersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllImageServersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllImageServersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllImageServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllImageServersDataSource_Read_InvalidJSON exercises LoadAllImageServersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllImageServersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllImageServersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllImageServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
