package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestFetchInlineSslProfileListInfoDataSource_Read_Happy exercises FetchInlineSslProfileListInfoDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestFetchInlineSslProfileListInfoDataSource_Read_Happy(t *testing.T) {
	r := &FetchInlineSslProfileListInfoDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := FetchInlineSslProfileListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFetchInlineSslProfileListInfoDataSource_Read_NilClient exercises FetchInlineSslProfileListInfoDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFetchInlineSslProfileListInfoDataSource_Read_NilClient(t *testing.T) {
	r := &FetchInlineSslProfileListInfoDataSource{}
	m := FetchInlineSslProfileListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFetchInlineSslProfileListInfoDataSource_Read_BuildError exercises FetchInlineSslProfileListInfoDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFetchInlineSslProfileListInfoDataSource_Read_BuildError(t *testing.T) {
	r := &FetchInlineSslProfileListInfoDataSource{client: newMalformedBaseURLClient(t)}
	m := FetchInlineSslProfileListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFetchInlineSslProfileListInfoDataSource_Read_SendError exercises FetchInlineSslProfileListInfoDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestFetchInlineSslProfileListInfoDataSource_Read_SendError(t *testing.T) {
	r := &FetchInlineSslProfileListInfoDataSource{client: newTransportErrorClient(t)}
	m := FetchInlineSslProfileListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFetchInlineSslProfileListInfoDataSource_Read_NotFound exercises FetchInlineSslProfileListInfoDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestFetchInlineSslProfileListInfoDataSource_Read_NotFound(t *testing.T) {
	r := &FetchInlineSslProfileListInfoDataSource{client: newMockClientStatus(t, 404, "")}
	m := FetchInlineSslProfileListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestFetchInlineSslProfileListInfoDataSource_Read_APIError exercises FetchInlineSslProfileListInfoDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFetchInlineSslProfileListInfoDataSource_Read_APIError(t *testing.T) {
	r := &FetchInlineSslProfileListInfoDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FetchInlineSslProfileListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_fetch_inline_ssl_profile_list_info")
}

// TestFetchInlineSslProfileListInfoDataSource_Read_APIErrorReadBody exercises FetchInlineSslProfileListInfoDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFetchInlineSslProfileListInfoDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &FetchInlineSslProfileListInfoDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := FetchInlineSslProfileListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFetchInlineSslProfileListInfoDataSource_Read_InvalidJSON exercises FetchInlineSslProfileListInfoDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFetchInlineSslProfileListInfoDataSource_Read_InvalidJSON(t *testing.T) {
	r := &FetchInlineSslProfileListInfoDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := FetchInlineSslProfileListInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
