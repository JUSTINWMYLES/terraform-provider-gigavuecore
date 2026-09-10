package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteFilterTemplatesAction_Invoke_Happy exercises DeleteFilterTemplatesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteFilterTemplatesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteFilterTemplatesAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteFilterTemplatesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteFilterTemplatesAction_Invoke_NilClient exercises DeleteFilterTemplatesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteFilterTemplatesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteFilterTemplatesAction{}
	m := DeleteFilterTemplatesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteFilterTemplatesAction_Invoke_BuildError exercises DeleteFilterTemplatesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteFilterTemplatesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteFilterTemplatesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteFilterTemplatesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteFilterTemplatesAction_Invoke_SendError exercises DeleteFilterTemplatesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteFilterTemplatesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteFilterTemplatesAction{client: newTransportErrorClient(t)}
	m := DeleteFilterTemplatesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteFilterTemplatesAction_Invoke_APIError exercises DeleteFilterTemplatesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteFilterTemplatesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteFilterTemplatesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteFilterTemplatesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_filter_templates")
}

// TestDeleteFilterTemplatesAction_Invoke_APIErrorReadBody exercises DeleteFilterTemplatesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteFilterTemplatesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteFilterTemplatesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteFilterTemplatesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
