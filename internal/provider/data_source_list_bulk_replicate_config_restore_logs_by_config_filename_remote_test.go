package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource_Read_Happy exercises ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource_Read_Happy(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource{client: newMockClientStatus(t, 200, "{\"bulkReplicateConfigRestoreLogs\":[]}")}
	m := ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource_Read_NilClient exercises ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource_Read_NilClient(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource{}
	m := ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource_Read_BuildError exercises ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource_Read_BuildError(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource{client: newMalformedBaseURLClient(t)}
	m := ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource_Read_SendError exercises ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource_Read_SendError(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource{client: newTransportErrorClient(t)}
	m := ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource_Read_InvalidJSON exercises ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
