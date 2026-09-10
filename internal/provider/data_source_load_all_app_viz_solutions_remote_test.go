package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllAppVizSolutionsDataSource_Read_Happy exercises LoadAllAppVizSolutionsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllAppVizSolutionsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllAppVizSolutionsDataSource{client: newMockClientStatus(t, 200, "{\"appVzbilitySolutions\":[]}")}
	m := LoadAllAppVizSolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllAppVizSolutionsDataSource_Read_NilClient exercises LoadAllAppVizSolutionsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllAppVizSolutionsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllAppVizSolutionsDataSource{}
	m := LoadAllAppVizSolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllAppVizSolutionsDataSource_Read_BuildError exercises LoadAllAppVizSolutionsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllAppVizSolutionsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllAppVizSolutionsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllAppVizSolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllAppVizSolutionsDataSource_Read_SendError exercises LoadAllAppVizSolutionsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllAppVizSolutionsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllAppVizSolutionsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllAppVizSolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllAppVizSolutionsDataSource_Read_InvalidJSON exercises LoadAllAppVizSolutionsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllAppVizSolutionsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllAppVizSolutionsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllAppVizSolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
