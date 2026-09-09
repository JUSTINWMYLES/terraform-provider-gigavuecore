package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadInlineSslUrlStatusDataSource_Read_Happy exercises LoadInlineSslUrlStatusDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadInlineSslUrlStatusDataSource_Read_Happy(t *testing.T) {
	r := &LoadInlineSslUrlStatusDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadInlineSslUrlStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadInlineSslUrlStatusDataSource_Read_NilClient exercises LoadInlineSslUrlStatusDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadInlineSslUrlStatusDataSource_Read_NilClient(t *testing.T) {
	r := &LoadInlineSslUrlStatusDataSource{}
	m := LoadInlineSslUrlStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadInlineSslUrlStatusDataSource_Read_BuildError exercises LoadInlineSslUrlStatusDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadInlineSslUrlStatusDataSource_Read_BuildError(t *testing.T) {
	r := &LoadInlineSslUrlStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadInlineSslUrlStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadInlineSslUrlStatusDataSource_Read_SendError exercises LoadInlineSslUrlStatusDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadInlineSslUrlStatusDataSource_Read_SendError(t *testing.T) {
	r := &LoadInlineSslUrlStatusDataSource{client: newTransportErrorClient(t)}
	m := LoadInlineSslUrlStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadInlineSslUrlStatusDataSource_Read_NotFound exercises LoadInlineSslUrlStatusDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadInlineSslUrlStatusDataSource_Read_NotFound(t *testing.T) {
	r := &LoadInlineSslUrlStatusDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadInlineSslUrlStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadInlineSslUrlStatusDataSource_Read_APIError exercises LoadInlineSslUrlStatusDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadInlineSslUrlStatusDataSource_Read_APIError(t *testing.T) {
	r := &LoadInlineSslUrlStatusDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadInlineSslUrlStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_inline_ssl_url_status")
}

// TestLoadInlineSslUrlStatusDataSource_Read_APIErrorReadBody exercises LoadInlineSslUrlStatusDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadInlineSslUrlStatusDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadInlineSslUrlStatusDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadInlineSslUrlStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadInlineSslUrlStatusDataSource_Read_InvalidJSON exercises LoadInlineSslUrlStatusDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadInlineSslUrlStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadInlineSslUrlStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadInlineSslUrlStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
