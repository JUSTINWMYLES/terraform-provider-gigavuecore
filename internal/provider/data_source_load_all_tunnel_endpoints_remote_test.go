package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllTunnelEndpointsDataSource_Read_Happy exercises LoadAllTunnelEndpointsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllTunnelEndpointsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllTunnelEndpointsDataSource{client: newMockClientStatus(t, 200, "{\"tunnelLbEndpoints\":[]}")}
	m := LoadAllTunnelEndpointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllTunnelEndpointsDataSource_Read_NilClient exercises LoadAllTunnelEndpointsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllTunnelEndpointsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllTunnelEndpointsDataSource{}
	m := LoadAllTunnelEndpointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllTunnelEndpointsDataSource_Read_BuildError exercises LoadAllTunnelEndpointsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllTunnelEndpointsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllTunnelEndpointsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllTunnelEndpointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllTunnelEndpointsDataSource_Read_SendError exercises LoadAllTunnelEndpointsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllTunnelEndpointsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllTunnelEndpointsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllTunnelEndpointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllTunnelEndpointsDataSource_Read_InvalidJSON exercises LoadAllTunnelEndpointsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllTunnelEndpointsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllTunnelEndpointsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllTunnelEndpointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
