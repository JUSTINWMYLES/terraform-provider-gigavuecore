package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllForeignMastersDataSource_Read_Happy exercises GetAllForeignMastersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllForeignMastersDataSource_Read_Happy(t *testing.T) {
	r := &GetAllForeignMastersDataSource{client: newMockClientStatus(t, 200, "{\"foreignMasters\":[]}")}
	m := GetAllForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllForeignMastersDataSource_Read_NilClient exercises GetAllForeignMastersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllForeignMastersDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllForeignMastersDataSource{}
	m := GetAllForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllForeignMastersDataSource_Read_BuildError exercises GetAllForeignMastersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllForeignMastersDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllForeignMastersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllForeignMastersDataSource_Read_SendError exercises GetAllForeignMastersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllForeignMastersDataSource_Read_SendError(t *testing.T) {
	r := &GetAllForeignMastersDataSource{client: newTransportErrorClient(t)}
	m := GetAllForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllForeignMastersDataSource_Read_InvalidJSON exercises GetAllForeignMastersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllForeignMastersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllForeignMastersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
