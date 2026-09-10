package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFmTemplatesHierarchialConfigDataSource_Read_Happy exercises LoadAllFmTemplatesHierarchialConfigDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllFmTemplatesHierarchialConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllFmTemplatesHierarchialConfigDataSource{client: newMockClientStatus(t, 200, "{\"fmTemplateHierarchicalConfigs\":[]}")}
	m := LoadAllFmTemplatesHierarchialConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllFmTemplatesHierarchialConfigDataSource_Read_NilClient exercises LoadAllFmTemplatesHierarchialConfigDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllFmTemplatesHierarchialConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllFmTemplatesHierarchialConfigDataSource{}
	m := LoadAllFmTemplatesHierarchialConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllFmTemplatesHierarchialConfigDataSource_Read_BuildError exercises LoadAllFmTemplatesHierarchialConfigDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllFmTemplatesHierarchialConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllFmTemplatesHierarchialConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllFmTemplatesHierarchialConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllFmTemplatesHierarchialConfigDataSource_Read_SendError exercises LoadAllFmTemplatesHierarchialConfigDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllFmTemplatesHierarchialConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllFmTemplatesHierarchialConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadAllFmTemplatesHierarchialConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllFmTemplatesHierarchialConfigDataSource_Read_InvalidJSON exercises LoadAllFmTemplatesHierarchialConfigDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllFmTemplatesHierarchialConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllFmTemplatesHierarchialConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllFmTemplatesHierarchialConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
