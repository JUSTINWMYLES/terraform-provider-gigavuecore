package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateEventNotificationConfigurationAction_Invoke_Happy exercises CreateEventNotificationConfigurationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateEventNotificationConfigurationAction_Invoke_Happy(t *testing.T) {
	r := &CreateEventNotificationConfigurationAction{client: newMockClientStatus(t, 200, "{}")}
	m := CreateEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateEventNotificationConfigurationAction_Invoke_NilClient exercises CreateEventNotificationConfigurationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateEventNotificationConfigurationAction_Invoke_NilClient(t *testing.T) {
	r := &CreateEventNotificationConfigurationAction{}
	m := CreateEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateEventNotificationConfigurationAction_Invoke_BuildError exercises CreateEventNotificationConfigurationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateEventNotificationConfigurationAction_Invoke_BuildError(t *testing.T) {
	r := &CreateEventNotificationConfigurationAction{client: newMalformedBaseURLClient(t)}
	m := CreateEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateEventNotificationConfigurationAction_Invoke_SendError exercises CreateEventNotificationConfigurationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateEventNotificationConfigurationAction_Invoke_SendError(t *testing.T) {
	r := &CreateEventNotificationConfigurationAction{client: newTransportErrorClient(t)}
	m := CreateEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateEventNotificationConfigurationAction_Invoke_APIError exercises CreateEventNotificationConfigurationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateEventNotificationConfigurationAction_Invoke_APIError(t *testing.T) {
	r := &CreateEventNotificationConfigurationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_event_notification_configuration")
}

// TestCreateEventNotificationConfigurationAction_Invoke_APIErrorReadBody exercises CreateEventNotificationConfigurationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateEventNotificationConfigurationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateEventNotificationConfigurationAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
