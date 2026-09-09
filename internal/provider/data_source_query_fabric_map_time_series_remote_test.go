package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryFabricMapTimeSeriesDataSource_Read_Happy exercises QueryFabricMapTimeSeriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryFabricMapTimeSeriesDataSource_Read_Happy(t *testing.T) {
	r := &QueryFabricMapTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryFabricMapTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryFabricMapTimeSeriesDataSource_Read_NilClient exercises QueryFabricMapTimeSeriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryFabricMapTimeSeriesDataSource_Read_NilClient(t *testing.T) {
	r := &QueryFabricMapTimeSeriesDataSource{}
	m := QueryFabricMapTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryFabricMapTimeSeriesDataSource_Read_BuildError exercises QueryFabricMapTimeSeriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryFabricMapTimeSeriesDataSource_Read_BuildError(t *testing.T) {
	r := &QueryFabricMapTimeSeriesDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryFabricMapTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryFabricMapTimeSeriesDataSource_Read_SendError exercises QueryFabricMapTimeSeriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryFabricMapTimeSeriesDataSource_Read_SendError(t *testing.T) {
	r := &QueryFabricMapTimeSeriesDataSource{client: newTransportErrorClient(t)}
	m := QueryFabricMapTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryFabricMapTimeSeriesDataSource_Read_NotFound exercises QueryFabricMapTimeSeriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryFabricMapTimeSeriesDataSource_Read_NotFound(t *testing.T) {
	r := &QueryFabricMapTimeSeriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryFabricMapTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryFabricMapTimeSeriesDataSource_Read_APIError exercises QueryFabricMapTimeSeriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryFabricMapTimeSeriesDataSource_Read_APIError(t *testing.T) {
	r := &QueryFabricMapTimeSeriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryFabricMapTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_fabric_map_time_series")
}

// TestQueryFabricMapTimeSeriesDataSource_Read_APIErrorReadBody exercises QueryFabricMapTimeSeriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryFabricMapTimeSeriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryFabricMapTimeSeriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryFabricMapTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryFabricMapTimeSeriesDataSource_Read_InvalidJSON exercises QueryFabricMapTimeSeriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryFabricMapTimeSeriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryFabricMapTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryFabricMapTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
