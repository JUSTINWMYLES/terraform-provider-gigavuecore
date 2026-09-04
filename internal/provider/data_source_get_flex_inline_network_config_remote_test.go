package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlexInlineNetworkConfigDataSource_Read_Happy exercises GetFlexInlineNetworkConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFlexInlineNetworkConfigDataSource_Read_Happy(t *testing.T) {
	r := &GetFlexInlineNetworkConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFlexInlineNetworkConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlexInlineNetworkConfigDataSource_Read_NilClient exercises GetFlexInlineNetworkConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlexInlineNetworkConfigDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlexInlineNetworkConfigDataSource{}
	m := GetFlexInlineNetworkConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlexInlineNetworkConfigDataSource_Read_BuildError exercises GetFlexInlineNetworkConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFlexInlineNetworkConfigDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlexInlineNetworkConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlexInlineNetworkConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFlexInlineNetworkConfigDataSource_Read_SendError exercises GetFlexInlineNetworkConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFlexInlineNetworkConfigDataSource_Read_SendError(t *testing.T) {
	r := &GetFlexInlineNetworkConfigDataSource{client: newTransportErrorClient(t)}
	m := GetFlexInlineNetworkConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFlexInlineNetworkConfigDataSource_Read_NotFound exercises GetFlexInlineNetworkConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFlexInlineNetworkConfigDataSource_Read_NotFound(t *testing.T) {
	r := &GetFlexInlineNetworkConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFlexInlineNetworkConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFlexInlineNetworkConfigDataSource_Read_APIError exercises GetFlexInlineNetworkConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFlexInlineNetworkConfigDataSource_Read_APIError(t *testing.T) {
	r := &GetFlexInlineNetworkConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFlexInlineNetworkConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_flex_inline_network_config")
}

// TestGetFlexInlineNetworkConfigDataSource_Read_APIErrorReadBody exercises GetFlexInlineNetworkConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFlexInlineNetworkConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFlexInlineNetworkConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFlexInlineNetworkConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFlexInlineNetworkConfigDataSource_Read_InvalidJSON exercises GetFlexInlineNetworkConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFlexInlineNetworkConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlexInlineNetworkConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlexInlineNetworkConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
