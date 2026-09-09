package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestFmUnlockUserConfigAction_Invoke_Happy exercises FmUnlockUserConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestFmUnlockUserConfigAction_Invoke_Happy(t *testing.T) {
	r := &FmUnlockUserConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := FmUnlockUserConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmUnlockUserConfigAction_Invoke_NilClient exercises FmUnlockUserConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmUnlockUserConfigAction_Invoke_NilClient(t *testing.T) {
	r := &FmUnlockUserConfigAction{}
	m := FmUnlockUserConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmUnlockUserConfigAction_Invoke_BuildError exercises FmUnlockUserConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmUnlockUserConfigAction_Invoke_BuildError(t *testing.T) {
	r := &FmUnlockUserConfigAction{client: newMalformedBaseURLClient(t)}
	m := FmUnlockUserConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmUnlockUserConfigAction_Invoke_SendError exercises FmUnlockUserConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmUnlockUserConfigAction_Invoke_SendError(t *testing.T) {
	r := &FmUnlockUserConfigAction{client: newTransportErrorClient(t)}
	m := FmUnlockUserConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmUnlockUserConfigAction_Invoke_APIError exercises FmUnlockUserConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmUnlockUserConfigAction_Invoke_APIError(t *testing.T) {
	r := &FmUnlockUserConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmUnlockUserConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_fm_unlock_user_config")
}

// TestFmUnlockUserConfigAction_Invoke_APIErrorReadBody exercises FmUnlockUserConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmUnlockUserConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &FmUnlockUserConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := FmUnlockUserConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
