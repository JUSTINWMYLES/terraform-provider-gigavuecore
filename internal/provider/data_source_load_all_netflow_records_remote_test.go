package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllNetflowRecordsDataSource_Read_Happy exercises LoadAllNetflowRecordsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllNetflowRecordsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllNetflowRecordsDataSource{client: newMockClientStatus(t, 200, "{\"nfRecords\":[]}")}
	m := LoadAllNetflowRecordsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllNetflowRecordsDataSource_Read_NilClient exercises LoadAllNetflowRecordsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllNetflowRecordsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllNetflowRecordsDataSource{}
	m := LoadAllNetflowRecordsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllNetflowRecordsDataSource_Read_BuildError exercises LoadAllNetflowRecordsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllNetflowRecordsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllNetflowRecordsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllNetflowRecordsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNetflowRecordsDataSource_Read_SendError exercises LoadAllNetflowRecordsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllNetflowRecordsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllNetflowRecordsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllNetflowRecordsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNetflowRecordsDataSource_Read_InvalidJSON exercises LoadAllNetflowRecordsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllNetflowRecordsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllNetflowRecordsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllNetflowRecordsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
