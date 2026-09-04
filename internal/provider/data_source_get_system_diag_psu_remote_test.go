package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemDiagPsuDataSource_Read_Happy exercises GetSystemDiagPsuDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetSystemDiagPsuDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemDiagPsuDataSource{client: newMockClientStatus(t, 200, "{\"systemPsuDiagDetails\":[]}")}
	m := GetSystemDiagPsuDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemDiagPsuDataSource_Read_NilClient exercises GetSystemDiagPsuDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemDiagPsuDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemDiagPsuDataSource{}
	m := GetSystemDiagPsuDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemDiagPsuDataSource_Read_BuildError exercises GetSystemDiagPsuDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetSystemDiagPsuDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemDiagPsuDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemDiagPsuDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSystemDiagPsuDataSource_Read_SendError exercises GetSystemDiagPsuDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetSystemDiagPsuDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemDiagPsuDataSource{client: newTransportErrorClient(t)}
	m := GetSystemDiagPsuDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSystemDiagPsuDataSource_Read_InvalidJSON exercises GetSystemDiagPsuDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetSystemDiagPsuDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemDiagPsuDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemDiagPsuDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
