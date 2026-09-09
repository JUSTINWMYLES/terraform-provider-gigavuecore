package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetProcessedVolumesDataSource_Read_Happy exercises GetProcessedVolumesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetProcessedVolumesDataSource_Read_Happy(t *testing.T) {
	r := &GetProcessedVolumesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetProcessedVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetProcessedVolumesDataSource_Read_NilClient exercises GetProcessedVolumesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetProcessedVolumesDataSource_Read_NilClient(t *testing.T) {
	r := &GetProcessedVolumesDataSource{}
	m := GetProcessedVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetProcessedVolumesDataSource_Read_BuildError exercises GetProcessedVolumesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetProcessedVolumesDataSource_Read_BuildError(t *testing.T) {
	r := &GetProcessedVolumesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetProcessedVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetProcessedVolumesDataSource_Read_SendError exercises GetProcessedVolumesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetProcessedVolumesDataSource_Read_SendError(t *testing.T) {
	r := &GetProcessedVolumesDataSource{client: newTransportErrorClient(t)}
	m := GetProcessedVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetProcessedVolumesDataSource_Read_NotFound exercises GetProcessedVolumesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetProcessedVolumesDataSource_Read_NotFound(t *testing.T) {
	r := &GetProcessedVolumesDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetProcessedVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetProcessedVolumesDataSource_Read_APIError exercises GetProcessedVolumesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetProcessedVolumesDataSource_Read_APIError(t *testing.T) {
	r := &GetProcessedVolumesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetProcessedVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_processed_volumes")
}

// TestGetProcessedVolumesDataSource_Read_APIErrorReadBody exercises GetProcessedVolumesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetProcessedVolumesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetProcessedVolumesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetProcessedVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetProcessedVolumesDataSource_Read_InvalidJSON exercises GetProcessedVolumesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetProcessedVolumesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetProcessedVolumesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetProcessedVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
