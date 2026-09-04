package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowDiameterReportsS6ASummaryDataSource_Read_Happy exercises GetFlowDiameterReportsS6ASummaryDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetFlowDiameterReportsS6ASummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetFlowDiameterReportsS6ASummaryDataSource{client: newMockClientStatus(t, 200, "{\"flowDiameterS6aReportsSummary\":[]}")}
	m := GetFlowDiameterReportsS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlowDiameterReportsS6ASummaryDataSource_Read_NilClient exercises GetFlowDiameterReportsS6ASummaryDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlowDiameterReportsS6ASummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlowDiameterReportsS6ASummaryDataSource{}
	m := GetFlowDiameterReportsS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlowDiameterReportsS6ASummaryDataSource_Read_BuildError exercises GetFlowDiameterReportsS6ASummaryDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetFlowDiameterReportsS6ASummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlowDiameterReportsS6ASummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlowDiameterReportsS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetFlowDiameterReportsS6ASummaryDataSource_Read_SendError exercises GetFlowDiameterReportsS6ASummaryDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetFlowDiameterReportsS6ASummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetFlowDiameterReportsS6ASummaryDataSource{client: newTransportErrorClient(t)}
	m := GetFlowDiameterReportsS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetFlowDiameterReportsS6ASummaryDataSource_Read_InvalidJSON exercises GetFlowDiameterReportsS6ASummaryDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetFlowDiameterReportsS6ASummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlowDiameterReportsS6ASummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlowDiameterReportsS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
