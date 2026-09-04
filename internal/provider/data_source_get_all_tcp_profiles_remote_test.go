package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTcpProfilesDataSource_Read_Happy exercises GetAllTcpProfilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllTcpProfilesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllTcpProfilesDataSource{client: newMockClientStatus(t, 200, "{\"appsTcpProfiles\":[]}")}
	m := GetAllTcpProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllTcpProfilesDataSource_Read_NilClient exercises GetAllTcpProfilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllTcpProfilesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllTcpProfilesDataSource{}
	m := GetAllTcpProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllTcpProfilesDataSource_Read_BuildError exercises GetAllTcpProfilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllTcpProfilesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllTcpProfilesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllTcpProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTcpProfilesDataSource_Read_SendError exercises GetAllTcpProfilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllTcpProfilesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllTcpProfilesDataSource{client: newTransportErrorClient(t)}
	m := GetAllTcpProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTcpProfilesDataSource_Read_InvalidJSON exercises GetAllTcpProfilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllTcpProfilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllTcpProfilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllTcpProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
