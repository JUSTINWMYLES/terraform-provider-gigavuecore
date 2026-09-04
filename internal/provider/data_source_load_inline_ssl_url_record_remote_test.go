package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadInlineSslUrlRecordDataSource_Read_Happy exercises LoadInlineSslUrlRecordDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadInlineSslUrlRecordDataSource_Read_Happy(t *testing.T) {
	r := &LoadInlineSslUrlRecordDataSource{client: newMockClientStatus(t, 200, "{\"records\":[]}")}
	m := LoadInlineSslUrlRecordDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadInlineSslUrlRecordDataSource_Read_NilClient exercises LoadInlineSslUrlRecordDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadInlineSslUrlRecordDataSource_Read_NilClient(t *testing.T) {
	r := &LoadInlineSslUrlRecordDataSource{}
	m := LoadInlineSslUrlRecordDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadInlineSslUrlRecordDataSource_Read_BuildError exercises LoadInlineSslUrlRecordDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadInlineSslUrlRecordDataSource_Read_BuildError(t *testing.T) {
	r := &LoadInlineSslUrlRecordDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadInlineSslUrlRecordDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadInlineSslUrlRecordDataSource_Read_SendError exercises LoadInlineSslUrlRecordDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadInlineSslUrlRecordDataSource_Read_SendError(t *testing.T) {
	r := &LoadInlineSslUrlRecordDataSource{client: newTransportErrorClient(t)}
	m := LoadInlineSslUrlRecordDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadInlineSslUrlRecordDataSource_Read_InvalidJSON exercises LoadInlineSslUrlRecordDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadInlineSslUrlRecordDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadInlineSslUrlRecordDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadInlineSslUrlRecordDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
