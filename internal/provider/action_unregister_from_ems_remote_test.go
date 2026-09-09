package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUnregisterFromEmsAction_Invoke_Happy exercises UnregisterFromEmsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUnregisterFromEmsAction_Invoke_Happy(t *testing.T) {
	r := &UnregisterFromEmsAction{client: newMockClientStatus(t, 200, "{}")}
	m := UnregisterFromEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUnregisterFromEmsAction_Invoke_NilClient exercises UnregisterFromEmsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUnregisterFromEmsAction_Invoke_NilClient(t *testing.T) {
	r := &UnregisterFromEmsAction{}
	m := UnregisterFromEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUnregisterFromEmsAction_Invoke_BuildError exercises UnregisterFromEmsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUnregisterFromEmsAction_Invoke_BuildError(t *testing.T) {
	r := &UnregisterFromEmsAction{client: newMalformedBaseURLClient(t)}
	m := UnregisterFromEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUnregisterFromEmsAction_Invoke_SendError exercises UnregisterFromEmsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUnregisterFromEmsAction_Invoke_SendError(t *testing.T) {
	r := &UnregisterFromEmsAction{client: newTransportErrorClient(t)}
	m := UnregisterFromEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUnregisterFromEmsAction_Invoke_APIError exercises UnregisterFromEmsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUnregisterFromEmsAction_Invoke_APIError(t *testing.T) {
	r := &UnregisterFromEmsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UnregisterFromEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_unregister_from_ems")
}

// TestUnregisterFromEmsAction_Invoke_APIErrorReadBody exercises UnregisterFromEmsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUnregisterFromEmsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UnregisterFromEmsAction{client: newMockClientReadErrorBody(t, 501)}
	m := UnregisterFromEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
