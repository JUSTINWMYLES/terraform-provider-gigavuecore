package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPolicyNrtStatsDataSource_Read_Happy exercises GetPolicyNrtStatsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetPolicyNrtStatsDataSource_Read_Happy(t *testing.T) {
	r := &GetPolicyNrtStatsDataSource{client: newMockClientStatus(t, 200, "{\"fmNearRealTimeCachedStatisticsList\":[]}")}
	m := GetPolicyNrtStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPolicyNrtStatsDataSource_Read_NilClient exercises GetPolicyNrtStatsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPolicyNrtStatsDataSource_Read_NilClient(t *testing.T) {
	r := &GetPolicyNrtStatsDataSource{}
	m := GetPolicyNrtStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPolicyNrtStatsDataSource_Read_BuildError exercises GetPolicyNrtStatsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetPolicyNrtStatsDataSource_Read_BuildError(t *testing.T) {
	r := &GetPolicyNrtStatsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPolicyNrtStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetPolicyNrtStatsDataSource_Read_SendError exercises GetPolicyNrtStatsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetPolicyNrtStatsDataSource_Read_SendError(t *testing.T) {
	r := &GetPolicyNrtStatsDataSource{client: newTransportErrorClient(t)}
	m := GetPolicyNrtStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetPolicyNrtStatsDataSource_Read_InvalidJSON exercises GetPolicyNrtStatsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetPolicyNrtStatsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPolicyNrtStatsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPolicyNrtStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
