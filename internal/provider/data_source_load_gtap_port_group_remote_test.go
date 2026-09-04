package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGtapPortGroupDataSource_Read_Happy exercises LoadGtapPortGroupDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadGtapPortGroupDataSource_Read_Happy(t *testing.T) {
	r := &LoadGtapPortGroupDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadGtapPortGroupDataSource_Read_NilClient exercises LoadGtapPortGroupDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadGtapPortGroupDataSource_Read_NilClient(t *testing.T) {
	r := &LoadGtapPortGroupDataSource{}
	m := LoadGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadGtapPortGroupDataSource_Read_BuildError exercises LoadGtapPortGroupDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadGtapPortGroupDataSource_Read_BuildError(t *testing.T) {
	r := &LoadGtapPortGroupDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadGtapPortGroupDataSource_Read_SendError exercises LoadGtapPortGroupDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadGtapPortGroupDataSource_Read_SendError(t *testing.T) {
	r := &LoadGtapPortGroupDataSource{client: newTransportErrorClient(t)}
	m := LoadGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadGtapPortGroupDataSource_Read_NotFound exercises LoadGtapPortGroupDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadGtapPortGroupDataSource_Read_NotFound(t *testing.T) {
	r := &LoadGtapPortGroupDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadGtapPortGroupDataSource_Read_APIError exercises LoadGtapPortGroupDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadGtapPortGroupDataSource_Read_APIError(t *testing.T) {
	r := &LoadGtapPortGroupDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_gtap_port_group")
}

// TestLoadGtapPortGroupDataSource_Read_APIErrorReadBody exercises LoadGtapPortGroupDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadGtapPortGroupDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadGtapPortGroupDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadGtapPortGroupDataSource_Read_InvalidJSON exercises LoadGtapPortGroupDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadGtapPortGroupDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadGtapPortGroupDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadGtapPortGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
