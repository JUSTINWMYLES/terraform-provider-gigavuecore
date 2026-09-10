package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllGsEngineInterfaceAction_Invoke_Happy exercises DeleteAllGsEngineInterfaceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllGsEngineInterfaceAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllGsEngineInterfaceAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllGsEngineInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllGsEngineInterfaceAction_Invoke_NilClient exercises DeleteAllGsEngineInterfaceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllGsEngineInterfaceAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllGsEngineInterfaceAction{}
	m := DeleteAllGsEngineInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllGsEngineInterfaceAction_Invoke_BuildError exercises DeleteAllGsEngineInterfaceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllGsEngineInterfaceAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllGsEngineInterfaceAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllGsEngineInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllGsEngineInterfaceAction_Invoke_SendError exercises DeleteAllGsEngineInterfaceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllGsEngineInterfaceAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllGsEngineInterfaceAction{client: newTransportErrorClient(t)}
	m := DeleteAllGsEngineInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllGsEngineInterfaceAction_Invoke_APIError exercises DeleteAllGsEngineInterfaceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllGsEngineInterfaceAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllGsEngineInterfaceAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllGsEngineInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_gs_engine_interface")
}

// TestDeleteAllGsEngineInterfaceAction_Invoke_APIErrorReadBody exercises DeleteAllGsEngineInterfaceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllGsEngineInterfaceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllGsEngineInterfaceAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllGsEngineInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
