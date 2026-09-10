package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddAvPolicyActionAction_Invoke_Happy exercises AddAvPolicyActionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddAvPolicyActionAction_Invoke_Happy(t *testing.T) {
	r := &AddAvPolicyActionAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddAvPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddAvPolicyActionAction_Invoke_NilClient exercises AddAvPolicyActionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddAvPolicyActionAction_Invoke_NilClient(t *testing.T) {
	r := &AddAvPolicyActionAction{}
	m := AddAvPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddAvPolicyActionAction_Invoke_BuildError exercises AddAvPolicyActionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddAvPolicyActionAction_Invoke_BuildError(t *testing.T) {
	r := &AddAvPolicyActionAction{client: newMalformedBaseURLClient(t)}
	m := AddAvPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddAvPolicyActionAction_Invoke_SendError exercises AddAvPolicyActionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddAvPolicyActionAction_Invoke_SendError(t *testing.T) {
	r := &AddAvPolicyActionAction{client: newTransportErrorClient(t)}
	m := AddAvPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddAvPolicyActionAction_Invoke_APIError exercises AddAvPolicyActionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddAvPolicyActionAction_Invoke_APIError(t *testing.T) {
	r := &AddAvPolicyActionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddAvPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_av_policy_action")
}

// TestAddAvPolicyActionAction_Invoke_APIErrorReadBody exercises AddAvPolicyActionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddAvPolicyActionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddAvPolicyActionAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddAvPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
