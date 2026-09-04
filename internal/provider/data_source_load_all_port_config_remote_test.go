package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllPortConfigDataSource_Read_Happy exercises LoadAllPortConfigDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllPortConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllPortConfigDataSource{client: newMockClientStatus(t, 200, "{\"portConfigs\":[]}")}
	m := LoadAllPortConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllPortConfigDataSource_Read_NilClient exercises LoadAllPortConfigDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllPortConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllPortConfigDataSource{}
	m := LoadAllPortConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllPortConfigDataSource_Read_BuildError exercises LoadAllPortConfigDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllPortConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllPortConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllPortConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllPortConfigDataSource_Read_SendError exercises LoadAllPortConfigDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllPortConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllPortConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadAllPortConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllPortConfigDataSource_Read_InvalidJSON exercises LoadAllPortConfigDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllPortConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllPortConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllPortConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
