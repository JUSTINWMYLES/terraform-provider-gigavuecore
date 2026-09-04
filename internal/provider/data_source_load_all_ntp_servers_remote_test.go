package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllNtpServersDataSource_Read_Happy exercises LoadAllNtpServersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllNtpServersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllNtpServersDataSource{client: newMockClientStatus(t, 200, "{\"ntpServers\":[]}")}
	m := LoadAllNtpServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllNtpServersDataSource_Read_NilClient exercises LoadAllNtpServersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllNtpServersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllNtpServersDataSource{}
	m := LoadAllNtpServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllNtpServersDataSource_Read_BuildError exercises LoadAllNtpServersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllNtpServersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllNtpServersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllNtpServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNtpServersDataSource_Read_SendError exercises LoadAllNtpServersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllNtpServersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllNtpServersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllNtpServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNtpServersDataSource_Read_InvalidJSON exercises LoadAllNtpServersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllNtpServersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllNtpServersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllNtpServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
