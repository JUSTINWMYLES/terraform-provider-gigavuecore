package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadBatteryOptimizationDataSource_Read_Happy exercises LoadBatteryOptimizationDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadBatteryOptimizationDataSource_Read_Happy(t *testing.T) {
	r := &LoadBatteryOptimizationDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadBatteryOptimizationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadBatteryOptimizationDataSource_Read_NilClient exercises LoadBatteryOptimizationDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadBatteryOptimizationDataSource_Read_NilClient(t *testing.T) {
	r := &LoadBatteryOptimizationDataSource{}
	m := LoadBatteryOptimizationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadBatteryOptimizationDataSource_Read_BuildError exercises LoadBatteryOptimizationDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadBatteryOptimizationDataSource_Read_BuildError(t *testing.T) {
	r := &LoadBatteryOptimizationDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadBatteryOptimizationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadBatteryOptimizationDataSource_Read_SendError exercises LoadBatteryOptimizationDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadBatteryOptimizationDataSource_Read_SendError(t *testing.T) {
	r := &LoadBatteryOptimizationDataSource{client: newTransportErrorClient(t)}
	m := LoadBatteryOptimizationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadBatteryOptimizationDataSource_Read_NotFound exercises LoadBatteryOptimizationDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadBatteryOptimizationDataSource_Read_NotFound(t *testing.T) {
	r := &LoadBatteryOptimizationDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadBatteryOptimizationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadBatteryOptimizationDataSource_Read_APIError exercises LoadBatteryOptimizationDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadBatteryOptimizationDataSource_Read_APIError(t *testing.T) {
	r := &LoadBatteryOptimizationDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadBatteryOptimizationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_battery_optimization")
}

// TestLoadBatteryOptimizationDataSource_Read_APIErrorReadBody exercises LoadBatteryOptimizationDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadBatteryOptimizationDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadBatteryOptimizationDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadBatteryOptimizationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadBatteryOptimizationDataSource_Read_InvalidJSON exercises LoadBatteryOptimizationDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadBatteryOptimizationDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadBatteryOptimizationDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadBatteryOptimizationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
