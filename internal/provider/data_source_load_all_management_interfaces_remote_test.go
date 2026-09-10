package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllManagementInterfacesDataSource_Read_Happy exercises LoadAllManagementInterfacesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllManagementInterfacesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllManagementInterfacesDataSource{client: newMockClientStatus(t, 200, "{\"mgmtInterfaces\":[]}")}
	m := LoadAllManagementInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllManagementInterfacesDataSource_Read_NilClient exercises LoadAllManagementInterfacesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllManagementInterfacesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllManagementInterfacesDataSource{}
	m := LoadAllManagementInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllManagementInterfacesDataSource_Read_BuildError exercises LoadAllManagementInterfacesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllManagementInterfacesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllManagementInterfacesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllManagementInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllManagementInterfacesDataSource_Read_SendError exercises LoadAllManagementInterfacesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllManagementInterfacesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllManagementInterfacesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllManagementInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllManagementInterfacesDataSource_Read_InvalidJSON exercises LoadAllManagementInterfacesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllManagementInterfacesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllManagementInterfacesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllManagementInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
