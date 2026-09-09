package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowSamplingSummaryDataSource_Read_Happy exercises GetFlowSamplingSummaryDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFlowSamplingSummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetFlowSamplingSummaryDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFlowSamplingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlowSamplingSummaryDataSource_Read_NilClient exercises GetFlowSamplingSummaryDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlowSamplingSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlowSamplingSummaryDataSource{}
	m := GetFlowSamplingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlowSamplingSummaryDataSource_Read_BuildError exercises GetFlowSamplingSummaryDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFlowSamplingSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlowSamplingSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlowSamplingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFlowSamplingSummaryDataSource_Read_SendError exercises GetFlowSamplingSummaryDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFlowSamplingSummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetFlowSamplingSummaryDataSource{client: newTransportErrorClient(t)}
	m := GetFlowSamplingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFlowSamplingSummaryDataSource_Read_NotFound exercises GetFlowSamplingSummaryDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFlowSamplingSummaryDataSource_Read_NotFound(t *testing.T) {
	r := &GetFlowSamplingSummaryDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFlowSamplingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFlowSamplingSummaryDataSource_Read_APIError exercises GetFlowSamplingSummaryDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFlowSamplingSummaryDataSource_Read_APIError(t *testing.T) {
	r := &GetFlowSamplingSummaryDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFlowSamplingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_flow_sampling_summary")
}

// TestGetFlowSamplingSummaryDataSource_Read_APIErrorReadBody exercises GetFlowSamplingSummaryDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFlowSamplingSummaryDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFlowSamplingSummaryDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFlowSamplingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFlowSamplingSummaryDataSource_Read_InvalidJSON exercises GetFlowSamplingSummaryDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFlowSamplingSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlowSamplingSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlowSamplingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
