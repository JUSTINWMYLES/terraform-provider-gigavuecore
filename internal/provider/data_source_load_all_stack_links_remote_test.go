package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllStackLinksDataSource_Read_Happy exercises LoadAllStackLinksDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllStackLinksDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllStackLinksDataSource{client: newMockClientStatus(t, 200, "{\"stackLinks\":[]}")}
	m := LoadAllStackLinksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllStackLinksDataSource_Read_NilClient exercises LoadAllStackLinksDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllStackLinksDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllStackLinksDataSource{}
	m := LoadAllStackLinksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllStackLinksDataSource_Read_BuildError exercises LoadAllStackLinksDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllStackLinksDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllStackLinksDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllStackLinksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllStackLinksDataSource_Read_SendError exercises LoadAllStackLinksDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllStackLinksDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllStackLinksDataSource{client: newTransportErrorClient(t)}
	m := LoadAllStackLinksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllStackLinksDataSource_Read_InvalidJSON exercises LoadAllStackLinksDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllStackLinksDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllStackLinksDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllStackLinksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
