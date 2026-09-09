package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadDeviceSyslogDataSource_Read_Happy exercises LoadDeviceSyslogDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadDeviceSyslogDataSource_Read_Happy(t *testing.T) {
	r := &LoadDeviceSyslogDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadDeviceSyslogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadDeviceSyslogDataSource_Read_NilClient exercises LoadDeviceSyslogDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadDeviceSyslogDataSource_Read_NilClient(t *testing.T) {
	r := &LoadDeviceSyslogDataSource{}
	m := LoadDeviceSyslogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadDeviceSyslogDataSource_Read_BuildError exercises LoadDeviceSyslogDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadDeviceSyslogDataSource_Read_BuildError(t *testing.T) {
	r := &LoadDeviceSyslogDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadDeviceSyslogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadDeviceSyslogDataSource_Read_SendError exercises LoadDeviceSyslogDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadDeviceSyslogDataSource_Read_SendError(t *testing.T) {
	r := &LoadDeviceSyslogDataSource{client: newTransportErrorClient(t)}
	m := LoadDeviceSyslogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadDeviceSyslogDataSource_Read_NotFound exercises LoadDeviceSyslogDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadDeviceSyslogDataSource_Read_NotFound(t *testing.T) {
	r := &LoadDeviceSyslogDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadDeviceSyslogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadDeviceSyslogDataSource_Read_APIError exercises LoadDeviceSyslogDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadDeviceSyslogDataSource_Read_APIError(t *testing.T) {
	r := &LoadDeviceSyslogDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadDeviceSyslogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_device_syslog")
}

// TestLoadDeviceSyslogDataSource_Read_APIErrorReadBody exercises LoadDeviceSyslogDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadDeviceSyslogDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadDeviceSyslogDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadDeviceSyslogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadDeviceSyslogDataSource_Read_InvalidJSON exercises LoadDeviceSyslogDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadDeviceSyslogDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadDeviceSyslogDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadDeviceSyslogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
