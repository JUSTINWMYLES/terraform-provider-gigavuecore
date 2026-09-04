package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTroubleshootClusterConfigDataSource_Read_Happy exercises GetTroubleshootClusterConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetTroubleshootClusterConfigDataSource_Read_Happy(t *testing.T) {
	r := &GetTroubleshootClusterConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetTroubleshootClusterConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTroubleshootClusterConfigDataSource_Read_NilClient exercises GetTroubleshootClusterConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTroubleshootClusterConfigDataSource_Read_NilClient(t *testing.T) {
	r := &GetTroubleshootClusterConfigDataSource{}
	m := GetTroubleshootClusterConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTroubleshootClusterConfigDataSource_Read_BuildError exercises GetTroubleshootClusterConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetTroubleshootClusterConfigDataSource_Read_BuildError(t *testing.T) {
	r := &GetTroubleshootClusterConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTroubleshootClusterConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetTroubleshootClusterConfigDataSource_Read_SendError exercises GetTroubleshootClusterConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetTroubleshootClusterConfigDataSource_Read_SendError(t *testing.T) {
	r := &GetTroubleshootClusterConfigDataSource{client: newTransportErrorClient(t)}
	m := GetTroubleshootClusterConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetTroubleshootClusterConfigDataSource_Read_NotFound exercises GetTroubleshootClusterConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetTroubleshootClusterConfigDataSource_Read_NotFound(t *testing.T) {
	r := &GetTroubleshootClusterConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetTroubleshootClusterConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetTroubleshootClusterConfigDataSource_Read_APIError exercises GetTroubleshootClusterConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetTroubleshootClusterConfigDataSource_Read_APIError(t *testing.T) {
	r := &GetTroubleshootClusterConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetTroubleshootClusterConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_troubleshoot_cluster_config")
}

// TestGetTroubleshootClusterConfigDataSource_Read_APIErrorReadBody exercises GetTroubleshootClusterConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetTroubleshootClusterConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetTroubleshootClusterConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetTroubleshootClusterConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetTroubleshootClusterConfigDataSource_Read_InvalidJSON exercises GetTroubleshootClusterConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetTroubleshootClusterConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTroubleshootClusterConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTroubleshootClusterConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
