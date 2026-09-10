package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetInlineSslConfigDataSource_Read_Happy exercises GetInlineSslConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetInlineSslConfigDataSource_Read_Happy(t *testing.T) {
	r := &GetInlineSslConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetInlineSslConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetInlineSslConfigDataSource_Read_NilClient exercises GetInlineSslConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetInlineSslConfigDataSource_Read_NilClient(t *testing.T) {
	r := &GetInlineSslConfigDataSource{}
	m := GetInlineSslConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetInlineSslConfigDataSource_Read_BuildError exercises GetInlineSslConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetInlineSslConfigDataSource_Read_BuildError(t *testing.T) {
	r := &GetInlineSslConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := GetInlineSslConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetInlineSslConfigDataSource_Read_SendError exercises GetInlineSslConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetInlineSslConfigDataSource_Read_SendError(t *testing.T) {
	r := &GetInlineSslConfigDataSource{client: newTransportErrorClient(t)}
	m := GetInlineSslConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetInlineSslConfigDataSource_Read_NotFound exercises GetInlineSslConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetInlineSslConfigDataSource_Read_NotFound(t *testing.T) {
	r := &GetInlineSslConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetInlineSslConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetInlineSslConfigDataSource_Read_APIError exercises GetInlineSslConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetInlineSslConfigDataSource_Read_APIError(t *testing.T) {
	r := &GetInlineSslConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetInlineSslConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_inline_ssl_config")
}

// TestGetInlineSslConfigDataSource_Read_APIErrorReadBody exercises GetInlineSslConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetInlineSslConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetInlineSslConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetInlineSslConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetInlineSslConfigDataSource_Read_InvalidJSON exercises GetInlineSslConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetInlineSslConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetInlineSslConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetInlineSslConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
