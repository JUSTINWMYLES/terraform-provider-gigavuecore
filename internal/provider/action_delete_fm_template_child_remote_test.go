package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteFmTemplateChildAction_Invoke_Happy exercises DeleteFmTemplateChildAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteFmTemplateChildAction_Invoke_Happy(t *testing.T) {
	r := &DeleteFmTemplateChildAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteFmTemplateChildAction_Invoke_NilClient exercises DeleteFmTemplateChildAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteFmTemplateChildAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteFmTemplateChildAction{}
	m := DeleteFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteFmTemplateChildAction_Invoke_BuildError exercises DeleteFmTemplateChildAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteFmTemplateChildAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteFmTemplateChildAction{client: newMalformedBaseURLClient(t)}
	m := DeleteFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteFmTemplateChildAction_Invoke_SendError exercises DeleteFmTemplateChildAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteFmTemplateChildAction_Invoke_SendError(t *testing.T) {
	r := &DeleteFmTemplateChildAction{client: newTransportErrorClient(t)}
	m := DeleteFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteFmTemplateChildAction_Invoke_APIError exercises DeleteFmTemplateChildAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteFmTemplateChildAction_Invoke_APIError(t *testing.T) {
	r := &DeleteFmTemplateChildAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_fm_template_child")
}

// TestDeleteFmTemplateChildAction_Invoke_APIErrorReadBody exercises DeleteFmTemplateChildAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteFmTemplateChildAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteFmTemplateChildAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
