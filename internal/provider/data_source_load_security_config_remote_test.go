package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSecurityConfigDataSource_Read_Happy exercises LoadSecurityConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadSecurityConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadSecurityConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadSecurityConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadSecurityConfigDataSource_Read_NilClient exercises LoadSecurityConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadSecurityConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadSecurityConfigDataSource{}
	m := LoadSecurityConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadSecurityConfigDataSource_Read_BuildError exercises LoadSecurityConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadSecurityConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadSecurityConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadSecurityConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadSecurityConfigDataSource_Read_SendError exercises LoadSecurityConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadSecurityConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadSecurityConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadSecurityConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadSecurityConfigDataSource_Read_NotFound exercises LoadSecurityConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadSecurityConfigDataSource_Read_NotFound(t *testing.T) {
	r := &LoadSecurityConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadSecurityConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadSecurityConfigDataSource_Read_APIError exercises LoadSecurityConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadSecurityConfigDataSource_Read_APIError(t *testing.T) {
	r := &LoadSecurityConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadSecurityConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_security_config")
}

// TestLoadSecurityConfigDataSource_Read_APIErrorReadBody exercises LoadSecurityConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadSecurityConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadSecurityConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadSecurityConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadSecurityConfigDataSource_Read_InvalidJSON exercises LoadSecurityConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadSecurityConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadSecurityConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadSecurityConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
