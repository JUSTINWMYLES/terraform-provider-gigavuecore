package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAlarmByIdDataSource_Read_Happy exercises GetAlarmByIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAlarmByIdDataSource_Read_Happy(t *testing.T) {
	r := &GetAlarmByIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAlarmByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAlarmByIdDataSource_Read_NilClient exercises GetAlarmByIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAlarmByIdDataSource_Read_NilClient(t *testing.T) {
	r := &GetAlarmByIdDataSource{}
	m := GetAlarmByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAlarmByIdDataSource_Read_BuildError exercises GetAlarmByIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAlarmByIdDataSource_Read_BuildError(t *testing.T) {
	r := &GetAlarmByIdDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAlarmByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAlarmByIdDataSource_Read_SendError exercises GetAlarmByIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAlarmByIdDataSource_Read_SendError(t *testing.T) {
	r := &GetAlarmByIdDataSource{client: newTransportErrorClient(t)}
	m := GetAlarmByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAlarmByIdDataSource_Read_NotFound exercises GetAlarmByIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAlarmByIdDataSource_Read_NotFound(t *testing.T) {
	r := &GetAlarmByIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAlarmByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAlarmByIdDataSource_Read_APIError exercises GetAlarmByIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAlarmByIdDataSource_Read_APIError(t *testing.T) {
	r := &GetAlarmByIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAlarmByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_alarm_by_id")
}

// TestGetAlarmByIdDataSource_Read_APIErrorReadBody exercises GetAlarmByIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAlarmByIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAlarmByIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAlarmByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAlarmByIdDataSource_Read_InvalidJSON exercises GetAlarmByIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAlarmByIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAlarmByIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAlarmByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
