package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllAppVisibilitySolutionsDataSource_Read_Happy exercises LoadAllAppVisibilitySolutionsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllAppVisibilitySolutionsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllAppVisibilitySolutionsDataSource{client: newMockClientStatus(t, 200, "{\"appsVisibilitySolutions\":[]}")}
	m := LoadAllAppVisibilitySolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllAppVisibilitySolutionsDataSource_Read_NilClient exercises LoadAllAppVisibilitySolutionsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllAppVisibilitySolutionsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllAppVisibilitySolutionsDataSource{}
	m := LoadAllAppVisibilitySolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllAppVisibilitySolutionsDataSource_Read_BuildError exercises LoadAllAppVisibilitySolutionsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllAppVisibilitySolutionsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllAppVisibilitySolutionsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllAppVisibilitySolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllAppVisibilitySolutionsDataSource_Read_SendError exercises LoadAllAppVisibilitySolutionsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllAppVisibilitySolutionsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllAppVisibilitySolutionsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllAppVisibilitySolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllAppVisibilitySolutionsDataSource_Read_InvalidJSON exercises LoadAllAppVisibilitySolutionsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllAppVisibilitySolutionsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllAppVisibilitySolutionsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllAppVisibilitySolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
