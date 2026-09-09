package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestActivetDeviceConfigFileAction_Invoke_Happy exercises ActivetDeviceConfigFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestActivetDeviceConfigFileAction_Invoke_Happy(t *testing.T) {
	r := &ActivetDeviceConfigFileAction{client: newMockClientStatus(t, 200, "{}")}
	m := ActivetDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestActivetDeviceConfigFileAction_Invoke_NilClient exercises ActivetDeviceConfigFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestActivetDeviceConfigFileAction_Invoke_NilClient(t *testing.T) {
	r := &ActivetDeviceConfigFileAction{}
	m := ActivetDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestActivetDeviceConfigFileAction_Invoke_BuildError exercises ActivetDeviceConfigFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestActivetDeviceConfigFileAction_Invoke_BuildError(t *testing.T) {
	r := &ActivetDeviceConfigFileAction{client: newMalformedBaseURLClient(t)}
	m := ActivetDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestActivetDeviceConfigFileAction_Invoke_SendError exercises ActivetDeviceConfigFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestActivetDeviceConfigFileAction_Invoke_SendError(t *testing.T) {
	r := &ActivetDeviceConfigFileAction{client: newTransportErrorClient(t)}
	m := ActivetDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestActivetDeviceConfigFileAction_Invoke_APIError exercises ActivetDeviceConfigFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestActivetDeviceConfigFileAction_Invoke_APIError(t *testing.T) {
	r := &ActivetDeviceConfigFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ActivetDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_activet_device_config_file")
}

// TestActivetDeviceConfigFileAction_Invoke_APIErrorReadBody exercises ActivetDeviceConfigFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestActivetDeviceConfigFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ActivetDeviceConfigFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := ActivetDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
