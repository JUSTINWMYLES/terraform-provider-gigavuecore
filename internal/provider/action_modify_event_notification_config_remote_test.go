package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestModifyEventNotificationConfigAction_Invoke_Happy exercises ModifyEventNotificationConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestModifyEventNotificationConfigAction_Invoke_Happy(t *testing.T) {
	r := &ModifyEventNotificationConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := ModifyEventNotificationConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestModifyEventNotificationConfigAction_Invoke_NilClient exercises ModifyEventNotificationConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestModifyEventNotificationConfigAction_Invoke_NilClient(t *testing.T) {
	r := &ModifyEventNotificationConfigAction{}
	m := ModifyEventNotificationConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestModifyEventNotificationConfigAction_Invoke_BuildError exercises ModifyEventNotificationConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestModifyEventNotificationConfigAction_Invoke_BuildError(t *testing.T) {
	r := &ModifyEventNotificationConfigAction{client: newMalformedBaseURLClient(t)}
	m := ModifyEventNotificationConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestModifyEventNotificationConfigAction_Invoke_SendError exercises ModifyEventNotificationConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestModifyEventNotificationConfigAction_Invoke_SendError(t *testing.T) {
	r := &ModifyEventNotificationConfigAction{client: newTransportErrorClient(t)}
	m := ModifyEventNotificationConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestModifyEventNotificationConfigAction_Invoke_APIError exercises ModifyEventNotificationConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestModifyEventNotificationConfigAction_Invoke_APIError(t *testing.T) {
	r := &ModifyEventNotificationConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ModifyEventNotificationConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_modify_event_notification_config")
}

// TestModifyEventNotificationConfigAction_Invoke_APIErrorReadBody exercises ModifyEventNotificationConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestModifyEventNotificationConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ModifyEventNotificationConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := ModifyEventNotificationConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
