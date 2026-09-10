package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeJobInfoDataSource_Read_Happy exercises GetUpgradeJobInfoDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetUpgradeJobInfoDataSource_Read_Happy(t *testing.T) {
	r := &GetUpgradeJobInfoDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetUpgradeJobInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetUpgradeJobInfoDataSource_Read_NilClient exercises GetUpgradeJobInfoDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetUpgradeJobInfoDataSource_Read_NilClient(t *testing.T) {
	r := &GetUpgradeJobInfoDataSource{}
	m := GetUpgradeJobInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetUpgradeJobInfoDataSource_Read_BuildError exercises GetUpgradeJobInfoDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetUpgradeJobInfoDataSource_Read_BuildError(t *testing.T) {
	r := &GetUpgradeJobInfoDataSource{client: newMalformedBaseURLClient(t)}
	m := GetUpgradeJobInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetUpgradeJobInfoDataSource_Read_SendError exercises GetUpgradeJobInfoDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetUpgradeJobInfoDataSource_Read_SendError(t *testing.T) {
	r := &GetUpgradeJobInfoDataSource{client: newTransportErrorClient(t)}
	m := GetUpgradeJobInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetUpgradeJobInfoDataSource_Read_NotFound exercises GetUpgradeJobInfoDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetUpgradeJobInfoDataSource_Read_NotFound(t *testing.T) {
	r := &GetUpgradeJobInfoDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetUpgradeJobInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetUpgradeJobInfoDataSource_Read_APIError exercises GetUpgradeJobInfoDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetUpgradeJobInfoDataSource_Read_APIError(t *testing.T) {
	r := &GetUpgradeJobInfoDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetUpgradeJobInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_upgrade_job_info")
}

// TestGetUpgradeJobInfoDataSource_Read_APIErrorReadBody exercises GetUpgradeJobInfoDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetUpgradeJobInfoDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetUpgradeJobInfoDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetUpgradeJobInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetUpgradeJobInfoDataSource_Read_InvalidJSON exercises GetUpgradeJobInfoDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetUpgradeJobInfoDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetUpgradeJobInfoDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetUpgradeJobInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
