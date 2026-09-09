package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllInlineToolsDataSource_Read_Happy exercises LoadAllInlineToolsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllInlineToolsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllInlineToolsDataSource{client: newMockClientStatus(t, 200, "{\"inlineTools\":[]}")}
	m := LoadAllInlineToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllInlineToolsDataSource_Read_NilClient exercises LoadAllInlineToolsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllInlineToolsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllInlineToolsDataSource{}
	m := LoadAllInlineToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllInlineToolsDataSource_Read_BuildError exercises LoadAllInlineToolsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllInlineToolsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllInlineToolsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllInlineToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllInlineToolsDataSource_Read_SendError exercises LoadAllInlineToolsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllInlineToolsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllInlineToolsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllInlineToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllInlineToolsDataSource_Read_InvalidJSON exercises LoadAllInlineToolsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllInlineToolsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllInlineToolsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllInlineToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
