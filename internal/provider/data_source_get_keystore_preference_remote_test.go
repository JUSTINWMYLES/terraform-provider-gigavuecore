package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetKeystorePreferenceDataSource_Read_Happy exercises GetKeystorePreferenceDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetKeystorePreferenceDataSource_Read_Happy(t *testing.T) {
	r := &GetKeystorePreferenceDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetKeystorePreferenceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetKeystorePreferenceDataSource_Read_NilClient exercises GetKeystorePreferenceDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetKeystorePreferenceDataSource_Read_NilClient(t *testing.T) {
	r := &GetKeystorePreferenceDataSource{}
	m := GetKeystorePreferenceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetKeystorePreferenceDataSource_Read_BuildError exercises GetKeystorePreferenceDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetKeystorePreferenceDataSource_Read_BuildError(t *testing.T) {
	r := &GetKeystorePreferenceDataSource{client: newMalformedBaseURLClient(t)}
	m := GetKeystorePreferenceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetKeystorePreferenceDataSource_Read_SendError exercises GetKeystorePreferenceDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetKeystorePreferenceDataSource_Read_SendError(t *testing.T) {
	r := &GetKeystorePreferenceDataSource{client: newTransportErrorClient(t)}
	m := GetKeystorePreferenceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetKeystorePreferenceDataSource_Read_NotFound exercises GetKeystorePreferenceDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetKeystorePreferenceDataSource_Read_NotFound(t *testing.T) {
	r := &GetKeystorePreferenceDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetKeystorePreferenceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetKeystorePreferenceDataSource_Read_APIError exercises GetKeystorePreferenceDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetKeystorePreferenceDataSource_Read_APIError(t *testing.T) {
	r := &GetKeystorePreferenceDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetKeystorePreferenceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_keystore_preference")
}

// TestGetKeystorePreferenceDataSource_Read_APIErrorReadBody exercises GetKeystorePreferenceDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetKeystorePreferenceDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetKeystorePreferenceDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetKeystorePreferenceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetKeystorePreferenceDataSource_Read_InvalidJSON exercises GetKeystorePreferenceDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetKeystorePreferenceDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetKeystorePreferenceDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetKeystorePreferenceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
