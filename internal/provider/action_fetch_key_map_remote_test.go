package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestFetchKeyMapAction_Invoke_Happy exercises FetchKeyMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestFetchKeyMapAction_Invoke_Happy(t *testing.T) {
	r := &FetchKeyMapAction{client: newMockClientStatus(t, 201, "{}")}
	m := FetchKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFetchKeyMapAction_Invoke_NilClient exercises FetchKeyMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFetchKeyMapAction_Invoke_NilClient(t *testing.T) {
	r := &FetchKeyMapAction{}
	m := FetchKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFetchKeyMapAction_Invoke_BuildError exercises FetchKeyMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFetchKeyMapAction_Invoke_BuildError(t *testing.T) {
	r := &FetchKeyMapAction{client: newMalformedBaseURLClient(t)}
	m := FetchKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFetchKeyMapAction_Invoke_SendError exercises FetchKeyMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestFetchKeyMapAction_Invoke_SendError(t *testing.T) {
	r := &FetchKeyMapAction{client: newTransportErrorClient(t)}
	m := FetchKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFetchKeyMapAction_Invoke_APIError exercises FetchKeyMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFetchKeyMapAction_Invoke_APIError(t *testing.T) {
	r := &FetchKeyMapAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FetchKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_fetch_key_map")
}

// TestFetchKeyMapAction_Invoke_APIErrorReadBody exercises FetchKeyMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFetchKeyMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &FetchKeyMapAction{client: newMockClientReadErrorBody(t, 501)}
	m := FetchKeyMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
