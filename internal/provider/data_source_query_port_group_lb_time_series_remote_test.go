package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryPortGroupLbTimeSeriesDataSource_Read_Happy exercises QueryPortGroupLbTimeSeriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestQueryPortGroupLbTimeSeriesDataSource_Read_Happy(t *testing.T) {
	r := &QueryPortGroupLbTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := QueryPortGroupLbTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryPortGroupLbTimeSeriesDataSource_Read_NilClient exercises QueryPortGroupLbTimeSeriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryPortGroupLbTimeSeriesDataSource_Read_NilClient(t *testing.T) {
	r := &QueryPortGroupLbTimeSeriesDataSource{}
	m := QueryPortGroupLbTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryPortGroupLbTimeSeriesDataSource_Read_BuildError exercises QueryPortGroupLbTimeSeriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestQueryPortGroupLbTimeSeriesDataSource_Read_BuildError(t *testing.T) {
	r := &QueryPortGroupLbTimeSeriesDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryPortGroupLbTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestQueryPortGroupLbTimeSeriesDataSource_Read_SendError exercises QueryPortGroupLbTimeSeriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestQueryPortGroupLbTimeSeriesDataSource_Read_SendError(t *testing.T) {
	r := &QueryPortGroupLbTimeSeriesDataSource{client: newTransportErrorClient(t)}
	m := QueryPortGroupLbTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestQueryPortGroupLbTimeSeriesDataSource_Read_NotFound exercises QueryPortGroupLbTimeSeriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestQueryPortGroupLbTimeSeriesDataSource_Read_NotFound(t *testing.T) {
	r := &QueryPortGroupLbTimeSeriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := QueryPortGroupLbTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestQueryPortGroupLbTimeSeriesDataSource_Read_APIError exercises QueryPortGroupLbTimeSeriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestQueryPortGroupLbTimeSeriesDataSource_Read_APIError(t *testing.T) {
	r := &QueryPortGroupLbTimeSeriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := QueryPortGroupLbTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_query_port_group_lb_time_series")
}

// TestQueryPortGroupLbTimeSeriesDataSource_Read_APIErrorReadBody exercises QueryPortGroupLbTimeSeriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestQueryPortGroupLbTimeSeriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &QueryPortGroupLbTimeSeriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := QueryPortGroupLbTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestQueryPortGroupLbTimeSeriesDataSource_Read_InvalidJSON exercises QueryPortGroupLbTimeSeriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestQueryPortGroupLbTimeSeriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryPortGroupLbTimeSeriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryPortGroupLbTimeSeriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
