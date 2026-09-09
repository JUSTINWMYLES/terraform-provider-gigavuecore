package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource_Read_Happy exercises GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource_Read_Happy(t *testing.T) {
	r := &GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource{client: newMockClientStatus(t, 200, "{\"trafficPolicyGraphEndpointInterfaceMappings\":[]}")}
	m := GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource_Read_NilClient exercises GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource_Read_NilClient(t *testing.T) {
	r := &GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource{}
	m := GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource_Read_BuildError exercises GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource_Read_BuildError(t *testing.T) {
	r := &GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource_Read_SendError exercises GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource_Read_SendError(t *testing.T) {
	r := &GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource{client: newTransportErrorClient(t)}
	m := GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource_Read_InvalidJSON exercises GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
