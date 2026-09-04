package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFilterTemplatesDataSource_Read_Happy exercises LoadAllFilterTemplatesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllFilterTemplatesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllFilterTemplatesDataSource{client: newMockClientStatus(t, 200, "{\"filterTemplates\":[]}")}
	m := LoadAllFilterTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllFilterTemplatesDataSource_Read_NilClient exercises LoadAllFilterTemplatesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllFilterTemplatesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllFilterTemplatesDataSource{}
	m := LoadAllFilterTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllFilterTemplatesDataSource_Read_BuildError exercises LoadAllFilterTemplatesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllFilterTemplatesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllFilterTemplatesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllFilterTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllFilterTemplatesDataSource_Read_SendError exercises LoadAllFilterTemplatesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllFilterTemplatesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllFilterTemplatesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllFilterTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllFilterTemplatesDataSource_Read_InvalidJSON exercises LoadAllFilterTemplatesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllFilterTemplatesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllFilterTemplatesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllFilterTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
