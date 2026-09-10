package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGdpNodeReportsDataSource_Read_Happy exercises LoadGdpNodeReportsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadGdpNodeReportsDataSource_Read_Happy(t *testing.T) {
	r := &LoadGdpNodeReportsDataSource{client: newMockClientStatus(t, 200, "{\"gdpNodeReports\":[]}")}
	m := LoadGdpNodeReportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadGdpNodeReportsDataSource_Read_NilClient exercises LoadGdpNodeReportsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadGdpNodeReportsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadGdpNodeReportsDataSource{}
	m := LoadGdpNodeReportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadGdpNodeReportsDataSource_Read_BuildError exercises LoadGdpNodeReportsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadGdpNodeReportsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadGdpNodeReportsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadGdpNodeReportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadGdpNodeReportsDataSource_Read_SendError exercises LoadGdpNodeReportsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadGdpNodeReportsDataSource_Read_SendError(t *testing.T) {
	r := &LoadGdpNodeReportsDataSource{client: newTransportErrorClient(t)}
	m := LoadGdpNodeReportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadGdpNodeReportsDataSource_Read_InvalidJSON exercises LoadGdpNodeReportsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadGdpNodeReportsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadGdpNodeReportsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadGdpNodeReportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
