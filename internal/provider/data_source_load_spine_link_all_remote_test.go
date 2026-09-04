package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSpineLinkAllDataSource_Read_Happy exercises LoadSpineLinkAllDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadSpineLinkAllDataSource_Read_Happy(t *testing.T) {
	r := &LoadSpineLinkAllDataSource{client: newMockClientStatus(t, 200, "{\"spineLinks\":[]}")}
	m := LoadSpineLinkAllDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadSpineLinkAllDataSource_Read_NilClient exercises LoadSpineLinkAllDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadSpineLinkAllDataSource_Read_NilClient(t *testing.T) {
	r := &LoadSpineLinkAllDataSource{}
	m := LoadSpineLinkAllDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadSpineLinkAllDataSource_Read_BuildError exercises LoadSpineLinkAllDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadSpineLinkAllDataSource_Read_BuildError(t *testing.T) {
	r := &LoadSpineLinkAllDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadSpineLinkAllDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadSpineLinkAllDataSource_Read_SendError exercises LoadSpineLinkAllDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadSpineLinkAllDataSource_Read_SendError(t *testing.T) {
	r := &LoadSpineLinkAllDataSource{client: newTransportErrorClient(t)}
	m := LoadSpineLinkAllDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadSpineLinkAllDataSource_Read_InvalidJSON exercises LoadSpineLinkAllDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadSpineLinkAllDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadSpineLinkAllDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadSpineLinkAllDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
