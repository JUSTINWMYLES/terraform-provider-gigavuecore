package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadIpDestinationStatusDataSource_Read_Happy exercises LoadIpDestinationStatusDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadIpDestinationStatusDataSource_Read_Happy(t *testing.T) {
	r := &LoadIpDestinationStatusDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadIpDestinationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadIpDestinationStatusDataSource_Read_NilClient exercises LoadIpDestinationStatusDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadIpDestinationStatusDataSource_Read_NilClient(t *testing.T) {
	r := &LoadIpDestinationStatusDataSource{}
	m := LoadIpDestinationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadIpDestinationStatusDataSource_Read_BuildError exercises LoadIpDestinationStatusDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadIpDestinationStatusDataSource_Read_BuildError(t *testing.T) {
	r := &LoadIpDestinationStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadIpDestinationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadIpDestinationStatusDataSource_Read_SendError exercises LoadIpDestinationStatusDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadIpDestinationStatusDataSource_Read_SendError(t *testing.T) {
	r := &LoadIpDestinationStatusDataSource{client: newTransportErrorClient(t)}
	m := LoadIpDestinationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadIpDestinationStatusDataSource_Read_NotFound exercises LoadIpDestinationStatusDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadIpDestinationStatusDataSource_Read_NotFound(t *testing.T) {
	r := &LoadIpDestinationStatusDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadIpDestinationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadIpDestinationStatusDataSource_Read_APIError exercises LoadIpDestinationStatusDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadIpDestinationStatusDataSource_Read_APIError(t *testing.T) {
	r := &LoadIpDestinationStatusDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadIpDestinationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_ip_destination_status")
}

// TestLoadIpDestinationStatusDataSource_Read_APIErrorReadBody exercises LoadIpDestinationStatusDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadIpDestinationStatusDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadIpDestinationStatusDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadIpDestinationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadIpDestinationStatusDataSource_Read_InvalidJSON exercises LoadIpDestinationStatusDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadIpDestinationStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadIpDestinationStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadIpDestinationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
