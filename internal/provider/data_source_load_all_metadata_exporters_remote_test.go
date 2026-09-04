package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMetadataExportersDataSource_Read_Happy exercises LoadAllMetadataExportersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllMetadataExportersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllMetadataExportersDataSource{client: newMockClientStatus(t, 200, "{\"metadataExporters\":[]}")}
	m := LoadAllMetadataExportersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllMetadataExportersDataSource_Read_NilClient exercises LoadAllMetadataExportersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllMetadataExportersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllMetadataExportersDataSource{}
	m := LoadAllMetadataExportersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllMetadataExportersDataSource_Read_BuildError exercises LoadAllMetadataExportersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllMetadataExportersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllMetadataExportersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllMetadataExportersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMetadataExportersDataSource_Read_SendError exercises LoadAllMetadataExportersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllMetadataExportersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllMetadataExportersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllMetadataExportersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMetadataExportersDataSource_Read_InvalidJSON exercises LoadAllMetadataExportersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllMetadataExportersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllMetadataExportersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllMetadataExportersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
