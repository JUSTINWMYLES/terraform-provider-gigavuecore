package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetInternalFabricMapByAliasDataSource_Read_Happy exercises GetInternalFabricMapByAliasDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetInternalFabricMapByAliasDataSource_Read_Happy(t *testing.T) {
	r := &GetInternalFabricMapByAliasDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetInternalFabricMapByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetInternalFabricMapByAliasDataSource_Read_NilClient exercises GetInternalFabricMapByAliasDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetInternalFabricMapByAliasDataSource_Read_NilClient(t *testing.T) {
	r := &GetInternalFabricMapByAliasDataSource{}
	m := GetInternalFabricMapByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetInternalFabricMapByAliasDataSource_Read_BuildError exercises GetInternalFabricMapByAliasDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetInternalFabricMapByAliasDataSource_Read_BuildError(t *testing.T) {
	r := &GetInternalFabricMapByAliasDataSource{client: newMalformedBaseURLClient(t)}
	m := GetInternalFabricMapByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetInternalFabricMapByAliasDataSource_Read_SendError exercises GetInternalFabricMapByAliasDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetInternalFabricMapByAliasDataSource_Read_SendError(t *testing.T) {
	r := &GetInternalFabricMapByAliasDataSource{client: newTransportErrorClient(t)}
	m := GetInternalFabricMapByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetInternalFabricMapByAliasDataSource_Read_NotFound exercises GetInternalFabricMapByAliasDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetInternalFabricMapByAliasDataSource_Read_NotFound(t *testing.T) {
	r := &GetInternalFabricMapByAliasDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetInternalFabricMapByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetInternalFabricMapByAliasDataSource_Read_APIError exercises GetInternalFabricMapByAliasDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetInternalFabricMapByAliasDataSource_Read_APIError(t *testing.T) {
	r := &GetInternalFabricMapByAliasDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetInternalFabricMapByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_internal_fabric_map_by_alias")
}

// TestGetInternalFabricMapByAliasDataSource_Read_APIErrorReadBody exercises GetInternalFabricMapByAliasDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetInternalFabricMapByAliasDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetInternalFabricMapByAliasDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetInternalFabricMapByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetInternalFabricMapByAliasDataSource_Read_InvalidJSON exercises GetInternalFabricMapByAliasDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetInternalFabricMapByAliasDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetInternalFabricMapByAliasDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetInternalFabricMapByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
