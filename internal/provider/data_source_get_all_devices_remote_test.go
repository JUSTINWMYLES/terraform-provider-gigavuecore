package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllDevicesDataSource_Read_Happy exercises GetAllDevicesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllDevicesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllDevicesDataSource{client: newMockClientStatus(t, 200, "{\"devices\":[]}")}
	m := GetAllDevicesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllDevicesDataSource_Read_NilClient exercises GetAllDevicesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllDevicesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllDevicesDataSource{}
	m := GetAllDevicesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllDevicesDataSource_Read_BuildError exercises GetAllDevicesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllDevicesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllDevicesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllDevicesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllDevicesDataSource_Read_SendError exercises GetAllDevicesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllDevicesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllDevicesDataSource{client: newTransportErrorClient(t)}
	m := GetAllDevicesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllDevicesDataSource_Read_InvalidJSON exercises GetAllDevicesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllDevicesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllDevicesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllDevicesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
