package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllNameServersGroupedByInterfaceDataSource_Read_Happy exercises LoadAllNameServersGroupedByInterfaceDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllNameServersGroupedByInterfaceDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllNameServersGroupedByInterfaceDataSource{client: newMockClientStatus(t, 200, "{\"nameServers\":[]}")}
	m := LoadAllNameServersGroupedByInterfaceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllNameServersGroupedByInterfaceDataSource_Read_NilClient exercises LoadAllNameServersGroupedByInterfaceDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllNameServersGroupedByInterfaceDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllNameServersGroupedByInterfaceDataSource{}
	m := LoadAllNameServersGroupedByInterfaceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllNameServersGroupedByInterfaceDataSource_Read_BuildError exercises LoadAllNameServersGroupedByInterfaceDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllNameServersGroupedByInterfaceDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllNameServersGroupedByInterfaceDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllNameServersGroupedByInterfaceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNameServersGroupedByInterfaceDataSource_Read_SendError exercises LoadAllNameServersGroupedByInterfaceDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllNameServersGroupedByInterfaceDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllNameServersGroupedByInterfaceDataSource{client: newTransportErrorClient(t)}
	m := LoadAllNameServersGroupedByInterfaceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNameServersGroupedByInterfaceDataSource_Read_InvalidJSON exercises LoadAllNameServersGroupedByInterfaceDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllNameServersGroupedByInterfaceDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllNameServersGroupedByInterfaceDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllNameServersGroupedByInterfaceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
