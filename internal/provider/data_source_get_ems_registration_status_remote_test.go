package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEmsRegistrationStatusDataSource_Read_Happy exercises GetEmsRegistrationStatusDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetEmsRegistrationStatusDataSource_Read_Happy(t *testing.T) {
	r := &GetEmsRegistrationStatusDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetEmsRegistrationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEmsRegistrationStatusDataSource_Read_NilClient exercises GetEmsRegistrationStatusDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEmsRegistrationStatusDataSource_Read_NilClient(t *testing.T) {
	r := &GetEmsRegistrationStatusDataSource{}
	m := GetEmsRegistrationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEmsRegistrationStatusDataSource_Read_BuildError exercises GetEmsRegistrationStatusDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetEmsRegistrationStatusDataSource_Read_BuildError(t *testing.T) {
	r := &GetEmsRegistrationStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEmsRegistrationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetEmsRegistrationStatusDataSource_Read_SendError exercises GetEmsRegistrationStatusDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetEmsRegistrationStatusDataSource_Read_SendError(t *testing.T) {
	r := &GetEmsRegistrationStatusDataSource{client: newTransportErrorClient(t)}
	m := GetEmsRegistrationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetEmsRegistrationStatusDataSource_Read_NotFound exercises GetEmsRegistrationStatusDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetEmsRegistrationStatusDataSource_Read_NotFound(t *testing.T) {
	r := &GetEmsRegistrationStatusDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetEmsRegistrationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetEmsRegistrationStatusDataSource_Read_APIError exercises GetEmsRegistrationStatusDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetEmsRegistrationStatusDataSource_Read_APIError(t *testing.T) {
	r := &GetEmsRegistrationStatusDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetEmsRegistrationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ems_registration_status")
}

// TestGetEmsRegistrationStatusDataSource_Read_APIErrorReadBody exercises GetEmsRegistrationStatusDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetEmsRegistrationStatusDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetEmsRegistrationStatusDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetEmsRegistrationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetEmsRegistrationStatusDataSource_Read_InvalidJSON exercises GetEmsRegistrationStatusDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetEmsRegistrationStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEmsRegistrationStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEmsRegistrationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
