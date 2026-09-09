package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTopovizExternalDevicesDataSource_Read_Happy exercises GetTopovizExternalDevicesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetTopovizExternalDevicesDataSource_Read_Happy(t *testing.T) {
	r := &GetTopovizExternalDevicesDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetTopovizExternalDevicesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTopovizExternalDevicesDataSource_Read_NilClient exercises GetTopovizExternalDevicesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTopovizExternalDevicesDataSource_Read_NilClient(t *testing.T) {
	r := &GetTopovizExternalDevicesDataSource{}
	m := GetTopovizExternalDevicesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTopovizExternalDevicesDataSource_Read_BuildError exercises GetTopovizExternalDevicesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetTopovizExternalDevicesDataSource_Read_BuildError(t *testing.T) {
	r := &GetTopovizExternalDevicesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTopovizExternalDevicesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetTopovizExternalDevicesDataSource_Read_SendError exercises GetTopovizExternalDevicesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetTopovizExternalDevicesDataSource_Read_SendError(t *testing.T) {
	r := &GetTopovizExternalDevicesDataSource{client: newTransportErrorClient(t)}
	m := GetTopovizExternalDevicesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetTopovizExternalDevicesDataSource_Read_InvalidJSON exercises GetTopovizExternalDevicesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetTopovizExternalDevicesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTopovizExternalDevicesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTopovizExternalDevicesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
