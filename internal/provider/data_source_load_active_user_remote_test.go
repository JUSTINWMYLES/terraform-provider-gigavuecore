package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadActiveUserDataSource_Read_Happy exercises LoadActiveUserDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadActiveUserDataSource_Read_Happy(t *testing.T) {
	r := &LoadActiveUserDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadActiveUserDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadActiveUserDataSource_Read_NilClient exercises LoadActiveUserDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadActiveUserDataSource_Read_NilClient(t *testing.T) {
	r := &LoadActiveUserDataSource{}
	m := LoadActiveUserDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadActiveUserDataSource_Read_BuildError exercises LoadActiveUserDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadActiveUserDataSource_Read_BuildError(t *testing.T) {
	r := &LoadActiveUserDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadActiveUserDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadActiveUserDataSource_Read_SendError exercises LoadActiveUserDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadActiveUserDataSource_Read_SendError(t *testing.T) {
	r := &LoadActiveUserDataSource{client: newTransportErrorClient(t)}
	m := LoadActiveUserDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadActiveUserDataSource_Read_NotFound exercises LoadActiveUserDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadActiveUserDataSource_Read_NotFound(t *testing.T) {
	r := &LoadActiveUserDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadActiveUserDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadActiveUserDataSource_Read_APIError exercises LoadActiveUserDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadActiveUserDataSource_Read_APIError(t *testing.T) {
	r := &LoadActiveUserDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadActiveUserDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_active_user")
}

// TestLoadActiveUserDataSource_Read_APIErrorReadBody exercises LoadActiveUserDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadActiveUserDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadActiveUserDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadActiveUserDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadActiveUserDataSource_Read_InvalidJSON exercises LoadActiveUserDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadActiveUserDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadActiveUserDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadActiveUserDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
