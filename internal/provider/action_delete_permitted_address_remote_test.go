package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeletePermittedAddressAction_Invoke_Happy exercises DeletePermittedAddressAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeletePermittedAddressAction_Invoke_Happy(t *testing.T) {
	r := &DeletePermittedAddressAction{client: newMockClientStatus(t, 200, "{}")}
	m := DeletePermittedAddressActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeletePermittedAddressAction_Invoke_NilClient exercises DeletePermittedAddressAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeletePermittedAddressAction_Invoke_NilClient(t *testing.T) {
	r := &DeletePermittedAddressAction{}
	m := DeletePermittedAddressActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeletePermittedAddressAction_Invoke_BuildError exercises DeletePermittedAddressAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeletePermittedAddressAction_Invoke_BuildError(t *testing.T) {
	r := &DeletePermittedAddressAction{client: newMalformedBaseURLClient(t)}
	m := DeletePermittedAddressActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeletePermittedAddressAction_Invoke_SendError exercises DeletePermittedAddressAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeletePermittedAddressAction_Invoke_SendError(t *testing.T) {
	r := &DeletePermittedAddressAction{client: newTransportErrorClient(t)}
	m := DeletePermittedAddressActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeletePermittedAddressAction_Invoke_APIError exercises DeletePermittedAddressAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeletePermittedAddressAction_Invoke_APIError(t *testing.T) {
	r := &DeletePermittedAddressAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeletePermittedAddressActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_permitted_address")
}

// TestDeletePermittedAddressAction_Invoke_APIErrorReadBody exercises DeletePermittedAddressAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeletePermittedAddressAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeletePermittedAddressAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeletePermittedAddressActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
