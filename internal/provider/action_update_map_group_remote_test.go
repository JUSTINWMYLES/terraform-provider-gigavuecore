package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapGroupAction_Invoke_Happy exercises UpdateMapGroupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapGroupAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapGroupAction{client: newMockClientStatus(t, 202, "{}")}
	m := UpdateMapGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapGroupAction_Invoke_NilClient exercises UpdateMapGroupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapGroupAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapGroupAction{}
	m := UpdateMapGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapGroupAction_Invoke_BuildError exercises UpdateMapGroupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapGroupAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapGroupAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapGroupAction_Invoke_SendError exercises UpdateMapGroupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapGroupAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapGroupAction{client: newTransportErrorClient(t)}
	m := UpdateMapGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapGroupAction_Invoke_APIError exercises UpdateMapGroupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapGroupAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapGroupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_group")
}

// TestUpdateMapGroupAction_Invoke_APIErrorReadBody exercises UpdateMapGroupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapGroupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapGroupAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
