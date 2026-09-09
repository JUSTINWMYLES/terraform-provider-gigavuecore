package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetLastNPeriodsSummaryWithUnitDataSource_Read_Happy exercises GetLastNPeriodsSummaryWithUnitDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetLastNPeriodsSummaryWithUnitDataSource_Read_Happy(t *testing.T) {
	r := &GetLastNPeriodsSummaryWithUnitDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetLastNPeriodsSummaryWithUnitDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetLastNPeriodsSummaryWithUnitDataSource_Read_NilClient exercises GetLastNPeriodsSummaryWithUnitDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetLastNPeriodsSummaryWithUnitDataSource_Read_NilClient(t *testing.T) {
	r := &GetLastNPeriodsSummaryWithUnitDataSource{}
	m := GetLastNPeriodsSummaryWithUnitDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetLastNPeriodsSummaryWithUnitDataSource_Read_BuildError exercises GetLastNPeriodsSummaryWithUnitDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetLastNPeriodsSummaryWithUnitDataSource_Read_BuildError(t *testing.T) {
	r := &GetLastNPeriodsSummaryWithUnitDataSource{client: newMalformedBaseURLClient(t)}
	m := GetLastNPeriodsSummaryWithUnitDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetLastNPeriodsSummaryWithUnitDataSource_Read_SendError exercises GetLastNPeriodsSummaryWithUnitDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetLastNPeriodsSummaryWithUnitDataSource_Read_SendError(t *testing.T) {
	r := &GetLastNPeriodsSummaryWithUnitDataSource{client: newTransportErrorClient(t)}
	m := GetLastNPeriodsSummaryWithUnitDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetLastNPeriodsSummaryWithUnitDataSource_Read_NotFound exercises GetLastNPeriodsSummaryWithUnitDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetLastNPeriodsSummaryWithUnitDataSource_Read_NotFound(t *testing.T) {
	r := &GetLastNPeriodsSummaryWithUnitDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetLastNPeriodsSummaryWithUnitDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetLastNPeriodsSummaryWithUnitDataSource_Read_APIError exercises GetLastNPeriodsSummaryWithUnitDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetLastNPeriodsSummaryWithUnitDataSource_Read_APIError(t *testing.T) {
	r := &GetLastNPeriodsSummaryWithUnitDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetLastNPeriodsSummaryWithUnitDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_last_n_periods_summary_with_unit")
}

// TestGetLastNPeriodsSummaryWithUnitDataSource_Read_APIErrorReadBody exercises GetLastNPeriodsSummaryWithUnitDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetLastNPeriodsSummaryWithUnitDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetLastNPeriodsSummaryWithUnitDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetLastNPeriodsSummaryWithUnitDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetLastNPeriodsSummaryWithUnitDataSource_Read_InvalidJSON exercises GetLastNPeriodsSummaryWithUnitDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetLastNPeriodsSummaryWithUnitDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetLastNPeriodsSummaryWithUnitDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetLastNPeriodsSummaryWithUnitDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
