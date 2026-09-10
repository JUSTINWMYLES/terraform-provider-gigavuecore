package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineFabricAdvHashAction_Invoke_Happy exercises RedefineFabricAdvHashAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineFabricAdvHashAction_Invoke_Happy(t *testing.T) {
	r := &RedefineFabricAdvHashAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineFabricAdvHashActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineFabricAdvHashAction_Invoke_NilClient exercises RedefineFabricAdvHashAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineFabricAdvHashAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineFabricAdvHashAction{}
	m := RedefineFabricAdvHashActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineFabricAdvHashAction_Invoke_BuildError exercises RedefineFabricAdvHashAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineFabricAdvHashAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineFabricAdvHashAction{client: newMalformedBaseURLClient(t)}
	m := RedefineFabricAdvHashActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineFabricAdvHashAction_Invoke_SendError exercises RedefineFabricAdvHashAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineFabricAdvHashAction_Invoke_SendError(t *testing.T) {
	r := &RedefineFabricAdvHashAction{client: newTransportErrorClient(t)}
	m := RedefineFabricAdvHashActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineFabricAdvHashAction_Invoke_APIError exercises RedefineFabricAdvHashAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineFabricAdvHashAction_Invoke_APIError(t *testing.T) {
	r := &RedefineFabricAdvHashAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineFabricAdvHashActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_fabric_adv_hash")
}

// TestRedefineFabricAdvHashAction_Invoke_APIErrorReadBody exercises RedefineFabricAdvHashAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineFabricAdvHashAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineFabricAdvHashAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineFabricAdvHashActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
