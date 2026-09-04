package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeJobsDataSource_Read_Happy exercises GetUpgradeJobsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetUpgradeJobsDataSource_Read_Happy(t *testing.T) {
	r := &GetUpgradeJobsDataSource{client: newMockClientStatus(t, 200, "{\"deviceUpgradeTaskInfoRecord\":[]}")}
	m := GetUpgradeJobsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetUpgradeJobsDataSource_Read_NilClient exercises GetUpgradeJobsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetUpgradeJobsDataSource_Read_NilClient(t *testing.T) {
	r := &GetUpgradeJobsDataSource{}
	m := GetUpgradeJobsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetUpgradeJobsDataSource_Read_BuildError exercises GetUpgradeJobsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetUpgradeJobsDataSource_Read_BuildError(t *testing.T) {
	r := &GetUpgradeJobsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetUpgradeJobsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetUpgradeJobsDataSource_Read_SendError exercises GetUpgradeJobsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetUpgradeJobsDataSource_Read_SendError(t *testing.T) {
	r := &GetUpgradeJobsDataSource{client: newTransportErrorClient(t)}
	m := GetUpgradeJobsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetUpgradeJobsDataSource_Read_InvalidJSON exercises GetUpgradeJobsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetUpgradeJobsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetUpgradeJobsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetUpgradeJobsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
