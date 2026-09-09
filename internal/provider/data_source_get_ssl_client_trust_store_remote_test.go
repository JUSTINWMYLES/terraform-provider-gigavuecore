package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslClientTrustStoreDataSource_Read_Happy exercises GetSslClientTrustStoreDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSslClientTrustStoreDataSource_Read_Happy(t *testing.T) {
	r := &GetSslClientTrustStoreDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSslClientTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSslClientTrustStoreDataSource_Read_NilClient exercises GetSslClientTrustStoreDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSslClientTrustStoreDataSource_Read_NilClient(t *testing.T) {
	r := &GetSslClientTrustStoreDataSource{}
	m := GetSslClientTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSslClientTrustStoreDataSource_Read_BuildError exercises GetSslClientTrustStoreDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSslClientTrustStoreDataSource_Read_BuildError(t *testing.T) {
	r := &GetSslClientTrustStoreDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSslClientTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSslClientTrustStoreDataSource_Read_SendError exercises GetSslClientTrustStoreDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSslClientTrustStoreDataSource_Read_SendError(t *testing.T) {
	r := &GetSslClientTrustStoreDataSource{client: newTransportErrorClient(t)}
	m := GetSslClientTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSslClientTrustStoreDataSource_Read_NotFound exercises GetSslClientTrustStoreDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSslClientTrustStoreDataSource_Read_NotFound(t *testing.T) {
	r := &GetSslClientTrustStoreDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSslClientTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSslClientTrustStoreDataSource_Read_APIError exercises GetSslClientTrustStoreDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSslClientTrustStoreDataSource_Read_APIError(t *testing.T) {
	r := &GetSslClientTrustStoreDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSslClientTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ssl_client_trust_store")
}

// TestGetSslClientTrustStoreDataSource_Read_APIErrorReadBody exercises GetSslClientTrustStoreDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSslClientTrustStoreDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSslClientTrustStoreDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSslClientTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSslClientTrustStoreDataSource_Read_InvalidJSON exercises GetSslClientTrustStoreDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSslClientTrustStoreDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSslClientTrustStoreDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSslClientTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
