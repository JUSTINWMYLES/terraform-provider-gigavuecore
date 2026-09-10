package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteBindingsAction_Invoke_Happy exercises DeleteBindingsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteBindingsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteBindingsAction{client: newMockClientStatus(t, 200, "{}")}
	m := DeleteBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteBindingsAction_Invoke_NilClient exercises DeleteBindingsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteBindingsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteBindingsAction{}
	m := DeleteBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteBindingsAction_Invoke_BuildError exercises DeleteBindingsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteBindingsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteBindingsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteBindingsAction_Invoke_SendError exercises DeleteBindingsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteBindingsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteBindingsAction{client: newTransportErrorClient(t)}
	m := DeleteBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteBindingsAction_Invoke_APIError exercises DeleteBindingsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteBindingsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteBindingsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_bindings")
}

// TestDeleteBindingsAction_Invoke_APIErrorReadBody exercises DeleteBindingsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteBindingsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteBindingsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
