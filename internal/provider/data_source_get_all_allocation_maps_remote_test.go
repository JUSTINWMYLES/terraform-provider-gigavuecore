package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllAllocationMapsDataSource_Read_Happy exercises GetAllAllocationMapsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllAllocationMapsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllAllocationMapsDataSource{client: newMockClientStatus(t, 200, "{\"allocationMaps\":[]}")}
	m := GetAllAllocationMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllAllocationMapsDataSource_Read_NilClient exercises GetAllAllocationMapsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllAllocationMapsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllAllocationMapsDataSource{}
	m := GetAllAllocationMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllAllocationMapsDataSource_Read_BuildError exercises GetAllAllocationMapsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllAllocationMapsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllAllocationMapsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllAllocationMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllAllocationMapsDataSource_Read_SendError exercises GetAllAllocationMapsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllAllocationMapsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllAllocationMapsDataSource{client: newTransportErrorClient(t)}
	m := GetAllAllocationMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllAllocationMapsDataSource_Read_InvalidJSON exercises GetAllAllocationMapsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllAllocationMapsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllAllocationMapsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllAllocationMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
