package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllPortGroupDataSource_Read_Happy exercises LoadAllPortGroupDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllPortGroupDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllPortGroupDataSource{client: newMockClientStatus(t, 200, "{\"portGroups\":[]}")}
	m := LoadAllPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllPortGroupDataSource_Read_NilClient exercises LoadAllPortGroupDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllPortGroupDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllPortGroupDataSource{}
	m := LoadAllPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllPortGroupDataSource_Read_BuildError exercises LoadAllPortGroupDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllPortGroupDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllPortGroupDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllPortGroupDataSource_Read_SendError exercises LoadAllPortGroupDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllPortGroupDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllPortGroupDataSource{client: newTransportErrorClient(t)}
	m := LoadAllPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllPortGroupDataSource_Read_InvalidJSON exercises LoadAllPortGroupDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllPortGroupDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllPortGroupDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
