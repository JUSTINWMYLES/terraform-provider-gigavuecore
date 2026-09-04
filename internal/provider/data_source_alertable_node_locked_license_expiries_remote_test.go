package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestAlertableNodeLockedLicenseExpiriesDataSource_Read_Happy exercises AlertableNodeLockedLicenseExpiriesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestAlertableNodeLockedLicenseExpiriesDataSource_Read_Happy(t *testing.T) {
	r := &AlertableNodeLockedLicenseExpiriesDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := AlertableNodeLockedLicenseExpiriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAlertableNodeLockedLicenseExpiriesDataSource_Read_NilClient exercises AlertableNodeLockedLicenseExpiriesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAlertableNodeLockedLicenseExpiriesDataSource_Read_NilClient(t *testing.T) {
	r := &AlertableNodeLockedLicenseExpiriesDataSource{}
	m := AlertableNodeLockedLicenseExpiriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAlertableNodeLockedLicenseExpiriesDataSource_Read_BuildError exercises AlertableNodeLockedLicenseExpiriesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestAlertableNodeLockedLicenseExpiriesDataSource_Read_BuildError(t *testing.T) {
	r := &AlertableNodeLockedLicenseExpiriesDataSource{client: newMalformedBaseURLClient(t)}
	m := AlertableNodeLockedLicenseExpiriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestAlertableNodeLockedLicenseExpiriesDataSource_Read_SendError exercises AlertableNodeLockedLicenseExpiriesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestAlertableNodeLockedLicenseExpiriesDataSource_Read_SendError(t *testing.T) {
	r := &AlertableNodeLockedLicenseExpiriesDataSource{client: newTransportErrorClient(t)}
	m := AlertableNodeLockedLicenseExpiriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestAlertableNodeLockedLicenseExpiriesDataSource_Read_InvalidJSON exercises AlertableNodeLockedLicenseExpiriesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestAlertableNodeLockedLicenseExpiriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &AlertableNodeLockedLicenseExpiriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := AlertableNodeLockedLicenseExpiriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
