package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFabricAdvHashDataSource_Read_Happy exercises LoadFabricAdvHashDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadFabricAdvHashDataSource_Read_Happy(t *testing.T) {
	r := &LoadFabricAdvHashDataSource{client: newMockClientStatus(t, 200, "{\"fabricAdvHash\":[]}")}
	m := LoadFabricAdvHashDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadFabricAdvHashDataSource_Read_NilClient exercises LoadFabricAdvHashDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadFabricAdvHashDataSource_Read_NilClient(t *testing.T) {
	r := &LoadFabricAdvHashDataSource{}
	m := LoadFabricAdvHashDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadFabricAdvHashDataSource_Read_BuildError exercises LoadFabricAdvHashDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadFabricAdvHashDataSource_Read_BuildError(t *testing.T) {
	r := &LoadFabricAdvHashDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadFabricAdvHashDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadFabricAdvHashDataSource_Read_SendError exercises LoadFabricAdvHashDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadFabricAdvHashDataSource_Read_SendError(t *testing.T) {
	r := &LoadFabricAdvHashDataSource{client: newTransportErrorClient(t)}
	m := LoadFabricAdvHashDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadFabricAdvHashDataSource_Read_InvalidJSON exercises LoadFabricAdvHashDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadFabricAdvHashDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadFabricAdvHashDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadFabricAdvHashDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
