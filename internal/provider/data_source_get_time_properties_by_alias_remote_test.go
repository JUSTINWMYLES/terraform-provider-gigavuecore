package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTimePropertiesByAliasDataSource_Read_Happy exercises GetTimePropertiesByAliasDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetTimePropertiesByAliasDataSource_Read_Happy(t *testing.T) {
	r := &GetTimePropertiesByAliasDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetTimePropertiesByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTimePropertiesByAliasDataSource_Read_NilClient exercises GetTimePropertiesByAliasDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTimePropertiesByAliasDataSource_Read_NilClient(t *testing.T) {
	r := &GetTimePropertiesByAliasDataSource{}
	m := GetTimePropertiesByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTimePropertiesByAliasDataSource_Read_BuildError exercises GetTimePropertiesByAliasDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetTimePropertiesByAliasDataSource_Read_BuildError(t *testing.T) {
	r := &GetTimePropertiesByAliasDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTimePropertiesByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetTimePropertiesByAliasDataSource_Read_SendError exercises GetTimePropertiesByAliasDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetTimePropertiesByAliasDataSource_Read_SendError(t *testing.T) {
	r := &GetTimePropertiesByAliasDataSource{client: newTransportErrorClient(t)}
	m := GetTimePropertiesByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetTimePropertiesByAliasDataSource_Read_NotFound exercises GetTimePropertiesByAliasDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetTimePropertiesByAliasDataSource_Read_NotFound(t *testing.T) {
	r := &GetTimePropertiesByAliasDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetTimePropertiesByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetTimePropertiesByAliasDataSource_Read_APIError exercises GetTimePropertiesByAliasDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetTimePropertiesByAliasDataSource_Read_APIError(t *testing.T) {
	r := &GetTimePropertiesByAliasDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetTimePropertiesByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_time_properties_by_alias")
}

// TestGetTimePropertiesByAliasDataSource_Read_APIErrorReadBody exercises GetTimePropertiesByAliasDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetTimePropertiesByAliasDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetTimePropertiesByAliasDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetTimePropertiesByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetTimePropertiesByAliasDataSource_Read_InvalidJSON exercises GetTimePropertiesByAliasDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetTimePropertiesByAliasDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTimePropertiesByAliasDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTimePropertiesByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
