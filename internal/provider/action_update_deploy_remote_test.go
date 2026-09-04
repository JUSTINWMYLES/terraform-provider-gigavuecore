package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateDeployAction_Invoke_Happy exercises UpdateDeployAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateDeployAction_Invoke_Happy(t *testing.T) {
	r := &UpdateDeployAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateDeployAction_Invoke_NilClient exercises UpdateDeployAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateDeployAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateDeployAction{}
	m := UpdateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateDeployAction_Invoke_BuildError exercises UpdateDeployAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateDeployAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateDeployAction{client: newMalformedBaseURLClient(t)}
	m := UpdateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateDeployAction_Invoke_SendError exercises UpdateDeployAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateDeployAction_Invoke_SendError(t *testing.T) {
	r := &UpdateDeployAction{client: newTransportErrorClient(t)}
	m := UpdateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateDeployAction_Invoke_APIError exercises UpdateDeployAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateDeployAction_Invoke_APIError(t *testing.T) {
	r := &UpdateDeployAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_deploy")
}

// TestUpdateDeployAction_Invoke_APIErrorReadBody exercises UpdateDeployAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateDeployAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateDeployAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
