package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFmBackupArchiveServersDataSource_Read_Happy exercises LoadFmBackupArchiveServersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadFmBackupArchiveServersDataSource_Read_Happy(t *testing.T) {
	r := &LoadFmBackupArchiveServersDataSource{client: newMockClientStatus(t, 200, "{\"fmBackupArchiveServers\":[]}")}
	m := LoadFmBackupArchiveServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadFmBackupArchiveServersDataSource_Read_NilClient exercises LoadFmBackupArchiveServersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadFmBackupArchiveServersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadFmBackupArchiveServersDataSource{}
	m := LoadFmBackupArchiveServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadFmBackupArchiveServersDataSource_Read_BuildError exercises LoadFmBackupArchiveServersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadFmBackupArchiveServersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadFmBackupArchiveServersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadFmBackupArchiveServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadFmBackupArchiveServersDataSource_Read_SendError exercises LoadFmBackupArchiveServersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadFmBackupArchiveServersDataSource_Read_SendError(t *testing.T) {
	r := &LoadFmBackupArchiveServersDataSource{client: newTransportErrorClient(t)}
	m := LoadFmBackupArchiveServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadFmBackupArchiveServersDataSource_Read_InvalidJSON exercises LoadFmBackupArchiveServersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadFmBackupArchiveServersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadFmBackupArchiveServersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadFmBackupArchiveServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
