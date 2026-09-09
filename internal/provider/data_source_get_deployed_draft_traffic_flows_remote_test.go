package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetDeployedDraftTrafficFlowsDataSource_Read_Happy exercises GetDeployedDraftTrafficFlowsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetDeployedDraftTrafficFlowsDataSource_Read_Happy(t *testing.T) {
	r := &GetDeployedDraftTrafficFlowsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetDeployedDraftTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetDeployedDraftTrafficFlowsDataSource_Read_NilClient exercises GetDeployedDraftTrafficFlowsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetDeployedDraftTrafficFlowsDataSource_Read_NilClient(t *testing.T) {
	r := &GetDeployedDraftTrafficFlowsDataSource{}
	m := GetDeployedDraftTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetDeployedDraftTrafficFlowsDataSource_Read_BuildError exercises GetDeployedDraftTrafficFlowsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetDeployedDraftTrafficFlowsDataSource_Read_BuildError(t *testing.T) {
	r := &GetDeployedDraftTrafficFlowsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetDeployedDraftTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetDeployedDraftTrafficFlowsDataSource_Read_SendError exercises GetDeployedDraftTrafficFlowsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetDeployedDraftTrafficFlowsDataSource_Read_SendError(t *testing.T) {
	r := &GetDeployedDraftTrafficFlowsDataSource{client: newTransportErrorClient(t)}
	m := GetDeployedDraftTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetDeployedDraftTrafficFlowsDataSource_Read_NotFound exercises GetDeployedDraftTrafficFlowsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetDeployedDraftTrafficFlowsDataSource_Read_NotFound(t *testing.T) {
	r := &GetDeployedDraftTrafficFlowsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetDeployedDraftTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetDeployedDraftTrafficFlowsDataSource_Read_APIError exercises GetDeployedDraftTrafficFlowsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetDeployedDraftTrafficFlowsDataSource_Read_APIError(t *testing.T) {
	r := &GetDeployedDraftTrafficFlowsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetDeployedDraftTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_deployed_draft_traffic_flows")
}

// TestGetDeployedDraftTrafficFlowsDataSource_Read_APIErrorReadBody exercises GetDeployedDraftTrafficFlowsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetDeployedDraftTrafficFlowsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetDeployedDraftTrafficFlowsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetDeployedDraftTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetDeployedDraftTrafficFlowsDataSource_Read_InvalidJSON exercises GetDeployedDraftTrafficFlowsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetDeployedDraftTrafficFlowsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetDeployedDraftTrafficFlowsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetDeployedDraftTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
