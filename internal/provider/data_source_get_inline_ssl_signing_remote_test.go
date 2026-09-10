package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetInlineSslSigningDataSource_Read_Happy exercises GetInlineSslSigningDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetInlineSslSigningDataSource_Read_Happy(t *testing.T) {
	r := &GetInlineSslSigningDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetInlineSslSigningDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetInlineSslSigningDataSource_Read_NilClient exercises GetInlineSslSigningDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetInlineSslSigningDataSource_Read_NilClient(t *testing.T) {
	r := &GetInlineSslSigningDataSource{}
	m := GetInlineSslSigningDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetInlineSslSigningDataSource_Read_BuildError exercises GetInlineSslSigningDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetInlineSslSigningDataSource_Read_BuildError(t *testing.T) {
	r := &GetInlineSslSigningDataSource{client: newMalformedBaseURLClient(t)}
	m := GetInlineSslSigningDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetInlineSslSigningDataSource_Read_SendError exercises GetInlineSslSigningDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetInlineSslSigningDataSource_Read_SendError(t *testing.T) {
	r := &GetInlineSslSigningDataSource{client: newTransportErrorClient(t)}
	m := GetInlineSslSigningDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetInlineSslSigningDataSource_Read_NotFound exercises GetInlineSslSigningDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetInlineSslSigningDataSource_Read_NotFound(t *testing.T) {
	r := &GetInlineSslSigningDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetInlineSslSigningDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetInlineSslSigningDataSource_Read_APIError exercises GetInlineSslSigningDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetInlineSslSigningDataSource_Read_APIError(t *testing.T) {
	r := &GetInlineSslSigningDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetInlineSslSigningDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_inline_ssl_signing")
}

// TestGetInlineSslSigningDataSource_Read_APIErrorReadBody exercises GetInlineSslSigningDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetInlineSslSigningDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetInlineSslSigningDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetInlineSslSigningDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetInlineSslSigningDataSource_Read_InvalidJSON exercises GetInlineSslSigningDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetInlineSslSigningDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetInlineSslSigningDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetInlineSslSigningDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
