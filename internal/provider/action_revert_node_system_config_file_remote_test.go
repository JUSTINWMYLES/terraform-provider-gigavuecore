package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRevertNodeSystemConfigFileAction_Invoke_Happy exercises RevertNodeSystemConfigFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRevertNodeSystemConfigFileAction_Invoke_Happy(t *testing.T) {
	r := &RevertNodeSystemConfigFileAction{client: newMockClientStatus(t, 201, "{}")}
	m := RevertNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRevertNodeSystemConfigFileAction_Invoke_NilClient exercises RevertNodeSystemConfigFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRevertNodeSystemConfigFileAction_Invoke_NilClient(t *testing.T) {
	r := &RevertNodeSystemConfigFileAction{}
	m := RevertNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRevertNodeSystemConfigFileAction_Invoke_BuildError exercises RevertNodeSystemConfigFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRevertNodeSystemConfigFileAction_Invoke_BuildError(t *testing.T) {
	r := &RevertNodeSystemConfigFileAction{client: newMalformedBaseURLClient(t)}
	m := RevertNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRevertNodeSystemConfigFileAction_Invoke_SendError exercises RevertNodeSystemConfigFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRevertNodeSystemConfigFileAction_Invoke_SendError(t *testing.T) {
	r := &RevertNodeSystemConfigFileAction{client: newTransportErrorClient(t)}
	m := RevertNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRevertNodeSystemConfigFileAction_Invoke_APIError exercises RevertNodeSystemConfigFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRevertNodeSystemConfigFileAction_Invoke_APIError(t *testing.T) {
	r := &RevertNodeSystemConfigFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RevertNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_revert_node_system_config_file")
}

// TestRevertNodeSystemConfigFileAction_Invoke_APIErrorReadBody exercises RevertNodeSystemConfigFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRevertNodeSystemConfigFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RevertNodeSystemConfigFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := RevertNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
