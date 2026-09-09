package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListFmStoredClustersConfigBackupSnapshotsDataSource_Read_Happy exercises ListFmStoredClustersConfigBackupSnapshotsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestListFmStoredClustersConfigBackupSnapshotsDataSource_Read_Happy(t *testing.T) {
	r := &ListFmStoredClustersConfigBackupSnapshotsDataSource{client: newMockClientStatus(t, 200, "{\"clustersConfigBackups\":[]}")}
	m := ListFmStoredClustersConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListFmStoredClustersConfigBackupSnapshotsDataSource_Read_NilClient exercises ListFmStoredClustersConfigBackupSnapshotsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListFmStoredClustersConfigBackupSnapshotsDataSource_Read_NilClient(t *testing.T) {
	r := &ListFmStoredClustersConfigBackupSnapshotsDataSource{}
	m := ListFmStoredClustersConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListFmStoredClustersConfigBackupSnapshotsDataSource_Read_BuildError exercises ListFmStoredClustersConfigBackupSnapshotsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestListFmStoredClustersConfigBackupSnapshotsDataSource_Read_BuildError(t *testing.T) {
	r := &ListFmStoredClustersConfigBackupSnapshotsDataSource{client: newMalformedBaseURLClient(t)}
	m := ListFmStoredClustersConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestListFmStoredClustersConfigBackupSnapshotsDataSource_Read_SendError exercises ListFmStoredClustersConfigBackupSnapshotsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestListFmStoredClustersConfigBackupSnapshotsDataSource_Read_SendError(t *testing.T) {
	r := &ListFmStoredClustersConfigBackupSnapshotsDataSource{client: newTransportErrorClient(t)}
	m := ListFmStoredClustersConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestListFmStoredClustersConfigBackupSnapshotsDataSource_Read_InvalidJSON exercises ListFmStoredClustersConfigBackupSnapshotsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestListFmStoredClustersConfigBackupSnapshotsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ListFmStoredClustersConfigBackupSnapshotsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ListFmStoredClustersConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
