package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMetadataApplicationTemplatesDataSource_Read_Happy exercises LoadAllMetadataApplicationTemplatesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllMetadataApplicationTemplatesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllMetadataApplicationTemplatesDataSource{client: newMockClientStatus(t, 200, "{\"metadataApplicationTemplates\":[]}")}
	m := LoadAllMetadataApplicationTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllMetadataApplicationTemplatesDataSource_Read_NilClient exercises LoadAllMetadataApplicationTemplatesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllMetadataApplicationTemplatesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllMetadataApplicationTemplatesDataSource{}
	m := LoadAllMetadataApplicationTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllMetadataApplicationTemplatesDataSource_Read_BuildError exercises LoadAllMetadataApplicationTemplatesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllMetadataApplicationTemplatesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllMetadataApplicationTemplatesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllMetadataApplicationTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMetadataApplicationTemplatesDataSource_Read_SendError exercises LoadAllMetadataApplicationTemplatesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllMetadataApplicationTemplatesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllMetadataApplicationTemplatesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllMetadataApplicationTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMetadataApplicationTemplatesDataSource_Read_InvalidJSON exercises LoadAllMetadataApplicationTemplatesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllMetadataApplicationTemplatesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllMetadataApplicationTemplatesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllMetadataApplicationTemplatesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
