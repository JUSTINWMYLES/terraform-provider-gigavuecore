package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadNrtStatsFabricMapDataSource_Read_Happy exercises LoadNrtStatsFabricMapDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadNrtStatsFabricMapDataSource_Read_Happy(t *testing.T) {
	r := &LoadNrtStatsFabricMapDataSource{client: newMockClientStatus(t, 200, "{\"fmNearRealTimeCachedStatisticsList\":[]}")}
	m := LoadNrtStatsFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadNrtStatsFabricMapDataSource_Read_NilClient exercises LoadNrtStatsFabricMapDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadNrtStatsFabricMapDataSource_Read_NilClient(t *testing.T) {
	r := &LoadNrtStatsFabricMapDataSource{}
	m := LoadNrtStatsFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadNrtStatsFabricMapDataSource_Read_BuildError exercises LoadNrtStatsFabricMapDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadNrtStatsFabricMapDataSource_Read_BuildError(t *testing.T) {
	r := &LoadNrtStatsFabricMapDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadNrtStatsFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadNrtStatsFabricMapDataSource_Read_SendError exercises LoadNrtStatsFabricMapDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadNrtStatsFabricMapDataSource_Read_SendError(t *testing.T) {
	r := &LoadNrtStatsFabricMapDataSource{client: newTransportErrorClient(t)}
	m := LoadNrtStatsFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadNrtStatsFabricMapDataSource_Read_InvalidJSON exercises LoadNrtStatsFabricMapDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadNrtStatsFabricMapDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadNrtStatsFabricMapDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadNrtStatsFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
