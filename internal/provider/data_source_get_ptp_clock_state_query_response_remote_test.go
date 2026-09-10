package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPtpClockStateQueryResponseDataSource_Read_Happy exercises GetPtpClockStateQueryResponseDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetPtpClockStateQueryResponseDataSource_Read_Happy(t *testing.T) {
	r := &GetPtpClockStateQueryResponseDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetPtpClockStateQueryResponseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPtpClockStateQueryResponseDataSource_Read_NilClient exercises GetPtpClockStateQueryResponseDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPtpClockStateQueryResponseDataSource_Read_NilClient(t *testing.T) {
	r := &GetPtpClockStateQueryResponseDataSource{}
	m := GetPtpClockStateQueryResponseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPtpClockStateQueryResponseDataSource_Read_BuildError exercises GetPtpClockStateQueryResponseDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPtpClockStateQueryResponseDataSource_Read_BuildError(t *testing.T) {
	r := &GetPtpClockStateQueryResponseDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPtpClockStateQueryResponseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPtpClockStateQueryResponseDataSource_Read_SendError exercises GetPtpClockStateQueryResponseDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPtpClockStateQueryResponseDataSource_Read_SendError(t *testing.T) {
	r := &GetPtpClockStateQueryResponseDataSource{client: newTransportErrorClient(t)}
	m := GetPtpClockStateQueryResponseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPtpClockStateQueryResponseDataSource_Read_NotFound exercises GetPtpClockStateQueryResponseDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetPtpClockStateQueryResponseDataSource_Read_NotFound(t *testing.T) {
	r := &GetPtpClockStateQueryResponseDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetPtpClockStateQueryResponseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetPtpClockStateQueryResponseDataSource_Read_APIError exercises GetPtpClockStateQueryResponseDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPtpClockStateQueryResponseDataSource_Read_APIError(t *testing.T) {
	r := &GetPtpClockStateQueryResponseDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetPtpClockStateQueryResponseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ptp_clock_state_query_response")
}

// TestGetPtpClockStateQueryResponseDataSource_Read_APIErrorReadBody exercises GetPtpClockStateQueryResponseDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPtpClockStateQueryResponseDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetPtpClockStateQueryResponseDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetPtpClockStateQueryResponseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetPtpClockStateQueryResponseDataSource_Read_InvalidJSON exercises GetPtpClockStateQueryResponseDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetPtpClockStateQueryResponseDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPtpClockStateQueryResponseDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPtpClockStateQueryResponseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
