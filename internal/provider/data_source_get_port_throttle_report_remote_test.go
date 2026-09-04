package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPortThrottleReportDataSource_Read_Happy exercises GetPortThrottleReportDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetPortThrottleReportDataSource_Read_Happy(t *testing.T) {
	r := &GetPortThrottleReportDataSource{client: newMockClientStatus(t, 200, "{\"portThrottleReports\":[]}")}
	m := GetPortThrottleReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPortThrottleReportDataSource_Read_NilClient exercises GetPortThrottleReportDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPortThrottleReportDataSource_Read_NilClient(t *testing.T) {
	r := &GetPortThrottleReportDataSource{}
	m := GetPortThrottleReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPortThrottleReportDataSource_Read_BuildError exercises GetPortThrottleReportDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetPortThrottleReportDataSource_Read_BuildError(t *testing.T) {
	r := &GetPortThrottleReportDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPortThrottleReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetPortThrottleReportDataSource_Read_SendError exercises GetPortThrottleReportDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetPortThrottleReportDataSource_Read_SendError(t *testing.T) {
	r := &GetPortThrottleReportDataSource{client: newTransportErrorClient(t)}
	m := GetPortThrottleReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetPortThrottleReportDataSource_Read_InvalidJSON exercises GetPortThrottleReportDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetPortThrottleReportDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPortThrottleReportDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPortThrottleReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
