package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllGtaProfilesDataSource_Read_Happy exercises GetAllGtaProfilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllGtaProfilesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllGtaProfilesDataSource{client: newMockClientStatus(t, 200, "{\"gtaProfiles\":[]}")}
	m := GetAllGtaProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllGtaProfilesDataSource_Read_NilClient exercises GetAllGtaProfilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllGtaProfilesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllGtaProfilesDataSource{}
	m := GetAllGtaProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllGtaProfilesDataSource_Read_BuildError exercises GetAllGtaProfilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllGtaProfilesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllGtaProfilesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllGtaProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllGtaProfilesDataSource_Read_SendError exercises GetAllGtaProfilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllGtaProfilesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllGtaProfilesDataSource{client: newTransportErrorClient(t)}
	m := GetAllGtaProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllGtaProfilesDataSource_Read_InvalidJSON exercises GetAllGtaProfilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllGtaProfilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllGtaProfilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllGtaProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
