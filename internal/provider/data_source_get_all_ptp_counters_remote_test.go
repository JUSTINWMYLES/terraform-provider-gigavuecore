package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllPtpCountersDataSource_Read_Happy exercises GetAllPtpCountersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllPtpCountersDataSource_Read_Happy(t *testing.T) {
	r := &GetAllPtpCountersDataSource{client: newMockClientStatus(t, 200, "{\"counters\":[]}")}
	m := GetAllPtpCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllPtpCountersDataSource_Read_NilClient exercises GetAllPtpCountersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllPtpCountersDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllPtpCountersDataSource{}
	m := GetAllPtpCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllPtpCountersDataSource_Read_BuildError exercises GetAllPtpCountersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllPtpCountersDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllPtpCountersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllPtpCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllPtpCountersDataSource_Read_SendError exercises GetAllPtpCountersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllPtpCountersDataSource_Read_SendError(t *testing.T) {
	r := &GetAllPtpCountersDataSource{client: newTransportErrorClient(t)}
	m := GetAllPtpCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllPtpCountersDataSource_Read_InvalidJSON exercises GetAllPtpCountersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllPtpCountersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllPtpCountersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllPtpCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
