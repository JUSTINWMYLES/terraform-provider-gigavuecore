package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestImportPolicyAction_Invoke_Happy exercises ImportPolicyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestImportPolicyAction_Invoke_Happy(t *testing.T) {
	r := &ImportPolicyAction{client: newMockClientStatus(t, 201, "{}")}
	m := ImportPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestImportPolicyAction_Invoke_NilClient exercises ImportPolicyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestImportPolicyAction_Invoke_NilClient(t *testing.T) {
	r := &ImportPolicyAction{}
	m := ImportPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestImportPolicyAction_Invoke_BuildError exercises ImportPolicyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestImportPolicyAction_Invoke_BuildError(t *testing.T) {
	r := &ImportPolicyAction{client: newMalformedBaseURLClient(t)}
	m := ImportPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestImportPolicyAction_Invoke_SendError exercises ImportPolicyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestImportPolicyAction_Invoke_SendError(t *testing.T) {
	r := &ImportPolicyAction{client: newTransportErrorClient(t)}
	m := ImportPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestImportPolicyAction_Invoke_APIError exercises ImportPolicyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestImportPolicyAction_Invoke_APIError(t *testing.T) {
	r := &ImportPolicyAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ImportPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_import_policy")
}

// TestImportPolicyAction_Invoke_APIErrorReadBody exercises ImportPolicyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestImportPolicyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ImportPolicyAction{client: newMockClientReadErrorBody(t, 500)}
	m := ImportPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
