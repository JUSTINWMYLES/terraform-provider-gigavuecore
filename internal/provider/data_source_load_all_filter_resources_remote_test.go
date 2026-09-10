package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFilterResourcesDataSource_Read_Happy exercises LoadAllFilterResourcesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllFilterResourcesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllFilterResourcesDataSource{client: newMockClientStatus(t, 200, "{\"filterResources\":[]}")}
	m := LoadAllFilterResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllFilterResourcesDataSource_Read_NilClient exercises LoadAllFilterResourcesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllFilterResourcesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllFilterResourcesDataSource{}
	m := LoadAllFilterResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllFilterResourcesDataSource_Read_BuildError exercises LoadAllFilterResourcesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllFilterResourcesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllFilterResourcesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllFilterResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllFilterResourcesDataSource_Read_SendError exercises LoadAllFilterResourcesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllFilterResourcesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllFilterResourcesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllFilterResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllFilterResourcesDataSource_Read_InvalidJSON exercises LoadAllFilterResourcesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllFilterResourcesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllFilterResourcesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllFilterResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
