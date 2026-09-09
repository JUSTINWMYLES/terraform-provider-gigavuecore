package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_Happy exercises ListFmStoredClusterConfigBackupSnapshotDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_Happy(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := ListFmStoredClusterConfigBackupSnapshotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_NilClient exercises ListFmStoredClusterConfigBackupSnapshotDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_NilClient(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotDataSource{}
	m := ListFmStoredClusterConfigBackupSnapshotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_BuildError exercises ListFmStoredClusterConfigBackupSnapshotDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_BuildError(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotDataSource{client: newMalformedBaseURLClient(t)}
	m := ListFmStoredClusterConfigBackupSnapshotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_SendError exercises ListFmStoredClusterConfigBackupSnapshotDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_SendError(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotDataSource{client: newTransportErrorClient(t)}
	m := ListFmStoredClusterConfigBackupSnapshotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_NotFound exercises ListFmStoredClusterConfigBackupSnapshotDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_NotFound(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotDataSource{client: newMockClientStatus(t, 404, "")}
	m := ListFmStoredClusterConfigBackupSnapshotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_APIError exercises ListFmStoredClusterConfigBackupSnapshotDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_APIError(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ListFmStoredClusterConfigBackupSnapshotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot")
}

// TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_APIErrorReadBody exercises ListFmStoredClusterConfigBackupSnapshotDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := ListFmStoredClusterConfigBackupSnapshotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_InvalidJSON exercises ListFmStoredClusterConfigBackupSnapshotDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestListFmStoredClusterConfigBackupSnapshotDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupSnapshotDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ListFmStoredClusterConfigBackupSnapshotDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
