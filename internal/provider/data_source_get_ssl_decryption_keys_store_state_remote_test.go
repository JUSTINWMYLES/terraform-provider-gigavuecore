package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslDecryptionKeysStoreStateDataSource_Read_Happy exercises GetSslDecryptionKeysStoreStateDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSslDecryptionKeysStoreStateDataSource_Read_Happy(t *testing.T) {
	r := &GetSslDecryptionKeysStoreStateDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSslDecryptionKeysStoreStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSslDecryptionKeysStoreStateDataSource_Read_NilClient exercises GetSslDecryptionKeysStoreStateDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSslDecryptionKeysStoreStateDataSource_Read_NilClient(t *testing.T) {
	r := &GetSslDecryptionKeysStoreStateDataSource{}
	m := GetSslDecryptionKeysStoreStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSslDecryptionKeysStoreStateDataSource_Read_BuildError exercises GetSslDecryptionKeysStoreStateDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSslDecryptionKeysStoreStateDataSource_Read_BuildError(t *testing.T) {
	r := &GetSslDecryptionKeysStoreStateDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSslDecryptionKeysStoreStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSslDecryptionKeysStoreStateDataSource_Read_SendError exercises GetSslDecryptionKeysStoreStateDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSslDecryptionKeysStoreStateDataSource_Read_SendError(t *testing.T) {
	r := &GetSslDecryptionKeysStoreStateDataSource{client: newTransportErrorClient(t)}
	m := GetSslDecryptionKeysStoreStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSslDecryptionKeysStoreStateDataSource_Read_NotFound exercises GetSslDecryptionKeysStoreStateDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSslDecryptionKeysStoreStateDataSource_Read_NotFound(t *testing.T) {
	r := &GetSslDecryptionKeysStoreStateDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSslDecryptionKeysStoreStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSslDecryptionKeysStoreStateDataSource_Read_APIError exercises GetSslDecryptionKeysStoreStateDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSslDecryptionKeysStoreStateDataSource_Read_APIError(t *testing.T) {
	r := &GetSslDecryptionKeysStoreStateDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSslDecryptionKeysStoreStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ssl_decryption_keys_store_state")
}

// TestGetSslDecryptionKeysStoreStateDataSource_Read_APIErrorReadBody exercises GetSslDecryptionKeysStoreStateDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSslDecryptionKeysStoreStateDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSslDecryptionKeysStoreStateDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSslDecryptionKeysStoreStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSslDecryptionKeysStoreStateDataSource_Read_InvalidJSON exercises GetSslDecryptionKeysStoreStateDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSslDecryptionKeysStoreStateDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSslDecryptionKeysStoreStateDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSslDecryptionKeysStoreStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
