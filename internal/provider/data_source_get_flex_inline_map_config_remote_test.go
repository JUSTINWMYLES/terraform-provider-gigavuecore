package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlexInlineMapConfigDataSource_Read_Happy exercises GetFlexInlineMapConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFlexInlineMapConfigDataSource_Read_Happy(t *testing.T) {
	r := &GetFlexInlineMapConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFlexInlineMapConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlexInlineMapConfigDataSource_Read_NilClient exercises GetFlexInlineMapConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlexInlineMapConfigDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlexInlineMapConfigDataSource{}
	m := GetFlexInlineMapConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlexInlineMapConfigDataSource_Read_BuildError exercises GetFlexInlineMapConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFlexInlineMapConfigDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlexInlineMapConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlexInlineMapConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFlexInlineMapConfigDataSource_Read_SendError exercises GetFlexInlineMapConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFlexInlineMapConfigDataSource_Read_SendError(t *testing.T) {
	r := &GetFlexInlineMapConfigDataSource{client: newTransportErrorClient(t)}
	m := GetFlexInlineMapConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFlexInlineMapConfigDataSource_Read_NotFound exercises GetFlexInlineMapConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFlexInlineMapConfigDataSource_Read_NotFound(t *testing.T) {
	r := &GetFlexInlineMapConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFlexInlineMapConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFlexInlineMapConfigDataSource_Read_APIError exercises GetFlexInlineMapConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFlexInlineMapConfigDataSource_Read_APIError(t *testing.T) {
	r := &GetFlexInlineMapConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFlexInlineMapConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_flex_inline_map_config")
}

// TestGetFlexInlineMapConfigDataSource_Read_APIErrorReadBody exercises GetFlexInlineMapConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFlexInlineMapConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFlexInlineMapConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFlexInlineMapConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFlexInlineMapConfigDataSource_Read_InvalidJSON exercises GetFlexInlineMapConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFlexInlineMapConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlexInlineMapConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlexInlineMapConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
