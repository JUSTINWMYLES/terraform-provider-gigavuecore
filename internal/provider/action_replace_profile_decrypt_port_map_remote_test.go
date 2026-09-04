package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReplaceProfileDecryptPortMapAction_Invoke_Happy exercises ReplaceProfileDecryptPortMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReplaceProfileDecryptPortMapAction_Invoke_Happy(t *testing.T) {
	r := &ReplaceProfileDecryptPortMapAction{client: newMockClientStatus(t, 200, "{}")}
	m := ReplaceProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReplaceProfileDecryptPortMapAction_Invoke_NilClient exercises ReplaceProfileDecryptPortMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReplaceProfileDecryptPortMapAction_Invoke_NilClient(t *testing.T) {
	r := &ReplaceProfileDecryptPortMapAction{}
	m := ReplaceProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReplaceProfileDecryptPortMapAction_Invoke_BuildError exercises ReplaceProfileDecryptPortMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReplaceProfileDecryptPortMapAction_Invoke_BuildError(t *testing.T) {
	r := &ReplaceProfileDecryptPortMapAction{client: newMalformedBaseURLClient(t)}
	m := ReplaceProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReplaceProfileDecryptPortMapAction_Invoke_SendError exercises ReplaceProfileDecryptPortMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReplaceProfileDecryptPortMapAction_Invoke_SendError(t *testing.T) {
	r := &ReplaceProfileDecryptPortMapAction{client: newTransportErrorClient(t)}
	m := ReplaceProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReplaceProfileDecryptPortMapAction_Invoke_APIError exercises ReplaceProfileDecryptPortMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReplaceProfileDecryptPortMapAction_Invoke_APIError(t *testing.T) {
	r := &ReplaceProfileDecryptPortMapAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReplaceProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_replace_profile_decrypt_port_map")
}

// TestReplaceProfileDecryptPortMapAction_Invoke_APIErrorReadBody exercises ReplaceProfileDecryptPortMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReplaceProfileDecryptPortMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReplaceProfileDecryptPortMapAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReplaceProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
