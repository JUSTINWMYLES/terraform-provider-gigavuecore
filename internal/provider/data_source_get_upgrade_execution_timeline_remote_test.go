package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeExecutionTimelineDataSource_Read_Happy exercises GetUpgradeExecutionTimelineDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetUpgradeExecutionTimelineDataSource_Read_Happy(t *testing.T) {
	r := &GetUpgradeExecutionTimelineDataSource{client: newMockClientStatus(t, 200, "{\"clustersStatus\":[]}")}
	m := GetUpgradeExecutionTimelineDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetUpgradeExecutionTimelineDataSource_Read_NilClient exercises GetUpgradeExecutionTimelineDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetUpgradeExecutionTimelineDataSource_Read_NilClient(t *testing.T) {
	r := &GetUpgradeExecutionTimelineDataSource{}
	m := GetUpgradeExecutionTimelineDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetUpgradeExecutionTimelineDataSource_Read_BuildError exercises GetUpgradeExecutionTimelineDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetUpgradeExecutionTimelineDataSource_Read_BuildError(t *testing.T) {
	r := &GetUpgradeExecutionTimelineDataSource{client: newMalformedBaseURLClient(t)}
	m := GetUpgradeExecutionTimelineDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetUpgradeExecutionTimelineDataSource_Read_SendError exercises GetUpgradeExecutionTimelineDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetUpgradeExecutionTimelineDataSource_Read_SendError(t *testing.T) {
	r := &GetUpgradeExecutionTimelineDataSource{client: newTransportErrorClient(t)}
	m := GetUpgradeExecutionTimelineDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetUpgradeExecutionTimelineDataSource_Read_InvalidJSON exercises GetUpgradeExecutionTimelineDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetUpgradeExecutionTimelineDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetUpgradeExecutionTimelineDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetUpgradeExecutionTimelineDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
