package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllGpfcpProfileDataSource_Read_Happy exercises GetAllGpfcpProfileDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllGpfcpProfileDataSource_Read_Happy(t *testing.T) {
	r := &GetAllGpfcpProfileDataSource{client: newMockClientStatus(t, 200, "{\"gpfcpProfiles\":[]}")}
	m := GetAllGpfcpProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllGpfcpProfileDataSource_Read_NilClient exercises GetAllGpfcpProfileDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllGpfcpProfileDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllGpfcpProfileDataSource{}
	m := GetAllGpfcpProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllGpfcpProfileDataSource_Read_BuildError exercises GetAllGpfcpProfileDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllGpfcpProfileDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllGpfcpProfileDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllGpfcpProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllGpfcpProfileDataSource_Read_SendError exercises GetAllGpfcpProfileDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllGpfcpProfileDataSource_Read_SendError(t *testing.T) {
	r := &GetAllGpfcpProfileDataSource{client: newTransportErrorClient(t)}
	m := GetAllGpfcpProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllGpfcpProfileDataSource_Read_InvalidJSON exercises GetAllGpfcpProfileDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllGpfcpProfileDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllGpfcpProfileDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllGpfcpProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
