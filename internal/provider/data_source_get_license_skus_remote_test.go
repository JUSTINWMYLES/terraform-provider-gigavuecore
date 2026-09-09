package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetLicenseSkusDataSource_Read_Happy exercises GetLicenseSkusDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetLicenseSkusDataSource_Read_Happy(t *testing.T) {
	r := &GetLicenseSkusDataSource{client: newMockClientStatus(t, 200, "{\"devices\":[]}")}
	m := GetLicenseSkusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetLicenseSkusDataSource_Read_NilClient exercises GetLicenseSkusDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetLicenseSkusDataSource_Read_NilClient(t *testing.T) {
	r := &GetLicenseSkusDataSource{}
	m := GetLicenseSkusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetLicenseSkusDataSource_Read_BuildError exercises GetLicenseSkusDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetLicenseSkusDataSource_Read_BuildError(t *testing.T) {
	r := &GetLicenseSkusDataSource{client: newMalformedBaseURLClient(t)}
	m := GetLicenseSkusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetLicenseSkusDataSource_Read_SendError exercises GetLicenseSkusDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetLicenseSkusDataSource_Read_SendError(t *testing.T) {
	r := &GetLicenseSkusDataSource{client: newTransportErrorClient(t)}
	m := GetLicenseSkusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetLicenseSkusDataSource_Read_InvalidJSON exercises GetLicenseSkusDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetLicenseSkusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetLicenseSkusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetLicenseSkusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
