package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemArpEntriesDataSource_Read_Happy exercises GetSystemArpEntriesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetSystemArpEntriesDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemArpEntriesDataSource{client: newMockClientStatus(t, 200, "{\"arpEntries\":[]}")}
	m := GetSystemArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemArpEntriesDataSource_Read_NilClient exercises GetSystemArpEntriesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemArpEntriesDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemArpEntriesDataSource{}
	m := GetSystemArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemArpEntriesDataSource_Read_BuildError exercises GetSystemArpEntriesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetSystemArpEntriesDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemArpEntriesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSystemArpEntriesDataSource_Read_SendError exercises GetSystemArpEntriesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetSystemArpEntriesDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemArpEntriesDataSource{client: newTransportErrorClient(t)}
	m := GetSystemArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSystemArpEntriesDataSource_Read_InvalidJSON exercises GetSystemArpEntriesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetSystemArpEntriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemArpEntriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
