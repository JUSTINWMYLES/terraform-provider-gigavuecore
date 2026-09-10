package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllInlineSerialToolGroupsDataSource_Read_Happy exercises LoadAllInlineSerialToolGroupsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllInlineSerialToolGroupsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllInlineSerialToolGroupsDataSource{client: newMockClientStatus(t, 200, "{\"inlineSerialToolGroups\":[]}")}
	m := LoadAllInlineSerialToolGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllInlineSerialToolGroupsDataSource_Read_NilClient exercises LoadAllInlineSerialToolGroupsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllInlineSerialToolGroupsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllInlineSerialToolGroupsDataSource{}
	m := LoadAllInlineSerialToolGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllInlineSerialToolGroupsDataSource_Read_BuildError exercises LoadAllInlineSerialToolGroupsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllInlineSerialToolGroupsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllInlineSerialToolGroupsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllInlineSerialToolGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllInlineSerialToolGroupsDataSource_Read_SendError exercises LoadAllInlineSerialToolGroupsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllInlineSerialToolGroupsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllInlineSerialToolGroupsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllInlineSerialToolGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllInlineSerialToolGroupsDataSource_Read_InvalidJSON exercises LoadAllInlineSerialToolGroupsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllInlineSerialToolGroupsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllInlineSerialToolGroupsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllInlineSerialToolGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
