package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSshCiphersDataSource_Read_Happy exercises GetSshCiphersDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSshCiphersDataSource_Read_Happy(t *testing.T) {
	r := &GetSshCiphersDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSshCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSshCiphersDataSource_Read_NilClient exercises GetSshCiphersDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSshCiphersDataSource_Read_NilClient(t *testing.T) {
	r := &GetSshCiphersDataSource{}
	m := GetSshCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSshCiphersDataSource_Read_BuildError exercises GetSshCiphersDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSshCiphersDataSource_Read_BuildError(t *testing.T) {
	r := &GetSshCiphersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSshCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSshCiphersDataSource_Read_SendError exercises GetSshCiphersDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSshCiphersDataSource_Read_SendError(t *testing.T) {
	r := &GetSshCiphersDataSource{client: newTransportErrorClient(t)}
	m := GetSshCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSshCiphersDataSource_Read_NotFound exercises GetSshCiphersDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSshCiphersDataSource_Read_NotFound(t *testing.T) {
	r := &GetSshCiphersDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSshCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSshCiphersDataSource_Read_APIError exercises GetSshCiphersDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSshCiphersDataSource_Read_APIError(t *testing.T) {
	r := &GetSshCiphersDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSshCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ssh_ciphers")
}

// TestGetSshCiphersDataSource_Read_APIErrorReadBody exercises GetSshCiphersDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSshCiphersDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSshCiphersDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSshCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSshCiphersDataSource_Read_InvalidJSON exercises GetSshCiphersDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSshCiphersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSshCiphersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSshCiphersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
