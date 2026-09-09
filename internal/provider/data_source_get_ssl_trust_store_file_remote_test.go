package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslTrustStoreFileDataSource_Read_Happy exercises GetSslTrustStoreFileDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSslTrustStoreFileDataSource_Read_Happy(t *testing.T) {
	r := &GetSslTrustStoreFileDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSslTrustStoreFileDataSource_Read_NilClient exercises GetSslTrustStoreFileDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSslTrustStoreFileDataSource_Read_NilClient(t *testing.T) {
	r := &GetSslTrustStoreFileDataSource{}
	m := GetSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSslTrustStoreFileDataSource_Read_BuildError exercises GetSslTrustStoreFileDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSslTrustStoreFileDataSource_Read_BuildError(t *testing.T) {
	r := &GetSslTrustStoreFileDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSslTrustStoreFileDataSource_Read_SendError exercises GetSslTrustStoreFileDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSslTrustStoreFileDataSource_Read_SendError(t *testing.T) {
	r := &GetSslTrustStoreFileDataSource{client: newTransportErrorClient(t)}
	m := GetSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSslTrustStoreFileDataSource_Read_NotFound exercises GetSslTrustStoreFileDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSslTrustStoreFileDataSource_Read_NotFound(t *testing.T) {
	r := &GetSslTrustStoreFileDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSslTrustStoreFileDataSource_Read_APIError exercises GetSslTrustStoreFileDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSslTrustStoreFileDataSource_Read_APIError(t *testing.T) {
	r := &GetSslTrustStoreFileDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ssl_trust_store_file")
}

// TestGetSslTrustStoreFileDataSource_Read_APIErrorReadBody exercises GetSslTrustStoreFileDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSslTrustStoreFileDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSslTrustStoreFileDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSslTrustStoreFileDataSource_Read_InvalidJSON exercises GetSslTrustStoreFileDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSslTrustStoreFileDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSslTrustStoreFileDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
