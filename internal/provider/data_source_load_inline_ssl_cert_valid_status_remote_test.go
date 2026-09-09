package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadInlineSslCertValidStatusDataSource_Read_Happy exercises LoadInlineSslCertValidStatusDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadInlineSslCertValidStatusDataSource_Read_Happy(t *testing.T) {
	r := &LoadInlineSslCertValidStatusDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadInlineSslCertValidStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadInlineSslCertValidStatusDataSource_Read_NilClient exercises LoadInlineSslCertValidStatusDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadInlineSslCertValidStatusDataSource_Read_NilClient(t *testing.T) {
	r := &LoadInlineSslCertValidStatusDataSource{}
	m := LoadInlineSslCertValidStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadInlineSslCertValidStatusDataSource_Read_BuildError exercises LoadInlineSslCertValidStatusDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadInlineSslCertValidStatusDataSource_Read_BuildError(t *testing.T) {
	r := &LoadInlineSslCertValidStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadInlineSslCertValidStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadInlineSslCertValidStatusDataSource_Read_SendError exercises LoadInlineSslCertValidStatusDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadInlineSslCertValidStatusDataSource_Read_SendError(t *testing.T) {
	r := &LoadInlineSslCertValidStatusDataSource{client: newTransportErrorClient(t)}
	m := LoadInlineSslCertValidStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadInlineSslCertValidStatusDataSource_Read_NotFound exercises LoadInlineSslCertValidStatusDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadInlineSslCertValidStatusDataSource_Read_NotFound(t *testing.T) {
	r := &LoadInlineSslCertValidStatusDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadInlineSslCertValidStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadInlineSslCertValidStatusDataSource_Read_APIError exercises LoadInlineSslCertValidStatusDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadInlineSslCertValidStatusDataSource_Read_APIError(t *testing.T) {
	r := &LoadInlineSslCertValidStatusDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadInlineSslCertValidStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_inline_ssl_cert_valid_status")
}

// TestLoadInlineSslCertValidStatusDataSource_Read_APIErrorReadBody exercises LoadInlineSslCertValidStatusDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadInlineSslCertValidStatusDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadInlineSslCertValidStatusDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadInlineSslCertValidStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadInlineSslCertValidStatusDataSource_Read_InvalidJSON exercises LoadInlineSslCertValidStatusDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadInlineSslCertValidStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadInlineSslCertValidStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadInlineSslCertValidStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
