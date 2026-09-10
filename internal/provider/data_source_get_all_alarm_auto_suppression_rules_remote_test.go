package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllAlarmAutoSuppressionRulesDataSource_Read_Happy exercises GetAllAlarmAutoSuppressionRulesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAllAlarmAutoSuppressionRulesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllAlarmAutoSuppressionRulesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAllAlarmAutoSuppressionRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllAlarmAutoSuppressionRulesDataSource_Read_NilClient exercises GetAllAlarmAutoSuppressionRulesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllAlarmAutoSuppressionRulesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllAlarmAutoSuppressionRulesDataSource{}
	m := GetAllAlarmAutoSuppressionRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllAlarmAutoSuppressionRulesDataSource_Read_BuildError exercises GetAllAlarmAutoSuppressionRulesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAllAlarmAutoSuppressionRulesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllAlarmAutoSuppressionRulesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllAlarmAutoSuppressionRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAllAlarmAutoSuppressionRulesDataSource_Read_SendError exercises GetAllAlarmAutoSuppressionRulesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAllAlarmAutoSuppressionRulesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllAlarmAutoSuppressionRulesDataSource{client: newTransportErrorClient(t)}
	m := GetAllAlarmAutoSuppressionRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAllAlarmAutoSuppressionRulesDataSource_Read_NotFound exercises GetAllAlarmAutoSuppressionRulesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAllAlarmAutoSuppressionRulesDataSource_Read_NotFound(t *testing.T) {
	r := &GetAllAlarmAutoSuppressionRulesDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAllAlarmAutoSuppressionRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAllAlarmAutoSuppressionRulesDataSource_Read_APIError exercises GetAllAlarmAutoSuppressionRulesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAllAlarmAutoSuppressionRulesDataSource_Read_APIError(t *testing.T) {
	r := &GetAllAlarmAutoSuppressionRulesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAllAlarmAutoSuppressionRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_all_alarm_auto_suppression_rules")
}

// TestGetAllAlarmAutoSuppressionRulesDataSource_Read_APIErrorReadBody exercises GetAllAlarmAutoSuppressionRulesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAllAlarmAutoSuppressionRulesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAllAlarmAutoSuppressionRulesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAllAlarmAutoSuppressionRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAllAlarmAutoSuppressionRulesDataSource_Read_InvalidJSON exercises GetAllAlarmAutoSuppressionRulesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAllAlarmAutoSuppressionRulesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllAlarmAutoSuppressionRulesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllAlarmAutoSuppressionRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
