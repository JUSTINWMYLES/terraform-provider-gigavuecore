package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateEnvAction_Invoke_Happy exercises CreateEnvAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateEnvAction_Invoke_Happy(t *testing.T) {
	r := &CreateEnvAction{client: newMockClientStatus(t, 200, "{}")}
	m := CreateEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateEnvAction_Invoke_NilClient exercises CreateEnvAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateEnvAction_Invoke_NilClient(t *testing.T) {
	r := &CreateEnvAction{}
	m := CreateEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateEnvAction_Invoke_BuildError exercises CreateEnvAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateEnvAction_Invoke_BuildError(t *testing.T) {
	r := &CreateEnvAction{client: newMalformedBaseURLClient(t)}
	m := CreateEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateEnvAction_Invoke_SendError exercises CreateEnvAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateEnvAction_Invoke_SendError(t *testing.T) {
	r := &CreateEnvAction{client: newTransportErrorClient(t)}
	m := CreateEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateEnvAction_Invoke_APIError exercises CreateEnvAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateEnvAction_Invoke_APIError(t *testing.T) {
	r := &CreateEnvAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_env")
}

// TestCreateEnvAction_Invoke_APIErrorReadBody exercises CreateEnvAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateEnvAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateEnvAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateEnvActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
