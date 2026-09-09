package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_Happy exercises GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_Happy(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newMockClientStatus(t, 200, "{\"upgradeTaskStatus\":[]}")}
	m := GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_NilClient exercises GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{}
	m := GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_BuildError exercises GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_SendError exercises GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_SendError(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newTransportErrorClient(t)}
	m := GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_InvalidJSON exercises GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
