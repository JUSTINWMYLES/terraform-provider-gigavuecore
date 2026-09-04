package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetInterfacesDataSource_Read_Happy exercises GetInterfacesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetInterfacesDataSource_Read_Happy(t *testing.T) {
	r := &GetInterfacesDataSource{client: newMockClientStatus(t, 200, "{\"env\":[]}")}
	m := GetInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetInterfacesDataSource_Read_NilClient exercises GetInterfacesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetInterfacesDataSource_Read_NilClient(t *testing.T) {
	r := &GetInterfacesDataSource{}
	m := GetInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetInterfacesDataSource_Read_BuildError exercises GetInterfacesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetInterfacesDataSource_Read_BuildError(t *testing.T) {
	r := &GetInterfacesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetInterfacesDataSource_Read_SendError exercises GetInterfacesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetInterfacesDataSource_Read_SendError(t *testing.T) {
	r := &GetInterfacesDataSource{client: newTransportErrorClient(t)}
	m := GetInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetInterfacesDataSource_Read_InvalidJSON exercises GetInterfacesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetInterfacesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetInterfacesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
