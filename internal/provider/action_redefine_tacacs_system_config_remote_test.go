package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineTacacsSystemConfigAction_Invoke_Happy exercises RedefineTacacsSystemConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineTacacsSystemConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineTacacsSystemConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineTacacsSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineTacacsSystemConfigAction_Invoke_NilClient exercises RedefineTacacsSystemConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineTacacsSystemConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineTacacsSystemConfigAction{}
	m := RedefineTacacsSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineTacacsSystemConfigAction_Invoke_BuildError exercises RedefineTacacsSystemConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineTacacsSystemConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineTacacsSystemConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineTacacsSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineTacacsSystemConfigAction_Invoke_SendError exercises RedefineTacacsSystemConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineTacacsSystemConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineTacacsSystemConfigAction{client: newTransportErrorClient(t)}
	m := RedefineTacacsSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineTacacsSystemConfigAction_Invoke_APIError exercises RedefineTacacsSystemConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineTacacsSystemConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineTacacsSystemConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineTacacsSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_tacacs_system_config")
}

// TestRedefineTacacsSystemConfigAction_Invoke_APIErrorReadBody exercises RedefineTacacsSystemConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineTacacsSystemConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineTacacsSystemConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineTacacsSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
