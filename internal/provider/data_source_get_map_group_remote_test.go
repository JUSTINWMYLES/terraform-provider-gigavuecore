package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMapGroupDataSource_Read_Happy exercises GetMapGroupDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetMapGroupDataSource_Read_Happy(t *testing.T) {
	r := &GetMapGroupDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetMapGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMapGroupDataSource_Read_NilClient exercises GetMapGroupDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMapGroupDataSource_Read_NilClient(t *testing.T) {
	r := &GetMapGroupDataSource{}
	m := GetMapGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMapGroupDataSource_Read_BuildError exercises GetMapGroupDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetMapGroupDataSource_Read_BuildError(t *testing.T) {
	r := &GetMapGroupDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMapGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetMapGroupDataSource_Read_SendError exercises GetMapGroupDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetMapGroupDataSource_Read_SendError(t *testing.T) {
	r := &GetMapGroupDataSource{client: newTransportErrorClient(t)}
	m := GetMapGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetMapGroupDataSource_Read_NotFound exercises GetMapGroupDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetMapGroupDataSource_Read_NotFound(t *testing.T) {
	r := &GetMapGroupDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetMapGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetMapGroupDataSource_Read_APIError exercises GetMapGroupDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetMapGroupDataSource_Read_APIError(t *testing.T) {
	r := &GetMapGroupDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetMapGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_map_group")
}

// TestGetMapGroupDataSource_Read_APIErrorReadBody exercises GetMapGroupDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetMapGroupDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetMapGroupDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetMapGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetMapGroupDataSource_Read_InvalidJSON exercises GetMapGroupDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetMapGroupDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMapGroupDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMapGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
