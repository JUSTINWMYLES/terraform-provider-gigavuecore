package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUserTokensForPrivilegeUsersDataSource_Read_Happy exercises GetUserTokensForPrivilegeUsersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetUserTokensForPrivilegeUsersDataSource_Read_Happy(t *testing.T) {
	r := &GetUserTokensForPrivilegeUsersDataSource{client: newMockClientStatus(t, 200, "{\"fmUserTokenEntities\":[]}")}
	m := GetUserTokensForPrivilegeUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetUserTokensForPrivilegeUsersDataSource_Read_NilClient exercises GetUserTokensForPrivilegeUsersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetUserTokensForPrivilegeUsersDataSource_Read_NilClient(t *testing.T) {
	r := &GetUserTokensForPrivilegeUsersDataSource{}
	m := GetUserTokensForPrivilegeUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetUserTokensForPrivilegeUsersDataSource_Read_BuildError exercises GetUserTokensForPrivilegeUsersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetUserTokensForPrivilegeUsersDataSource_Read_BuildError(t *testing.T) {
	r := &GetUserTokensForPrivilegeUsersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetUserTokensForPrivilegeUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetUserTokensForPrivilegeUsersDataSource_Read_SendError exercises GetUserTokensForPrivilegeUsersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetUserTokensForPrivilegeUsersDataSource_Read_SendError(t *testing.T) {
	r := &GetUserTokensForPrivilegeUsersDataSource{client: newTransportErrorClient(t)}
	m := GetUserTokensForPrivilegeUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetUserTokensForPrivilegeUsersDataSource_Read_InvalidJSON exercises GetUserTokensForPrivilegeUsersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetUserTokensForPrivilegeUsersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetUserTokensForPrivilegeUsersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetUserTokensForPrivilegeUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
