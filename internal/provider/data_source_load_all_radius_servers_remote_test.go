package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllRadiusServersDataSource_Read_Happy exercises LoadAllRadiusServersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllRadiusServersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllRadiusServersDataSource{client: newMockClientStatus(t, 200, "{\"radiusServers\":[]}")}
	m := LoadAllRadiusServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllRadiusServersDataSource_Read_NilClient exercises LoadAllRadiusServersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllRadiusServersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllRadiusServersDataSource{}
	m := LoadAllRadiusServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllRadiusServersDataSource_Read_BuildError exercises LoadAllRadiusServersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllRadiusServersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllRadiusServersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllRadiusServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllRadiusServersDataSource_Read_SendError exercises LoadAllRadiusServersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllRadiusServersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllRadiusServersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllRadiusServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllRadiusServersDataSource_Read_InvalidJSON exercises LoadAllRadiusServersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllRadiusServersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllRadiusServersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllRadiusServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
