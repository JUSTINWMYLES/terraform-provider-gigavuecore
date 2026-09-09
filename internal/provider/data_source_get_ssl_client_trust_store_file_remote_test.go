package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslClientTrustStoreFileDataSource_Read_Happy exercises GetSslClientTrustStoreFileDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSslClientTrustStoreFileDataSource_Read_Happy(t *testing.T) {
	r := &GetSslClientTrustStoreFileDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSslClientTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSslClientTrustStoreFileDataSource_Read_NilClient exercises GetSslClientTrustStoreFileDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSslClientTrustStoreFileDataSource_Read_NilClient(t *testing.T) {
	r := &GetSslClientTrustStoreFileDataSource{}
	m := GetSslClientTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSslClientTrustStoreFileDataSource_Read_BuildError exercises GetSslClientTrustStoreFileDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSslClientTrustStoreFileDataSource_Read_BuildError(t *testing.T) {
	r := &GetSslClientTrustStoreFileDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSslClientTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSslClientTrustStoreFileDataSource_Read_SendError exercises GetSslClientTrustStoreFileDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSslClientTrustStoreFileDataSource_Read_SendError(t *testing.T) {
	r := &GetSslClientTrustStoreFileDataSource{client: newTransportErrorClient(t)}
	m := GetSslClientTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSslClientTrustStoreFileDataSource_Read_NotFound exercises GetSslClientTrustStoreFileDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSslClientTrustStoreFileDataSource_Read_NotFound(t *testing.T) {
	r := &GetSslClientTrustStoreFileDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSslClientTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSslClientTrustStoreFileDataSource_Read_APIError exercises GetSslClientTrustStoreFileDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSslClientTrustStoreFileDataSource_Read_APIError(t *testing.T) {
	r := &GetSslClientTrustStoreFileDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSslClientTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ssl_client_trust_store_file")
}

// TestGetSslClientTrustStoreFileDataSource_Read_APIErrorReadBody exercises GetSslClientTrustStoreFileDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSslClientTrustStoreFileDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSslClientTrustStoreFileDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSslClientTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSslClientTrustStoreFileDataSource_Read_InvalidJSON exercises GetSslClientTrustStoreFileDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSslClientTrustStoreFileDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSslClientTrustStoreFileDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSslClientTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
