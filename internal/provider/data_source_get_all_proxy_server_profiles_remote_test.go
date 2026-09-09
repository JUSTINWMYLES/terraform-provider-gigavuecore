package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllProxyServerProfilesDataSource_Read_Happy exercises GetAllProxyServerProfilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllProxyServerProfilesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllProxyServerProfilesDataSource{client: newMockClientStatus(t, 200, "{\"appsProxyServerProfiles\":[]}")}
	m := GetAllProxyServerProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllProxyServerProfilesDataSource_Read_NilClient exercises GetAllProxyServerProfilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllProxyServerProfilesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllProxyServerProfilesDataSource{}
	m := GetAllProxyServerProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllProxyServerProfilesDataSource_Read_BuildError exercises GetAllProxyServerProfilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllProxyServerProfilesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllProxyServerProfilesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllProxyServerProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllProxyServerProfilesDataSource_Read_SendError exercises GetAllProxyServerProfilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllProxyServerProfilesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllProxyServerProfilesDataSource{client: newTransportErrorClient(t)}
	m := GetAllProxyServerProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllProxyServerProfilesDataSource_Read_InvalidJSON exercises GetAllProxyServerProfilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllProxyServerProfilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllProxyServerProfilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllProxyServerProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
