package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteInlineSslProfileListAction_Invoke_Happy exercises DeleteInlineSslProfileListAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteInlineSslProfileListAction_Invoke_Happy(t *testing.T) {
	r := &DeleteInlineSslProfileListAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteInlineSslProfileListAction_Invoke_NilClient exercises DeleteInlineSslProfileListAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteInlineSslProfileListAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteInlineSslProfileListAction{}
	m := DeleteInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteInlineSslProfileListAction_Invoke_BuildError exercises DeleteInlineSslProfileListAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteInlineSslProfileListAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteInlineSslProfileListAction{client: newMalformedBaseURLClient(t)}
	m := DeleteInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteInlineSslProfileListAction_Invoke_SendError exercises DeleteInlineSslProfileListAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteInlineSslProfileListAction_Invoke_SendError(t *testing.T) {
	r := &DeleteInlineSslProfileListAction{client: newTransportErrorClient(t)}
	m := DeleteInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteInlineSslProfileListAction_Invoke_APIError exercises DeleteInlineSslProfileListAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteInlineSslProfileListAction_Invoke_APIError(t *testing.T) {
	r := &DeleteInlineSslProfileListAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_inline_ssl_profile_list")
}

// TestDeleteInlineSslProfileListAction_Invoke_APIErrorReadBody exercises DeleteInlineSslProfileListAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteInlineSslProfileListAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteInlineSslProfileListAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
