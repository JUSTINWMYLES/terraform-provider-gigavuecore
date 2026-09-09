package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetToolsDataSource_Read_Happy exercises GetToolsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetToolsDataSource_Read_Happy(t *testing.T) {
	r := &GetToolsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetToolsDataSource_Read_NilClient exercises GetToolsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetToolsDataSource_Read_NilClient(t *testing.T) {
	r := &GetToolsDataSource{}
	m := GetToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetToolsDataSource_Read_BuildError exercises GetToolsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetToolsDataSource_Read_BuildError(t *testing.T) {
	r := &GetToolsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetToolsDataSource_Read_SendError exercises GetToolsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetToolsDataSource_Read_SendError(t *testing.T) {
	r := &GetToolsDataSource{client: newTransportErrorClient(t)}
	m := GetToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetToolsDataSource_Read_NotFound exercises GetToolsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetToolsDataSource_Read_NotFound(t *testing.T) {
	r := &GetToolsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetToolsDataSource_Read_APIError exercises GetToolsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetToolsDataSource_Read_APIError(t *testing.T) {
	r := &GetToolsDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_tools")
}

// TestGetToolsDataSource_Read_APIErrorReadBody exercises GetToolsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetToolsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetToolsDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetToolsDataSource_Read_InvalidJSON exercises GetToolsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetToolsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetToolsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetToolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
