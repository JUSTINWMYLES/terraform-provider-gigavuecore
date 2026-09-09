package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllGtapPortGroupDataSource_Read_Happy exercises LoadAllGtapPortGroupDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllGtapPortGroupDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllGtapPortGroupDataSource{client: newMockClientStatus(t, 200, "{\"gtapPortGroups\":[]}")}
	m := LoadAllGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllGtapPortGroupDataSource_Read_NilClient exercises LoadAllGtapPortGroupDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllGtapPortGroupDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllGtapPortGroupDataSource{}
	m := LoadAllGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllGtapPortGroupDataSource_Read_BuildError exercises LoadAllGtapPortGroupDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllGtapPortGroupDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllGtapPortGroupDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllGtapPortGroupDataSource_Read_SendError exercises LoadAllGtapPortGroupDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllGtapPortGroupDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllGtapPortGroupDataSource{client: newTransportErrorClient(t)}
	m := LoadAllGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllGtapPortGroupDataSource_Read_InvalidJSON exercises LoadAllGtapPortGroupDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllGtapPortGroupDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllGtapPortGroupDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
