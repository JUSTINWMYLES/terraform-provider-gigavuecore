package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSource_Read_Happy exercises GetSerialNumbersOfChassisRequiringGvosLicenseDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSource_Read_Happy(t *testing.T) {
	r := &GetSerialNumbersOfChassisRequiringGvosLicenseDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetSerialNumbersOfChassisRequiringGvosLicenseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSource_Read_NilClient exercises GetSerialNumbersOfChassisRequiringGvosLicenseDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSource_Read_NilClient(t *testing.T) {
	r := &GetSerialNumbersOfChassisRequiringGvosLicenseDataSource{}
	m := GetSerialNumbersOfChassisRequiringGvosLicenseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSource_Read_BuildError exercises GetSerialNumbersOfChassisRequiringGvosLicenseDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSource_Read_BuildError(t *testing.T) {
	r := &GetSerialNumbersOfChassisRequiringGvosLicenseDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSerialNumbersOfChassisRequiringGvosLicenseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSource_Read_SendError exercises GetSerialNumbersOfChassisRequiringGvosLicenseDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSource_Read_SendError(t *testing.T) {
	r := &GetSerialNumbersOfChassisRequiringGvosLicenseDataSource{client: newTransportErrorClient(t)}
	m := GetSerialNumbersOfChassisRequiringGvosLicenseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSource_Read_InvalidJSON exercises GetSerialNumbersOfChassisRequiringGvosLicenseDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSerialNumbersOfChassisRequiringGvosLicenseDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSerialNumbersOfChassisRequiringGvosLicenseDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
