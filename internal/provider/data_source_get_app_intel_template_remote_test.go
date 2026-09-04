package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAppIntelTemplateDataSource_Read_Happy exercises GetAppIntelTemplateDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAppIntelTemplateDataSource_Read_Happy(t *testing.T) {
	r := &GetAppIntelTemplateDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAppIntelTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAppIntelTemplateDataSource_Read_NilClient exercises GetAppIntelTemplateDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAppIntelTemplateDataSource_Read_NilClient(t *testing.T) {
	r := &GetAppIntelTemplateDataSource{}
	m := GetAppIntelTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAppIntelTemplateDataSource_Read_BuildError exercises GetAppIntelTemplateDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAppIntelTemplateDataSource_Read_BuildError(t *testing.T) {
	r := &GetAppIntelTemplateDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAppIntelTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAppIntelTemplateDataSource_Read_SendError exercises GetAppIntelTemplateDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAppIntelTemplateDataSource_Read_SendError(t *testing.T) {
	r := &GetAppIntelTemplateDataSource{client: newTransportErrorClient(t)}
	m := GetAppIntelTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAppIntelTemplateDataSource_Read_NotFound exercises GetAppIntelTemplateDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAppIntelTemplateDataSource_Read_NotFound(t *testing.T) {
	r := &GetAppIntelTemplateDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAppIntelTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAppIntelTemplateDataSource_Read_APIError exercises GetAppIntelTemplateDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAppIntelTemplateDataSource_Read_APIError(t *testing.T) {
	r := &GetAppIntelTemplateDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAppIntelTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_app_intel_template")
}

// TestGetAppIntelTemplateDataSource_Read_APIErrorReadBody exercises GetAppIntelTemplateDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAppIntelTemplateDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAppIntelTemplateDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAppIntelTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAppIntelTemplateDataSource_Read_InvalidJSON exercises GetAppIntelTemplateDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAppIntelTemplateDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAppIntelTemplateDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAppIntelTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
