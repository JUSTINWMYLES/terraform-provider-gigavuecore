package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddProfileKeyMapAction_Invoke_Happy exercises AddProfileKeyMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddProfileKeyMapAction_Invoke_Happy(t *testing.T) {
	r := &AddProfileKeyMapAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddProfileKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddProfileKeyMapAction_Invoke_NilClient exercises AddProfileKeyMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddProfileKeyMapAction_Invoke_NilClient(t *testing.T) {
	r := &AddProfileKeyMapAction{}
	m := AddProfileKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddProfileKeyMapAction_Invoke_BuildError exercises AddProfileKeyMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddProfileKeyMapAction_Invoke_BuildError(t *testing.T) {
	r := &AddProfileKeyMapAction{client: newMalformedBaseURLClient(t)}
	m := AddProfileKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddProfileKeyMapAction_Invoke_SendError exercises AddProfileKeyMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddProfileKeyMapAction_Invoke_SendError(t *testing.T) {
	r := &AddProfileKeyMapAction{client: newTransportErrorClient(t)}
	m := AddProfileKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddProfileKeyMapAction_Invoke_APIError exercises AddProfileKeyMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddProfileKeyMapAction_Invoke_APIError(t *testing.T) {
	r := &AddProfileKeyMapAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddProfileKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_profile_key_map")
}

// TestAddProfileKeyMapAction_Invoke_APIErrorReadBody exercises AddProfileKeyMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddProfileKeyMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddProfileKeyMapAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddProfileKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
