package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource_Read_Happy exercises ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource_Read_Happy(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource{client: newMockClientStatus(t, 200, "{\"configFiles\":[]}")}
	m := ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource_Read_NilClient exercises ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource_Read_NilClient(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource{}
	m := ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource_Read_BuildError exercises ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource_Read_BuildError(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource{client: newMalformedBaseURLClient(t)}
	m := ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource_Read_SendError exercises ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource_Read_SendError(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource{client: newTransportErrorClient(t)}
	m := ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource_Read_InvalidJSON exercises ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
