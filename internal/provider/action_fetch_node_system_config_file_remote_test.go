package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestFetchNodeSystemConfigFileAction_Invoke_Happy exercises FetchNodeSystemConfigFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestFetchNodeSystemConfigFileAction_Invoke_Happy(t *testing.T) {
	r := &FetchNodeSystemConfigFileAction{client: newMockClientStatus(t, 201, "{}")}
	m := FetchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFetchNodeSystemConfigFileAction_Invoke_NilClient exercises FetchNodeSystemConfigFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFetchNodeSystemConfigFileAction_Invoke_NilClient(t *testing.T) {
	r := &FetchNodeSystemConfigFileAction{}
	m := FetchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFetchNodeSystemConfigFileAction_Invoke_BuildError exercises FetchNodeSystemConfigFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFetchNodeSystemConfigFileAction_Invoke_BuildError(t *testing.T) {
	r := &FetchNodeSystemConfigFileAction{client: newMalformedBaseURLClient(t)}
	m := FetchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFetchNodeSystemConfigFileAction_Invoke_SendError exercises FetchNodeSystemConfigFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestFetchNodeSystemConfigFileAction_Invoke_SendError(t *testing.T) {
	r := &FetchNodeSystemConfigFileAction{client: newTransportErrorClient(t)}
	m := FetchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFetchNodeSystemConfigFileAction_Invoke_APIError exercises FetchNodeSystemConfigFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFetchNodeSystemConfigFileAction_Invoke_APIError(t *testing.T) {
	r := &FetchNodeSystemConfigFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FetchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_fetch_node_system_config_file")
}

// TestFetchNodeSystemConfigFileAction_Invoke_APIErrorReadBody exercises FetchNodeSystemConfigFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFetchNodeSystemConfigFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &FetchNodeSystemConfigFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := FetchNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
