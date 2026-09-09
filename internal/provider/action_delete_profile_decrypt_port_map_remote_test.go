package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteProfileDecryptPortMapAction_Invoke_Happy exercises DeleteProfileDecryptPortMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteProfileDecryptPortMapAction_Invoke_Happy(t *testing.T) {
	r := &DeleteProfileDecryptPortMapAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteProfileDecryptPortMapAction_Invoke_NilClient exercises DeleteProfileDecryptPortMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteProfileDecryptPortMapAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteProfileDecryptPortMapAction{}
	m := DeleteProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteProfileDecryptPortMapAction_Invoke_BuildError exercises DeleteProfileDecryptPortMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteProfileDecryptPortMapAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteProfileDecryptPortMapAction{client: newMalformedBaseURLClient(t)}
	m := DeleteProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteProfileDecryptPortMapAction_Invoke_SendError exercises DeleteProfileDecryptPortMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteProfileDecryptPortMapAction_Invoke_SendError(t *testing.T) {
	r := &DeleteProfileDecryptPortMapAction{client: newTransportErrorClient(t)}
	m := DeleteProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteProfileDecryptPortMapAction_Invoke_APIError exercises DeleteProfileDecryptPortMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteProfileDecryptPortMapAction_Invoke_APIError(t *testing.T) {
	r := &DeleteProfileDecryptPortMapAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_profile_decrypt_port_map")
}

// TestDeleteProfileDecryptPortMapAction_Invoke_APIErrorReadBody exercises DeleteProfileDecryptPortMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteProfileDecryptPortMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteProfileDecryptPortMapAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
