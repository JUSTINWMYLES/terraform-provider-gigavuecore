package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCorrelatedAlarmsCountByParameterDataSource_Read_Happy exercises GetCorrelatedAlarmsCountByParameterDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetCorrelatedAlarmsCountByParameterDataSource_Read_Happy(t *testing.T) {
	r := &GetCorrelatedAlarmsCountByParameterDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetCorrelatedAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCorrelatedAlarmsCountByParameterDataSource_Read_NilClient exercises GetCorrelatedAlarmsCountByParameterDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCorrelatedAlarmsCountByParameterDataSource_Read_NilClient(t *testing.T) {
	r := &GetCorrelatedAlarmsCountByParameterDataSource{}
	m := GetCorrelatedAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCorrelatedAlarmsCountByParameterDataSource_Read_BuildError exercises GetCorrelatedAlarmsCountByParameterDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetCorrelatedAlarmsCountByParameterDataSource_Read_BuildError(t *testing.T) {
	r := &GetCorrelatedAlarmsCountByParameterDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCorrelatedAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetCorrelatedAlarmsCountByParameterDataSource_Read_SendError exercises GetCorrelatedAlarmsCountByParameterDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetCorrelatedAlarmsCountByParameterDataSource_Read_SendError(t *testing.T) {
	r := &GetCorrelatedAlarmsCountByParameterDataSource{client: newTransportErrorClient(t)}
	m := GetCorrelatedAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetCorrelatedAlarmsCountByParameterDataSource_Read_NotFound exercises GetCorrelatedAlarmsCountByParameterDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetCorrelatedAlarmsCountByParameterDataSource_Read_NotFound(t *testing.T) {
	r := &GetCorrelatedAlarmsCountByParameterDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetCorrelatedAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetCorrelatedAlarmsCountByParameterDataSource_Read_APIError exercises GetCorrelatedAlarmsCountByParameterDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetCorrelatedAlarmsCountByParameterDataSource_Read_APIError(t *testing.T) {
	r := &GetCorrelatedAlarmsCountByParameterDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetCorrelatedAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_correlated_alarms_count_by_parameter")
}

// TestGetCorrelatedAlarmsCountByParameterDataSource_Read_APIErrorReadBody exercises GetCorrelatedAlarmsCountByParameterDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetCorrelatedAlarmsCountByParameterDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetCorrelatedAlarmsCountByParameterDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetCorrelatedAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetCorrelatedAlarmsCountByParameterDataSource_Read_InvalidJSON exercises GetCorrelatedAlarmsCountByParameterDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetCorrelatedAlarmsCountByParameterDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCorrelatedAlarmsCountByParameterDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCorrelatedAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
