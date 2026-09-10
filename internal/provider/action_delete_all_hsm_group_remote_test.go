package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllHsmGroupAction_Invoke_Happy exercises DeleteAllHsmGroupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllHsmGroupAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllHsmGroupAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllHsmGroupAction_Invoke_NilClient exercises DeleteAllHsmGroupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllHsmGroupAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllHsmGroupAction{}
	m := DeleteAllHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllHsmGroupAction_Invoke_BuildError exercises DeleteAllHsmGroupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllHsmGroupAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllHsmGroupAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllHsmGroupAction_Invoke_SendError exercises DeleteAllHsmGroupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllHsmGroupAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllHsmGroupAction{client: newTransportErrorClient(t)}
	m := DeleteAllHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllHsmGroupAction_Invoke_APIError exercises DeleteAllHsmGroupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllHsmGroupAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllHsmGroupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_hsm_group")
}

// TestDeleteAllHsmGroupAction_Invoke_APIErrorReadBody exercises DeleteAllHsmGroupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllHsmGroupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllHsmGroupAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
