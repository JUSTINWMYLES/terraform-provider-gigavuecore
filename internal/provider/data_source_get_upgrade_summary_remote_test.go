package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeSummaryDataSource_Read_Happy exercises GetUpgradeSummaryDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetUpgradeSummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetUpgradeSummaryDataSource{client: newMockClientStatus(t, 200, "{\"deviceUpgradeSummary\":[]}")}
	m := GetUpgradeSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetUpgradeSummaryDataSource_Read_NilClient exercises GetUpgradeSummaryDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetUpgradeSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetUpgradeSummaryDataSource{}
	m := GetUpgradeSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetUpgradeSummaryDataSource_Read_BuildError exercises GetUpgradeSummaryDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetUpgradeSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetUpgradeSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetUpgradeSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetUpgradeSummaryDataSource_Read_SendError exercises GetUpgradeSummaryDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetUpgradeSummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetUpgradeSummaryDataSource{client: newTransportErrorClient(t)}
	m := GetUpgradeSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetUpgradeSummaryDataSource_Read_InvalidJSON exercises GetUpgradeSummaryDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetUpgradeSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetUpgradeSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetUpgradeSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
