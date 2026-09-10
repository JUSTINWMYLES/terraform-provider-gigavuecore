package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowFilteringSummaryDataSource_Read_Happy exercises GetFlowFilteringSummaryDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFlowFilteringSummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetFlowFilteringSummaryDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlowFilteringSummaryDataSource_Read_NilClient exercises GetFlowFilteringSummaryDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlowFilteringSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlowFilteringSummaryDataSource{}
	m := GetFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlowFilteringSummaryDataSource_Read_BuildError exercises GetFlowFilteringSummaryDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFlowFilteringSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlowFilteringSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFlowFilteringSummaryDataSource_Read_SendError exercises GetFlowFilteringSummaryDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFlowFilteringSummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetFlowFilteringSummaryDataSource{client: newTransportErrorClient(t)}
	m := GetFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFlowFilteringSummaryDataSource_Read_NotFound exercises GetFlowFilteringSummaryDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFlowFilteringSummaryDataSource_Read_NotFound(t *testing.T) {
	r := &GetFlowFilteringSummaryDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFlowFilteringSummaryDataSource_Read_APIError exercises GetFlowFilteringSummaryDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFlowFilteringSummaryDataSource_Read_APIError(t *testing.T) {
	r := &GetFlowFilteringSummaryDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_flow_filtering_summary")
}

// TestGetFlowFilteringSummaryDataSource_Read_APIErrorReadBody exercises GetFlowFilteringSummaryDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFlowFilteringSummaryDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFlowFilteringSummaryDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFlowFilteringSummaryDataSource_Read_InvalidJSON exercises GetFlowFilteringSummaryDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFlowFilteringSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlowFilteringSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
