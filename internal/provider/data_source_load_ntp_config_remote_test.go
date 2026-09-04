package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadNtpConfigDataSource_Read_Happy exercises LoadNtpConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadNtpConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadNtpConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadNtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadNtpConfigDataSource_Read_NilClient exercises LoadNtpConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadNtpConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadNtpConfigDataSource{}
	m := LoadNtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadNtpConfigDataSource_Read_BuildError exercises LoadNtpConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadNtpConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadNtpConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadNtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadNtpConfigDataSource_Read_SendError exercises LoadNtpConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadNtpConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadNtpConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadNtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadNtpConfigDataSource_Read_NotFound exercises LoadNtpConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadNtpConfigDataSource_Read_NotFound(t *testing.T) {
	r := &LoadNtpConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadNtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadNtpConfigDataSource_Read_APIError exercises LoadNtpConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadNtpConfigDataSource_Read_APIError(t *testing.T) {
	r := &LoadNtpConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadNtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_ntp_config")
}

// TestLoadNtpConfigDataSource_Read_APIErrorReadBody exercises LoadNtpConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadNtpConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadNtpConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadNtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadNtpConfigDataSource_Read_InvalidJSON exercises LoadNtpConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadNtpConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadNtpConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadNtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
