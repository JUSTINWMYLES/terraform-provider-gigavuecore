package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllPortThrottlesDataSource_Read_Happy exercises LoadAllPortThrottlesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllPortThrottlesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllPortThrottlesDataSource{client: newMockClientStatus(t, 200, "{\"portThrottles\":[]}")}
	m := LoadAllPortThrottlesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllPortThrottlesDataSource_Read_NilClient exercises LoadAllPortThrottlesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllPortThrottlesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllPortThrottlesDataSource{}
	m := LoadAllPortThrottlesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllPortThrottlesDataSource_Read_BuildError exercises LoadAllPortThrottlesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllPortThrottlesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllPortThrottlesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllPortThrottlesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllPortThrottlesDataSource_Read_SendError exercises LoadAllPortThrottlesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllPortThrottlesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllPortThrottlesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllPortThrottlesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllPortThrottlesDataSource_Read_InvalidJSON exercises LoadAllPortThrottlesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllPortThrottlesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllPortThrottlesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllPortThrottlesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
