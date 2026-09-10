package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetInlineSslTrustStoreFileDataSource_Read_Happy exercises GetInlineSslTrustStoreFileDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetInlineSslTrustStoreFileDataSource_Read_Happy(t *testing.T) {
	r := &GetInlineSslTrustStoreFileDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetInlineSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetInlineSslTrustStoreFileDataSource_Read_NilClient exercises GetInlineSslTrustStoreFileDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetInlineSslTrustStoreFileDataSource_Read_NilClient(t *testing.T) {
	r := &GetInlineSslTrustStoreFileDataSource{}
	m := GetInlineSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetInlineSslTrustStoreFileDataSource_Read_BuildError exercises GetInlineSslTrustStoreFileDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetInlineSslTrustStoreFileDataSource_Read_BuildError(t *testing.T) {
	r := &GetInlineSslTrustStoreFileDataSource{client: newMalformedBaseURLClient(t)}
	m := GetInlineSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetInlineSslTrustStoreFileDataSource_Read_SendError exercises GetInlineSslTrustStoreFileDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetInlineSslTrustStoreFileDataSource_Read_SendError(t *testing.T) {
	r := &GetInlineSslTrustStoreFileDataSource{client: newTransportErrorClient(t)}
	m := GetInlineSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetInlineSslTrustStoreFileDataSource_Read_NotFound exercises GetInlineSslTrustStoreFileDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetInlineSslTrustStoreFileDataSource_Read_NotFound(t *testing.T) {
	r := &GetInlineSslTrustStoreFileDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetInlineSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetInlineSslTrustStoreFileDataSource_Read_APIError exercises GetInlineSslTrustStoreFileDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetInlineSslTrustStoreFileDataSource_Read_APIError(t *testing.T) {
	r := &GetInlineSslTrustStoreFileDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetInlineSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_inline_ssl_trust_store_file")
}

// TestGetInlineSslTrustStoreFileDataSource_Read_APIErrorReadBody exercises GetInlineSslTrustStoreFileDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetInlineSslTrustStoreFileDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetInlineSslTrustStoreFileDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetInlineSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetInlineSslTrustStoreFileDataSource_Read_InvalidJSON exercises GetInlineSslTrustStoreFileDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetInlineSslTrustStoreFileDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetInlineSslTrustStoreFileDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetInlineSslTrustStoreFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
