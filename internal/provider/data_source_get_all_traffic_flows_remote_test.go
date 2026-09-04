package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTrafficFlowsDataSource_Read_Happy exercises GetAllTrafficFlowsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAllTrafficFlowsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllTrafficFlowsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAllTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllTrafficFlowsDataSource_Read_NilClient exercises GetAllTrafficFlowsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllTrafficFlowsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllTrafficFlowsDataSource{}
	m := GetAllTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllTrafficFlowsDataSource_Read_BuildError exercises GetAllTrafficFlowsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAllTrafficFlowsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllTrafficFlowsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAllTrafficFlowsDataSource_Read_SendError exercises GetAllTrafficFlowsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAllTrafficFlowsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllTrafficFlowsDataSource{client: newTransportErrorClient(t)}
	m := GetAllTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAllTrafficFlowsDataSource_Read_NotFound exercises GetAllTrafficFlowsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAllTrafficFlowsDataSource_Read_NotFound(t *testing.T) {
	r := &GetAllTrafficFlowsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAllTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAllTrafficFlowsDataSource_Read_APIError exercises GetAllTrafficFlowsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAllTrafficFlowsDataSource_Read_APIError(t *testing.T) {
	r := &GetAllTrafficFlowsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAllTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_all_traffic_flows")
}

// TestGetAllTrafficFlowsDataSource_Read_APIErrorReadBody exercises GetAllTrafficFlowsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAllTrafficFlowsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAllTrafficFlowsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAllTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAllTrafficFlowsDataSource_Read_InvalidJSON exercises GetAllTrafficFlowsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAllTrafficFlowsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllTrafficFlowsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
