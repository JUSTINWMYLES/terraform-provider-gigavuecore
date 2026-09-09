package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllPlatformFilterLimitsDataSource_Read_Happy exercises LoadAllPlatformFilterLimitsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllPlatformFilterLimitsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllPlatformFilterLimitsDataSource{client: newMockClientStatus(t, 200, "{\"platformsFilterLimits\":[]}")}
	m := LoadAllPlatformFilterLimitsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllPlatformFilterLimitsDataSource_Read_NilClient exercises LoadAllPlatformFilterLimitsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllPlatformFilterLimitsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllPlatformFilterLimitsDataSource{}
	m := LoadAllPlatformFilterLimitsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllPlatformFilterLimitsDataSource_Read_BuildError exercises LoadAllPlatformFilterLimitsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllPlatformFilterLimitsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllPlatformFilterLimitsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllPlatformFilterLimitsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllPlatformFilterLimitsDataSource_Read_SendError exercises LoadAllPlatformFilterLimitsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllPlatformFilterLimitsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllPlatformFilterLimitsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllPlatformFilterLimitsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllPlatformFilterLimitsDataSource_Read_InvalidJSON exercises LoadAllPlatformFilterLimitsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllPlatformFilterLimitsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllPlatformFilterLimitsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllPlatformFilterLimitsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
