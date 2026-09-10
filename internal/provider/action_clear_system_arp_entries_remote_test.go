package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearSystemArpEntriesAction_Invoke_Happy exercises ClearSystemArpEntriesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearSystemArpEntriesAction_Invoke_Happy(t *testing.T) {
	r := &ClearSystemArpEntriesAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearSystemArpEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearSystemArpEntriesAction_Invoke_NilClient exercises ClearSystemArpEntriesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearSystemArpEntriesAction_Invoke_NilClient(t *testing.T) {
	r := &ClearSystemArpEntriesAction{}
	m := ClearSystemArpEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearSystemArpEntriesAction_Invoke_BuildError exercises ClearSystemArpEntriesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearSystemArpEntriesAction_Invoke_BuildError(t *testing.T) {
	r := &ClearSystemArpEntriesAction{client: newMalformedBaseURLClient(t)}
	m := ClearSystemArpEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearSystemArpEntriesAction_Invoke_SendError exercises ClearSystemArpEntriesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearSystemArpEntriesAction_Invoke_SendError(t *testing.T) {
	r := &ClearSystemArpEntriesAction{client: newTransportErrorClient(t)}
	m := ClearSystemArpEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearSystemArpEntriesAction_Invoke_APIError exercises ClearSystemArpEntriesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearSystemArpEntriesAction_Invoke_APIError(t *testing.T) {
	r := &ClearSystemArpEntriesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearSystemArpEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_system_arp_entries")
}

// TestClearSystemArpEntriesAction_Invoke_APIErrorReadBody exercises ClearSystemArpEntriesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearSystemArpEntriesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearSystemArpEntriesAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearSystemArpEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
