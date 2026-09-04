package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGtpWhitelistEntryDataSource_Read_Happy exercises LoadGtpWhitelistEntryDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadGtpWhitelistEntryDataSource_Read_Happy(t *testing.T) {
	r := &LoadGtpWhitelistEntryDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadGtpWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadGtpWhitelistEntryDataSource_Read_NilClient exercises LoadGtpWhitelistEntryDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadGtpWhitelistEntryDataSource_Read_NilClient(t *testing.T) {
	r := &LoadGtpWhitelistEntryDataSource{}
	m := LoadGtpWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadGtpWhitelistEntryDataSource_Read_BuildError exercises LoadGtpWhitelistEntryDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadGtpWhitelistEntryDataSource_Read_BuildError(t *testing.T) {
	r := &LoadGtpWhitelistEntryDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadGtpWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadGtpWhitelistEntryDataSource_Read_SendError exercises LoadGtpWhitelistEntryDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadGtpWhitelistEntryDataSource_Read_SendError(t *testing.T) {
	r := &LoadGtpWhitelistEntryDataSource{client: newTransportErrorClient(t)}
	m := LoadGtpWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadGtpWhitelistEntryDataSource_Read_NotFound exercises LoadGtpWhitelistEntryDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadGtpWhitelistEntryDataSource_Read_NotFound(t *testing.T) {
	r := &LoadGtpWhitelistEntryDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadGtpWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadGtpWhitelistEntryDataSource_Read_APIError exercises LoadGtpWhitelistEntryDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadGtpWhitelistEntryDataSource_Read_APIError(t *testing.T) {
	r := &LoadGtpWhitelistEntryDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadGtpWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_gtp_whitelist_entry")
}

// TestLoadGtpWhitelistEntryDataSource_Read_APIErrorReadBody exercises LoadGtpWhitelistEntryDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadGtpWhitelistEntryDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadGtpWhitelistEntryDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadGtpWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadGtpWhitelistEntryDataSource_Read_InvalidJSON exercises LoadGtpWhitelistEntryDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadGtpWhitelistEntryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadGtpWhitelistEntryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadGtpWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
