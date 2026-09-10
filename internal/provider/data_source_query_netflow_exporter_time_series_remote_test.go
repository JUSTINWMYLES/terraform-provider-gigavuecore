package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryNetflowExporterTimeSeriesDataSource_Read_Happy exercises QueryNetflowExporterTimeSeriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryNetflowExporterTimeSeriesDataSource_Read_Happy(t *testing.T) {
	r := &QueryNetflowExporterTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryNetflowExporterTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryNetflowExporterTimeSeriesDataSource_Read_NilClient exercises QueryNetflowExporterTimeSeriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryNetflowExporterTimeSeriesDataSource_Read_NilClient(t *testing.T) {
	r := &QueryNetflowExporterTimeSeriesDataSource{}
	m := QueryNetflowExporterTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryNetflowExporterTimeSeriesDataSource_Read_BuildError exercises QueryNetflowExporterTimeSeriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryNetflowExporterTimeSeriesDataSource_Read_BuildError(t *testing.T) {
	r := &QueryNetflowExporterTimeSeriesDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryNetflowExporterTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryNetflowExporterTimeSeriesDataSource_Read_SendError exercises QueryNetflowExporterTimeSeriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryNetflowExporterTimeSeriesDataSource_Read_SendError(t *testing.T) {
	r := &QueryNetflowExporterTimeSeriesDataSource{client: newTransportErrorClient(t)}
	m := QueryNetflowExporterTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryNetflowExporterTimeSeriesDataSource_Read_NotFound exercises QueryNetflowExporterTimeSeriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryNetflowExporterTimeSeriesDataSource_Read_NotFound(t *testing.T) {
	r := &QueryNetflowExporterTimeSeriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryNetflowExporterTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryNetflowExporterTimeSeriesDataSource_Read_APIError exercises QueryNetflowExporterTimeSeriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryNetflowExporterTimeSeriesDataSource_Read_APIError(t *testing.T) {
	r := &QueryNetflowExporterTimeSeriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryNetflowExporterTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_netflow_exporter_time_series")
}

// TestQueryNetflowExporterTimeSeriesDataSource_Read_APIErrorReadBody exercises QueryNetflowExporterTimeSeriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryNetflowExporterTimeSeriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryNetflowExporterTimeSeriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryNetflowExporterTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryNetflowExporterTimeSeriesDataSource_Read_InvalidJSON exercises QueryNetflowExporterTimeSeriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryNetflowExporterTimeSeriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryNetflowExporterTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryNetflowExporterTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
