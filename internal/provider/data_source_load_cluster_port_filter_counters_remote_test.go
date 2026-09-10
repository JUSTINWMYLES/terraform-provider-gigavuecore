package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadClusterPortFilterCountersDataSource_Read_Happy exercises LoadClusterPortFilterCountersDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadClusterPortFilterCountersDataSource_Read_Happy(t *testing.T) {
	r := &LoadClusterPortFilterCountersDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadClusterPortFilterCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadClusterPortFilterCountersDataSource_Read_NilClient exercises LoadClusterPortFilterCountersDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadClusterPortFilterCountersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadClusterPortFilterCountersDataSource{}
	m := LoadClusterPortFilterCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadClusterPortFilterCountersDataSource_Read_BuildError exercises LoadClusterPortFilterCountersDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadClusterPortFilterCountersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadClusterPortFilterCountersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadClusterPortFilterCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadClusterPortFilterCountersDataSource_Read_SendError exercises LoadClusterPortFilterCountersDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadClusterPortFilterCountersDataSource_Read_SendError(t *testing.T) {
	r := &LoadClusterPortFilterCountersDataSource{client: newTransportErrorClient(t)}
	m := LoadClusterPortFilterCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadClusterPortFilterCountersDataSource_Read_NotFound exercises LoadClusterPortFilterCountersDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadClusterPortFilterCountersDataSource_Read_NotFound(t *testing.T) {
	r := &LoadClusterPortFilterCountersDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadClusterPortFilterCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadClusterPortFilterCountersDataSource_Read_APIError exercises LoadClusterPortFilterCountersDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadClusterPortFilterCountersDataSource_Read_APIError(t *testing.T) {
	r := &LoadClusterPortFilterCountersDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadClusterPortFilterCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_cluster_port_filter_counters")
}

// TestLoadClusterPortFilterCountersDataSource_Read_APIErrorReadBody exercises LoadClusterPortFilterCountersDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadClusterPortFilterCountersDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadClusterPortFilterCountersDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadClusterPortFilterCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadClusterPortFilterCountersDataSource_Read_InvalidJSON exercises LoadClusterPortFilterCountersDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadClusterPortFilterCountersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadClusterPortFilterCountersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadClusterPortFilterCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
