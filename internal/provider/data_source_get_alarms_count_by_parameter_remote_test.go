package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAlarmsCountByParameterDataSource_Read_Happy exercises GetAlarmsCountByParameterDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAlarmsCountByParameterDataSource_Read_Happy(t *testing.T) {
	r := &GetAlarmsCountByParameterDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAlarmsCountByParameterDataSource_Read_NilClient exercises GetAlarmsCountByParameterDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAlarmsCountByParameterDataSource_Read_NilClient(t *testing.T) {
	r := &GetAlarmsCountByParameterDataSource{}
	m := GetAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAlarmsCountByParameterDataSource_Read_BuildError exercises GetAlarmsCountByParameterDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAlarmsCountByParameterDataSource_Read_BuildError(t *testing.T) {
	r := &GetAlarmsCountByParameterDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAlarmsCountByParameterDataSource_Read_SendError exercises GetAlarmsCountByParameterDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAlarmsCountByParameterDataSource_Read_SendError(t *testing.T) {
	r := &GetAlarmsCountByParameterDataSource{client: newTransportErrorClient(t)}
	m := GetAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAlarmsCountByParameterDataSource_Read_NotFound exercises GetAlarmsCountByParameterDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAlarmsCountByParameterDataSource_Read_NotFound(t *testing.T) {
	r := &GetAlarmsCountByParameterDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAlarmsCountByParameterDataSource_Read_APIError exercises GetAlarmsCountByParameterDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAlarmsCountByParameterDataSource_Read_APIError(t *testing.T) {
	r := &GetAlarmsCountByParameterDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_alarms_count_by_parameter")
}

// TestGetAlarmsCountByParameterDataSource_Read_APIErrorReadBody exercises GetAlarmsCountByParameterDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAlarmsCountByParameterDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAlarmsCountByParameterDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAlarmsCountByParameterDataSource_Read_InvalidJSON exercises GetAlarmsCountByParameterDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAlarmsCountByParameterDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAlarmsCountByParameterDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAlarmsCountByParameterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
