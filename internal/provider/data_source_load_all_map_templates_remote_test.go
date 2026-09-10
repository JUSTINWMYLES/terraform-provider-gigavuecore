package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMapTemplatesDataSource_Read_Happy exercises LoadAllMapTemplatesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllMapTemplatesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllMapTemplatesDataSource{client: newMockClientStatus(t, 200, "{\"mapTemplates\":[]}")}
	m := LoadAllMapTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllMapTemplatesDataSource_Read_NilClient exercises LoadAllMapTemplatesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllMapTemplatesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllMapTemplatesDataSource{}
	m := LoadAllMapTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllMapTemplatesDataSource_Read_BuildError exercises LoadAllMapTemplatesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllMapTemplatesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllMapTemplatesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllMapTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMapTemplatesDataSource_Read_SendError exercises LoadAllMapTemplatesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllMapTemplatesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllMapTemplatesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllMapTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMapTemplatesDataSource_Read_InvalidJSON exercises LoadAllMapTemplatesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllMapTemplatesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllMapTemplatesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllMapTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
