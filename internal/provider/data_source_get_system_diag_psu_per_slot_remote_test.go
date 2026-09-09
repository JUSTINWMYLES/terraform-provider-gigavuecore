package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemDiagPsuPerSlotDataSource_Read_Happy exercises GetSystemDiagPsuPerSlotDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSystemDiagPsuPerSlotDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemDiagPsuPerSlotDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSystemDiagPsuPerSlotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemDiagPsuPerSlotDataSource_Read_NilClient exercises GetSystemDiagPsuPerSlotDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemDiagPsuPerSlotDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemDiagPsuPerSlotDataSource{}
	m := GetSystemDiagPsuPerSlotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemDiagPsuPerSlotDataSource_Read_BuildError exercises GetSystemDiagPsuPerSlotDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSystemDiagPsuPerSlotDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemDiagPsuPerSlotDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemDiagPsuPerSlotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSystemDiagPsuPerSlotDataSource_Read_SendError exercises GetSystemDiagPsuPerSlotDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSystemDiagPsuPerSlotDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemDiagPsuPerSlotDataSource{client: newTransportErrorClient(t)}
	m := GetSystemDiagPsuPerSlotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSystemDiagPsuPerSlotDataSource_Read_NotFound exercises GetSystemDiagPsuPerSlotDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSystemDiagPsuPerSlotDataSource_Read_NotFound(t *testing.T) {
	r := &GetSystemDiagPsuPerSlotDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSystemDiagPsuPerSlotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSystemDiagPsuPerSlotDataSource_Read_APIError exercises GetSystemDiagPsuPerSlotDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSystemDiagPsuPerSlotDataSource_Read_APIError(t *testing.T) {
	r := &GetSystemDiagPsuPerSlotDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetSystemDiagPsuPerSlotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_system_diag_psu_per_slot")
}

// TestGetSystemDiagPsuPerSlotDataSource_Read_APIErrorReadBody exercises GetSystemDiagPsuPerSlotDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSystemDiagPsuPerSlotDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSystemDiagPsuPerSlotDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetSystemDiagPsuPerSlotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSystemDiagPsuPerSlotDataSource_Read_InvalidJSON exercises GetSystemDiagPsuPerSlotDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSystemDiagPsuPerSlotDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemDiagPsuPerSlotDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemDiagPsuPerSlotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
