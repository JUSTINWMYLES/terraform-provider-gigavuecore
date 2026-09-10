package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllNetflowMonitorsDataSource_Read_Happy exercises LoadAllNetflowMonitorsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllNetflowMonitorsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllNetflowMonitorsDataSource{client: newMockClientStatus(t, 200, "{\"nfMonitors\":[]}")}
	m := LoadAllNetflowMonitorsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllNetflowMonitorsDataSource_Read_NilClient exercises LoadAllNetflowMonitorsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllNetflowMonitorsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllNetflowMonitorsDataSource{}
	m := LoadAllNetflowMonitorsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllNetflowMonitorsDataSource_Read_BuildError exercises LoadAllNetflowMonitorsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllNetflowMonitorsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllNetflowMonitorsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllNetflowMonitorsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNetflowMonitorsDataSource_Read_SendError exercises LoadAllNetflowMonitorsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllNetflowMonitorsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllNetflowMonitorsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllNetflowMonitorsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNetflowMonitorsDataSource_Read_InvalidJSON exercises LoadAllNetflowMonitorsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllNetflowMonitorsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllNetflowMonitorsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllNetflowMonitorsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
