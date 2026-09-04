package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPeriodVolumesDataSource_Read_Happy exercises GetPeriodVolumesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetPeriodVolumesDataSource_Read_Happy(t *testing.T) {
	r := &GetPeriodVolumesDataSource{client: newMockClientStatus(t, 200, "{\"volumes\":[]}")}
	m := GetPeriodVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPeriodVolumesDataSource_Read_NilClient exercises GetPeriodVolumesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPeriodVolumesDataSource_Read_NilClient(t *testing.T) {
	r := &GetPeriodVolumesDataSource{}
	m := GetPeriodVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPeriodVolumesDataSource_Read_BuildError exercises GetPeriodVolumesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetPeriodVolumesDataSource_Read_BuildError(t *testing.T) {
	r := &GetPeriodVolumesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPeriodVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetPeriodVolumesDataSource_Read_SendError exercises GetPeriodVolumesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetPeriodVolumesDataSource_Read_SendError(t *testing.T) {
	r := &GetPeriodVolumesDataSource{client: newTransportErrorClient(t)}
	m := GetPeriodVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetPeriodVolumesDataSource_Read_InvalidJSON exercises GetPeriodVolumesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetPeriodVolumesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPeriodVolumesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPeriodVolumesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
