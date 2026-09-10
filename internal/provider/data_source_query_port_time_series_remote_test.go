package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryPortTimeSeriesDataSource_Read_Happy exercises QueryPortTimeSeriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryPortTimeSeriesDataSource_Read_Happy(t *testing.T) {
	r := &QueryPortTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryPortTimeSeriesDataSource_Read_NilClient exercises QueryPortTimeSeriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryPortTimeSeriesDataSource_Read_NilClient(t *testing.T) {
	r := &QueryPortTimeSeriesDataSource{}
	m := QueryPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryPortTimeSeriesDataSource_Read_BuildError exercises QueryPortTimeSeriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryPortTimeSeriesDataSource_Read_BuildError(t *testing.T) {
	r := &QueryPortTimeSeriesDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryPortTimeSeriesDataSource_Read_SendError exercises QueryPortTimeSeriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryPortTimeSeriesDataSource_Read_SendError(t *testing.T) {
	r := &QueryPortTimeSeriesDataSource{client: newTransportErrorClient(t)}
	m := QueryPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryPortTimeSeriesDataSource_Read_NotFound exercises QueryPortTimeSeriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryPortTimeSeriesDataSource_Read_NotFound(t *testing.T) {
	r := &QueryPortTimeSeriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryPortTimeSeriesDataSource_Read_APIError exercises QueryPortTimeSeriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryPortTimeSeriesDataSource_Read_APIError(t *testing.T) {
	r := &QueryPortTimeSeriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_port_time_series")
}

// TestQueryPortTimeSeriesDataSource_Read_APIErrorReadBody exercises QueryPortTimeSeriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryPortTimeSeriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryPortTimeSeriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryPortTimeSeriesDataSource_Read_InvalidJSON exercises QueryPortTimeSeriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryPortTimeSeriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryPortTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
