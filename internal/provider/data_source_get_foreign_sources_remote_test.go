package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetForeignSourcesDataSource_Read_Happy exercises GetForeignSourcesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetForeignSourcesDataSource_Read_Happy(t *testing.T) {
	r := &GetForeignSourcesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetForeignSourcesDataSource_Read_NilClient exercises GetForeignSourcesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetForeignSourcesDataSource_Read_NilClient(t *testing.T) {
	r := &GetForeignSourcesDataSource{}
	m := GetForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetForeignSourcesDataSource_Read_BuildError exercises GetForeignSourcesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetForeignSourcesDataSource_Read_BuildError(t *testing.T) {
	r := &GetForeignSourcesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetForeignSourcesDataSource_Read_SendError exercises GetForeignSourcesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetForeignSourcesDataSource_Read_SendError(t *testing.T) {
	r := &GetForeignSourcesDataSource{client: newTransportErrorClient(t)}
	m := GetForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetForeignSourcesDataSource_Read_NotFound exercises GetForeignSourcesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetForeignSourcesDataSource_Read_NotFound(t *testing.T) {
	r := &GetForeignSourcesDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetForeignSourcesDataSource_Read_APIError exercises GetForeignSourcesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetForeignSourcesDataSource_Read_APIError(t *testing.T) {
	r := &GetForeignSourcesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_foreign_sources")
}

// TestGetForeignSourcesDataSource_Read_APIErrorReadBody exercises GetForeignSourcesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetForeignSourcesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetForeignSourcesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetForeignSourcesDataSource_Read_InvalidJSON exercises GetForeignSourcesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetForeignSourcesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetForeignSourcesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
