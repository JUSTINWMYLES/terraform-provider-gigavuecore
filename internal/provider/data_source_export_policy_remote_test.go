package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestExportPolicyDataSource_Read_Happy exercises ExportPolicyDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestExportPolicyDataSource_Read_Happy(t *testing.T) {
	r := &ExportPolicyDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := ExportPolicyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExportPolicyDataSource_Read_NilClient exercises ExportPolicyDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExportPolicyDataSource_Read_NilClient(t *testing.T) {
	r := &ExportPolicyDataSource{}
	m := ExportPolicyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExportPolicyDataSource_Read_BuildError exercises ExportPolicyDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExportPolicyDataSource_Read_BuildError(t *testing.T) {
	r := &ExportPolicyDataSource{client: newMalformedBaseURLClient(t)}
	m := ExportPolicyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExportPolicyDataSource_Read_SendError exercises ExportPolicyDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestExportPolicyDataSource_Read_SendError(t *testing.T) {
	r := &ExportPolicyDataSource{client: newTransportErrorClient(t)}
	m := ExportPolicyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExportPolicyDataSource_Read_NotFound exercises ExportPolicyDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestExportPolicyDataSource_Read_NotFound(t *testing.T) {
	r := &ExportPolicyDataSource{client: newMockClientStatus(t, 404, "")}
	m := ExportPolicyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestExportPolicyDataSource_Read_APIError exercises ExportPolicyDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExportPolicyDataSource_Read_APIError(t *testing.T) {
	r := &ExportPolicyDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ExportPolicyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_export_policy")
}

// TestExportPolicyDataSource_Read_APIErrorReadBody exercises ExportPolicyDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExportPolicyDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &ExportPolicyDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := ExportPolicyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestExportPolicyDataSource_Read_InvalidJSON exercises ExportPolicyDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestExportPolicyDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ExportPolicyDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ExportPolicyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
