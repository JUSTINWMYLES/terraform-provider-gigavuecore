package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetKeystoreKeyCertDataSource_Read_Happy exercises GetKeystoreKeyCertDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetKeystoreKeyCertDataSource_Read_Happy(t *testing.T) {
	r := &GetKeystoreKeyCertDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetKeystoreKeyCertDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetKeystoreKeyCertDataSource_Read_NilClient exercises GetKeystoreKeyCertDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetKeystoreKeyCertDataSource_Read_NilClient(t *testing.T) {
	r := &GetKeystoreKeyCertDataSource{}
	m := GetKeystoreKeyCertDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetKeystoreKeyCertDataSource_Read_BuildError exercises GetKeystoreKeyCertDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetKeystoreKeyCertDataSource_Read_BuildError(t *testing.T) {
	r := &GetKeystoreKeyCertDataSource{client: newMalformedBaseURLClient(t)}
	m := GetKeystoreKeyCertDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetKeystoreKeyCertDataSource_Read_SendError exercises GetKeystoreKeyCertDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetKeystoreKeyCertDataSource_Read_SendError(t *testing.T) {
	r := &GetKeystoreKeyCertDataSource{client: newTransportErrorClient(t)}
	m := GetKeystoreKeyCertDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetKeystoreKeyCertDataSource_Read_NotFound exercises GetKeystoreKeyCertDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetKeystoreKeyCertDataSource_Read_NotFound(t *testing.T) {
	r := &GetKeystoreKeyCertDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetKeystoreKeyCertDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetKeystoreKeyCertDataSource_Read_APIError exercises GetKeystoreKeyCertDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetKeystoreKeyCertDataSource_Read_APIError(t *testing.T) {
	r := &GetKeystoreKeyCertDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetKeystoreKeyCertDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_keystore_key_cert")
}

// TestGetKeystoreKeyCertDataSource_Read_APIErrorReadBody exercises GetKeystoreKeyCertDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetKeystoreKeyCertDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetKeystoreKeyCertDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetKeystoreKeyCertDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetKeystoreKeyCertDataSource_Read_InvalidJSON exercises GetKeystoreKeyCertDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetKeystoreKeyCertDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetKeystoreKeyCertDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetKeystoreKeyCertDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
