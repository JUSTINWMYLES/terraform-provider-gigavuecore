package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteProfileDecryptPortMapByIdAction_Invoke_Happy exercises DeleteProfileDecryptPortMapByIdAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteProfileDecryptPortMapByIdAction_Invoke_Happy(t *testing.T) {
	r := &DeleteProfileDecryptPortMapByIdAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteProfileDecryptPortMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteProfileDecryptPortMapByIdAction_Invoke_NilClient exercises DeleteProfileDecryptPortMapByIdAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteProfileDecryptPortMapByIdAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteProfileDecryptPortMapByIdAction{}
	m := DeleteProfileDecryptPortMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteProfileDecryptPortMapByIdAction_Invoke_BuildError exercises DeleteProfileDecryptPortMapByIdAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteProfileDecryptPortMapByIdAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteProfileDecryptPortMapByIdAction{client: newMalformedBaseURLClient(t)}
	m := DeleteProfileDecryptPortMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteProfileDecryptPortMapByIdAction_Invoke_SendError exercises DeleteProfileDecryptPortMapByIdAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteProfileDecryptPortMapByIdAction_Invoke_SendError(t *testing.T) {
	r := &DeleteProfileDecryptPortMapByIdAction{client: newTransportErrorClient(t)}
	m := DeleteProfileDecryptPortMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteProfileDecryptPortMapByIdAction_Invoke_APIError exercises DeleteProfileDecryptPortMapByIdAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteProfileDecryptPortMapByIdAction_Invoke_APIError(t *testing.T) {
	r := &DeleteProfileDecryptPortMapByIdAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteProfileDecryptPortMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_profile_decrypt_port_map_by_id")
}

// TestDeleteProfileDecryptPortMapByIdAction_Invoke_APIErrorReadBody exercises DeleteProfileDecryptPortMapByIdAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteProfileDecryptPortMapByIdAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteProfileDecryptPortMapByIdAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteProfileDecryptPortMapByIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
