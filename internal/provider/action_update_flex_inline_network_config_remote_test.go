package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateFlexInlineNetworkConfigAction_Invoke_Happy exercises UpdateFlexInlineNetworkConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateFlexInlineNetworkConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateFlexInlineNetworkConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateFlexInlineNetworkConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateFlexInlineNetworkConfigAction_Invoke_NilClient exercises UpdateFlexInlineNetworkConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateFlexInlineNetworkConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateFlexInlineNetworkConfigAction{}
	m := UpdateFlexInlineNetworkConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateFlexInlineNetworkConfigAction_Invoke_BuildError exercises UpdateFlexInlineNetworkConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateFlexInlineNetworkConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateFlexInlineNetworkConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateFlexInlineNetworkConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateFlexInlineNetworkConfigAction_Invoke_SendError exercises UpdateFlexInlineNetworkConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateFlexInlineNetworkConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateFlexInlineNetworkConfigAction{client: newTransportErrorClient(t)}
	m := UpdateFlexInlineNetworkConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateFlexInlineNetworkConfigAction_Invoke_APIError exercises UpdateFlexInlineNetworkConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateFlexInlineNetworkConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateFlexInlineNetworkConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateFlexInlineNetworkConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_flex_inline_network_config")
}

// TestUpdateFlexInlineNetworkConfigAction_Invoke_APIErrorReadBody exercises UpdateFlexInlineNetworkConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateFlexInlineNetworkConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateFlexInlineNetworkConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateFlexInlineNetworkConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
