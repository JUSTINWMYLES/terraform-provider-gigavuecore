package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAaaAuthConfigDataSource_Read_Happy exercises LoadAaaAuthConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadAaaAuthConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadAaaAuthConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadAaaAuthConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAaaAuthConfigDataSource_Read_NilClient exercises LoadAaaAuthConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAaaAuthConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAaaAuthConfigDataSource{}
	m := LoadAaaAuthConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAaaAuthConfigDataSource_Read_BuildError exercises LoadAaaAuthConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadAaaAuthConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAaaAuthConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAaaAuthConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadAaaAuthConfigDataSource_Read_SendError exercises LoadAaaAuthConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadAaaAuthConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadAaaAuthConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadAaaAuthConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadAaaAuthConfigDataSource_Read_NotFound exercises LoadAaaAuthConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadAaaAuthConfigDataSource_Read_NotFound(t *testing.T) {
	r := &LoadAaaAuthConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadAaaAuthConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadAaaAuthConfigDataSource_Read_APIError exercises LoadAaaAuthConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadAaaAuthConfigDataSource_Read_APIError(t *testing.T) {
	r := &LoadAaaAuthConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadAaaAuthConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_aaa_auth_config")
}

// TestLoadAaaAuthConfigDataSource_Read_APIErrorReadBody exercises LoadAaaAuthConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadAaaAuthConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadAaaAuthConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadAaaAuthConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadAaaAuthConfigDataSource_Read_InvalidJSON exercises LoadAaaAuthConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadAaaAuthConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAaaAuthConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAaaAuthConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
