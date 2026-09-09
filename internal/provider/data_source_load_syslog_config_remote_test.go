package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSyslogConfigDataSource_Read_Happy exercises LoadSyslogConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadSyslogConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadSyslogConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadSyslogConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadSyslogConfigDataSource_Read_NilClient exercises LoadSyslogConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadSyslogConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadSyslogConfigDataSource{}
	m := LoadSyslogConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadSyslogConfigDataSource_Read_BuildError exercises LoadSyslogConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadSyslogConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadSyslogConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadSyslogConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadSyslogConfigDataSource_Read_SendError exercises LoadSyslogConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadSyslogConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadSyslogConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadSyslogConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadSyslogConfigDataSource_Read_NotFound exercises LoadSyslogConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadSyslogConfigDataSource_Read_NotFound(t *testing.T) {
	r := &LoadSyslogConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadSyslogConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadSyslogConfigDataSource_Read_APIError exercises LoadSyslogConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadSyslogConfigDataSource_Read_APIError(t *testing.T) {
	r := &LoadSyslogConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadSyslogConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_syslog_config")
}

// TestLoadSyslogConfigDataSource_Read_APIErrorReadBody exercises LoadSyslogConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadSyslogConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadSyslogConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadSyslogConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadSyslogConfigDataSource_Read_InvalidJSON exercises LoadSyslogConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadSyslogConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadSyslogConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadSyslogConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
