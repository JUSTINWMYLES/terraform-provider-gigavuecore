package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUserSelectedCiphersDataSource_Read_Happy exercises GetUserSelectedCiphersDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetUserSelectedCiphersDataSource_Read_Happy(t *testing.T) {
	r := &GetUserSelectedCiphersDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetUserSelectedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetUserSelectedCiphersDataSource_Read_NilClient exercises GetUserSelectedCiphersDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetUserSelectedCiphersDataSource_Read_NilClient(t *testing.T) {
	r := &GetUserSelectedCiphersDataSource{}
	m := GetUserSelectedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetUserSelectedCiphersDataSource_Read_BuildError exercises GetUserSelectedCiphersDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetUserSelectedCiphersDataSource_Read_BuildError(t *testing.T) {
	r := &GetUserSelectedCiphersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetUserSelectedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetUserSelectedCiphersDataSource_Read_SendError exercises GetUserSelectedCiphersDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetUserSelectedCiphersDataSource_Read_SendError(t *testing.T) {
	r := &GetUserSelectedCiphersDataSource{client: newTransportErrorClient(t)}
	m := GetUserSelectedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetUserSelectedCiphersDataSource_Read_NotFound exercises GetUserSelectedCiphersDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetUserSelectedCiphersDataSource_Read_NotFound(t *testing.T) {
	r := &GetUserSelectedCiphersDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetUserSelectedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetUserSelectedCiphersDataSource_Read_APIError exercises GetUserSelectedCiphersDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetUserSelectedCiphersDataSource_Read_APIError(t *testing.T) {
	r := &GetUserSelectedCiphersDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetUserSelectedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_user_selected_ciphers")
}

// TestGetUserSelectedCiphersDataSource_Read_APIErrorReadBody exercises GetUserSelectedCiphersDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetUserSelectedCiphersDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetUserSelectedCiphersDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetUserSelectedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetUserSelectedCiphersDataSource_Read_InvalidJSON exercises GetUserSelectedCiphersDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetUserSelectedCiphersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetUserSelectedCiphersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetUserSelectedCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
