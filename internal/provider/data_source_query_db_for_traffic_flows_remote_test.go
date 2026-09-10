package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryDbForTrafficFlowsDataSource_Read_Happy exercises QueryDbForTrafficFlowsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryDbForTrafficFlowsDataSource_Read_Happy(t *testing.T) {
	r := &QueryDbForTrafficFlowsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryDbForTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryDbForTrafficFlowsDataSource_Read_NilClient exercises QueryDbForTrafficFlowsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryDbForTrafficFlowsDataSource_Read_NilClient(t *testing.T) {
	r := &QueryDbForTrafficFlowsDataSource{}
	m := QueryDbForTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryDbForTrafficFlowsDataSource_Read_BuildError exercises QueryDbForTrafficFlowsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryDbForTrafficFlowsDataSource_Read_BuildError(t *testing.T) {
	r := &QueryDbForTrafficFlowsDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryDbForTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryDbForTrafficFlowsDataSource_Read_SendError exercises QueryDbForTrafficFlowsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryDbForTrafficFlowsDataSource_Read_SendError(t *testing.T) {
	r := &QueryDbForTrafficFlowsDataSource{client: newTransportErrorClient(t)}
	m := QueryDbForTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryDbForTrafficFlowsDataSource_Read_NotFound exercises QueryDbForTrafficFlowsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryDbForTrafficFlowsDataSource_Read_NotFound(t *testing.T) {
	r := &QueryDbForTrafficFlowsDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryDbForTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryDbForTrafficFlowsDataSource_Read_APIError exercises QueryDbForTrafficFlowsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryDbForTrafficFlowsDataSource_Read_APIError(t *testing.T) {
	r := &QueryDbForTrafficFlowsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryDbForTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_db_for_traffic_flows")
}

// TestQueryDbForTrafficFlowsDataSource_Read_APIErrorReadBody exercises QueryDbForTrafficFlowsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryDbForTrafficFlowsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryDbForTrafficFlowsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryDbForTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryDbForTrafficFlowsDataSource_Read_InvalidJSON exercises QueryDbForTrafficFlowsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryDbForTrafficFlowsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryDbForTrafficFlowsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryDbForTrafficFlowsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
