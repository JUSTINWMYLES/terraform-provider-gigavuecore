package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapApRulesAction_Invoke_Happy exercises RemoveAllMapApRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapApRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapApRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapApRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapApRulesAction_Invoke_NilClient exercises RemoveAllMapApRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapApRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapApRulesAction{}
	m := RemoveAllMapApRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapApRulesAction_Invoke_BuildError exercises RemoveAllMapApRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapApRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapApRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapApRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapApRulesAction_Invoke_SendError exercises RemoveAllMapApRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapApRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapApRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapApRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapApRulesAction_Invoke_APIError exercises RemoveAllMapApRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapApRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapApRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapApRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_ap_rules")
}

// TestRemoveAllMapApRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapApRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapApRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapApRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapApRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
