package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllPortPairDataSource_Read_Happy exercises LoadAllPortPairDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllPortPairDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllPortPairDataSource{client: newMockClientStatus(t, 200, "{\"portPairs\":[]}")}
	m := LoadAllPortPairDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllPortPairDataSource_Read_NilClient exercises LoadAllPortPairDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllPortPairDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllPortPairDataSource{}
	m := LoadAllPortPairDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllPortPairDataSource_Read_BuildError exercises LoadAllPortPairDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllPortPairDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllPortPairDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllPortPairDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllPortPairDataSource_Read_SendError exercises LoadAllPortPairDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllPortPairDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllPortPairDataSource{client: newTransportErrorClient(t)}
	m := LoadAllPortPairDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllPortPairDataSource_Read_InvalidJSON exercises LoadAllPortPairDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllPortPairDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllPortPairDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllPortPairDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
