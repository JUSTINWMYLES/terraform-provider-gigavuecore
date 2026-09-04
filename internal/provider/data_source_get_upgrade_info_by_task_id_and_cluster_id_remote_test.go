package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_Happy exercises GetUpgradeInfoByTaskIdAndClusterIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_Happy(t *testing.T) {
	r := &GetUpgradeInfoByTaskIdAndClusterIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_NilClient exercises GetUpgradeInfoByTaskIdAndClusterIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_NilClient(t *testing.T) {
	r := &GetUpgradeInfoByTaskIdAndClusterIdDataSource{}
	m := GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_BuildError exercises GetUpgradeInfoByTaskIdAndClusterIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_BuildError(t *testing.T) {
	r := &GetUpgradeInfoByTaskIdAndClusterIdDataSource{client: newMalformedBaseURLClient(t)}
	m := GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_SendError exercises GetUpgradeInfoByTaskIdAndClusterIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_SendError(t *testing.T) {
	r := &GetUpgradeInfoByTaskIdAndClusterIdDataSource{client: newTransportErrorClient(t)}
	m := GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_NotFound exercises GetUpgradeInfoByTaskIdAndClusterIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_NotFound(t *testing.T) {
	r := &GetUpgradeInfoByTaskIdAndClusterIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_APIError exercises GetUpgradeInfoByTaskIdAndClusterIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_APIError(t *testing.T) {
	r := &GetUpgradeInfoByTaskIdAndClusterIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id")
}

// TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_APIErrorReadBody exercises GetUpgradeInfoByTaskIdAndClusterIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetUpgradeInfoByTaskIdAndClusterIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_InvalidJSON exercises GetUpgradeInfoByTaskIdAndClusterIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetUpgradeInfoByTaskIdAndClusterIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetUpgradeInfoByTaskIdAndClusterIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
