package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_Happy exercises GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_Happy(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_NilClient exercises GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_NilClient(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{}
	m := GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_BuildError exercises GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_BuildError(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newMalformedBaseURLClient(t)}
	m := GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_SendError exercises GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_SendError(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newTransportErrorClient(t)}
	m := GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_NotFound exercises GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_NotFound(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_APIError exercises GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_APIError(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_cluster_config_image_upgrade_status_by_task_group_id")
}

// TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_APIErrorReadBody exercises GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_InvalidJSON exercises GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
