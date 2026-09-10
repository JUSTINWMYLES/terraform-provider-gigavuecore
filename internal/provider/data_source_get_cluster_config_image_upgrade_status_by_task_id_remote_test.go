package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_Happy exercises GetClusterConfigImageUpgradeStatusByTaskIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_Happy(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetClusterConfigImageUpgradeStatusByTaskIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_NilClient exercises GetClusterConfigImageUpgradeStatusByTaskIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_NilClient(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskIdDataSource{}
	m := GetClusterConfigImageUpgradeStatusByTaskIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_BuildError exercises GetClusterConfigImageUpgradeStatusByTaskIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_BuildError(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskIdDataSource{client: newMalformedBaseURLClient(t)}
	m := GetClusterConfigImageUpgradeStatusByTaskIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_SendError exercises GetClusterConfigImageUpgradeStatusByTaskIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_SendError(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskIdDataSource{client: newTransportErrorClient(t)}
	m := GetClusterConfigImageUpgradeStatusByTaskIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_NotFound exercises GetClusterConfigImageUpgradeStatusByTaskIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_NotFound(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetClusterConfigImageUpgradeStatusByTaskIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_APIError exercises GetClusterConfigImageUpgradeStatusByTaskIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_APIError(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetClusterConfigImageUpgradeStatusByTaskIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_cluster_config_image_upgrade_status_by_task_id")
}

// TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_APIErrorReadBody exercises GetClusterConfigImageUpgradeStatusByTaskIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetClusterConfigImageUpgradeStatusByTaskIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_InvalidJSON exercises GetClusterConfigImageUpgradeStatusByTaskIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetClusterConfigImageUpgradeStatusByTaskIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetClusterConfigImageUpgradeStatusByTaskIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
