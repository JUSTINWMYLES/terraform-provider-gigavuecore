package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllPtpPortsCountersDataSource_Read_Happy exercises GetAllPtpPortsCountersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllPtpPortsCountersDataSource_Read_Happy(t *testing.T) {
	r := &GetAllPtpPortsCountersDataSource{client: newMockClientStatus(t, 200, "{\"ports\":[]}")}
	m := GetAllPtpPortsCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllPtpPortsCountersDataSource_Read_NilClient exercises GetAllPtpPortsCountersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllPtpPortsCountersDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllPtpPortsCountersDataSource{}
	m := GetAllPtpPortsCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllPtpPortsCountersDataSource_Read_BuildError exercises GetAllPtpPortsCountersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllPtpPortsCountersDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllPtpPortsCountersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllPtpPortsCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllPtpPortsCountersDataSource_Read_SendError exercises GetAllPtpPortsCountersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllPtpPortsCountersDataSource_Read_SendError(t *testing.T) {
	r := &GetAllPtpPortsCountersDataSource{client: newTransportErrorClient(t)}
	m := GetAllPtpPortsCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllPtpPortsCountersDataSource_Read_InvalidJSON exercises GetAllPtpPortsCountersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllPtpPortsCountersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllPtpPortsCountersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllPtpPortsCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
