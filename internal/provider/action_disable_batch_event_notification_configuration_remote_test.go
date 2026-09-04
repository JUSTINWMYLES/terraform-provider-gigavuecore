package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDisableBatchEventNotificationConfigurationAction_Invoke_Happy exercises DisableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDisableBatchEventNotificationConfigurationAction_Invoke_Happy(t *testing.T) {
	r := &DisableBatchEventNotificationConfigurationAction{client: newMockClientStatus(t, 200, "{}")}
	m := DisableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDisableBatchEventNotificationConfigurationAction_Invoke_NilClient exercises DisableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDisableBatchEventNotificationConfigurationAction_Invoke_NilClient(t *testing.T) {
	r := &DisableBatchEventNotificationConfigurationAction{}
	m := DisableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDisableBatchEventNotificationConfigurationAction_Invoke_BuildError exercises DisableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDisableBatchEventNotificationConfigurationAction_Invoke_BuildError(t *testing.T) {
	r := &DisableBatchEventNotificationConfigurationAction{client: newMalformedBaseURLClient(t)}
	m := DisableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDisableBatchEventNotificationConfigurationAction_Invoke_SendError exercises DisableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDisableBatchEventNotificationConfigurationAction_Invoke_SendError(t *testing.T) {
	r := &DisableBatchEventNotificationConfigurationAction{client: newTransportErrorClient(t)}
	m := DisableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDisableBatchEventNotificationConfigurationAction_Invoke_APIError exercises DisableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDisableBatchEventNotificationConfigurationAction_Invoke_APIError(t *testing.T) {
	r := &DisableBatchEventNotificationConfigurationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DisableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_disable_batch_event_notification_configuration")
}

// TestDisableBatchEventNotificationConfigurationAction_Invoke_APIErrorReadBody exercises DisableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDisableBatchEventNotificationConfigurationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DisableBatchEventNotificationConfigurationAction{client: newMockClientReadErrorBody(t, 501)}
	m := DisableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
