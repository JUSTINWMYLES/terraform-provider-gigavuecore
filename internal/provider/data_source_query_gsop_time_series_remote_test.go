package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryGsopTimeSeriesDataSource_Read_Happy exercises QueryGsopTimeSeriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryGsopTimeSeriesDataSource_Read_Happy(t *testing.T) {
	r := &QueryGsopTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryGsopTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryGsopTimeSeriesDataSource_Read_NilClient exercises QueryGsopTimeSeriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryGsopTimeSeriesDataSource_Read_NilClient(t *testing.T) {
	r := &QueryGsopTimeSeriesDataSource{}
	m := QueryGsopTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryGsopTimeSeriesDataSource_Read_BuildError exercises QueryGsopTimeSeriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryGsopTimeSeriesDataSource_Read_BuildError(t *testing.T) {
	r := &QueryGsopTimeSeriesDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryGsopTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryGsopTimeSeriesDataSource_Read_SendError exercises QueryGsopTimeSeriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryGsopTimeSeriesDataSource_Read_SendError(t *testing.T) {
	r := &QueryGsopTimeSeriesDataSource{client: newTransportErrorClient(t)}
	m := QueryGsopTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryGsopTimeSeriesDataSource_Read_NotFound exercises QueryGsopTimeSeriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryGsopTimeSeriesDataSource_Read_NotFound(t *testing.T) {
	r := &QueryGsopTimeSeriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryGsopTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryGsopTimeSeriesDataSource_Read_APIError exercises QueryGsopTimeSeriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryGsopTimeSeriesDataSource_Read_APIError(t *testing.T) {
	r := &QueryGsopTimeSeriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryGsopTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_gsop_time_series")
}

// TestQueryGsopTimeSeriesDataSource_Read_APIErrorReadBody exercises QueryGsopTimeSeriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryGsopTimeSeriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryGsopTimeSeriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryGsopTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryGsopTimeSeriesDataSource_Read_InvalidJSON exercises QueryGsopTimeSeriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryGsopTimeSeriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryGsopTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryGsopTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
