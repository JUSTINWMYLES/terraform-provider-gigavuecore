package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFmCryptoStatusDataSource_Read_Happy exercises GetFmCryptoStatusDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFmCryptoStatusDataSource_Read_Happy(t *testing.T) {
	r := &GetFmCryptoStatusDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFmCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFmCryptoStatusDataSource_Read_NilClient exercises GetFmCryptoStatusDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFmCryptoStatusDataSource_Read_NilClient(t *testing.T) {
	r := &GetFmCryptoStatusDataSource{}
	m := GetFmCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFmCryptoStatusDataSource_Read_BuildError exercises GetFmCryptoStatusDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFmCryptoStatusDataSource_Read_BuildError(t *testing.T) {
	r := &GetFmCryptoStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFmCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFmCryptoStatusDataSource_Read_SendError exercises GetFmCryptoStatusDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFmCryptoStatusDataSource_Read_SendError(t *testing.T) {
	r := &GetFmCryptoStatusDataSource{client: newTransportErrorClient(t)}
	m := GetFmCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFmCryptoStatusDataSource_Read_NotFound exercises GetFmCryptoStatusDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFmCryptoStatusDataSource_Read_NotFound(t *testing.T) {
	r := &GetFmCryptoStatusDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFmCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFmCryptoStatusDataSource_Read_APIError exercises GetFmCryptoStatusDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFmCryptoStatusDataSource_Read_APIError(t *testing.T) {
	r := &GetFmCryptoStatusDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFmCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_fm_crypto_status")
}

// TestGetFmCryptoStatusDataSource_Read_APIErrorReadBody exercises GetFmCryptoStatusDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFmCryptoStatusDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFmCryptoStatusDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFmCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFmCryptoStatusDataSource_Read_InvalidJSON exercises GetFmCryptoStatusDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFmCryptoStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFmCryptoStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFmCryptoStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
