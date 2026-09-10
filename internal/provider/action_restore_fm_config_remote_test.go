package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRestoreFmConfigAction_Invoke_Happy exercises RestoreFmConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRestoreFmConfigAction_Invoke_Happy(t *testing.T) {
	r := &RestoreFmConfigAction{client: newMockClientStatus(t, 201, "{}")}
	m := RestoreFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRestoreFmConfigAction_Invoke_NilClient exercises RestoreFmConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRestoreFmConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RestoreFmConfigAction{}
	m := RestoreFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRestoreFmConfigAction_Invoke_BuildError exercises RestoreFmConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRestoreFmConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RestoreFmConfigAction{client: newMalformedBaseURLClient(t)}
	m := RestoreFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRestoreFmConfigAction_Invoke_SendError exercises RestoreFmConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRestoreFmConfigAction_Invoke_SendError(t *testing.T) {
	r := &RestoreFmConfigAction{client: newTransportErrorClient(t)}
	m := RestoreFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRestoreFmConfigAction_Invoke_APIError exercises RestoreFmConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRestoreFmConfigAction_Invoke_APIError(t *testing.T) {
	r := &RestoreFmConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RestoreFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_restore_fm_config")
}

// TestRestoreFmConfigAction_Invoke_APIErrorReadBody exercises RestoreFmConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRestoreFmConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RestoreFmConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RestoreFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
