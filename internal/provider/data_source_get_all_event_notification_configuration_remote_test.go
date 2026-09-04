package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllEventNotificationConfigurationDataSource_Read_Happy exercises GetAllEventNotificationConfigurationDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAllEventNotificationConfigurationDataSource_Read_Happy(t *testing.T) {
	r := &GetAllEventNotificationConfigurationDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAllEventNotificationConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllEventNotificationConfigurationDataSource_Read_NilClient exercises GetAllEventNotificationConfigurationDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllEventNotificationConfigurationDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllEventNotificationConfigurationDataSource{}
	m := GetAllEventNotificationConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllEventNotificationConfigurationDataSource_Read_BuildError exercises GetAllEventNotificationConfigurationDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAllEventNotificationConfigurationDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllEventNotificationConfigurationDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllEventNotificationConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAllEventNotificationConfigurationDataSource_Read_SendError exercises GetAllEventNotificationConfigurationDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAllEventNotificationConfigurationDataSource_Read_SendError(t *testing.T) {
	r := &GetAllEventNotificationConfigurationDataSource{client: newTransportErrorClient(t)}
	m := GetAllEventNotificationConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAllEventNotificationConfigurationDataSource_Read_NotFound exercises GetAllEventNotificationConfigurationDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAllEventNotificationConfigurationDataSource_Read_NotFound(t *testing.T) {
	r := &GetAllEventNotificationConfigurationDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAllEventNotificationConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAllEventNotificationConfigurationDataSource_Read_APIError exercises GetAllEventNotificationConfigurationDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAllEventNotificationConfigurationDataSource_Read_APIError(t *testing.T) {
	r := &GetAllEventNotificationConfigurationDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAllEventNotificationConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_all_event_notification_configuration")
}

// TestGetAllEventNotificationConfigurationDataSource_Read_APIErrorReadBody exercises GetAllEventNotificationConfigurationDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAllEventNotificationConfigurationDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAllEventNotificationConfigurationDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAllEventNotificationConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAllEventNotificationConfigurationDataSource_Read_InvalidJSON exercises GetAllEventNotificationConfigurationDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAllEventNotificationConfigurationDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllEventNotificationConfigurationDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllEventNotificationConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
