package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllEnhancedSlicingsDataSource_Read_Happy exercises LoadAllEnhancedSlicingsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllEnhancedSlicingsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllEnhancedSlicingsDataSource{client: newMockClientStatus(t, 200, "{\"enhancedSlicings\":[]}")}
	m := LoadAllEnhancedSlicingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllEnhancedSlicingsDataSource_Read_NilClient exercises LoadAllEnhancedSlicingsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllEnhancedSlicingsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllEnhancedSlicingsDataSource{}
	m := LoadAllEnhancedSlicingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllEnhancedSlicingsDataSource_Read_BuildError exercises LoadAllEnhancedSlicingsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllEnhancedSlicingsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllEnhancedSlicingsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllEnhancedSlicingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllEnhancedSlicingsDataSource_Read_SendError exercises LoadAllEnhancedSlicingsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllEnhancedSlicingsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllEnhancedSlicingsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllEnhancedSlicingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllEnhancedSlicingsDataSource_Read_InvalidJSON exercises LoadAllEnhancedSlicingsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllEnhancedSlicingsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllEnhancedSlicingsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllEnhancedSlicingsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
