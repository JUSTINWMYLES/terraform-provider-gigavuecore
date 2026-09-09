package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestLeaveHaInstanceAction_Invoke_Happy exercises LeaveHaInstanceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestLeaveHaInstanceAction_Invoke_Happy(t *testing.T) {
	r := &LeaveHaInstanceAction{client: newMockClientStatus(t, 201, "{}")}
	m := LeaveHaInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLeaveHaInstanceAction_Invoke_NilClient exercises LeaveHaInstanceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLeaveHaInstanceAction_Invoke_NilClient(t *testing.T) {
	r := &LeaveHaInstanceAction{}
	m := LeaveHaInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLeaveHaInstanceAction_Invoke_BuildError exercises LeaveHaInstanceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLeaveHaInstanceAction_Invoke_BuildError(t *testing.T) {
	r := &LeaveHaInstanceAction{client: newMalformedBaseURLClient(t)}
	m := LeaveHaInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLeaveHaInstanceAction_Invoke_SendError exercises LeaveHaInstanceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestLeaveHaInstanceAction_Invoke_SendError(t *testing.T) {
	r := &LeaveHaInstanceAction{client: newTransportErrorClient(t)}
	m := LeaveHaInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLeaveHaInstanceAction_Invoke_APIError exercises LeaveHaInstanceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLeaveHaInstanceAction_Invoke_APIError(t *testing.T) {
	r := &LeaveHaInstanceAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LeaveHaInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_leave_ha_instance")
}

// TestLeaveHaInstanceAction_Invoke_APIErrorReadBody exercises LeaveHaInstanceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLeaveHaInstanceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &LeaveHaInstanceAction{client: newMockClientReadErrorBody(t, 501)}
	m := LeaveHaInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
