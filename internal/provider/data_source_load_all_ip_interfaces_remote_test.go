package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllIpInterfacesDataSource_Read_Happy exercises LoadAllIpInterfacesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllIpInterfacesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllIpInterfacesDataSource{client: newMockClientStatus(t, 200, "{\"ipInterfaces\":[]}")}
	m := LoadAllIpInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllIpInterfacesDataSource_Read_NilClient exercises LoadAllIpInterfacesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllIpInterfacesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllIpInterfacesDataSource{}
	m := LoadAllIpInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllIpInterfacesDataSource_Read_BuildError exercises LoadAllIpInterfacesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllIpInterfacesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllIpInterfacesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllIpInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllIpInterfacesDataSource_Read_SendError exercises LoadAllIpInterfacesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllIpInterfacesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllIpInterfacesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllIpInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllIpInterfacesDataSource_Read_InvalidJSON exercises LoadAllIpInterfacesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllIpInterfacesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllIpInterfacesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllIpInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
