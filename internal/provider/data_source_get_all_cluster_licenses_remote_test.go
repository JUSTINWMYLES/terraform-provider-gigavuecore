package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterLicensesDataSource_Read_Happy exercises GetAllClusterLicensesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterLicensesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllClusterLicensesDataSource{client: newMockClientStatus(t, 200, "{\"clusterLicensesList\":[]}")}
	m := GetAllClusterLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllClusterLicensesDataSource_Read_NilClient exercises GetAllClusterLicensesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterLicensesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllClusterLicensesDataSource{}
	m := GetAllClusterLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllClusterLicensesDataSource_Read_BuildError exercises GetAllClusterLicensesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterLicensesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllClusterLicensesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllClusterLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterLicensesDataSource_Read_SendError exercises GetAllClusterLicensesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterLicensesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllClusterLicensesDataSource{client: newTransportErrorClient(t)}
	m := GetAllClusterLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterLicensesDataSource_Read_InvalidJSON exercises GetAllClusterLicensesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterLicensesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllClusterLicensesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllClusterLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
