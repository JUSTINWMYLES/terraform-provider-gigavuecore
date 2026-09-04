package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetKeystoreCountersDataSource_Read_Happy exercises GetKeystoreCountersDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetKeystoreCountersDataSource_Read_Happy(t *testing.T) {
	r := &GetKeystoreCountersDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetKeystoreCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetKeystoreCountersDataSource_Read_NilClient exercises GetKeystoreCountersDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetKeystoreCountersDataSource_Read_NilClient(t *testing.T) {
	r := &GetKeystoreCountersDataSource{}
	m := GetKeystoreCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetKeystoreCountersDataSource_Read_BuildError exercises GetKeystoreCountersDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetKeystoreCountersDataSource_Read_BuildError(t *testing.T) {
	r := &GetKeystoreCountersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetKeystoreCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetKeystoreCountersDataSource_Read_SendError exercises GetKeystoreCountersDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetKeystoreCountersDataSource_Read_SendError(t *testing.T) {
	r := &GetKeystoreCountersDataSource{client: newTransportErrorClient(t)}
	m := GetKeystoreCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetKeystoreCountersDataSource_Read_NotFound exercises GetKeystoreCountersDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetKeystoreCountersDataSource_Read_NotFound(t *testing.T) {
	r := &GetKeystoreCountersDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetKeystoreCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetKeystoreCountersDataSource_Read_APIError exercises GetKeystoreCountersDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetKeystoreCountersDataSource_Read_APIError(t *testing.T) {
	r := &GetKeystoreCountersDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetKeystoreCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_keystore_counters")
}

// TestGetKeystoreCountersDataSource_Read_APIErrorReadBody exercises GetKeystoreCountersDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetKeystoreCountersDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetKeystoreCountersDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetKeystoreCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetKeystoreCountersDataSource_Read_InvalidJSON exercises GetKeystoreCountersDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetKeystoreCountersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetKeystoreCountersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetKeystoreCountersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
