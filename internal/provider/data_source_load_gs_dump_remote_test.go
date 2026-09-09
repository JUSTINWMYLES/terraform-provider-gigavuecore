package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGsDumpDataSource_Read_Happy exercises LoadGsDumpDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadGsDumpDataSource_Read_Happy(t *testing.T) {
	r := &LoadGsDumpDataSource{client: newMockClientStatus(t, 200, "{\"gsDumpFiles\":[]}")}
	m := LoadGsDumpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadGsDumpDataSource_Read_NilClient exercises LoadGsDumpDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadGsDumpDataSource_Read_NilClient(t *testing.T) {
	r := &LoadGsDumpDataSource{}
	m := LoadGsDumpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadGsDumpDataSource_Read_BuildError exercises LoadGsDumpDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadGsDumpDataSource_Read_BuildError(t *testing.T) {
	r := &LoadGsDumpDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadGsDumpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadGsDumpDataSource_Read_SendError exercises LoadGsDumpDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadGsDumpDataSource_Read_SendError(t *testing.T) {
	r := &LoadGsDumpDataSource{client: newTransportErrorClient(t)}
	m := LoadGsDumpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadGsDumpDataSource_Read_InvalidJSON exercises LoadGsDumpDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadGsDumpDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadGsDumpDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadGsDumpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
