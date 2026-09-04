package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFlexInlineSslAppsDataSource_Read_Happy exercises GetAllFlexInlineSslAppsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllFlexInlineSslAppsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllFlexInlineSslAppsDataSource{client: newMockClientStatus(t, 200, "{\"gigaFlexInlineSslApps\":[]}")}
	m := GetAllFlexInlineSslAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllFlexInlineSslAppsDataSource_Read_NilClient exercises GetAllFlexInlineSslAppsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllFlexInlineSslAppsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllFlexInlineSslAppsDataSource{}
	m := GetAllFlexInlineSslAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllFlexInlineSslAppsDataSource_Read_BuildError exercises GetAllFlexInlineSslAppsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllFlexInlineSslAppsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllFlexInlineSslAppsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllFlexInlineSslAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFlexInlineSslAppsDataSource_Read_SendError exercises GetAllFlexInlineSslAppsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllFlexInlineSslAppsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllFlexInlineSslAppsDataSource{client: newTransportErrorClient(t)}
	m := GetAllFlexInlineSslAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFlexInlineSslAppsDataSource_Read_InvalidJSON exercises GetAllFlexInlineSslAppsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllFlexInlineSslAppsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllFlexInlineSslAppsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllFlexInlineSslAppsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
