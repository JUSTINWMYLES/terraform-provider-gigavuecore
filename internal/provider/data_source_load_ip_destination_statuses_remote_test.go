package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadIpDestinationStatusesDataSource_Read_Happy exercises LoadIpDestinationStatusesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadIpDestinationStatusesDataSource_Read_Happy(t *testing.T) {
	r := &LoadIpDestinationStatusesDataSource{client: newMockClientStatus(t, 200, "{\"ipInterfaces\":[]}")}
	m := LoadIpDestinationStatusesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadIpDestinationStatusesDataSource_Read_NilClient exercises LoadIpDestinationStatusesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadIpDestinationStatusesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadIpDestinationStatusesDataSource{}
	m := LoadIpDestinationStatusesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadIpDestinationStatusesDataSource_Read_BuildError exercises LoadIpDestinationStatusesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadIpDestinationStatusesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadIpDestinationStatusesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadIpDestinationStatusesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadIpDestinationStatusesDataSource_Read_SendError exercises LoadIpDestinationStatusesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadIpDestinationStatusesDataSource_Read_SendError(t *testing.T) {
	r := &LoadIpDestinationStatusesDataSource{client: newTransportErrorClient(t)}
	m := LoadIpDestinationStatusesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadIpDestinationStatusesDataSource_Read_InvalidJSON exercises LoadIpDestinationStatusesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadIpDestinationStatusesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadIpDestinationStatusesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadIpDestinationStatusesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
