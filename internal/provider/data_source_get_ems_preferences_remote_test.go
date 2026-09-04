package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEmsPreferencesDataSource_Read_Happy exercises GetEmsPreferencesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetEmsPreferencesDataSource_Read_Happy(t *testing.T) {
	r := &GetEmsPreferencesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetEmsPreferencesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEmsPreferencesDataSource_Read_NilClient exercises GetEmsPreferencesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEmsPreferencesDataSource_Read_NilClient(t *testing.T) {
	r := &GetEmsPreferencesDataSource{}
	m := GetEmsPreferencesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEmsPreferencesDataSource_Read_BuildError exercises GetEmsPreferencesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetEmsPreferencesDataSource_Read_BuildError(t *testing.T) {
	r := &GetEmsPreferencesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEmsPreferencesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetEmsPreferencesDataSource_Read_SendError exercises GetEmsPreferencesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetEmsPreferencesDataSource_Read_SendError(t *testing.T) {
	r := &GetEmsPreferencesDataSource{client: newTransportErrorClient(t)}
	m := GetEmsPreferencesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetEmsPreferencesDataSource_Read_NotFound exercises GetEmsPreferencesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetEmsPreferencesDataSource_Read_NotFound(t *testing.T) {
	r := &GetEmsPreferencesDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetEmsPreferencesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetEmsPreferencesDataSource_Read_APIError exercises GetEmsPreferencesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetEmsPreferencesDataSource_Read_APIError(t *testing.T) {
	r := &GetEmsPreferencesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetEmsPreferencesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ems_preferences")
}

// TestGetEmsPreferencesDataSource_Read_APIErrorReadBody exercises GetEmsPreferencesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetEmsPreferencesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetEmsPreferencesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetEmsPreferencesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetEmsPreferencesDataSource_Read_InvalidJSON exercises GetEmsPreferencesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetEmsPreferencesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEmsPreferencesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEmsPreferencesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
