package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetHsmGroupStatusReportDataSource_Read_Happy exercises GetHsmGroupStatusReportDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetHsmGroupStatusReportDataSource_Read_Happy(t *testing.T) {
	r := &GetHsmGroupStatusReportDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetHsmGroupStatusReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetHsmGroupStatusReportDataSource_Read_NilClient exercises GetHsmGroupStatusReportDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetHsmGroupStatusReportDataSource_Read_NilClient(t *testing.T) {
	r := &GetHsmGroupStatusReportDataSource{}
	m := GetHsmGroupStatusReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetHsmGroupStatusReportDataSource_Read_BuildError exercises GetHsmGroupStatusReportDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetHsmGroupStatusReportDataSource_Read_BuildError(t *testing.T) {
	r := &GetHsmGroupStatusReportDataSource{client: newMalformedBaseURLClient(t)}
	m := GetHsmGroupStatusReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetHsmGroupStatusReportDataSource_Read_SendError exercises GetHsmGroupStatusReportDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetHsmGroupStatusReportDataSource_Read_SendError(t *testing.T) {
	r := &GetHsmGroupStatusReportDataSource{client: newTransportErrorClient(t)}
	m := GetHsmGroupStatusReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetHsmGroupStatusReportDataSource_Read_NotFound exercises GetHsmGroupStatusReportDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetHsmGroupStatusReportDataSource_Read_NotFound(t *testing.T) {
	r := &GetHsmGroupStatusReportDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetHsmGroupStatusReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetHsmGroupStatusReportDataSource_Read_APIError exercises GetHsmGroupStatusReportDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetHsmGroupStatusReportDataSource_Read_APIError(t *testing.T) {
	r := &GetHsmGroupStatusReportDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetHsmGroupStatusReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_hsm_group_status_report")
}

// TestGetHsmGroupStatusReportDataSource_Read_APIErrorReadBody exercises GetHsmGroupStatusReportDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetHsmGroupStatusReportDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetHsmGroupStatusReportDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetHsmGroupStatusReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetHsmGroupStatusReportDataSource_Read_InvalidJSON exercises GetHsmGroupStatusReportDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetHsmGroupStatusReportDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetHsmGroupStatusReportDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetHsmGroupStatusReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
