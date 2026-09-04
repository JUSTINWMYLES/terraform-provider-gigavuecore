package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlexInlineConfigDataSource_Read_Happy exercises GetFlexInlineConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFlexInlineConfigDataSource_Read_Happy(t *testing.T) {
	r := &GetFlexInlineConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFlexInlineConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlexInlineConfigDataSource_Read_NilClient exercises GetFlexInlineConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlexInlineConfigDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlexInlineConfigDataSource{}
	m := GetFlexInlineConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlexInlineConfigDataSource_Read_BuildError exercises GetFlexInlineConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFlexInlineConfigDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlexInlineConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlexInlineConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFlexInlineConfigDataSource_Read_SendError exercises GetFlexInlineConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFlexInlineConfigDataSource_Read_SendError(t *testing.T) {
	r := &GetFlexInlineConfigDataSource{client: newTransportErrorClient(t)}
	m := GetFlexInlineConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFlexInlineConfigDataSource_Read_NotFound exercises GetFlexInlineConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFlexInlineConfigDataSource_Read_NotFound(t *testing.T) {
	r := &GetFlexInlineConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFlexInlineConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFlexInlineConfigDataSource_Read_APIError exercises GetFlexInlineConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFlexInlineConfigDataSource_Read_APIError(t *testing.T) {
	r := &GetFlexInlineConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFlexInlineConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_flex_inline_config")
}

// TestGetFlexInlineConfigDataSource_Read_APIErrorReadBody exercises GetFlexInlineConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFlexInlineConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFlexInlineConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFlexInlineConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFlexInlineConfigDataSource_Read_InvalidJSON exercises GetFlexInlineConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFlexInlineConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlexInlineConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlexInlineConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
