package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNetworkInterfacesDataSource_Read_Happy exercises GetNetworkInterfacesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetNetworkInterfacesDataSource_Read_Happy(t *testing.T) {
	r := &GetNetworkInterfacesDataSource{client: newMockClientStatus(t, 200, "{\"env\":[]}")}
	m := GetNetworkInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetNetworkInterfacesDataSource_Read_NilClient exercises GetNetworkInterfacesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetNetworkInterfacesDataSource_Read_NilClient(t *testing.T) {
	r := &GetNetworkInterfacesDataSource{}
	m := GetNetworkInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetNetworkInterfacesDataSource_Read_BuildError exercises GetNetworkInterfacesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetNetworkInterfacesDataSource_Read_BuildError(t *testing.T) {
	r := &GetNetworkInterfacesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetNetworkInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNetworkInterfacesDataSource_Read_SendError exercises GetNetworkInterfacesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetNetworkInterfacesDataSource_Read_SendError(t *testing.T) {
	r := &GetNetworkInterfacesDataSource{client: newTransportErrorClient(t)}
	m := GetNetworkInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNetworkInterfacesDataSource_Read_InvalidJSON exercises GetNetworkInterfacesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetNetworkInterfacesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetNetworkInterfacesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetNetworkInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
