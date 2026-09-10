package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestEnableBatchEventNotificationConfigurationAction_Invoke_Happy exercises EnableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestEnableBatchEventNotificationConfigurationAction_Invoke_Happy(t *testing.T) {
	r := &EnableBatchEventNotificationConfigurationAction{client: newMockClientStatus(t, 200, "{}")}
	m := EnableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnableBatchEventNotificationConfigurationAction_Invoke_NilClient exercises EnableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEnableBatchEventNotificationConfigurationAction_Invoke_NilClient(t *testing.T) {
	r := &EnableBatchEventNotificationConfigurationAction{}
	m := EnableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEnableBatchEventNotificationConfigurationAction_Invoke_BuildError exercises EnableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEnableBatchEventNotificationConfigurationAction_Invoke_BuildError(t *testing.T) {
	r := &EnableBatchEventNotificationConfigurationAction{client: newMalformedBaseURLClient(t)}
	m := EnableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEnableBatchEventNotificationConfigurationAction_Invoke_SendError exercises EnableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestEnableBatchEventNotificationConfigurationAction_Invoke_SendError(t *testing.T) {
	r := &EnableBatchEventNotificationConfigurationAction{client: newTransportErrorClient(t)}
	m := EnableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEnableBatchEventNotificationConfigurationAction_Invoke_APIError exercises EnableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEnableBatchEventNotificationConfigurationAction_Invoke_APIError(t *testing.T) {
	r := &EnableBatchEventNotificationConfigurationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EnableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_enable_batch_event_notification_configuration")
}

// TestEnableBatchEventNotificationConfigurationAction_Invoke_APIErrorReadBody exercises EnableBatchEventNotificationConfigurationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEnableBatchEventNotificationConfigurationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &EnableBatchEventNotificationConfigurationAction{client: newMockClientReadErrorBody(t, 501)}
	m := EnableBatchEventNotificationConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
