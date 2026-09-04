package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestPasteRulesAction_Invoke_Happy exercises PasteRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestPasteRulesAction_Invoke_Happy(t *testing.T) {
	r := &PasteRulesAction{client: newMockClientStatus(t, 200, "{}")}
	m := PasteRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPasteRulesAction_Invoke_NilClient exercises PasteRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPasteRulesAction_Invoke_NilClient(t *testing.T) {
	r := &PasteRulesAction{}
	m := PasteRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPasteRulesAction_Invoke_BuildError exercises PasteRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPasteRulesAction_Invoke_BuildError(t *testing.T) {
	r := &PasteRulesAction{client: newMalformedBaseURLClient(t)}
	m := PasteRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPasteRulesAction_Invoke_SendError exercises PasteRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestPasteRulesAction_Invoke_SendError(t *testing.T) {
	r := &PasteRulesAction{client: newTransportErrorClient(t)}
	m := PasteRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPasteRulesAction_Invoke_APIError exercises PasteRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPasteRulesAction_Invoke_APIError(t *testing.T) {
	r := &PasteRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PasteRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_paste_rules")
}

// TestPasteRulesAction_Invoke_APIErrorReadBody exercises PasteRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPasteRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &PasteRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := PasteRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
