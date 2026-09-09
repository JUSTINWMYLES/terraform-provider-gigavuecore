package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryVportTimeSeriesDataSource_Read_Happy exercises QueryVportTimeSeriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryVportTimeSeriesDataSource_Read_Happy(t *testing.T) {
	r := &QueryVportTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryVportTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryVportTimeSeriesDataSource_Read_NilClient exercises QueryVportTimeSeriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryVportTimeSeriesDataSource_Read_NilClient(t *testing.T) {
	r := &QueryVportTimeSeriesDataSource{}
	m := QueryVportTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryVportTimeSeriesDataSource_Read_BuildError exercises QueryVportTimeSeriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryVportTimeSeriesDataSource_Read_BuildError(t *testing.T) {
	r := &QueryVportTimeSeriesDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryVportTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryVportTimeSeriesDataSource_Read_SendError exercises QueryVportTimeSeriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryVportTimeSeriesDataSource_Read_SendError(t *testing.T) {
	r := &QueryVportTimeSeriesDataSource{client: newTransportErrorClient(t)}
	m := QueryVportTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryVportTimeSeriesDataSource_Read_NotFound exercises QueryVportTimeSeriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryVportTimeSeriesDataSource_Read_NotFound(t *testing.T) {
	r := &QueryVportTimeSeriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryVportTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryVportTimeSeriesDataSource_Read_APIError exercises QueryVportTimeSeriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryVportTimeSeriesDataSource_Read_APIError(t *testing.T) {
	r := &QueryVportTimeSeriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryVportTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_vport_time_series")
}

// TestQueryVportTimeSeriesDataSource_Read_APIErrorReadBody exercises QueryVportTimeSeriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryVportTimeSeriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryVportTimeSeriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryVportTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryVportTimeSeriesDataSource_Read_InvalidJSON exercises QueryVportTimeSeriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryVportTimeSeriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryVportTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryVportTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
