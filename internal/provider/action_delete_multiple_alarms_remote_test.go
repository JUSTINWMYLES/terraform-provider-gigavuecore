package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteMultipleAlarmsAction_Invoke_Happy exercises DeleteMultipleAlarmsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteMultipleAlarmsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteMultipleAlarmsAction{client: newMockClientStatus(t, 200, "{}")}
	m := DeleteMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteMultipleAlarmsAction_Invoke_NilClient exercises DeleteMultipleAlarmsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteMultipleAlarmsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteMultipleAlarmsAction{}
	m := DeleteMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteMultipleAlarmsAction_Invoke_BuildError exercises DeleteMultipleAlarmsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteMultipleAlarmsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteMultipleAlarmsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteMultipleAlarmsAction_Invoke_SendError exercises DeleteMultipleAlarmsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteMultipleAlarmsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteMultipleAlarmsAction{client: newTransportErrorClient(t)}
	m := DeleteMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteMultipleAlarmsAction_Invoke_APIError exercises DeleteMultipleAlarmsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteMultipleAlarmsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteMultipleAlarmsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_multiple_alarms")
}

// TestDeleteMultipleAlarmsAction_Invoke_APIErrorReadBody exercises DeleteMultipleAlarmsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteMultipleAlarmsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteMultipleAlarmsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
