package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTimePropertiesDataSource_Read_Happy exercises GetAllTimePropertiesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllTimePropertiesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllTimePropertiesDataSource{client: newMockClientStatus(t, 200, "{\"timeProperties\":[]}")}
	m := GetAllTimePropertiesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllTimePropertiesDataSource_Read_NilClient exercises GetAllTimePropertiesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllTimePropertiesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllTimePropertiesDataSource{}
	m := GetAllTimePropertiesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllTimePropertiesDataSource_Read_BuildError exercises GetAllTimePropertiesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllTimePropertiesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllTimePropertiesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllTimePropertiesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTimePropertiesDataSource_Read_SendError exercises GetAllTimePropertiesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllTimePropertiesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllTimePropertiesDataSource{client: newTransportErrorClient(t)}
	m := GetAllTimePropertiesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTimePropertiesDataSource_Read_InvalidJSON exercises GetAllTimePropertiesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllTimePropertiesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllTimePropertiesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllTimePropertiesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
