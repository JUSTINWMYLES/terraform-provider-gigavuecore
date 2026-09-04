package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestFetchKeyHandlerRemoteAction_Invoke_Happy exercises FetchKeyHandlerRemoteAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestFetchKeyHandlerRemoteAction_Invoke_Happy(t *testing.T) {
	r := &FetchKeyHandlerRemoteAction{client: newMockClientStatus(t, 201, "{}")}
	m := FetchKeyHandlerRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFetchKeyHandlerRemoteAction_Invoke_NilClient exercises FetchKeyHandlerRemoteAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFetchKeyHandlerRemoteAction_Invoke_NilClient(t *testing.T) {
	r := &FetchKeyHandlerRemoteAction{}
	m := FetchKeyHandlerRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFetchKeyHandlerRemoteAction_Invoke_BuildError exercises FetchKeyHandlerRemoteAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFetchKeyHandlerRemoteAction_Invoke_BuildError(t *testing.T) {
	r := &FetchKeyHandlerRemoteAction{client: newMalformedBaseURLClient(t)}
	m := FetchKeyHandlerRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFetchKeyHandlerRemoteAction_Invoke_SendError exercises FetchKeyHandlerRemoteAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestFetchKeyHandlerRemoteAction_Invoke_SendError(t *testing.T) {
	r := &FetchKeyHandlerRemoteAction{client: newTransportErrorClient(t)}
	m := FetchKeyHandlerRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFetchKeyHandlerRemoteAction_Invoke_APIError exercises FetchKeyHandlerRemoteAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFetchKeyHandlerRemoteAction_Invoke_APIError(t *testing.T) {
	r := &FetchKeyHandlerRemoteAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FetchKeyHandlerRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_fetch_key_handler_remote")
}

// TestFetchKeyHandlerRemoteAction_Invoke_APIErrorReadBody exercises FetchKeyHandlerRemoteAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFetchKeyHandlerRemoteAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &FetchKeyHandlerRemoteAction{client: newMockClientReadErrorBody(t, 501)}
	m := FetchKeyHandlerRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
