package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslDecryptionSummaryDataSource_Read_Happy exercises GetSslDecryptionSummaryDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSslDecryptionSummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetSslDecryptionSummaryDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSslDecryptionSummaryDataSource_Read_NilClient exercises GetSslDecryptionSummaryDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSslDecryptionSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetSslDecryptionSummaryDataSource{}
	m := GetSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSslDecryptionSummaryDataSource_Read_BuildError exercises GetSslDecryptionSummaryDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSslDecryptionSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetSslDecryptionSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSslDecryptionSummaryDataSource_Read_SendError exercises GetSslDecryptionSummaryDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSslDecryptionSummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetSslDecryptionSummaryDataSource{client: newTransportErrorClient(t)}
	m := GetSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSslDecryptionSummaryDataSource_Read_NotFound exercises GetSslDecryptionSummaryDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSslDecryptionSummaryDataSource_Read_NotFound(t *testing.T) {
	r := &GetSslDecryptionSummaryDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSslDecryptionSummaryDataSource_Read_APIError exercises GetSslDecryptionSummaryDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSslDecryptionSummaryDataSource_Read_APIError(t *testing.T) {
	r := &GetSslDecryptionSummaryDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ssl_decryption_summary")
}

// TestGetSslDecryptionSummaryDataSource_Read_APIErrorReadBody exercises GetSslDecryptionSummaryDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSslDecryptionSummaryDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSslDecryptionSummaryDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSslDecryptionSummaryDataSource_Read_InvalidJSON exercises GetSslDecryptionSummaryDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSslDecryptionSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSslDecryptionSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
