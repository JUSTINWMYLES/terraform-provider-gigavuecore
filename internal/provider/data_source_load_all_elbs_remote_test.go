package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllElbsDataSource_Read_Happy exercises LoadAllElbsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllElbsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllElbsDataSource{client: newMockClientStatus(t, 200, "{\"elbs\":[]}")}
	m := LoadAllElbsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllElbsDataSource_Read_NilClient exercises LoadAllElbsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllElbsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllElbsDataSource{}
	m := LoadAllElbsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllElbsDataSource_Read_BuildError exercises LoadAllElbsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllElbsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllElbsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllElbsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllElbsDataSource_Read_SendError exercises LoadAllElbsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllElbsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllElbsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllElbsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllElbsDataSource_Read_InvalidJSON exercises LoadAllElbsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllElbsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllElbsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllElbsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
