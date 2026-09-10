package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadInlineSslCertValidRecordDataSource_Read_Happy exercises LoadInlineSslCertValidRecordDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadInlineSslCertValidRecordDataSource_Read_Happy(t *testing.T) {
	r := &LoadInlineSslCertValidRecordDataSource{client: newMockClientStatus(t, 200, "{\"records\":[]}")}
	m := LoadInlineSslCertValidRecordDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadInlineSslCertValidRecordDataSource_Read_NilClient exercises LoadInlineSslCertValidRecordDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadInlineSslCertValidRecordDataSource_Read_NilClient(t *testing.T) {
	r := &LoadInlineSslCertValidRecordDataSource{}
	m := LoadInlineSslCertValidRecordDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadInlineSslCertValidRecordDataSource_Read_BuildError exercises LoadInlineSslCertValidRecordDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadInlineSslCertValidRecordDataSource_Read_BuildError(t *testing.T) {
	r := &LoadInlineSslCertValidRecordDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadInlineSslCertValidRecordDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadInlineSslCertValidRecordDataSource_Read_SendError exercises LoadInlineSslCertValidRecordDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadInlineSslCertValidRecordDataSource_Read_SendError(t *testing.T) {
	r := &LoadInlineSslCertValidRecordDataSource{client: newTransportErrorClient(t)}
	m := LoadInlineSslCertValidRecordDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadInlineSslCertValidRecordDataSource_Read_InvalidJSON exercises LoadInlineSslCertValidRecordDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadInlineSslCertValidRecordDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadInlineSslCertValidRecordDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadInlineSslCertValidRecordDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
