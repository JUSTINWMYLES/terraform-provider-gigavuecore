package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadClusterEventNotificationStatusDataSource_Read_Happy exercises LoadClusterEventNotificationStatusDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadClusterEventNotificationStatusDataSource_Read_Happy(t *testing.T) {
	r := &LoadClusterEventNotificationStatusDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadClusterEventNotificationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadClusterEventNotificationStatusDataSource_Read_NilClient exercises LoadClusterEventNotificationStatusDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadClusterEventNotificationStatusDataSource_Read_NilClient(t *testing.T) {
	r := &LoadClusterEventNotificationStatusDataSource{}
	m := LoadClusterEventNotificationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadClusterEventNotificationStatusDataSource_Read_BuildError exercises LoadClusterEventNotificationStatusDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadClusterEventNotificationStatusDataSource_Read_BuildError(t *testing.T) {
	r := &LoadClusterEventNotificationStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadClusterEventNotificationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadClusterEventNotificationStatusDataSource_Read_SendError exercises LoadClusterEventNotificationStatusDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadClusterEventNotificationStatusDataSource_Read_SendError(t *testing.T) {
	r := &LoadClusterEventNotificationStatusDataSource{client: newTransportErrorClient(t)}
	m := LoadClusterEventNotificationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadClusterEventNotificationStatusDataSource_Read_NotFound exercises LoadClusterEventNotificationStatusDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadClusterEventNotificationStatusDataSource_Read_NotFound(t *testing.T) {
	r := &LoadClusterEventNotificationStatusDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadClusterEventNotificationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadClusterEventNotificationStatusDataSource_Read_APIError exercises LoadClusterEventNotificationStatusDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadClusterEventNotificationStatusDataSource_Read_APIError(t *testing.T) {
	r := &LoadClusterEventNotificationStatusDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadClusterEventNotificationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_cluster_event_notification_status")
}

// TestLoadClusterEventNotificationStatusDataSource_Read_APIErrorReadBody exercises LoadClusterEventNotificationStatusDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadClusterEventNotificationStatusDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadClusterEventNotificationStatusDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadClusterEventNotificationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadClusterEventNotificationStatusDataSource_Read_InvalidJSON exercises LoadClusterEventNotificationStatusDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadClusterEventNotificationStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadClusterEventNotificationStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadClusterEventNotificationStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
