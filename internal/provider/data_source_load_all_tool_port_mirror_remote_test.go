package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllToolPortMirrorDataSource_Read_Happy exercises LoadAllToolPortMirrorDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllToolPortMirrorDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllToolPortMirrorDataSource{client: newMockClientStatus(t, 200, "{\"toolPortMirrors\":[]}")}
	m := LoadAllToolPortMirrorDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllToolPortMirrorDataSource_Read_NilClient exercises LoadAllToolPortMirrorDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllToolPortMirrorDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllToolPortMirrorDataSource{}
	m := LoadAllToolPortMirrorDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllToolPortMirrorDataSource_Read_BuildError exercises LoadAllToolPortMirrorDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllToolPortMirrorDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllToolPortMirrorDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllToolPortMirrorDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllToolPortMirrorDataSource_Read_SendError exercises LoadAllToolPortMirrorDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllToolPortMirrorDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllToolPortMirrorDataSource{client: newTransportErrorClient(t)}
	m := LoadAllToolPortMirrorDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllToolPortMirrorDataSource_Read_InvalidJSON exercises LoadAllToolPortMirrorDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllToolPortMirrorDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllToolPortMirrorDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllToolPortMirrorDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
