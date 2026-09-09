package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadPermittedAddressesDataSource_Read_Happy exercises LoadPermittedAddressesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadPermittedAddressesDataSource_Read_Happy(t *testing.T) {
	r := &LoadPermittedAddressesDataSource{client: newMockClientStatus(t, 200, "{\"allowListAddresses\":[]}")}
	m := LoadPermittedAddressesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadPermittedAddressesDataSource_Read_NilClient exercises LoadPermittedAddressesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadPermittedAddressesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadPermittedAddressesDataSource{}
	m := LoadPermittedAddressesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadPermittedAddressesDataSource_Read_BuildError exercises LoadPermittedAddressesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadPermittedAddressesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadPermittedAddressesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadPermittedAddressesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadPermittedAddressesDataSource_Read_SendError exercises LoadPermittedAddressesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadPermittedAddressesDataSource_Read_SendError(t *testing.T) {
	r := &LoadPermittedAddressesDataSource{client: newTransportErrorClient(t)}
	m := LoadPermittedAddressesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadPermittedAddressesDataSource_Read_InvalidJSON exercises LoadPermittedAddressesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadPermittedAddressesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadPermittedAddressesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadPermittedAddressesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
