package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadIpv6NeighborsDataSource_Read_Happy exercises LoadIpv6NeighborsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadIpv6NeighborsDataSource_Read_Happy(t *testing.T) {
	r := &LoadIpv6NeighborsDataSource{client: newMockClientStatus(t, 200, "{\"arpEntries\":[]}")}
	m := LoadIpv6NeighborsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadIpv6NeighborsDataSource_Read_NilClient exercises LoadIpv6NeighborsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadIpv6NeighborsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadIpv6NeighborsDataSource{}
	m := LoadIpv6NeighborsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadIpv6NeighborsDataSource_Read_BuildError exercises LoadIpv6NeighborsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadIpv6NeighborsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadIpv6NeighborsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadIpv6NeighborsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadIpv6NeighborsDataSource_Read_SendError exercises LoadIpv6NeighborsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadIpv6NeighborsDataSource_Read_SendError(t *testing.T) {
	r := &LoadIpv6NeighborsDataSource{client: newTransportErrorClient(t)}
	m := LoadIpv6NeighborsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadIpv6NeighborsDataSource_Read_InvalidJSON exercises LoadIpv6NeighborsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadIpv6NeighborsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadIpv6NeighborsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadIpv6NeighborsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
