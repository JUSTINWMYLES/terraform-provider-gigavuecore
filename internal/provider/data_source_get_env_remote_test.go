package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEnvDataSource_Read_Happy exercises GetEnvDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetEnvDataSource_Read_Happy(t *testing.T) {
	r := &GetEnvDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEnvDataSource_Read_NilClient exercises GetEnvDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEnvDataSource_Read_NilClient(t *testing.T) {
	r := &GetEnvDataSource{}
	m := GetEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEnvDataSource_Read_BuildError exercises GetEnvDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetEnvDataSource_Read_BuildError(t *testing.T) {
	r := &GetEnvDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetEnvDataSource_Read_SendError exercises GetEnvDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetEnvDataSource_Read_SendError(t *testing.T) {
	r := &GetEnvDataSource{client: newTransportErrorClient(t)}
	m := GetEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetEnvDataSource_Read_NotFound exercises GetEnvDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetEnvDataSource_Read_NotFound(t *testing.T) {
	r := &GetEnvDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetEnvDataSource_Read_APIError exercises GetEnvDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetEnvDataSource_Read_APIError(t *testing.T) {
	r := &GetEnvDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_env")
}

// TestGetEnvDataSource_Read_APIErrorReadBody exercises GetEnvDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetEnvDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetEnvDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetEnvDataSource_Read_InvalidJSON exercises GetEnvDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetEnvDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEnvDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
