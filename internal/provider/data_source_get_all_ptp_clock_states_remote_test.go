package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllPtpClockStatesDataSource_Read_Happy exercises GetAllPtpClockStatesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllPtpClockStatesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllPtpClockStatesDataSource{client: newMockClientStatus(t, 200, "{\"clockStates\":[]}")}
	m := GetAllPtpClockStatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllPtpClockStatesDataSource_Read_NilClient exercises GetAllPtpClockStatesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllPtpClockStatesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllPtpClockStatesDataSource{}
	m := GetAllPtpClockStatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllPtpClockStatesDataSource_Read_BuildError exercises GetAllPtpClockStatesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllPtpClockStatesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllPtpClockStatesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllPtpClockStatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllPtpClockStatesDataSource_Read_SendError exercises GetAllPtpClockStatesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllPtpClockStatesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllPtpClockStatesDataSource{client: newTransportErrorClient(t)}
	m := GetAllPtpClockStatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllPtpClockStatesDataSource_Read_InvalidJSON exercises GetAllPtpClockStatesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllPtpClockStatesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllPtpClockStatesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllPtpClockStatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
