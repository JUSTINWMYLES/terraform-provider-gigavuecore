package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_Happy exercises ListFmStoredClusterConfigBackupSnapshotsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_Happy(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := ListFmStoredClusterConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_NilClient exercises ListFmStoredClusterConfigBackupSnapshotsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_NilClient(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotsDataSource{}
	m := ListFmStoredClusterConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_BuildError exercises ListFmStoredClusterConfigBackupSnapshotsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_BuildError(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotsDataSource{client: newMalformedBaseURLClient(t)}
	m := ListFmStoredClusterConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_SendError exercises ListFmStoredClusterConfigBackupSnapshotsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_SendError(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotsDataSource{client: newTransportErrorClient(t)}
	m := ListFmStoredClusterConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_NotFound exercises ListFmStoredClusterConfigBackupSnapshotsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_NotFound(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotsDataSource{client: newMockClientStatus(t, 404, "")}
	m := ListFmStoredClusterConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_APIError exercises ListFmStoredClusterConfigBackupSnapshotsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_APIError(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ListFmStoredClusterConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshots")
}

// TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_APIErrorReadBody exercises ListFmStoredClusterConfigBackupSnapshotsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := ListFmStoredClusterConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_InvalidJSON exercises ListFmStoredClusterConfigBackupSnapshotsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestListFmStoredClusterConfigBackupSnapshotsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ListFmStoredClusterConfigBackupSnapshotsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
