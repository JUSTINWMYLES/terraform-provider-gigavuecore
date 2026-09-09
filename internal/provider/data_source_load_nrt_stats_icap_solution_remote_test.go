package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadNrtStatsIcapSolutionDataSource_Read_Happy exercises LoadNrtStatsIcapSolutionDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadNrtStatsIcapSolutionDataSource_Read_Happy(t *testing.T) {
	r := &LoadNrtStatsIcapSolutionDataSource{client: newMockClientStatus(t, 200, "{\"fmNearRealTimeCachedStatisticsList\":[]}")}
	m := LoadNrtStatsIcapSolutionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadNrtStatsIcapSolutionDataSource_Read_NilClient exercises LoadNrtStatsIcapSolutionDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadNrtStatsIcapSolutionDataSource_Read_NilClient(t *testing.T) {
	r := &LoadNrtStatsIcapSolutionDataSource{}
	m := LoadNrtStatsIcapSolutionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadNrtStatsIcapSolutionDataSource_Read_BuildError exercises LoadNrtStatsIcapSolutionDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadNrtStatsIcapSolutionDataSource_Read_BuildError(t *testing.T) {
	r := &LoadNrtStatsIcapSolutionDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadNrtStatsIcapSolutionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadNrtStatsIcapSolutionDataSource_Read_SendError exercises LoadNrtStatsIcapSolutionDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadNrtStatsIcapSolutionDataSource_Read_SendError(t *testing.T) {
	r := &LoadNrtStatsIcapSolutionDataSource{client: newTransportErrorClient(t)}
	m := LoadNrtStatsIcapSolutionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadNrtStatsIcapSolutionDataSource_Read_InvalidJSON exercises LoadNrtStatsIcapSolutionDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadNrtStatsIcapSolutionDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadNrtStatsIcapSolutionDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadNrtStatsIcapSolutionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
