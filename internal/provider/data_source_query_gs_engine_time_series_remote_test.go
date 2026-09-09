package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryGsEngineTimeSeriesDataSource_Read_Happy exercises QueryGsEngineTimeSeriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryGsEngineTimeSeriesDataSource_Read_Happy(t *testing.T) {
	r := &QueryGsEngineTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryGsEngineTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryGsEngineTimeSeriesDataSource_Read_NilClient exercises QueryGsEngineTimeSeriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryGsEngineTimeSeriesDataSource_Read_NilClient(t *testing.T) {
	r := &QueryGsEngineTimeSeriesDataSource{}
	m := QueryGsEngineTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryGsEngineTimeSeriesDataSource_Read_BuildError exercises QueryGsEngineTimeSeriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryGsEngineTimeSeriesDataSource_Read_BuildError(t *testing.T) {
	r := &QueryGsEngineTimeSeriesDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryGsEngineTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryGsEngineTimeSeriesDataSource_Read_SendError exercises QueryGsEngineTimeSeriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryGsEngineTimeSeriesDataSource_Read_SendError(t *testing.T) {
	r := &QueryGsEngineTimeSeriesDataSource{client: newTransportErrorClient(t)}
	m := QueryGsEngineTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryGsEngineTimeSeriesDataSource_Read_NotFound exercises QueryGsEngineTimeSeriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryGsEngineTimeSeriesDataSource_Read_NotFound(t *testing.T) {
	r := &QueryGsEngineTimeSeriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryGsEngineTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryGsEngineTimeSeriesDataSource_Read_APIError exercises QueryGsEngineTimeSeriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryGsEngineTimeSeriesDataSource_Read_APIError(t *testing.T) {
	r := &QueryGsEngineTimeSeriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryGsEngineTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_gs_engine_time_series")
}

// TestQueryGsEngineTimeSeriesDataSource_Read_APIErrorReadBody exercises QueryGsEngineTimeSeriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryGsEngineTimeSeriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryGsEngineTimeSeriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryGsEngineTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryGsEngineTimeSeriesDataSource_Read_InvalidJSON exercises QueryGsEngineTimeSeriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryGsEngineTimeSeriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryGsEngineTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryGsEngineTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
