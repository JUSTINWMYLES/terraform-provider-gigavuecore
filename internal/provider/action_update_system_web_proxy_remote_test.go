package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSystemWebProxyAction_Invoke_Happy exercises UpdateSystemWebProxyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSystemWebProxyAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSystemWebProxyAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSystemWebProxyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSystemWebProxyAction_Invoke_NilClient exercises UpdateSystemWebProxyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSystemWebProxyAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSystemWebProxyAction{}
	m := UpdateSystemWebProxyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSystemWebProxyAction_Invoke_BuildError exercises UpdateSystemWebProxyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSystemWebProxyAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSystemWebProxyAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSystemWebProxyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSystemWebProxyAction_Invoke_SendError exercises UpdateSystemWebProxyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSystemWebProxyAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSystemWebProxyAction{client: newTransportErrorClient(t)}
	m := UpdateSystemWebProxyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSystemWebProxyAction_Invoke_APIError exercises UpdateSystemWebProxyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSystemWebProxyAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSystemWebProxyAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSystemWebProxyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_system_web_proxy")
}

// TestUpdateSystemWebProxyAction_Invoke_APIErrorReadBody exercises UpdateSystemWebProxyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSystemWebProxyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSystemWebProxyAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSystemWebProxyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
