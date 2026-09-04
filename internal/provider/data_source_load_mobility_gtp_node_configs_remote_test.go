package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilityGtpNodeConfigsDataSource_Read_Happy exercises LoadMobilityGtpNodeConfigsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadMobilityGtpNodeConfigsDataSource_Read_Happy(t *testing.T) {
	r := &LoadMobilityGtpNodeConfigsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadMobilityGtpNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadMobilityGtpNodeConfigsDataSource_Read_NilClient exercises LoadMobilityGtpNodeConfigsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadMobilityGtpNodeConfigsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadMobilityGtpNodeConfigsDataSource{}
	m := LoadMobilityGtpNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadMobilityGtpNodeConfigsDataSource_Read_BuildError exercises LoadMobilityGtpNodeConfigsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadMobilityGtpNodeConfigsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadMobilityGtpNodeConfigsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadMobilityGtpNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadMobilityGtpNodeConfigsDataSource_Read_SendError exercises LoadMobilityGtpNodeConfigsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadMobilityGtpNodeConfigsDataSource_Read_SendError(t *testing.T) {
	r := &LoadMobilityGtpNodeConfigsDataSource{client: newTransportErrorClient(t)}
	m := LoadMobilityGtpNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadMobilityGtpNodeConfigsDataSource_Read_NotFound exercises LoadMobilityGtpNodeConfigsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadMobilityGtpNodeConfigsDataSource_Read_NotFound(t *testing.T) {
	r := &LoadMobilityGtpNodeConfigsDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadMobilityGtpNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadMobilityGtpNodeConfigsDataSource_Read_APIError exercises LoadMobilityGtpNodeConfigsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadMobilityGtpNodeConfigsDataSource_Read_APIError(t *testing.T) {
	r := &LoadMobilityGtpNodeConfigsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadMobilityGtpNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_mobility_gtp_node_configs")
}

// TestLoadMobilityGtpNodeConfigsDataSource_Read_APIErrorReadBody exercises LoadMobilityGtpNodeConfigsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadMobilityGtpNodeConfigsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadMobilityGtpNodeConfigsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadMobilityGtpNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadMobilityGtpNodeConfigsDataSource_Read_InvalidJSON exercises LoadMobilityGtpNodeConfigsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadMobilityGtpNodeConfigsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadMobilityGtpNodeConfigsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadMobilityGtpNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
