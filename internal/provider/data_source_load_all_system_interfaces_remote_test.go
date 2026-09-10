package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllSystemInterfacesDataSource_Read_Happy exercises LoadAllSystemInterfacesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllSystemInterfacesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllSystemInterfacesDataSource{client: newMockClientStatus(t, 200, "{\"interfaces\":[]}")}
	m := LoadAllSystemInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllSystemInterfacesDataSource_Read_NilClient exercises LoadAllSystemInterfacesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllSystemInterfacesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllSystemInterfacesDataSource{}
	m := LoadAllSystemInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllSystemInterfacesDataSource_Read_BuildError exercises LoadAllSystemInterfacesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllSystemInterfacesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllSystemInterfacesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllSystemInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSystemInterfacesDataSource_Read_SendError exercises LoadAllSystemInterfacesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllSystemInterfacesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllSystemInterfacesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllSystemInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSystemInterfacesDataSource_Read_InvalidJSON exercises LoadAllSystemInterfacesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllSystemInterfacesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllSystemInterfacesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllSystemInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
