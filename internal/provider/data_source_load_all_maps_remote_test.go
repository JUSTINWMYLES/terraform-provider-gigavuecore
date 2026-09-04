package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMapsDataSource_Read_Happy exercises LoadAllMapsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllMapsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllMapsDataSource{client: newMockClientStatus(t, 200, "{\"maps\":[]}")}
	m := LoadAllMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllMapsDataSource_Read_NilClient exercises LoadAllMapsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllMapsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllMapsDataSource{}
	m := LoadAllMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllMapsDataSource_Read_BuildError exercises LoadAllMapsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllMapsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllMapsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMapsDataSource_Read_SendError exercises LoadAllMapsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllMapsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllMapsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMapsDataSource_Read_InvalidJSON exercises LoadAllMapsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllMapsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllMapsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
