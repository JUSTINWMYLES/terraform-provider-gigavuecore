package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetBundlesInPeriodDataSource_Read_Happy exercises GetBundlesInPeriodDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetBundlesInPeriodDataSource_Read_Happy(t *testing.T) {
	r := &GetBundlesInPeriodDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetBundlesInPeriodDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetBundlesInPeriodDataSource_Read_NilClient exercises GetBundlesInPeriodDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetBundlesInPeriodDataSource_Read_NilClient(t *testing.T) {
	r := &GetBundlesInPeriodDataSource{}
	m := GetBundlesInPeriodDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetBundlesInPeriodDataSource_Read_BuildError exercises GetBundlesInPeriodDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetBundlesInPeriodDataSource_Read_BuildError(t *testing.T) {
	r := &GetBundlesInPeriodDataSource{client: newMalformedBaseURLClient(t)}
	m := GetBundlesInPeriodDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetBundlesInPeriodDataSource_Read_SendError exercises GetBundlesInPeriodDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetBundlesInPeriodDataSource_Read_SendError(t *testing.T) {
	r := &GetBundlesInPeriodDataSource{client: newTransportErrorClient(t)}
	m := GetBundlesInPeriodDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetBundlesInPeriodDataSource_Read_InvalidJSON exercises GetBundlesInPeriodDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetBundlesInPeriodDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetBundlesInPeriodDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetBundlesInPeriodDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
