package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTunnelLogicalGroupsDataSource_Read_Happy exercises GetAllTunnelLogicalGroupsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllTunnelLogicalGroupsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllTunnelLogicalGroupsDataSource{client: newMockClientStatus(t, 200, "{\"tunnelLogicalGroups\":[]}")}
	m := GetAllTunnelLogicalGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllTunnelLogicalGroupsDataSource_Read_NilClient exercises GetAllTunnelLogicalGroupsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllTunnelLogicalGroupsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllTunnelLogicalGroupsDataSource{}
	m := GetAllTunnelLogicalGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllTunnelLogicalGroupsDataSource_Read_BuildError exercises GetAllTunnelLogicalGroupsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllTunnelLogicalGroupsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllTunnelLogicalGroupsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllTunnelLogicalGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTunnelLogicalGroupsDataSource_Read_SendError exercises GetAllTunnelLogicalGroupsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllTunnelLogicalGroupsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllTunnelLogicalGroupsDataSource{client: newTransportErrorClient(t)}
	m := GetAllTunnelLogicalGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTunnelLogicalGroupsDataSource_Read_InvalidJSON exercises GetAllTunnelLogicalGroupsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllTunnelLogicalGroupsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllTunnelLogicalGroupsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllTunnelLogicalGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
