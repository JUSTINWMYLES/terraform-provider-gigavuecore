package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMetadataCacheDataSource_Read_Happy exercises LoadAllMetadataCacheDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllMetadataCacheDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllMetadataCacheDataSource{client: newMockClientStatus(t, 200, "{\"metadataCaches\":[]}")}
	m := LoadAllMetadataCacheDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllMetadataCacheDataSource_Read_NilClient exercises LoadAllMetadataCacheDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllMetadataCacheDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllMetadataCacheDataSource{}
	m := LoadAllMetadataCacheDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllMetadataCacheDataSource_Read_BuildError exercises LoadAllMetadataCacheDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllMetadataCacheDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllMetadataCacheDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllMetadataCacheDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMetadataCacheDataSource_Read_SendError exercises LoadAllMetadataCacheDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllMetadataCacheDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllMetadataCacheDataSource{client: newTransportErrorClient(t)}
	m := LoadAllMetadataCacheDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMetadataCacheDataSource_Read_InvalidJSON exercises LoadAllMetadataCacheDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllMetadataCacheDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllMetadataCacheDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllMetadataCacheDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
