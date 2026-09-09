package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestSwitchNodeSystemConfigFileAction_Invoke_Happy exercises SwitchNodeSystemConfigFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestSwitchNodeSystemConfigFileAction_Invoke_Happy(t *testing.T) {
	r := &SwitchNodeSystemConfigFileAction{client: newMockClientStatus(t, 201, "{}")}
	m := SwitchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSwitchNodeSystemConfigFileAction_Invoke_NilClient exercises SwitchNodeSystemConfigFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSwitchNodeSystemConfigFileAction_Invoke_NilClient(t *testing.T) {
	r := &SwitchNodeSystemConfigFileAction{}
	m := SwitchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSwitchNodeSystemConfigFileAction_Invoke_BuildError exercises SwitchNodeSystemConfigFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSwitchNodeSystemConfigFileAction_Invoke_BuildError(t *testing.T) {
	r := &SwitchNodeSystemConfigFileAction{client: newMalformedBaseURLClient(t)}
	m := SwitchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSwitchNodeSystemConfigFileAction_Invoke_SendError exercises SwitchNodeSystemConfigFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestSwitchNodeSystemConfigFileAction_Invoke_SendError(t *testing.T) {
	r := &SwitchNodeSystemConfigFileAction{client: newTransportErrorClient(t)}
	m := SwitchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSwitchNodeSystemConfigFileAction_Invoke_APIError exercises SwitchNodeSystemConfigFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSwitchNodeSystemConfigFileAction_Invoke_APIError(t *testing.T) {
	r := &SwitchNodeSystemConfigFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SwitchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_switch_node_system_config_file")
}

// TestSwitchNodeSystemConfigFileAction_Invoke_APIErrorReadBody exercises SwitchNodeSystemConfigFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSwitchNodeSystemConfigFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &SwitchNodeSystemConfigFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := SwitchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
