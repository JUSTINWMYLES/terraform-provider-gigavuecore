package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllSffpProfileDataSource_Read_Happy exercises GetAllSffpProfileDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllSffpProfileDataSource_Read_Happy(t *testing.T) {
	r := &GetAllSffpProfileDataSource{client: newMockClientStatus(t, 200, "{\"sffpProfiles\":[]}")}
	m := GetAllSffpProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllSffpProfileDataSource_Read_NilClient exercises GetAllSffpProfileDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllSffpProfileDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllSffpProfileDataSource{}
	m := GetAllSffpProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllSffpProfileDataSource_Read_BuildError exercises GetAllSffpProfileDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllSffpProfileDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllSffpProfileDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllSffpProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllSffpProfileDataSource_Read_SendError exercises GetAllSffpProfileDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllSffpProfileDataSource_Read_SendError(t *testing.T) {
	r := &GetAllSffpProfileDataSource{client: newTransportErrorClient(t)}
	m := GetAllSffpProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllSffpProfileDataSource_Read_InvalidJSON exercises GetAllSffpProfileDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllSffpProfileDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllSffpProfileDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllSffpProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
