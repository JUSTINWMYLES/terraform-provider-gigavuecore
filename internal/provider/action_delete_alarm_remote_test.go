package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAlarmAction_Invoke_Happy exercises DeleteAlarmAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAlarmAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAlarmAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAlarmAction_Invoke_NilClient exercises DeleteAlarmAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAlarmAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAlarmAction{}
	m := DeleteAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAlarmAction_Invoke_BuildError exercises DeleteAlarmAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAlarmAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAlarmAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAlarmAction_Invoke_SendError exercises DeleteAlarmAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAlarmAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAlarmAction{client: newTransportErrorClient(t)}
	m := DeleteAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAlarmAction_Invoke_APIError exercises DeleteAlarmAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAlarmAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAlarmAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_alarm")
}

// TestDeleteAlarmAction_Invoke_APIErrorReadBody exercises DeleteAlarmAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAlarmAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAlarmAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
