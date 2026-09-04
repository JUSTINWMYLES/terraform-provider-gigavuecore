package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllGsopsDataSource_Read_Happy exercises LoadAllGsopsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllGsopsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllGsopsDataSource{client: newMockClientStatus(t, 200, "{\"gsops\":[]}")}
	m := LoadAllGsopsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllGsopsDataSource_Read_NilClient exercises LoadAllGsopsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllGsopsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllGsopsDataSource{}
	m := LoadAllGsopsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllGsopsDataSource_Read_BuildError exercises LoadAllGsopsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllGsopsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllGsopsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllGsopsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllGsopsDataSource_Read_SendError exercises LoadAllGsopsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllGsopsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllGsopsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllGsopsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllGsopsDataSource_Read_InvalidJSON exercises LoadAllGsopsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllGsopsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllGsopsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllGsopsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
