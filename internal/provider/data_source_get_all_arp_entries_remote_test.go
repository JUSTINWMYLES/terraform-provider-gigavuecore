package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllArpEntriesDataSource_Read_Happy exercises GetAllArpEntriesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllArpEntriesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllArpEntriesDataSource{client: newMockClientStatus(t, 200, "{\"arpResponses\":[]}")}
	m := GetAllArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllArpEntriesDataSource_Read_NilClient exercises GetAllArpEntriesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllArpEntriesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllArpEntriesDataSource{}
	m := GetAllArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllArpEntriesDataSource_Read_BuildError exercises GetAllArpEntriesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllArpEntriesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllArpEntriesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllArpEntriesDataSource_Read_SendError exercises GetAllArpEntriesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllArpEntriesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllArpEntriesDataSource{client: newTransportErrorClient(t)}
	m := GetAllArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllArpEntriesDataSource_Read_InvalidJSON exercises GetAllArpEntriesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllArpEntriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllArpEntriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
