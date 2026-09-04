package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestPurgeEventAction_Invoke_Happy exercises PurgeEventAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestPurgeEventAction_Invoke_Happy(t *testing.T) {
	r := &PurgeEventAction{client: newMockClientStatus(t, 204, "{}")}
	m := PurgeEventActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPurgeEventAction_Invoke_NilClient exercises PurgeEventAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPurgeEventAction_Invoke_NilClient(t *testing.T) {
	r := &PurgeEventAction{}
	m := PurgeEventActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPurgeEventAction_Invoke_BuildError exercises PurgeEventAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPurgeEventAction_Invoke_BuildError(t *testing.T) {
	r := &PurgeEventAction{client: newMalformedBaseURLClient(t)}
	m := PurgeEventActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPurgeEventAction_Invoke_SendError exercises PurgeEventAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestPurgeEventAction_Invoke_SendError(t *testing.T) {
	r := &PurgeEventAction{client: newTransportErrorClient(t)}
	m := PurgeEventActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPurgeEventAction_Invoke_APIError exercises PurgeEventAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPurgeEventAction_Invoke_APIError(t *testing.T) {
	r := &PurgeEventAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PurgeEventActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_purge_event")
}

// TestPurgeEventAction_Invoke_APIErrorReadBody exercises PurgeEventAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPurgeEventAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &PurgeEventAction{client: newMockClientReadErrorBody(t, 501)}
	m := PurgeEventActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
