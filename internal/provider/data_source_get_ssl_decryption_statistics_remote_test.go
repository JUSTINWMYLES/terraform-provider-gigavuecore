package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslDecryptionStatisticsDataSource_Read_Happy exercises GetSslDecryptionStatisticsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSslDecryptionStatisticsDataSource_Read_Happy(t *testing.T) {
	r := &GetSslDecryptionStatisticsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSslDecryptionStatisticsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSslDecryptionStatisticsDataSource_Read_NilClient exercises GetSslDecryptionStatisticsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSslDecryptionStatisticsDataSource_Read_NilClient(t *testing.T) {
	r := &GetSslDecryptionStatisticsDataSource{}
	m := GetSslDecryptionStatisticsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSslDecryptionStatisticsDataSource_Read_BuildError exercises GetSslDecryptionStatisticsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSslDecryptionStatisticsDataSource_Read_BuildError(t *testing.T) {
	r := &GetSslDecryptionStatisticsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSslDecryptionStatisticsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSslDecryptionStatisticsDataSource_Read_SendError exercises GetSslDecryptionStatisticsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSslDecryptionStatisticsDataSource_Read_SendError(t *testing.T) {
	r := &GetSslDecryptionStatisticsDataSource{client: newTransportErrorClient(t)}
	m := GetSslDecryptionStatisticsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSslDecryptionStatisticsDataSource_Read_NotFound exercises GetSslDecryptionStatisticsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSslDecryptionStatisticsDataSource_Read_NotFound(t *testing.T) {
	r := &GetSslDecryptionStatisticsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSslDecryptionStatisticsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSslDecryptionStatisticsDataSource_Read_APIError exercises GetSslDecryptionStatisticsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSslDecryptionStatisticsDataSource_Read_APIError(t *testing.T) {
	r := &GetSslDecryptionStatisticsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSslDecryptionStatisticsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ssl_decryption_statistics")
}

// TestGetSslDecryptionStatisticsDataSource_Read_APIErrorReadBody exercises GetSslDecryptionStatisticsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSslDecryptionStatisticsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSslDecryptionStatisticsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSslDecryptionStatisticsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSslDecryptionStatisticsDataSource_Read_InvalidJSON exercises GetSslDecryptionStatisticsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSslDecryptionStatisticsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSslDecryptionStatisticsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSslDecryptionStatisticsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
