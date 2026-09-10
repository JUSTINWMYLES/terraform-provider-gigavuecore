package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReplaceFlexInlineMapConfigAction_Invoke_Happy exercises ReplaceFlexInlineMapConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReplaceFlexInlineMapConfigAction_Invoke_Happy(t *testing.T) {
	r := &ReplaceFlexInlineMapConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := ReplaceFlexInlineMapConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReplaceFlexInlineMapConfigAction_Invoke_NilClient exercises ReplaceFlexInlineMapConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReplaceFlexInlineMapConfigAction_Invoke_NilClient(t *testing.T) {
	r := &ReplaceFlexInlineMapConfigAction{}
	m := ReplaceFlexInlineMapConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReplaceFlexInlineMapConfigAction_Invoke_BuildError exercises ReplaceFlexInlineMapConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReplaceFlexInlineMapConfigAction_Invoke_BuildError(t *testing.T) {
	r := &ReplaceFlexInlineMapConfigAction{client: newMalformedBaseURLClient(t)}
	m := ReplaceFlexInlineMapConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReplaceFlexInlineMapConfigAction_Invoke_SendError exercises ReplaceFlexInlineMapConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReplaceFlexInlineMapConfigAction_Invoke_SendError(t *testing.T) {
	r := &ReplaceFlexInlineMapConfigAction{client: newTransportErrorClient(t)}
	m := ReplaceFlexInlineMapConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReplaceFlexInlineMapConfigAction_Invoke_APIError exercises ReplaceFlexInlineMapConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReplaceFlexInlineMapConfigAction_Invoke_APIError(t *testing.T) {
	r := &ReplaceFlexInlineMapConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReplaceFlexInlineMapConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_replace_flex_inline_map_config")
}

// TestReplaceFlexInlineMapConfigAction_Invoke_APIErrorReadBody exercises ReplaceFlexInlineMapConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReplaceFlexInlineMapConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReplaceFlexInlineMapConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReplaceFlexInlineMapConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
