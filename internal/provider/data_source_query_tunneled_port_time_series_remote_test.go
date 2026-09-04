package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryTunneledPortTimeSeriesDataSource_Read_Happy exercises QueryTunneledPortTimeSeriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryTunneledPortTimeSeriesDataSource_Read_Happy(t *testing.T) {
	r := &QueryTunneledPortTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryTunneledPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryTunneledPortTimeSeriesDataSource_Read_NilClient exercises QueryTunneledPortTimeSeriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryTunneledPortTimeSeriesDataSource_Read_NilClient(t *testing.T) {
	r := &QueryTunneledPortTimeSeriesDataSource{}
	m := QueryTunneledPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryTunneledPortTimeSeriesDataSource_Read_BuildError exercises QueryTunneledPortTimeSeriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryTunneledPortTimeSeriesDataSource_Read_BuildError(t *testing.T) {
	r := &QueryTunneledPortTimeSeriesDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryTunneledPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryTunneledPortTimeSeriesDataSource_Read_SendError exercises QueryTunneledPortTimeSeriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryTunneledPortTimeSeriesDataSource_Read_SendError(t *testing.T) {
	r := &QueryTunneledPortTimeSeriesDataSource{client: newTransportErrorClient(t)}
	m := QueryTunneledPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryTunneledPortTimeSeriesDataSource_Read_NotFound exercises QueryTunneledPortTimeSeriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryTunneledPortTimeSeriesDataSource_Read_NotFound(t *testing.T) {
	r := &QueryTunneledPortTimeSeriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryTunneledPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryTunneledPortTimeSeriesDataSource_Read_APIError exercises QueryTunneledPortTimeSeriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryTunneledPortTimeSeriesDataSource_Read_APIError(t *testing.T) {
	r := &QueryTunneledPortTimeSeriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryTunneledPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_tunneled_port_time_series")
}

// TestQueryTunneledPortTimeSeriesDataSource_Read_APIErrorReadBody exercises QueryTunneledPortTimeSeriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryTunneledPortTimeSeriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryTunneledPortTimeSeriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryTunneledPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryTunneledPortTimeSeriesDataSource_Read_InvalidJSON exercises QueryTunneledPortTimeSeriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryTunneledPortTimeSeriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryTunneledPortTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryTunneledPortTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
