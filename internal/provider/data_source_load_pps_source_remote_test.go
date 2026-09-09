package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadPpsSourceDataSource_Read_Happy exercises LoadPpsSourceDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadPpsSourceDataSource_Read_Happy(t *testing.T) {
	r := &LoadPpsSourceDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadPpsSourceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadPpsSourceDataSource_Read_NilClient exercises LoadPpsSourceDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadPpsSourceDataSource_Read_NilClient(t *testing.T) {
	r := &LoadPpsSourceDataSource{}
	m := LoadPpsSourceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadPpsSourceDataSource_Read_BuildError exercises LoadPpsSourceDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadPpsSourceDataSource_Read_BuildError(t *testing.T) {
	r := &LoadPpsSourceDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadPpsSourceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadPpsSourceDataSource_Read_SendError exercises LoadPpsSourceDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadPpsSourceDataSource_Read_SendError(t *testing.T) {
	r := &LoadPpsSourceDataSource{client: newTransportErrorClient(t)}
	m := LoadPpsSourceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadPpsSourceDataSource_Read_NotFound exercises LoadPpsSourceDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadPpsSourceDataSource_Read_NotFound(t *testing.T) {
	r := &LoadPpsSourceDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadPpsSourceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadPpsSourceDataSource_Read_APIError exercises LoadPpsSourceDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadPpsSourceDataSource_Read_APIError(t *testing.T) {
	r := &LoadPpsSourceDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadPpsSourceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_pps_source")
}

// TestLoadPpsSourceDataSource_Read_APIErrorReadBody exercises LoadPpsSourceDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadPpsSourceDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadPpsSourceDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadPpsSourceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadPpsSourceDataSource_Read_InvalidJSON exercises LoadPpsSourceDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadPpsSourceDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadPpsSourceDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadPpsSourceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
