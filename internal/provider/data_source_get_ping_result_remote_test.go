package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPingResultDataSource_Read_Happy exercises GetPingResultDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetPingResultDataSource_Read_Happy(t *testing.T) {
	r := &GetPingResultDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetPingResultDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPingResultDataSource_Read_NilClient exercises GetPingResultDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPingResultDataSource_Read_NilClient(t *testing.T) {
	r := &GetPingResultDataSource{}
	m := GetPingResultDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPingResultDataSource_Read_BuildError exercises GetPingResultDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPingResultDataSource_Read_BuildError(t *testing.T) {
	r := &GetPingResultDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPingResultDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPingResultDataSource_Read_SendError exercises GetPingResultDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPingResultDataSource_Read_SendError(t *testing.T) {
	r := &GetPingResultDataSource{client: newTransportErrorClient(t)}
	m := GetPingResultDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPingResultDataSource_Read_NotFound exercises GetPingResultDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetPingResultDataSource_Read_NotFound(t *testing.T) {
	r := &GetPingResultDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetPingResultDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetPingResultDataSource_Read_APIError exercises GetPingResultDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPingResultDataSource_Read_APIError(t *testing.T) {
	r := &GetPingResultDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetPingResultDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ping_result")
}

// TestGetPingResultDataSource_Read_APIErrorReadBody exercises GetPingResultDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPingResultDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetPingResultDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetPingResultDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetPingResultDataSource_Read_InvalidJSON exercises GetPingResultDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetPingResultDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPingResultDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPingResultDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
