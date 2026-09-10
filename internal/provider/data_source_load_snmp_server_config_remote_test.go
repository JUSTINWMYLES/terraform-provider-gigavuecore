package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSnmpServerConfigDataSource_Read_Happy exercises LoadSnmpServerConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadSnmpServerConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadSnmpServerConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadSnmpServerConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadSnmpServerConfigDataSource_Read_NilClient exercises LoadSnmpServerConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadSnmpServerConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadSnmpServerConfigDataSource{}
	m := LoadSnmpServerConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadSnmpServerConfigDataSource_Read_BuildError exercises LoadSnmpServerConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadSnmpServerConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadSnmpServerConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadSnmpServerConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadSnmpServerConfigDataSource_Read_SendError exercises LoadSnmpServerConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadSnmpServerConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadSnmpServerConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadSnmpServerConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadSnmpServerConfigDataSource_Read_NotFound exercises LoadSnmpServerConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadSnmpServerConfigDataSource_Read_NotFound(t *testing.T) {
	r := &LoadSnmpServerConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadSnmpServerConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadSnmpServerConfigDataSource_Read_APIError exercises LoadSnmpServerConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadSnmpServerConfigDataSource_Read_APIError(t *testing.T) {
	r := &LoadSnmpServerConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadSnmpServerConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_snmp_server_config")
}

// TestLoadSnmpServerConfigDataSource_Read_APIErrorReadBody exercises LoadSnmpServerConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadSnmpServerConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadSnmpServerConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadSnmpServerConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadSnmpServerConfigDataSource_Read_InvalidJSON exercises LoadSnmpServerConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadSnmpServerConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadSnmpServerConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadSnmpServerConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
