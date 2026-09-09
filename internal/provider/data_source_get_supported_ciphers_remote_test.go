package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSupportedCiphersDataSource_Read_Happy exercises GetSupportedCiphersDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSupportedCiphersDataSource_Read_Happy(t *testing.T) {
	r := &GetSupportedCiphersDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSupportedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSupportedCiphersDataSource_Read_NilClient exercises GetSupportedCiphersDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSupportedCiphersDataSource_Read_NilClient(t *testing.T) {
	r := &GetSupportedCiphersDataSource{}
	m := GetSupportedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSupportedCiphersDataSource_Read_BuildError exercises GetSupportedCiphersDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSupportedCiphersDataSource_Read_BuildError(t *testing.T) {
	r := &GetSupportedCiphersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSupportedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSupportedCiphersDataSource_Read_SendError exercises GetSupportedCiphersDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSupportedCiphersDataSource_Read_SendError(t *testing.T) {
	r := &GetSupportedCiphersDataSource{client: newTransportErrorClient(t)}
	m := GetSupportedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSupportedCiphersDataSource_Read_NotFound exercises GetSupportedCiphersDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSupportedCiphersDataSource_Read_NotFound(t *testing.T) {
	r := &GetSupportedCiphersDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSupportedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSupportedCiphersDataSource_Read_APIError exercises GetSupportedCiphersDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSupportedCiphersDataSource_Read_APIError(t *testing.T) {
	r := &GetSupportedCiphersDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSupportedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_supported_ciphers")
}

// TestGetSupportedCiphersDataSource_Read_APIErrorReadBody exercises GetSupportedCiphersDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSupportedCiphersDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSupportedCiphersDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSupportedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSupportedCiphersDataSource_Read_InvalidJSON exercises GetSupportedCiphersDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSupportedCiphersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSupportedCiphersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSupportedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
