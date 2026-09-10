package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllUserTagsDataSource_Read_Happy exercises LoadAllUserTagsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllUserTagsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllUserTagsDataSource{client: newMockClientStatus(t, 200, "{\"tags\":[]}")}
	m := LoadAllUserTagsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllUserTagsDataSource_Read_NilClient exercises LoadAllUserTagsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllUserTagsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllUserTagsDataSource{}
	m := LoadAllUserTagsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllUserTagsDataSource_Read_BuildError exercises LoadAllUserTagsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllUserTagsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllUserTagsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllUserTagsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllUserTagsDataSource_Read_SendError exercises LoadAllUserTagsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllUserTagsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllUserTagsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllUserTagsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllUserTagsDataSource_Read_InvalidJSON exercises LoadAllUserTagsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllUserTagsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllUserTagsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllUserTagsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
