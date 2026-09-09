package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteProfileKeyMapsAction_Invoke_Happy exercises DeleteProfileKeyMapsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteProfileKeyMapsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteProfileKeyMapsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteProfileKeyMapsAction_Invoke_NilClient exercises DeleteProfileKeyMapsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteProfileKeyMapsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteProfileKeyMapsAction{}
	m := DeleteProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteProfileKeyMapsAction_Invoke_BuildError exercises DeleteProfileKeyMapsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteProfileKeyMapsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteProfileKeyMapsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteProfileKeyMapsAction_Invoke_SendError exercises DeleteProfileKeyMapsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteProfileKeyMapsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteProfileKeyMapsAction{client: newTransportErrorClient(t)}
	m := DeleteProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteProfileKeyMapsAction_Invoke_APIError exercises DeleteProfileKeyMapsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteProfileKeyMapsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteProfileKeyMapsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_profile_key_maps")
}

// TestDeleteProfileKeyMapsAction_Invoke_APIErrorReadBody exercises DeleteProfileKeyMapsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteProfileKeyMapsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteProfileKeyMapsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteProfileKeyMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
