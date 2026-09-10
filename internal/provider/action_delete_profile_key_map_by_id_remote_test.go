package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteProfileKeyMapByIdAction_Invoke_Happy exercises DeleteProfileKeyMapByIdAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteProfileKeyMapByIdAction_Invoke_Happy(t *testing.T) {
	r := &DeleteProfileKeyMapByIdAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteProfileKeyMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteProfileKeyMapByIdAction_Invoke_NilClient exercises DeleteProfileKeyMapByIdAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteProfileKeyMapByIdAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteProfileKeyMapByIdAction{}
	m := DeleteProfileKeyMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteProfileKeyMapByIdAction_Invoke_BuildError exercises DeleteProfileKeyMapByIdAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteProfileKeyMapByIdAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteProfileKeyMapByIdAction{client: newMalformedBaseURLClient(t)}
	m := DeleteProfileKeyMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteProfileKeyMapByIdAction_Invoke_SendError exercises DeleteProfileKeyMapByIdAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteProfileKeyMapByIdAction_Invoke_SendError(t *testing.T) {
	r := &DeleteProfileKeyMapByIdAction{client: newTransportErrorClient(t)}
	m := DeleteProfileKeyMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteProfileKeyMapByIdAction_Invoke_APIError exercises DeleteProfileKeyMapByIdAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteProfileKeyMapByIdAction_Invoke_APIError(t *testing.T) {
	r := &DeleteProfileKeyMapByIdAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteProfileKeyMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_profile_key_map_by_id")
}

// TestDeleteProfileKeyMapByIdAction_Invoke_APIErrorReadBody exercises DeleteProfileKeyMapByIdAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteProfileKeyMapByIdAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteProfileKeyMapByIdAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteProfileKeyMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
