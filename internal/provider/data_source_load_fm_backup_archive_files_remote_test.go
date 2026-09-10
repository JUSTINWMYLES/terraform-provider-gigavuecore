package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFmBackupArchiveFilesDataSource_Read_Happy exercises LoadFmBackupArchiveFilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadFmBackupArchiveFilesDataSource_Read_Happy(t *testing.T) {
	r := &LoadFmBackupArchiveFilesDataSource{client: newMockClientStatus(t, 200, "{\"fmBackupArchiveFiles\":[]}")}
	m := LoadFmBackupArchiveFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadFmBackupArchiveFilesDataSource_Read_NilClient exercises LoadFmBackupArchiveFilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadFmBackupArchiveFilesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadFmBackupArchiveFilesDataSource{}
	m := LoadFmBackupArchiveFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadFmBackupArchiveFilesDataSource_Read_BuildError exercises LoadFmBackupArchiveFilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadFmBackupArchiveFilesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadFmBackupArchiveFilesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadFmBackupArchiveFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadFmBackupArchiveFilesDataSource_Read_SendError exercises LoadFmBackupArchiveFilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadFmBackupArchiveFilesDataSource_Read_SendError(t *testing.T) {
	r := &LoadFmBackupArchiveFilesDataSource{client: newTransportErrorClient(t)}
	m := LoadFmBackupArchiveFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadFmBackupArchiveFilesDataSource_Read_InvalidJSON exercises LoadFmBackupArchiveFilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadFmBackupArchiveFilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadFmBackupArchiveFilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadFmBackupArchiveFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
