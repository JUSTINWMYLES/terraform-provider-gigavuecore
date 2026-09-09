package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTrafficFlowsOverviewDataSource_Read_Happy exercises GetTrafficFlowsOverviewDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetTrafficFlowsOverviewDataSource_Read_Happy(t *testing.T) {
	r := &GetTrafficFlowsOverviewDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetTrafficFlowsOverviewDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTrafficFlowsOverviewDataSource_Read_NilClient exercises GetTrafficFlowsOverviewDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTrafficFlowsOverviewDataSource_Read_NilClient(t *testing.T) {
	r := &GetTrafficFlowsOverviewDataSource{}
	m := GetTrafficFlowsOverviewDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTrafficFlowsOverviewDataSource_Read_BuildError exercises GetTrafficFlowsOverviewDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetTrafficFlowsOverviewDataSource_Read_BuildError(t *testing.T) {
	r := &GetTrafficFlowsOverviewDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTrafficFlowsOverviewDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetTrafficFlowsOverviewDataSource_Read_SendError exercises GetTrafficFlowsOverviewDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetTrafficFlowsOverviewDataSource_Read_SendError(t *testing.T) {
	r := &GetTrafficFlowsOverviewDataSource{client: newTransportErrorClient(t)}
	m := GetTrafficFlowsOverviewDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetTrafficFlowsOverviewDataSource_Read_NotFound exercises GetTrafficFlowsOverviewDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetTrafficFlowsOverviewDataSource_Read_NotFound(t *testing.T) {
	r := &GetTrafficFlowsOverviewDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetTrafficFlowsOverviewDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetTrafficFlowsOverviewDataSource_Read_APIError exercises GetTrafficFlowsOverviewDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetTrafficFlowsOverviewDataSource_Read_APIError(t *testing.T) {
	r := &GetTrafficFlowsOverviewDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetTrafficFlowsOverviewDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_traffic_flows_overview")
}

// TestGetTrafficFlowsOverviewDataSource_Read_APIErrorReadBody exercises GetTrafficFlowsOverviewDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetTrafficFlowsOverviewDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetTrafficFlowsOverviewDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetTrafficFlowsOverviewDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetTrafficFlowsOverviewDataSource_Read_InvalidJSON exercises GetTrafficFlowsOverviewDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetTrafficFlowsOverviewDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTrafficFlowsOverviewDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTrafficFlowsOverviewDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
