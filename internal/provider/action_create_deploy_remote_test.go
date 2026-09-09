package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateDeployAction_Invoke_Happy exercises CreateDeployAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateDeployAction_Invoke_Happy(t *testing.T) {
	r := &CreateDeployAction{client: newMockClientStatus(t, 200, "{}")}
	m := CreateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateDeployAction_Invoke_NilClient exercises CreateDeployAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateDeployAction_Invoke_NilClient(t *testing.T) {
	r := &CreateDeployAction{}
	m := CreateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateDeployAction_Invoke_BuildError exercises CreateDeployAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateDeployAction_Invoke_BuildError(t *testing.T) {
	r := &CreateDeployAction{client: newMalformedBaseURLClient(t)}
	m := CreateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateDeployAction_Invoke_SendError exercises CreateDeployAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateDeployAction_Invoke_SendError(t *testing.T) {
	r := &CreateDeployAction{client: newTransportErrorClient(t)}
	m := CreateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateDeployAction_Invoke_APIError exercises CreateDeployAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateDeployAction_Invoke_APIError(t *testing.T) {
	r := &CreateDeployAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_deploy")
}

// TestCreateDeployAction_Invoke_APIErrorReadBody exercises CreateDeployAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateDeployAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateDeployAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateDeployActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
