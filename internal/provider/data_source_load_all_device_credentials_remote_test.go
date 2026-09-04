package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllDeviceCredentialsDataSource_Read_Happy exercises LoadAllDeviceCredentialsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllDeviceCredentialsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllDeviceCredentialsDataSource{client: newMockClientStatus(t, 200, "{\"devCredsList\":[]}")}
	m := LoadAllDeviceCredentialsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllDeviceCredentialsDataSource_Read_NilClient exercises LoadAllDeviceCredentialsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllDeviceCredentialsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllDeviceCredentialsDataSource{}
	m := LoadAllDeviceCredentialsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllDeviceCredentialsDataSource_Read_BuildError exercises LoadAllDeviceCredentialsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllDeviceCredentialsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllDeviceCredentialsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllDeviceCredentialsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllDeviceCredentialsDataSource_Read_SendError exercises LoadAllDeviceCredentialsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllDeviceCredentialsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllDeviceCredentialsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllDeviceCredentialsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllDeviceCredentialsDataSource_Read_InvalidJSON exercises LoadAllDeviceCredentialsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllDeviceCredentialsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllDeviceCredentialsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllDeviceCredentialsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
