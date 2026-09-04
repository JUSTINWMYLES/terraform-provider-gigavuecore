package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllInternalFabricMapsDataSource_Read_Happy exercises GetAllInternalFabricMapsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllInternalFabricMapsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllInternalFabricMapsDataSource{client: newMockClientStatus(t, 200, "{\"fabricMaps\":[]}")}
	m := GetAllInternalFabricMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllInternalFabricMapsDataSource_Read_NilClient exercises GetAllInternalFabricMapsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllInternalFabricMapsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllInternalFabricMapsDataSource{}
	m := GetAllInternalFabricMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllInternalFabricMapsDataSource_Read_BuildError exercises GetAllInternalFabricMapsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllInternalFabricMapsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllInternalFabricMapsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllInternalFabricMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllInternalFabricMapsDataSource_Read_SendError exercises GetAllInternalFabricMapsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllInternalFabricMapsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllInternalFabricMapsDataSource{client: newTransportErrorClient(t)}
	m := GetAllInternalFabricMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllInternalFabricMapsDataSource_Read_InvalidJSON exercises GetAllInternalFabricMapsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllInternalFabricMapsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllInternalFabricMapsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllInternalFabricMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
