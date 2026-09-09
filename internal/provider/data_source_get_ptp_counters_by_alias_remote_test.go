package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPtpCountersByAliasDataSource_Read_Happy exercises GetPtpCountersByAliasDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetPtpCountersByAliasDataSource_Read_Happy(t *testing.T) {
	r := &GetPtpCountersByAliasDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetPtpCountersByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPtpCountersByAliasDataSource_Read_NilClient exercises GetPtpCountersByAliasDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPtpCountersByAliasDataSource_Read_NilClient(t *testing.T) {
	r := &GetPtpCountersByAliasDataSource{}
	m := GetPtpCountersByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPtpCountersByAliasDataSource_Read_BuildError exercises GetPtpCountersByAliasDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPtpCountersByAliasDataSource_Read_BuildError(t *testing.T) {
	r := &GetPtpCountersByAliasDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPtpCountersByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPtpCountersByAliasDataSource_Read_SendError exercises GetPtpCountersByAliasDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPtpCountersByAliasDataSource_Read_SendError(t *testing.T) {
	r := &GetPtpCountersByAliasDataSource{client: newTransportErrorClient(t)}
	m := GetPtpCountersByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPtpCountersByAliasDataSource_Read_NotFound exercises GetPtpCountersByAliasDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetPtpCountersByAliasDataSource_Read_NotFound(t *testing.T) {
	r := &GetPtpCountersByAliasDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetPtpCountersByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetPtpCountersByAliasDataSource_Read_APIError exercises GetPtpCountersByAliasDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPtpCountersByAliasDataSource_Read_APIError(t *testing.T) {
	r := &GetPtpCountersByAliasDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetPtpCountersByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ptp_counters_by_alias")
}

// TestGetPtpCountersByAliasDataSource_Read_APIErrorReadBody exercises GetPtpCountersByAliasDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPtpCountersByAliasDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetPtpCountersByAliasDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetPtpCountersByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetPtpCountersByAliasDataSource_Read_InvalidJSON exercises GetPtpCountersByAliasDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetPtpCountersByAliasDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPtpCountersByAliasDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPtpCountersByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
