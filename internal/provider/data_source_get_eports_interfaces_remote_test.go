package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEportsInterfacesDataSource_Read_Happy exercises GetEportsInterfacesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetEportsInterfacesDataSource_Read_Happy(t *testing.T) {
	r := &GetEportsInterfacesDataSource{client: newMockClientStatus(t, 200, "{\"gsEngineInterfaces\":[]}")}
	m := GetEportsInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEportsInterfacesDataSource_Read_NilClient exercises GetEportsInterfacesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEportsInterfacesDataSource_Read_NilClient(t *testing.T) {
	r := &GetEportsInterfacesDataSource{}
	m := GetEportsInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEportsInterfacesDataSource_Read_BuildError exercises GetEportsInterfacesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetEportsInterfacesDataSource_Read_BuildError(t *testing.T) {
	r := &GetEportsInterfacesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEportsInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetEportsInterfacesDataSource_Read_SendError exercises GetEportsInterfacesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetEportsInterfacesDataSource_Read_SendError(t *testing.T) {
	r := &GetEportsInterfacesDataSource{client: newTransportErrorClient(t)}
	m := GetEportsInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetEportsInterfacesDataSource_Read_InvalidJSON exercises GetEportsInterfacesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetEportsInterfacesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEportsInterfacesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEportsInterfacesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
