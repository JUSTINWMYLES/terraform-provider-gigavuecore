package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestBuildFlowPathAction_Invoke_Happy exercises BuildFlowPathAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestBuildFlowPathAction_Invoke_Happy(t *testing.T) {
	r := &BuildFlowPathAction{client: newMockClientStatus(t, 201, "{}")}
	m := BuildFlowPathActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestBuildFlowPathAction_Invoke_NilClient exercises BuildFlowPathAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestBuildFlowPathAction_Invoke_NilClient(t *testing.T) {
	r := &BuildFlowPathAction{}
	m := BuildFlowPathActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestBuildFlowPathAction_Invoke_BuildError exercises BuildFlowPathAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestBuildFlowPathAction_Invoke_BuildError(t *testing.T) {
	r := &BuildFlowPathAction{client: newMalformedBaseURLClient(t)}
	m := BuildFlowPathActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestBuildFlowPathAction_Invoke_SendError exercises BuildFlowPathAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestBuildFlowPathAction_Invoke_SendError(t *testing.T) {
	r := &BuildFlowPathAction{client: newTransportErrorClient(t)}
	m := BuildFlowPathActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestBuildFlowPathAction_Invoke_APIError exercises BuildFlowPathAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestBuildFlowPathAction_Invoke_APIError(t *testing.T) {
	r := &BuildFlowPathAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := BuildFlowPathActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_build_flow_path")
}

// TestBuildFlowPathAction_Invoke_APIErrorReadBody exercises BuildFlowPathAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestBuildFlowPathAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &BuildFlowPathAction{client: newMockClientReadErrorBody(t, 501)}
	m := BuildFlowPathActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
