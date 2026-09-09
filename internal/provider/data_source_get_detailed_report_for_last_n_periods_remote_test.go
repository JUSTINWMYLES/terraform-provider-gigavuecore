package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetDetailedReportForLastNPeriodsDataSource_Read_Happy exercises GetDetailedReportForLastNPeriodsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetDetailedReportForLastNPeriodsDataSource_Read_Happy(t *testing.T) {
	r := &GetDetailedReportForLastNPeriodsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetDetailedReportForLastNPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetDetailedReportForLastNPeriodsDataSource_Read_NilClient exercises GetDetailedReportForLastNPeriodsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetDetailedReportForLastNPeriodsDataSource_Read_NilClient(t *testing.T) {
	r := &GetDetailedReportForLastNPeriodsDataSource{}
	m := GetDetailedReportForLastNPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetDetailedReportForLastNPeriodsDataSource_Read_BuildError exercises GetDetailedReportForLastNPeriodsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetDetailedReportForLastNPeriodsDataSource_Read_BuildError(t *testing.T) {
	r := &GetDetailedReportForLastNPeriodsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetDetailedReportForLastNPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetDetailedReportForLastNPeriodsDataSource_Read_SendError exercises GetDetailedReportForLastNPeriodsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetDetailedReportForLastNPeriodsDataSource_Read_SendError(t *testing.T) {
	r := &GetDetailedReportForLastNPeriodsDataSource{client: newTransportErrorClient(t)}
	m := GetDetailedReportForLastNPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetDetailedReportForLastNPeriodsDataSource_Read_NotFound exercises GetDetailedReportForLastNPeriodsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetDetailedReportForLastNPeriodsDataSource_Read_NotFound(t *testing.T) {
	r := &GetDetailedReportForLastNPeriodsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetDetailedReportForLastNPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetDetailedReportForLastNPeriodsDataSource_Read_APIError exercises GetDetailedReportForLastNPeriodsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetDetailedReportForLastNPeriodsDataSource_Read_APIError(t *testing.T) {
	r := &GetDetailedReportForLastNPeriodsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetDetailedReportForLastNPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_detailed_report_for_last_n_periods")
}

// TestGetDetailedReportForLastNPeriodsDataSource_Read_APIErrorReadBody exercises GetDetailedReportForLastNPeriodsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetDetailedReportForLastNPeriodsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetDetailedReportForLastNPeriodsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetDetailedReportForLastNPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetDetailedReportForLastNPeriodsDataSource_Read_InvalidJSON exercises GetDetailedReportForLastNPeriodsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetDetailedReportForLastNPeriodsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetDetailedReportForLastNPeriodsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetDetailedReportForLastNPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
