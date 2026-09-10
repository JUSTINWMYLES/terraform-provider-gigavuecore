package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMapAliasChainsDataSource_Read_Happy exercises LoadAllMapAliasChainsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllMapAliasChainsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllMapAliasChainsDataSource{client: newMockClientStatus(t, 200, "{\"mapAliasChains\":[]}")}
	m := LoadAllMapAliasChainsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllMapAliasChainsDataSource_Read_NilClient exercises LoadAllMapAliasChainsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllMapAliasChainsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllMapAliasChainsDataSource{}
	m := LoadAllMapAliasChainsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllMapAliasChainsDataSource_Read_BuildError exercises LoadAllMapAliasChainsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllMapAliasChainsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllMapAliasChainsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllMapAliasChainsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMapAliasChainsDataSource_Read_SendError exercises LoadAllMapAliasChainsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllMapAliasChainsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllMapAliasChainsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllMapAliasChainsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMapAliasChainsDataSource_Read_InvalidJSON exercises LoadAllMapAliasChainsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllMapAliasChainsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllMapAliasChainsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllMapAliasChainsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
