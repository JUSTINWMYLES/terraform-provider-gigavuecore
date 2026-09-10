package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestRefreshAllClusterLicensesDataSource_Read_Happy exercises RefreshAllClusterLicensesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestRefreshAllClusterLicensesDataSource_Read_Happy(t *testing.T) {
	r := &RefreshAllClusterLicensesDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := RefreshAllClusterLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRefreshAllClusterLicensesDataSource_Read_NilClient exercises RefreshAllClusterLicensesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRefreshAllClusterLicensesDataSource_Read_NilClient(t *testing.T) {
	r := &RefreshAllClusterLicensesDataSource{}
	m := RefreshAllClusterLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRefreshAllClusterLicensesDataSource_Read_BuildError exercises RefreshAllClusterLicensesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestRefreshAllClusterLicensesDataSource_Read_BuildError(t *testing.T) {
	r := &RefreshAllClusterLicensesDataSource{client: newMalformedBaseURLClient(t)}
	m := RefreshAllClusterLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestRefreshAllClusterLicensesDataSource_Read_SendError exercises RefreshAllClusterLicensesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestRefreshAllClusterLicensesDataSource_Read_SendError(t *testing.T) {
	r := &RefreshAllClusterLicensesDataSource{client: newTransportErrorClient(t)}
	m := RefreshAllClusterLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestRefreshAllClusterLicensesDataSource_Read_InvalidJSON exercises RefreshAllClusterLicensesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestRefreshAllClusterLicensesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &RefreshAllClusterLicensesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := RefreshAllClusterLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
