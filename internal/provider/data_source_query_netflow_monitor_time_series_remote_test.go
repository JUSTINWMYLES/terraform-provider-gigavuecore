package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryNetflowMonitorTimeSeriesDataSource_Read_Happy exercises QueryNetflowMonitorTimeSeriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryNetflowMonitorTimeSeriesDataSource_Read_Happy(t *testing.T) {
	r := &QueryNetflowMonitorTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryNetflowMonitorTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryNetflowMonitorTimeSeriesDataSource_Read_NilClient exercises QueryNetflowMonitorTimeSeriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryNetflowMonitorTimeSeriesDataSource_Read_NilClient(t *testing.T) {
	r := &QueryNetflowMonitorTimeSeriesDataSource{}
	m := QueryNetflowMonitorTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryNetflowMonitorTimeSeriesDataSource_Read_BuildError exercises QueryNetflowMonitorTimeSeriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryNetflowMonitorTimeSeriesDataSource_Read_BuildError(t *testing.T) {
	r := &QueryNetflowMonitorTimeSeriesDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryNetflowMonitorTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryNetflowMonitorTimeSeriesDataSource_Read_SendError exercises QueryNetflowMonitorTimeSeriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryNetflowMonitorTimeSeriesDataSource_Read_SendError(t *testing.T) {
	r := &QueryNetflowMonitorTimeSeriesDataSource{client: newTransportErrorClient(t)}
	m := QueryNetflowMonitorTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryNetflowMonitorTimeSeriesDataSource_Read_NotFound exercises QueryNetflowMonitorTimeSeriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryNetflowMonitorTimeSeriesDataSource_Read_NotFound(t *testing.T) {
	r := &QueryNetflowMonitorTimeSeriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryNetflowMonitorTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryNetflowMonitorTimeSeriesDataSource_Read_APIError exercises QueryNetflowMonitorTimeSeriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryNetflowMonitorTimeSeriesDataSource_Read_APIError(t *testing.T) {
	r := &QueryNetflowMonitorTimeSeriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryNetflowMonitorTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_netflow_monitor_time_series")
}

// TestQueryNetflowMonitorTimeSeriesDataSource_Read_APIErrorReadBody exercises QueryNetflowMonitorTimeSeriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryNetflowMonitorTimeSeriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryNetflowMonitorTimeSeriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryNetflowMonitorTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryNetflowMonitorTimeSeriesDataSource_Read_InvalidJSON exercises QueryNetflowMonitorTimeSeriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryNetflowMonitorTimeSeriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryNetflowMonitorTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryNetflowMonitorTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
