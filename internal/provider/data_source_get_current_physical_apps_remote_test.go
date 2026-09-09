package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCurrentPhysicalAppsDataSource_Read_Happy exercises GetCurrentPhysicalAppsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetCurrentPhysicalAppsDataSource_Read_Happy(t *testing.T) {
	r := &GetCurrentPhysicalAppsDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetCurrentPhysicalAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCurrentPhysicalAppsDataSource_Read_NilClient exercises GetCurrentPhysicalAppsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCurrentPhysicalAppsDataSource_Read_NilClient(t *testing.T) {
	r := &GetCurrentPhysicalAppsDataSource{}
	m := GetCurrentPhysicalAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCurrentPhysicalAppsDataSource_Read_BuildError exercises GetCurrentPhysicalAppsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetCurrentPhysicalAppsDataSource_Read_BuildError(t *testing.T) {
	r := &GetCurrentPhysicalAppsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCurrentPhysicalAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCurrentPhysicalAppsDataSource_Read_SendError exercises GetCurrentPhysicalAppsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetCurrentPhysicalAppsDataSource_Read_SendError(t *testing.T) {
	r := &GetCurrentPhysicalAppsDataSource{client: newTransportErrorClient(t)}
	m := GetCurrentPhysicalAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCurrentPhysicalAppsDataSource_Read_InvalidJSON exercises GetCurrentPhysicalAppsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetCurrentPhysicalAppsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCurrentPhysicalAppsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCurrentPhysicalAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
