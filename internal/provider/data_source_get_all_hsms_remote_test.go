package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllHsmsDataSource_Read_Happy exercises GetAllHsmsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllHsmsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllHsmsDataSource{client: newMockClientStatus(t, 200, "{\"hsms\":[]}")}
	m := GetAllHsmsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllHsmsDataSource_Read_NilClient exercises GetAllHsmsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllHsmsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllHsmsDataSource{}
	m := GetAllHsmsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllHsmsDataSource_Read_BuildError exercises GetAllHsmsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllHsmsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllHsmsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllHsmsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllHsmsDataSource_Read_SendError exercises GetAllHsmsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllHsmsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllHsmsDataSource{client: newTransportErrorClient(t)}
	m := GetAllHsmsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllHsmsDataSource_Read_InvalidJSON exercises GetAllHsmsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllHsmsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllHsmsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllHsmsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
