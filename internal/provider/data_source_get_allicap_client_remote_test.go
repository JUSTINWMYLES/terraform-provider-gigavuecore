package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllicapClientDataSource_Read_Happy exercises GetAllicapClientDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllicapClientDataSource_Read_Happy(t *testing.T) {
	r := &GetAllicapClientDataSource{client: newMockClientStatus(t, 200, "{\"gigaIcapClients\":[]}")}
	m := GetAllicapClientDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllicapClientDataSource_Read_NilClient exercises GetAllicapClientDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllicapClientDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllicapClientDataSource{}
	m := GetAllicapClientDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllicapClientDataSource_Read_BuildError exercises GetAllicapClientDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllicapClientDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllicapClientDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllicapClientDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllicapClientDataSource_Read_SendError exercises GetAllicapClientDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllicapClientDataSource_Read_SendError(t *testing.T) {
	r := &GetAllicapClientDataSource{client: newTransportErrorClient(t)}
	m := GetAllicapClientDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllicapClientDataSource_Read_InvalidJSON exercises GetAllicapClientDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllicapClientDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllicapClientDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllicapClientDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
