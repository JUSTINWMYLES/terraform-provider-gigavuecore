package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestAlertableFloatingLicenseExpiriesDataSource_Read_Happy exercises AlertableFloatingLicenseExpiriesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestAlertableFloatingLicenseExpiriesDataSource_Read_Happy(t *testing.T) {
	r := &AlertableFloatingLicenseExpiriesDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := AlertableFloatingLicenseExpiriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAlertableFloatingLicenseExpiriesDataSource_Read_NilClient exercises AlertableFloatingLicenseExpiriesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAlertableFloatingLicenseExpiriesDataSource_Read_NilClient(t *testing.T) {
	r := &AlertableFloatingLicenseExpiriesDataSource{}
	m := AlertableFloatingLicenseExpiriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAlertableFloatingLicenseExpiriesDataSource_Read_BuildError exercises AlertableFloatingLicenseExpiriesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestAlertableFloatingLicenseExpiriesDataSource_Read_BuildError(t *testing.T) {
	r := &AlertableFloatingLicenseExpiriesDataSource{client: newMalformedBaseURLClient(t)}
	m := AlertableFloatingLicenseExpiriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestAlertableFloatingLicenseExpiriesDataSource_Read_SendError exercises AlertableFloatingLicenseExpiriesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestAlertableFloatingLicenseExpiriesDataSource_Read_SendError(t *testing.T) {
	r := &AlertableFloatingLicenseExpiriesDataSource{client: newTransportErrorClient(t)}
	m := AlertableFloatingLicenseExpiriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestAlertableFloatingLicenseExpiriesDataSource_Read_InvalidJSON exercises AlertableFloatingLicenseExpiriesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestAlertableFloatingLicenseExpiriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &AlertableFloatingLicenseExpiriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := AlertableFloatingLicenseExpiriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
