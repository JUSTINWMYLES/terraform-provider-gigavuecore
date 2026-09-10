package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllCorrelatedAlarmsDataSource_Read_Happy exercises GetAllCorrelatedAlarmsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllCorrelatedAlarmsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllCorrelatedAlarmsDataSource{client: newMockClientStatus(t, 200, "{\"alarms\":[]}")}
	m := GetAllCorrelatedAlarmsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllCorrelatedAlarmsDataSource_Read_NilClient exercises GetAllCorrelatedAlarmsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllCorrelatedAlarmsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllCorrelatedAlarmsDataSource{}
	m := GetAllCorrelatedAlarmsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllCorrelatedAlarmsDataSource_Read_BuildError exercises GetAllCorrelatedAlarmsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllCorrelatedAlarmsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllCorrelatedAlarmsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllCorrelatedAlarmsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllCorrelatedAlarmsDataSource_Read_SendError exercises GetAllCorrelatedAlarmsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllCorrelatedAlarmsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllCorrelatedAlarmsDataSource{client: newTransportErrorClient(t)}
	m := GetAllCorrelatedAlarmsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllCorrelatedAlarmsDataSource_Read_InvalidJSON exercises GetAllCorrelatedAlarmsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllCorrelatedAlarmsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllCorrelatedAlarmsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllCorrelatedAlarmsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
