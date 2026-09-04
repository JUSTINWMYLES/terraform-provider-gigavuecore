package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryForApplicationsDataSource_Read_Happy exercises QueryForApplicationsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestQueryForApplicationsDataSource_Read_Happy(t *testing.T) {
	r := &QueryForApplicationsDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := QueryForApplicationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestQueryForApplicationsDataSource_Read_NilClient exercises QueryForApplicationsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestQueryForApplicationsDataSource_Read_NilClient(t *testing.T) {
	r := &QueryForApplicationsDataSource{}
	m := QueryForApplicationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestQueryForApplicationsDataSource_Read_BuildError exercises QueryForApplicationsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestQueryForApplicationsDataSource_Read_BuildError(t *testing.T) {
	r := &QueryForApplicationsDataSource{client: newMalformedBaseURLClient(t)}
	m := QueryForApplicationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestQueryForApplicationsDataSource_Read_SendError exercises QueryForApplicationsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestQueryForApplicationsDataSource_Read_SendError(t *testing.T) {
	r := &QueryForApplicationsDataSource{client: newTransportErrorClient(t)}
	m := QueryForApplicationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestQueryForApplicationsDataSource_Read_InvalidJSON exercises QueryForApplicationsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestQueryForApplicationsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &QueryForApplicationsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := QueryForApplicationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
