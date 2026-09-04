package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateNodeSystemConfigFileAction_Invoke_Happy exercises CreateNodeSystemConfigFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateNodeSystemConfigFileAction_Invoke_Happy(t *testing.T) {
	r := &CreateNodeSystemConfigFileAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateNodeSystemConfigFileAction_Invoke_NilClient exercises CreateNodeSystemConfigFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateNodeSystemConfigFileAction_Invoke_NilClient(t *testing.T) {
	r := &CreateNodeSystemConfigFileAction{}
	m := CreateNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateNodeSystemConfigFileAction_Invoke_BuildError exercises CreateNodeSystemConfigFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateNodeSystemConfigFileAction_Invoke_BuildError(t *testing.T) {
	r := &CreateNodeSystemConfigFileAction{client: newMalformedBaseURLClient(t)}
	m := CreateNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateNodeSystemConfigFileAction_Invoke_SendError exercises CreateNodeSystemConfigFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateNodeSystemConfigFileAction_Invoke_SendError(t *testing.T) {
	r := &CreateNodeSystemConfigFileAction{client: newTransportErrorClient(t)}
	m := CreateNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateNodeSystemConfigFileAction_Invoke_APIError exercises CreateNodeSystemConfigFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateNodeSystemConfigFileAction_Invoke_APIError(t *testing.T) {
	r := &CreateNodeSystemConfigFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_node_system_config_file")
}

// TestCreateNodeSystemConfigFileAction_Invoke_APIErrorReadBody exercises CreateNodeSystemConfigFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateNodeSystemConfigFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateNodeSystemConfigFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateNodeSystemConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
