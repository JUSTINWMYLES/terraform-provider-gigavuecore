package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMapGroupsDataSource_Read_Happy exercises GetMapGroupsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetMapGroupsDataSource_Read_Happy(t *testing.T) {
	r := &GetMapGroupsDataSource{client: newMockClientStatus(t, 200, "{\"mapGroups\":[]}")}
	m := GetMapGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMapGroupsDataSource_Read_NilClient exercises GetMapGroupsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMapGroupsDataSource_Read_NilClient(t *testing.T) {
	r := &GetMapGroupsDataSource{}
	m := GetMapGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMapGroupsDataSource_Read_BuildError exercises GetMapGroupsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetMapGroupsDataSource_Read_BuildError(t *testing.T) {
	r := &GetMapGroupsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMapGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMapGroupsDataSource_Read_SendError exercises GetMapGroupsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetMapGroupsDataSource_Read_SendError(t *testing.T) {
	r := &GetMapGroupsDataSource{client: newTransportErrorClient(t)}
	m := GetMapGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMapGroupsDataSource_Read_InvalidJSON exercises GetMapGroupsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetMapGroupsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMapGroupsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMapGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
