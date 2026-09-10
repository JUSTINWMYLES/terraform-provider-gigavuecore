package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteEnvAction_Invoke_Happy exercises DeleteEnvAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteEnvAction_Invoke_Happy(t *testing.T) {
	r := &DeleteEnvAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteEnvAction_Invoke_NilClient exercises DeleteEnvAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteEnvAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteEnvAction{}
	m := DeleteEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteEnvAction_Invoke_BuildError exercises DeleteEnvAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteEnvAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteEnvAction{client: newMalformedBaseURLClient(t)}
	m := DeleteEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteEnvAction_Invoke_SendError exercises DeleteEnvAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteEnvAction_Invoke_SendError(t *testing.T) {
	r := &DeleteEnvAction{client: newTransportErrorClient(t)}
	m := DeleteEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteEnvAction_Invoke_APIError exercises DeleteEnvAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteEnvAction_Invoke_APIError(t *testing.T) {
	r := &DeleteEnvAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_env")
}

// TestDeleteEnvAction_Invoke_APIErrorReadBody exercises DeleteEnvAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteEnvAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteEnvAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
