package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCurrentAppsDataSource_Read_Happy exercises GetCurrentAppsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetCurrentAppsDataSource_Read_Happy(t *testing.T) {
	r := &GetCurrentAppsDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetCurrentAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCurrentAppsDataSource_Read_NilClient exercises GetCurrentAppsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCurrentAppsDataSource_Read_NilClient(t *testing.T) {
	r := &GetCurrentAppsDataSource{}
	m := GetCurrentAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCurrentAppsDataSource_Read_BuildError exercises GetCurrentAppsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetCurrentAppsDataSource_Read_BuildError(t *testing.T) {
	r := &GetCurrentAppsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCurrentAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCurrentAppsDataSource_Read_SendError exercises GetCurrentAppsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetCurrentAppsDataSource_Read_SendError(t *testing.T) {
	r := &GetCurrentAppsDataSource{client: newTransportErrorClient(t)}
	m := GetCurrentAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCurrentAppsDataSource_Read_InvalidJSON exercises GetCurrentAppsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetCurrentAppsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCurrentAppsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCurrentAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
