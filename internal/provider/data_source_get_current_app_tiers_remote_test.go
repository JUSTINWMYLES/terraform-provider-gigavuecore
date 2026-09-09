package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCurrentAppTiersDataSource_Read_Happy exercises GetCurrentAppTiersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetCurrentAppTiersDataSource_Read_Happy(t *testing.T) {
	r := &GetCurrentAppTiersDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetCurrentAppTiersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCurrentAppTiersDataSource_Read_NilClient exercises GetCurrentAppTiersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCurrentAppTiersDataSource_Read_NilClient(t *testing.T) {
	r := &GetCurrentAppTiersDataSource{}
	m := GetCurrentAppTiersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCurrentAppTiersDataSource_Read_BuildError exercises GetCurrentAppTiersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetCurrentAppTiersDataSource_Read_BuildError(t *testing.T) {
	r := &GetCurrentAppTiersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCurrentAppTiersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCurrentAppTiersDataSource_Read_SendError exercises GetCurrentAppTiersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetCurrentAppTiersDataSource_Read_SendError(t *testing.T) {
	r := &GetCurrentAppTiersDataSource{client: newTransportErrorClient(t)}
	m := GetCurrentAppTiersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCurrentAppTiersDataSource_Read_InvalidJSON exercises GetCurrentAppTiersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetCurrentAppTiersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCurrentAppTiersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCurrentAppTiersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
