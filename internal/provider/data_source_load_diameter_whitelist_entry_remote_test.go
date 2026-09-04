package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadDiameterWhitelistEntryDataSource_Read_Happy exercises LoadDiameterWhitelistEntryDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadDiameterWhitelistEntryDataSource_Read_Happy(t *testing.T) {
	r := &LoadDiameterWhitelistEntryDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadDiameterWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadDiameterWhitelistEntryDataSource_Read_NilClient exercises LoadDiameterWhitelistEntryDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadDiameterWhitelistEntryDataSource_Read_NilClient(t *testing.T) {
	r := &LoadDiameterWhitelistEntryDataSource{}
	m := LoadDiameterWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadDiameterWhitelistEntryDataSource_Read_BuildError exercises LoadDiameterWhitelistEntryDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadDiameterWhitelistEntryDataSource_Read_BuildError(t *testing.T) {
	r := &LoadDiameterWhitelistEntryDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadDiameterWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadDiameterWhitelistEntryDataSource_Read_SendError exercises LoadDiameterWhitelistEntryDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadDiameterWhitelistEntryDataSource_Read_SendError(t *testing.T) {
	r := &LoadDiameterWhitelistEntryDataSource{client: newTransportErrorClient(t)}
	m := LoadDiameterWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadDiameterWhitelistEntryDataSource_Read_NotFound exercises LoadDiameterWhitelistEntryDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadDiameterWhitelistEntryDataSource_Read_NotFound(t *testing.T) {
	r := &LoadDiameterWhitelistEntryDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadDiameterWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadDiameterWhitelistEntryDataSource_Read_APIError exercises LoadDiameterWhitelistEntryDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadDiameterWhitelistEntryDataSource_Read_APIError(t *testing.T) {
	r := &LoadDiameterWhitelistEntryDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadDiameterWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_diameter_whitelist_entry")
}

// TestLoadDiameterWhitelistEntryDataSource_Read_APIErrorReadBody exercises LoadDiameterWhitelistEntryDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadDiameterWhitelistEntryDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadDiameterWhitelistEntryDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadDiameterWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadDiameterWhitelistEntryDataSource_Read_InvalidJSON exercises LoadDiameterWhitelistEntryDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadDiameterWhitelistEntryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadDiameterWhitelistEntryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadDiameterWhitelistEntryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
