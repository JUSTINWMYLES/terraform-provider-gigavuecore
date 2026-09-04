package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllPtpPortsDataSource_Read_Happy exercises GetAllPtpPortsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllPtpPortsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllPtpPortsDataSource{client: newMockClientStatus(t, 200, "{\"portStates\":[]}")}
	m := GetAllPtpPortsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllPtpPortsDataSource_Read_NilClient exercises GetAllPtpPortsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllPtpPortsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllPtpPortsDataSource{}
	m := GetAllPtpPortsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllPtpPortsDataSource_Read_BuildError exercises GetAllPtpPortsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllPtpPortsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllPtpPortsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllPtpPortsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllPtpPortsDataSource_Read_SendError exercises GetAllPtpPortsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllPtpPortsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllPtpPortsDataSource{client: newTransportErrorClient(t)}
	m := GetAllPtpPortsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllPtpPortsDataSource_Read_InvalidJSON exercises GetAllPtpPortsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllPtpPortsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllPtpPortsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllPtpPortsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
