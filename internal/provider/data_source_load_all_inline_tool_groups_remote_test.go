package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllInlineToolGroupsDataSource_Read_Happy exercises LoadAllInlineToolGroupsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllInlineToolGroupsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllInlineToolGroupsDataSource{client: newMockClientStatus(t, 200, "{\"inlineToolGroups\":[]}")}
	m := LoadAllInlineToolGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllInlineToolGroupsDataSource_Read_NilClient exercises LoadAllInlineToolGroupsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllInlineToolGroupsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllInlineToolGroupsDataSource{}
	m := LoadAllInlineToolGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllInlineToolGroupsDataSource_Read_BuildError exercises LoadAllInlineToolGroupsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllInlineToolGroupsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllInlineToolGroupsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllInlineToolGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllInlineToolGroupsDataSource_Read_SendError exercises LoadAllInlineToolGroupsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllInlineToolGroupsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllInlineToolGroupsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllInlineToolGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllInlineToolGroupsDataSource_Read_InvalidJSON exercises LoadAllInlineToolGroupsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllInlineToolGroupsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllInlineToolGroupsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllInlineToolGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
