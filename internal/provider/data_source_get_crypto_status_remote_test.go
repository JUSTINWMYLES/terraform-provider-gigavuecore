package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCryptoStatusDataSource_Read_Happy exercises GetCryptoStatusDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetCryptoStatusDataSource_Read_Happy(t *testing.T) {
	r := &GetCryptoStatusDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCryptoStatusDataSource_Read_NilClient exercises GetCryptoStatusDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCryptoStatusDataSource_Read_NilClient(t *testing.T) {
	r := &GetCryptoStatusDataSource{}
	m := GetCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCryptoStatusDataSource_Read_BuildError exercises GetCryptoStatusDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetCryptoStatusDataSource_Read_BuildError(t *testing.T) {
	r := &GetCryptoStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetCryptoStatusDataSource_Read_SendError exercises GetCryptoStatusDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetCryptoStatusDataSource_Read_SendError(t *testing.T) {
	r := &GetCryptoStatusDataSource{client: newTransportErrorClient(t)}
	m := GetCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetCryptoStatusDataSource_Read_NotFound exercises GetCryptoStatusDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetCryptoStatusDataSource_Read_NotFound(t *testing.T) {
	r := &GetCryptoStatusDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetCryptoStatusDataSource_Read_APIError exercises GetCryptoStatusDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetCryptoStatusDataSource_Read_APIError(t *testing.T) {
	r := &GetCryptoStatusDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_crypto_status")
}

// TestGetCryptoStatusDataSource_Read_APIErrorReadBody exercises GetCryptoStatusDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetCryptoStatusDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetCryptoStatusDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetCryptoStatusDataSource_Read_InvalidJSON exercises GetCryptoStatusDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetCryptoStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCryptoStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
