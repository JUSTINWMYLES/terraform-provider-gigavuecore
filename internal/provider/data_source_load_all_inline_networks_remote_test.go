package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllInlineNetworksDataSource_Read_Happy exercises LoadAllInlineNetworksDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllInlineNetworksDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllInlineNetworksDataSource{client: newMockClientStatus(t, 200, "{\"inlineNetworks\":[]}")}
	m := LoadAllInlineNetworksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllInlineNetworksDataSource_Read_NilClient exercises LoadAllInlineNetworksDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllInlineNetworksDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllInlineNetworksDataSource{}
	m := LoadAllInlineNetworksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllInlineNetworksDataSource_Read_BuildError exercises LoadAllInlineNetworksDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllInlineNetworksDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllInlineNetworksDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllInlineNetworksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllInlineNetworksDataSource_Read_SendError exercises LoadAllInlineNetworksDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllInlineNetworksDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllInlineNetworksDataSource{client: newTransportErrorClient(t)}
	m := LoadAllInlineNetworksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllInlineNetworksDataSource_Read_InvalidJSON exercises LoadAllInlineNetworksDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllInlineNetworksDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllInlineNetworksDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllInlineNetworksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
