package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllSnmpNotifTargetsDataSource_Read_Happy exercises LoadAllSnmpNotifTargetsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllSnmpNotifTargetsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllSnmpNotifTargetsDataSource{client: newMockClientStatus(t, 200, "{\"notifTargets\":[]}")}
	m := LoadAllSnmpNotifTargetsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllSnmpNotifTargetsDataSource_Read_NilClient exercises LoadAllSnmpNotifTargetsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllSnmpNotifTargetsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllSnmpNotifTargetsDataSource{}
	m := LoadAllSnmpNotifTargetsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllSnmpNotifTargetsDataSource_Read_BuildError exercises LoadAllSnmpNotifTargetsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllSnmpNotifTargetsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllSnmpNotifTargetsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllSnmpNotifTargetsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSnmpNotifTargetsDataSource_Read_SendError exercises LoadAllSnmpNotifTargetsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllSnmpNotifTargetsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllSnmpNotifTargetsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllSnmpNotifTargetsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSnmpNotifTargetsDataSource_Read_InvalidJSON exercises LoadAllSnmpNotifTargetsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllSnmpNotifTargetsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllSnmpNotifTargetsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllSnmpNotifTargetsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
