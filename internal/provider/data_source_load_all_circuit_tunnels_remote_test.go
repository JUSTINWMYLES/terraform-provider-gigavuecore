package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllCircuitTunnelsDataSource_Read_Happy exercises LoadAllCircuitTunnelsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllCircuitTunnelsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllCircuitTunnelsDataSource{client: newMockClientStatus(t, 200, "{\"circuitTunnels\":[]}")}
	m := LoadAllCircuitTunnelsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllCircuitTunnelsDataSource_Read_NilClient exercises LoadAllCircuitTunnelsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllCircuitTunnelsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllCircuitTunnelsDataSource{}
	m := LoadAllCircuitTunnelsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllCircuitTunnelsDataSource_Read_BuildError exercises LoadAllCircuitTunnelsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllCircuitTunnelsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllCircuitTunnelsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllCircuitTunnelsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllCircuitTunnelsDataSource_Read_SendError exercises LoadAllCircuitTunnelsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllCircuitTunnelsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllCircuitTunnelsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllCircuitTunnelsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllCircuitTunnelsDataSource_Read_InvalidJSON exercises LoadAllCircuitTunnelsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllCircuitTunnelsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllCircuitTunnelsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllCircuitTunnelsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
