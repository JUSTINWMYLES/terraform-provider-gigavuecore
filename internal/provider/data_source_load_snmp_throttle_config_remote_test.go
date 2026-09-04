package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSnmpThrottleConfigDataSource_Read_Happy exercises LoadSnmpThrottleConfigDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadSnmpThrottleConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadSnmpThrottleConfigDataSource{client: newMockClientStatus(t, 200, "{\"throttleConfigDetails\":[]}")}
	m := LoadSnmpThrottleConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadSnmpThrottleConfigDataSource_Read_NilClient exercises LoadSnmpThrottleConfigDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadSnmpThrottleConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadSnmpThrottleConfigDataSource{}
	m := LoadSnmpThrottleConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadSnmpThrottleConfigDataSource_Read_BuildError exercises LoadSnmpThrottleConfigDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadSnmpThrottleConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadSnmpThrottleConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadSnmpThrottleConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadSnmpThrottleConfigDataSource_Read_SendError exercises LoadSnmpThrottleConfigDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadSnmpThrottleConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadSnmpThrottleConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadSnmpThrottleConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadSnmpThrottleConfigDataSource_Read_InvalidJSON exercises LoadSnmpThrottleConfigDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadSnmpThrottleConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadSnmpThrottleConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadSnmpThrottleConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
