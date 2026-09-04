package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadPtpConfigDataSource_Read_Happy exercises LoadPtpConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadPtpConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadPtpConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadPtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadPtpConfigDataSource_Read_NilClient exercises LoadPtpConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadPtpConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadPtpConfigDataSource{}
	m := LoadPtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadPtpConfigDataSource_Read_BuildError exercises LoadPtpConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadPtpConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadPtpConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadPtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadPtpConfigDataSource_Read_SendError exercises LoadPtpConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadPtpConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadPtpConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadPtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadPtpConfigDataSource_Read_NotFound exercises LoadPtpConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadPtpConfigDataSource_Read_NotFound(t *testing.T) {
	r := &LoadPtpConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadPtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadPtpConfigDataSource_Read_APIError exercises LoadPtpConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadPtpConfigDataSource_Read_APIError(t *testing.T) {
	r := &LoadPtpConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadPtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_ptp_config")
}

// TestLoadPtpConfigDataSource_Read_APIErrorReadBody exercises LoadPtpConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadPtpConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadPtpConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadPtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadPtpConfigDataSource_Read_InvalidJSON exercises LoadPtpConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadPtpConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadPtpConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadPtpConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
