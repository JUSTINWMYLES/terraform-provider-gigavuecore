package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSnmpThrottleStatusDataSource_Read_Happy exercises LoadSnmpThrottleStatusDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadSnmpThrottleStatusDataSource_Read_Happy(t *testing.T) {
	r := &LoadSnmpThrottleStatusDataSource{client: newMockClientStatus(t, 200, "{\"snmpThrottleStatus\":[]}")}
	m := LoadSnmpThrottleStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadSnmpThrottleStatusDataSource_Read_NilClient exercises LoadSnmpThrottleStatusDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadSnmpThrottleStatusDataSource_Read_NilClient(t *testing.T) {
	r := &LoadSnmpThrottleStatusDataSource{}
	m := LoadSnmpThrottleStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadSnmpThrottleStatusDataSource_Read_BuildError exercises LoadSnmpThrottleStatusDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadSnmpThrottleStatusDataSource_Read_BuildError(t *testing.T) {
	r := &LoadSnmpThrottleStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadSnmpThrottleStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadSnmpThrottleStatusDataSource_Read_SendError exercises LoadSnmpThrottleStatusDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadSnmpThrottleStatusDataSource_Read_SendError(t *testing.T) {
	r := &LoadSnmpThrottleStatusDataSource{client: newTransportErrorClient(t)}
	m := LoadSnmpThrottleStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadSnmpThrottleStatusDataSource_Read_InvalidJSON exercises LoadSnmpThrottleStatusDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadSnmpThrottleStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadSnmpThrottleStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadSnmpThrottleStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
