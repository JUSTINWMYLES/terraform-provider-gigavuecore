package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllAppsListenerDataSource_Read_Happy exercises GetAllAppsListenerDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllAppsListenerDataSource_Read_Happy(t *testing.T) {
	r := &GetAllAppsListenerDataSource{client: newMockClientStatus(t, 200, "{\"appsListeners\":[]}")}
	m := GetAllAppsListenerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllAppsListenerDataSource_Read_NilClient exercises GetAllAppsListenerDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllAppsListenerDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllAppsListenerDataSource{}
	m := GetAllAppsListenerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllAppsListenerDataSource_Read_BuildError exercises GetAllAppsListenerDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllAppsListenerDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllAppsListenerDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllAppsListenerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllAppsListenerDataSource_Read_SendError exercises GetAllAppsListenerDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllAppsListenerDataSource_Read_SendError(t *testing.T) {
	r := &GetAllAppsListenerDataSource{client: newTransportErrorClient(t)}
	m := GetAllAppsListenerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllAppsListenerDataSource_Read_InvalidJSON exercises GetAllAppsListenerDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllAppsListenerDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllAppsListenerDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllAppsListenerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
