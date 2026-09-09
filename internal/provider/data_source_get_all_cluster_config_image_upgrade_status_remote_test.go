package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterConfigImageUpgradeStatusDataSource_Read_Happy exercises GetAllClusterConfigImageUpgradeStatusDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterConfigImageUpgradeStatusDataSource_Read_Happy(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusDataSource{client: newMockClientStatus(t, 200, "{\"clustersStatus\":[]}")}
	m := GetAllClusterConfigImageUpgradeStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllClusterConfigImageUpgradeStatusDataSource_Read_NilClient exercises GetAllClusterConfigImageUpgradeStatusDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterConfigImageUpgradeStatusDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusDataSource{}
	m := GetAllClusterConfigImageUpgradeStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllClusterConfigImageUpgradeStatusDataSource_Read_BuildError exercises GetAllClusterConfigImageUpgradeStatusDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterConfigImageUpgradeStatusDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllClusterConfigImageUpgradeStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterConfigImageUpgradeStatusDataSource_Read_SendError exercises GetAllClusterConfigImageUpgradeStatusDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterConfigImageUpgradeStatusDataSource_Read_SendError(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusDataSource{client: newTransportErrorClient(t)}
	m := GetAllClusterConfigImageUpgradeStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterConfigImageUpgradeStatusDataSource_Read_InvalidJSON exercises GetAllClusterConfigImageUpgradeStatusDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterConfigImageUpgradeStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllClusterConfigImageUpgradeStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
