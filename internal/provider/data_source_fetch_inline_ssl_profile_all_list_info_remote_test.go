package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestFetchInlineSslProfileAllListInfoDataSource_Read_Happy exercises FetchInlineSslProfileAllListInfoDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestFetchInlineSslProfileAllListInfoDataSource_Read_Happy(t *testing.T) {
	r := &FetchInlineSslProfileAllListInfoDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := FetchInlineSslProfileAllListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFetchInlineSslProfileAllListInfoDataSource_Read_NilClient exercises FetchInlineSslProfileAllListInfoDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFetchInlineSslProfileAllListInfoDataSource_Read_NilClient(t *testing.T) {
	r := &FetchInlineSslProfileAllListInfoDataSource{}
	m := FetchInlineSslProfileAllListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFetchInlineSslProfileAllListInfoDataSource_Read_BuildError exercises FetchInlineSslProfileAllListInfoDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFetchInlineSslProfileAllListInfoDataSource_Read_BuildError(t *testing.T) {
	r := &FetchInlineSslProfileAllListInfoDataSource{client: newMalformedBaseURLClient(t)}
	m := FetchInlineSslProfileAllListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFetchInlineSslProfileAllListInfoDataSource_Read_SendError exercises FetchInlineSslProfileAllListInfoDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestFetchInlineSslProfileAllListInfoDataSource_Read_SendError(t *testing.T) {
	r := &FetchInlineSslProfileAllListInfoDataSource{client: newTransportErrorClient(t)}
	m := FetchInlineSslProfileAllListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFetchInlineSslProfileAllListInfoDataSource_Read_NotFound exercises FetchInlineSslProfileAllListInfoDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestFetchInlineSslProfileAllListInfoDataSource_Read_NotFound(t *testing.T) {
	r := &FetchInlineSslProfileAllListInfoDataSource{client: newMockClientStatus(t, 404, "")}
	m := FetchInlineSslProfileAllListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestFetchInlineSslProfileAllListInfoDataSource_Read_APIError exercises FetchInlineSslProfileAllListInfoDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFetchInlineSslProfileAllListInfoDataSource_Read_APIError(t *testing.T) {
	r := &FetchInlineSslProfileAllListInfoDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FetchInlineSslProfileAllListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info")
}

// TestFetchInlineSslProfileAllListInfoDataSource_Read_APIErrorReadBody exercises FetchInlineSslProfileAllListInfoDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFetchInlineSslProfileAllListInfoDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &FetchInlineSslProfileAllListInfoDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := FetchInlineSslProfileAllListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFetchInlineSslProfileAllListInfoDataSource_Read_InvalidJSON exercises FetchInlineSslProfileAllListInfoDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFetchInlineSslProfileAllListInfoDataSource_Read_InvalidJSON(t *testing.T) {
	r := &FetchInlineSslProfileAllListInfoDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := FetchInlineSslProfileAllListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
