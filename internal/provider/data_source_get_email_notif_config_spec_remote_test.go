package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEmailNotifConfigSpecDataSource_Read_Happy exercises GetEmailNotifConfigSpecDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetEmailNotifConfigSpecDataSource_Read_Happy(t *testing.T) {
	r := &GetEmailNotifConfigSpecDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetEmailNotifConfigSpecDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEmailNotifConfigSpecDataSource_Read_NilClient exercises GetEmailNotifConfigSpecDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEmailNotifConfigSpecDataSource_Read_NilClient(t *testing.T) {
	r := &GetEmailNotifConfigSpecDataSource{}
	m := GetEmailNotifConfigSpecDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEmailNotifConfigSpecDataSource_Read_BuildError exercises GetEmailNotifConfigSpecDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetEmailNotifConfigSpecDataSource_Read_BuildError(t *testing.T) {
	r := &GetEmailNotifConfigSpecDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEmailNotifConfigSpecDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetEmailNotifConfigSpecDataSource_Read_SendError exercises GetEmailNotifConfigSpecDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetEmailNotifConfigSpecDataSource_Read_SendError(t *testing.T) {
	r := &GetEmailNotifConfigSpecDataSource{client: newTransportErrorClient(t)}
	m := GetEmailNotifConfigSpecDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetEmailNotifConfigSpecDataSource_Read_NotFound exercises GetEmailNotifConfigSpecDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetEmailNotifConfigSpecDataSource_Read_NotFound(t *testing.T) {
	r := &GetEmailNotifConfigSpecDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetEmailNotifConfigSpecDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetEmailNotifConfigSpecDataSource_Read_APIError exercises GetEmailNotifConfigSpecDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetEmailNotifConfigSpecDataSource_Read_APIError(t *testing.T) {
	r := &GetEmailNotifConfigSpecDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetEmailNotifConfigSpecDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_email_notif_config_spec")
}

// TestGetEmailNotifConfigSpecDataSource_Read_APIErrorReadBody exercises GetEmailNotifConfigSpecDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetEmailNotifConfigSpecDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetEmailNotifConfigSpecDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetEmailNotifConfigSpecDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetEmailNotifConfigSpecDataSource_Read_InvalidJSON exercises GetEmailNotifConfigSpecDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetEmailNotifConfigSpecDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEmailNotifConfigSpecDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEmailNotifConfigSpecDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
