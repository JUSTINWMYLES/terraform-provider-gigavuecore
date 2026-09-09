package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCommonCiphersDataSource_Read_Happy exercises GetCommonCiphersDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetCommonCiphersDataSource_Read_Happy(t *testing.T) {
	r := &GetCommonCiphersDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetCommonCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCommonCiphersDataSource_Read_NilClient exercises GetCommonCiphersDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCommonCiphersDataSource_Read_NilClient(t *testing.T) {
	r := &GetCommonCiphersDataSource{}
	m := GetCommonCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCommonCiphersDataSource_Read_BuildError exercises GetCommonCiphersDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetCommonCiphersDataSource_Read_BuildError(t *testing.T) {
	r := &GetCommonCiphersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCommonCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetCommonCiphersDataSource_Read_SendError exercises GetCommonCiphersDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetCommonCiphersDataSource_Read_SendError(t *testing.T) {
	r := &GetCommonCiphersDataSource{client: newTransportErrorClient(t)}
	m := GetCommonCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetCommonCiphersDataSource_Read_NotFound exercises GetCommonCiphersDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetCommonCiphersDataSource_Read_NotFound(t *testing.T) {
	r := &GetCommonCiphersDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetCommonCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetCommonCiphersDataSource_Read_APIError exercises GetCommonCiphersDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetCommonCiphersDataSource_Read_APIError(t *testing.T) {
	r := &GetCommonCiphersDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetCommonCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_common_ciphers")
}

// TestGetCommonCiphersDataSource_Read_APIErrorReadBody exercises GetCommonCiphersDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetCommonCiphersDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetCommonCiphersDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetCommonCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetCommonCiphersDataSource_Read_InvalidJSON exercises GetCommonCiphersDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetCommonCiphersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCommonCiphersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCommonCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
