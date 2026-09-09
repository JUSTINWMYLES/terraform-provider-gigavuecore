package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTrafficFlowsGlobalSettingsDataSource_Read_Happy exercises GetTrafficFlowsGlobalSettingsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetTrafficFlowsGlobalSettingsDataSource_Read_Happy(t *testing.T) {
	r := &GetTrafficFlowsGlobalSettingsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetTrafficFlowsGlobalSettingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTrafficFlowsGlobalSettingsDataSource_Read_NilClient exercises GetTrafficFlowsGlobalSettingsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTrafficFlowsGlobalSettingsDataSource_Read_NilClient(t *testing.T) {
	r := &GetTrafficFlowsGlobalSettingsDataSource{}
	m := GetTrafficFlowsGlobalSettingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTrafficFlowsGlobalSettingsDataSource_Read_BuildError exercises GetTrafficFlowsGlobalSettingsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetTrafficFlowsGlobalSettingsDataSource_Read_BuildError(t *testing.T) {
	r := &GetTrafficFlowsGlobalSettingsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTrafficFlowsGlobalSettingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetTrafficFlowsGlobalSettingsDataSource_Read_SendError exercises GetTrafficFlowsGlobalSettingsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetTrafficFlowsGlobalSettingsDataSource_Read_SendError(t *testing.T) {
	r := &GetTrafficFlowsGlobalSettingsDataSource{client: newTransportErrorClient(t)}
	m := GetTrafficFlowsGlobalSettingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetTrafficFlowsGlobalSettingsDataSource_Read_NotFound exercises GetTrafficFlowsGlobalSettingsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetTrafficFlowsGlobalSettingsDataSource_Read_NotFound(t *testing.T) {
	r := &GetTrafficFlowsGlobalSettingsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetTrafficFlowsGlobalSettingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetTrafficFlowsGlobalSettingsDataSource_Read_APIError exercises GetTrafficFlowsGlobalSettingsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetTrafficFlowsGlobalSettingsDataSource_Read_APIError(t *testing.T) {
	r := &GetTrafficFlowsGlobalSettingsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetTrafficFlowsGlobalSettingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_traffic_flows_global_settings")
}

// TestGetTrafficFlowsGlobalSettingsDataSource_Read_APIErrorReadBody exercises GetTrafficFlowsGlobalSettingsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetTrafficFlowsGlobalSettingsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetTrafficFlowsGlobalSettingsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetTrafficFlowsGlobalSettingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetTrafficFlowsGlobalSettingsDataSource_Read_InvalidJSON exercises GetTrafficFlowsGlobalSettingsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetTrafficFlowsGlobalSettingsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTrafficFlowsGlobalSettingsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTrafficFlowsGlobalSettingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
