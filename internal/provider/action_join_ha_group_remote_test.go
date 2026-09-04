package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestJoinHaGroupAction_Invoke_Happy exercises JoinHaGroupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestJoinHaGroupAction_Invoke_Happy(t *testing.T) {
	r := &JoinHaGroupAction{client: newMockClientStatus(t, 201, "{}")}
	m := JoinHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestJoinHaGroupAction_Invoke_NilClient exercises JoinHaGroupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestJoinHaGroupAction_Invoke_NilClient(t *testing.T) {
	r := &JoinHaGroupAction{}
	m := JoinHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestJoinHaGroupAction_Invoke_BuildError exercises JoinHaGroupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestJoinHaGroupAction_Invoke_BuildError(t *testing.T) {
	r := &JoinHaGroupAction{client: newMalformedBaseURLClient(t)}
	m := JoinHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestJoinHaGroupAction_Invoke_SendError exercises JoinHaGroupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestJoinHaGroupAction_Invoke_SendError(t *testing.T) {
	r := &JoinHaGroupAction{client: newTransportErrorClient(t)}
	m := JoinHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestJoinHaGroupAction_Invoke_APIError exercises JoinHaGroupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestJoinHaGroupAction_Invoke_APIError(t *testing.T) {
	r := &JoinHaGroupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := JoinHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_join_ha_group")
}

// TestJoinHaGroupAction_Invoke_APIErrorReadBody exercises JoinHaGroupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestJoinHaGroupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &JoinHaGroupAction{client: newMockClientReadErrorBody(t, 501)}
	m := JoinHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
