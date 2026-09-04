package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryGsGroupTimeSeriesDataSource_Read_Happy exercises QueryGsGroupTimeSeriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryGsGroupTimeSeriesDataSource_Read_Happy(t *testing.T) {
	r := &QueryGsGroupTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryGsGroupTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryGsGroupTimeSeriesDataSource_Read_NilClient exercises QueryGsGroupTimeSeriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryGsGroupTimeSeriesDataSource_Read_NilClient(t *testing.T) {
	r := &QueryGsGroupTimeSeriesDataSource{}
	m := QueryGsGroupTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryGsGroupTimeSeriesDataSource_Read_BuildError exercises QueryGsGroupTimeSeriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryGsGroupTimeSeriesDataSource_Read_BuildError(t *testing.T) {
	r := &QueryGsGroupTimeSeriesDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryGsGroupTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryGsGroupTimeSeriesDataSource_Read_SendError exercises QueryGsGroupTimeSeriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryGsGroupTimeSeriesDataSource_Read_SendError(t *testing.T) {
	r := &QueryGsGroupTimeSeriesDataSource{client: newTransportErrorClient(t)}
	m := QueryGsGroupTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryGsGroupTimeSeriesDataSource_Read_NotFound exercises QueryGsGroupTimeSeriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryGsGroupTimeSeriesDataSource_Read_NotFound(t *testing.T) {
	r := &QueryGsGroupTimeSeriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryGsGroupTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryGsGroupTimeSeriesDataSource_Read_APIError exercises QueryGsGroupTimeSeriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryGsGroupTimeSeriesDataSource_Read_APIError(t *testing.T) {
	r := &QueryGsGroupTimeSeriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryGsGroupTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_gs_group_time_series")
}

// TestQueryGsGroupTimeSeriesDataSource_Read_APIErrorReadBody exercises QueryGsGroupTimeSeriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryGsGroupTimeSeriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryGsGroupTimeSeriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryGsGroupTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryGsGroupTimeSeriesDataSource_Read_InvalidJSON exercises QueryGsGroupTimeSeriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryGsGroupTimeSeriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryGsGroupTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryGsGroupTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
