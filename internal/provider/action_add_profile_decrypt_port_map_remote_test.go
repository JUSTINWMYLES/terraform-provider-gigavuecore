package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddProfileDecryptPortMapAction_Invoke_Happy exercises AddProfileDecryptPortMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddProfileDecryptPortMapAction_Invoke_Happy(t *testing.T) {
	r := &AddProfileDecryptPortMapAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddProfileDecryptPortMapAction_Invoke_NilClient exercises AddProfileDecryptPortMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddProfileDecryptPortMapAction_Invoke_NilClient(t *testing.T) {
	r := &AddProfileDecryptPortMapAction{}
	m := AddProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddProfileDecryptPortMapAction_Invoke_BuildError exercises AddProfileDecryptPortMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddProfileDecryptPortMapAction_Invoke_BuildError(t *testing.T) {
	r := &AddProfileDecryptPortMapAction{client: newMalformedBaseURLClient(t)}
	m := AddProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddProfileDecryptPortMapAction_Invoke_SendError exercises AddProfileDecryptPortMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddProfileDecryptPortMapAction_Invoke_SendError(t *testing.T) {
	r := &AddProfileDecryptPortMapAction{client: newTransportErrorClient(t)}
	m := AddProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddProfileDecryptPortMapAction_Invoke_APIError exercises AddProfileDecryptPortMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddProfileDecryptPortMapAction_Invoke_APIError(t *testing.T) {
	r := &AddProfileDecryptPortMapAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_profile_decrypt_port_map")
}

// TestAddProfileDecryptPortMapAction_Invoke_APIErrorReadBody exercises AddProfileDecryptPortMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddProfileDecryptPortMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddProfileDecryptPortMapAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddProfileDecryptPortMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
