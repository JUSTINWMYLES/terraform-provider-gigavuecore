package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRegisterGigaInsightNodeAction_Invoke_Happy exercises RegisterGigaInsightNodeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRegisterGigaInsightNodeAction_Invoke_Happy(t *testing.T) {
	r := &RegisterGigaInsightNodeAction{client: newMockClientStatus(t, 201, "{}")}
	m := RegisterGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRegisterGigaInsightNodeAction_Invoke_NilClient exercises RegisterGigaInsightNodeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRegisterGigaInsightNodeAction_Invoke_NilClient(t *testing.T) {
	r := &RegisterGigaInsightNodeAction{}
	m := RegisterGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRegisterGigaInsightNodeAction_Invoke_BuildError exercises RegisterGigaInsightNodeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRegisterGigaInsightNodeAction_Invoke_BuildError(t *testing.T) {
	r := &RegisterGigaInsightNodeAction{client: newMalformedBaseURLClient(t)}
	m := RegisterGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRegisterGigaInsightNodeAction_Invoke_SendError exercises RegisterGigaInsightNodeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRegisterGigaInsightNodeAction_Invoke_SendError(t *testing.T) {
	r := &RegisterGigaInsightNodeAction{client: newTransportErrorClient(t)}
	m := RegisterGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRegisterGigaInsightNodeAction_Invoke_APIError exercises RegisterGigaInsightNodeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRegisterGigaInsightNodeAction_Invoke_APIError(t *testing.T) {
	r := &RegisterGigaInsightNodeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RegisterGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_register_giga_insight_node")
}

// TestRegisterGigaInsightNodeAction_Invoke_APIErrorReadBody exercises RegisterGigaInsightNodeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRegisterGigaInsightNodeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RegisterGigaInsightNodeAction{client: newMockClientReadErrorBody(t, 501)}
	m := RegisterGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
