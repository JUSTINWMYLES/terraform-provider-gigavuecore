package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAvailablePcapFilenamesDataSource_Read_Happy exercises LoadAvailablePcapFilenamesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAvailablePcapFilenamesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAvailablePcapFilenamesDataSource{client: newMockClientStatus(t, 200, "{\"pcapFiles\":[]}")}
	m := LoadAvailablePcapFilenamesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAvailablePcapFilenamesDataSource_Read_NilClient exercises LoadAvailablePcapFilenamesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAvailablePcapFilenamesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAvailablePcapFilenamesDataSource{}
	m := LoadAvailablePcapFilenamesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAvailablePcapFilenamesDataSource_Read_BuildError exercises LoadAvailablePcapFilenamesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAvailablePcapFilenamesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAvailablePcapFilenamesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAvailablePcapFilenamesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAvailablePcapFilenamesDataSource_Read_SendError exercises LoadAvailablePcapFilenamesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAvailablePcapFilenamesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAvailablePcapFilenamesDataSource{client: newTransportErrorClient(t)}
	m := LoadAvailablePcapFilenamesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAvailablePcapFilenamesDataSource_Read_InvalidJSON exercises LoadAvailablePcapFilenamesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAvailablePcapFilenamesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAvailablePcapFilenamesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAvailablePcapFilenamesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
