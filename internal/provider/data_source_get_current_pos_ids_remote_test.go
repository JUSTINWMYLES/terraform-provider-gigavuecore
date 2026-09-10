package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCurrentPosIdsDataSource_Read_Happy exercises GetCurrentPosIdsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetCurrentPosIdsDataSource_Read_Happy(t *testing.T) {
	r := &GetCurrentPosIdsDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetCurrentPosIdsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCurrentPosIdsDataSource_Read_NilClient exercises GetCurrentPosIdsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCurrentPosIdsDataSource_Read_NilClient(t *testing.T) {
	r := &GetCurrentPosIdsDataSource{}
	m := GetCurrentPosIdsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCurrentPosIdsDataSource_Read_BuildError exercises GetCurrentPosIdsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetCurrentPosIdsDataSource_Read_BuildError(t *testing.T) {
	r := &GetCurrentPosIdsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCurrentPosIdsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCurrentPosIdsDataSource_Read_SendError exercises GetCurrentPosIdsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetCurrentPosIdsDataSource_Read_SendError(t *testing.T) {
	r := &GetCurrentPosIdsDataSource{client: newTransportErrorClient(t)}
	m := GetCurrentPosIdsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCurrentPosIdsDataSource_Read_InvalidJSON exercises GetCurrentPosIdsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetCurrentPosIdsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCurrentPosIdsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCurrentPosIdsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
