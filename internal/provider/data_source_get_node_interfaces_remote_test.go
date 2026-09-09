package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeInterfacesDataSource_Read_Happy exercises GetNodeInterfacesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetNodeInterfacesDataSource_Read_Happy(t *testing.T) {
	r := &GetNodeInterfacesDataSource{client: newMockClientStatus(t, 200, "{\"env\":[]}")}
	m := GetNodeInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetNodeInterfacesDataSource_Read_NilClient exercises GetNodeInterfacesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetNodeInterfacesDataSource_Read_NilClient(t *testing.T) {
	r := &GetNodeInterfacesDataSource{}
	m := GetNodeInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetNodeInterfacesDataSource_Read_BuildError exercises GetNodeInterfacesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetNodeInterfacesDataSource_Read_BuildError(t *testing.T) {
	r := &GetNodeInterfacesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetNodeInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNodeInterfacesDataSource_Read_SendError exercises GetNodeInterfacesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetNodeInterfacesDataSource_Read_SendError(t *testing.T) {
	r := &GetNodeInterfacesDataSource{client: newTransportErrorClient(t)}
	m := GetNodeInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNodeInterfacesDataSource_Read_InvalidJSON exercises GetNodeInterfacesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetNodeInterfacesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetNodeInterfacesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetNodeInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
