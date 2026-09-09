package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddFmTemplateChildAction_Invoke_Happy exercises AddFmTemplateChildAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddFmTemplateChildAction_Invoke_Happy(t *testing.T) {
	r := &AddFmTemplateChildAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddFmTemplateChildAction_Invoke_NilClient exercises AddFmTemplateChildAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddFmTemplateChildAction_Invoke_NilClient(t *testing.T) {
	r := &AddFmTemplateChildAction{}
	m := AddFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddFmTemplateChildAction_Invoke_BuildError exercises AddFmTemplateChildAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddFmTemplateChildAction_Invoke_BuildError(t *testing.T) {
	r := &AddFmTemplateChildAction{client: newMalformedBaseURLClient(t)}
	m := AddFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddFmTemplateChildAction_Invoke_SendError exercises AddFmTemplateChildAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddFmTemplateChildAction_Invoke_SendError(t *testing.T) {
	r := &AddFmTemplateChildAction{client: newTransportErrorClient(t)}
	m := AddFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddFmTemplateChildAction_Invoke_APIError exercises AddFmTemplateChildAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddFmTemplateChildAction_Invoke_APIError(t *testing.T) {
	r := &AddFmTemplateChildAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_fm_template_child")
}

// TestAddFmTemplateChildAction_Invoke_APIErrorReadBody exercises AddFmTemplateChildAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddFmTemplateChildAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddFmTemplateChildAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddFmTemplateChildActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
