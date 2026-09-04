package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListEnvDataSource_Read_Happy exercises ListEnvDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestListEnvDataSource_Read_Happy(t *testing.T) {
	r := &ListEnvDataSource{client: newMockClientStatus(t, 200, "{\"env\":[]}")}
	m := ListEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListEnvDataSource_Read_NilClient exercises ListEnvDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListEnvDataSource_Read_NilClient(t *testing.T) {
	r := &ListEnvDataSource{}
	m := ListEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListEnvDataSource_Read_BuildError exercises ListEnvDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestListEnvDataSource_Read_BuildError(t *testing.T) {
	r := &ListEnvDataSource{client: newMalformedBaseURLClient(t)}
	m := ListEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestListEnvDataSource_Read_SendError exercises ListEnvDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestListEnvDataSource_Read_SendError(t *testing.T) {
	r := &ListEnvDataSource{client: newTransportErrorClient(t)}
	m := ListEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestListEnvDataSource_Read_InvalidJSON exercises ListEnvDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestListEnvDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ListEnvDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ListEnvDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
