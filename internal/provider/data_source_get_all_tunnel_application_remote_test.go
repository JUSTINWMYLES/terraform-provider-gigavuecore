package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTunnelApplicationDataSource_Read_Happy exercises GetAllTunnelApplicationDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllTunnelApplicationDataSource_Read_Happy(t *testing.T) {
	r := &GetAllTunnelApplicationDataSource{client: newMockClientStatus(t, 200, "{\"tunnelApps\":[]}")}
	m := GetAllTunnelApplicationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllTunnelApplicationDataSource_Read_NilClient exercises GetAllTunnelApplicationDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllTunnelApplicationDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllTunnelApplicationDataSource{}
	m := GetAllTunnelApplicationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllTunnelApplicationDataSource_Read_BuildError exercises GetAllTunnelApplicationDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllTunnelApplicationDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllTunnelApplicationDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllTunnelApplicationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTunnelApplicationDataSource_Read_SendError exercises GetAllTunnelApplicationDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllTunnelApplicationDataSource_Read_SendError(t *testing.T) {
	r := &GetAllTunnelApplicationDataSource{client: newTransportErrorClient(t)}
	m := GetAllTunnelApplicationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTunnelApplicationDataSource_Read_InvalidJSON exercises GetAllTunnelApplicationDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllTunnelApplicationDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllTunnelApplicationDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllTunnelApplicationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
