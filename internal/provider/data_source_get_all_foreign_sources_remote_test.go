package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllForeignSourcesDataSource_Read_Happy exercises GetAllForeignSourcesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllForeignSourcesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllForeignSourcesDataSource{client: newMockClientStatus(t, 200, "{\"foreignSources\":[]}")}
	m := GetAllForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllForeignSourcesDataSource_Read_NilClient exercises GetAllForeignSourcesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllForeignSourcesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllForeignSourcesDataSource{}
	m := GetAllForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllForeignSourcesDataSource_Read_BuildError exercises GetAllForeignSourcesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllForeignSourcesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllForeignSourcesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllForeignSourcesDataSource_Read_SendError exercises GetAllForeignSourcesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllForeignSourcesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllForeignSourcesDataSource{client: newTransportErrorClient(t)}
	m := GetAllForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllForeignSourcesDataSource_Read_InvalidJSON exercises GetAllForeignSourcesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllForeignSourcesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllForeignSourcesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllForeignSourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
