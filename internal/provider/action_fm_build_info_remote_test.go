package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestFmBuildInfoAction_Invoke_Happy exercises FmBuildInfoAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestFmBuildInfoAction_Invoke_Happy(t *testing.T) {
	r := &FmBuildInfoAction{client: newMockClientStatus(t, 201, "{}")}
	m := FmBuildInfoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmBuildInfoAction_Invoke_NilClient exercises FmBuildInfoAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmBuildInfoAction_Invoke_NilClient(t *testing.T) {
	r := &FmBuildInfoAction{}
	m := FmBuildInfoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmBuildInfoAction_Invoke_BuildError exercises FmBuildInfoAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmBuildInfoAction_Invoke_BuildError(t *testing.T) {
	r := &FmBuildInfoAction{client: newMalformedBaseURLClient(t)}
	m := FmBuildInfoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmBuildInfoAction_Invoke_SendError exercises FmBuildInfoAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmBuildInfoAction_Invoke_SendError(t *testing.T) {
	r := &FmBuildInfoAction{client: newTransportErrorClient(t)}
	m := FmBuildInfoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmBuildInfoAction_Invoke_APIError exercises FmBuildInfoAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmBuildInfoAction_Invoke_APIError(t *testing.T) {
	r := &FmBuildInfoAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmBuildInfoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_fm_build_info")
}

// TestFmBuildInfoAction_Invoke_APIErrorReadBody exercises FmBuildInfoAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmBuildInfoAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &FmBuildInfoAction{client: newMockClientReadErrorBody(t, 501)}
	m := FmBuildInfoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
