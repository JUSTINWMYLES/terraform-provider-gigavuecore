package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReplaceProfileKeyMapsAction_Invoke_Happy exercises ReplaceProfileKeyMapsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReplaceProfileKeyMapsAction_Invoke_Happy(t *testing.T) {
	r := &ReplaceProfileKeyMapsAction{client: newMockClientStatus(t, 200, "{}")}
	m := ReplaceProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReplaceProfileKeyMapsAction_Invoke_NilClient exercises ReplaceProfileKeyMapsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReplaceProfileKeyMapsAction_Invoke_NilClient(t *testing.T) {
	r := &ReplaceProfileKeyMapsAction{}
	m := ReplaceProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReplaceProfileKeyMapsAction_Invoke_BuildError exercises ReplaceProfileKeyMapsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReplaceProfileKeyMapsAction_Invoke_BuildError(t *testing.T) {
	r := &ReplaceProfileKeyMapsAction{client: newMalformedBaseURLClient(t)}
	m := ReplaceProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReplaceProfileKeyMapsAction_Invoke_SendError exercises ReplaceProfileKeyMapsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReplaceProfileKeyMapsAction_Invoke_SendError(t *testing.T) {
	r := &ReplaceProfileKeyMapsAction{client: newTransportErrorClient(t)}
	m := ReplaceProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReplaceProfileKeyMapsAction_Invoke_APIError exercises ReplaceProfileKeyMapsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReplaceProfileKeyMapsAction_Invoke_APIError(t *testing.T) {
	r := &ReplaceProfileKeyMapsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReplaceProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_replace_profile_key_maps")
}

// TestReplaceProfileKeyMapsAction_Invoke_APIErrorReadBody exercises ReplaceProfileKeyMapsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReplaceProfileKeyMapsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReplaceProfileKeyMapsAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReplaceProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
