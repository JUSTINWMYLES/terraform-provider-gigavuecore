package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllCircuitTunnelGlobalDataSource_Read_Happy exercises LoadAllCircuitTunnelGlobalDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllCircuitTunnelGlobalDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllCircuitTunnelGlobalDataSource{client: newMockClientStatus(t, 200, "{\"circuitTunnelGlobals\":[]}")}
	m := LoadAllCircuitTunnelGlobalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllCircuitTunnelGlobalDataSource_Read_NilClient exercises LoadAllCircuitTunnelGlobalDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllCircuitTunnelGlobalDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllCircuitTunnelGlobalDataSource{}
	m := LoadAllCircuitTunnelGlobalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllCircuitTunnelGlobalDataSource_Read_BuildError exercises LoadAllCircuitTunnelGlobalDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllCircuitTunnelGlobalDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllCircuitTunnelGlobalDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllCircuitTunnelGlobalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllCircuitTunnelGlobalDataSource_Read_SendError exercises LoadAllCircuitTunnelGlobalDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllCircuitTunnelGlobalDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllCircuitTunnelGlobalDataSource{client: newTransportErrorClient(t)}
	m := LoadAllCircuitTunnelGlobalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllCircuitTunnelGlobalDataSource_Read_InvalidJSON exercises LoadAllCircuitTunnelGlobalDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllCircuitTunnelGlobalDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllCircuitTunnelGlobalDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllCircuitTunnelGlobalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
