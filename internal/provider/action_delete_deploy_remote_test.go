package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteDeployAction_Invoke_Happy exercises DeleteDeployAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteDeployAction_Invoke_Happy(t *testing.T) {
	r := &DeleteDeployAction{client: newMockClientStatus(t, 200, "{}")}
	m := DeleteDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteDeployAction_Invoke_NilClient exercises DeleteDeployAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteDeployAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteDeployAction{}
	m := DeleteDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteDeployAction_Invoke_BuildError exercises DeleteDeployAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteDeployAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteDeployAction{client: newMalformedBaseURLClient(t)}
	m := DeleteDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteDeployAction_Invoke_SendError exercises DeleteDeployAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteDeployAction_Invoke_SendError(t *testing.T) {
	r := &DeleteDeployAction{client: newTransportErrorClient(t)}
	m := DeleteDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteDeployAction_Invoke_APIError exercises DeleteDeployAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteDeployAction_Invoke_APIError(t *testing.T) {
	r := &DeleteDeployAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_deploy")
}

// TestDeleteDeployAction_Invoke_APIErrorReadBody exercises DeleteDeployAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteDeployAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteDeployAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
