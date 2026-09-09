package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetGtpBackupFilesDataSource_Read_Happy exercises GetGtpBackupFilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetGtpBackupFilesDataSource_Read_Happy(t *testing.T) {
	r := &GetGtpBackupFilesDataSource{client: newMockClientStatus(t, 200, "{\"gtpBackupFiles\":[]}")}
	m := GetGtpBackupFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetGtpBackupFilesDataSource_Read_NilClient exercises GetGtpBackupFilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetGtpBackupFilesDataSource_Read_NilClient(t *testing.T) {
	r := &GetGtpBackupFilesDataSource{}
	m := GetGtpBackupFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetGtpBackupFilesDataSource_Read_BuildError exercises GetGtpBackupFilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetGtpBackupFilesDataSource_Read_BuildError(t *testing.T) {
	r := &GetGtpBackupFilesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetGtpBackupFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetGtpBackupFilesDataSource_Read_SendError exercises GetGtpBackupFilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetGtpBackupFilesDataSource_Read_SendError(t *testing.T) {
	r := &GetGtpBackupFilesDataSource{client: newTransportErrorClient(t)}
	m := GetGtpBackupFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetGtpBackupFilesDataSource_Read_InvalidJSON exercises GetGtpBackupFilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetGtpBackupFilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetGtpBackupFilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetGtpBackupFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
