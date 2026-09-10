package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFabricMapsDataSource_Read_Happy exercises GetAllFabricMapsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllFabricMapsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllFabricMapsDataSource{client: newMockClientStatus(t, 200, "{\"fabricMaps\":[]}")}
	m := GetAllFabricMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllFabricMapsDataSource_Read_NilClient exercises GetAllFabricMapsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllFabricMapsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllFabricMapsDataSource{}
	m := GetAllFabricMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllFabricMapsDataSource_Read_BuildError exercises GetAllFabricMapsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllFabricMapsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllFabricMapsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllFabricMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFabricMapsDataSource_Read_SendError exercises GetAllFabricMapsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllFabricMapsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllFabricMapsDataSource{client: newTransportErrorClient(t)}
	m := GetAllFabricMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFabricMapsDataSource_Read_InvalidJSON exercises GetAllFabricMapsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllFabricMapsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllFabricMapsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllFabricMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
