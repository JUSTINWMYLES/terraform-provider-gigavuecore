package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilityGtpNodeDataSource_Read_Happy exercises LoadMobilityGtpNodeDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadMobilityGtpNodeDataSource_Read_Happy(t *testing.T) {
	r := &LoadMobilityGtpNodeDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadMobilityGtpNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadMobilityGtpNodeDataSource_Read_NilClient exercises LoadMobilityGtpNodeDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadMobilityGtpNodeDataSource_Read_NilClient(t *testing.T) {
	r := &LoadMobilityGtpNodeDataSource{}
	m := LoadMobilityGtpNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadMobilityGtpNodeDataSource_Read_BuildError exercises LoadMobilityGtpNodeDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadMobilityGtpNodeDataSource_Read_BuildError(t *testing.T) {
	r := &LoadMobilityGtpNodeDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadMobilityGtpNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadMobilityGtpNodeDataSource_Read_SendError exercises LoadMobilityGtpNodeDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadMobilityGtpNodeDataSource_Read_SendError(t *testing.T) {
	r := &LoadMobilityGtpNodeDataSource{client: newTransportErrorClient(t)}
	m := LoadMobilityGtpNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadMobilityGtpNodeDataSource_Read_NotFound exercises LoadMobilityGtpNodeDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadMobilityGtpNodeDataSource_Read_NotFound(t *testing.T) {
	r := &LoadMobilityGtpNodeDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadMobilityGtpNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadMobilityGtpNodeDataSource_Read_APIError exercises LoadMobilityGtpNodeDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadMobilityGtpNodeDataSource_Read_APIError(t *testing.T) {
	r := &LoadMobilityGtpNodeDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadMobilityGtpNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_mobility_gtp_node")
}

// TestLoadMobilityGtpNodeDataSource_Read_APIErrorReadBody exercises LoadMobilityGtpNodeDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadMobilityGtpNodeDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadMobilityGtpNodeDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadMobilityGtpNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadMobilityGtpNodeDataSource_Read_InvalidJSON exercises LoadMobilityGtpNodeDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadMobilityGtpNodeDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadMobilityGtpNodeDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadMobilityGtpNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
