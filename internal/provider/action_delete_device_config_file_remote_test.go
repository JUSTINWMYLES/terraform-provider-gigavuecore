package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteDeviceConfigFileAction_Invoke_Happy exercises DeleteDeviceConfigFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteDeviceConfigFileAction_Invoke_Happy(t *testing.T) {
	r := &DeleteDeviceConfigFileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteDeviceConfigFileAction_Invoke_NilClient exercises DeleteDeviceConfigFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteDeviceConfigFileAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteDeviceConfigFileAction{}
	m := DeleteDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteDeviceConfigFileAction_Invoke_BuildError exercises DeleteDeviceConfigFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteDeviceConfigFileAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteDeviceConfigFileAction{client: newMalformedBaseURLClient(t)}
	m := DeleteDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteDeviceConfigFileAction_Invoke_SendError exercises DeleteDeviceConfigFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteDeviceConfigFileAction_Invoke_SendError(t *testing.T) {
	r := &DeleteDeviceConfigFileAction{client: newTransportErrorClient(t)}
	m := DeleteDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteDeviceConfigFileAction_Invoke_APIError exercises DeleteDeviceConfigFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteDeviceConfigFileAction_Invoke_APIError(t *testing.T) {
	r := &DeleteDeviceConfigFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_device_config_file")
}

// TestDeleteDeviceConfigFileAction_Invoke_APIErrorReadBody exercises DeleteDeviceConfigFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteDeviceConfigFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteDeviceConfigFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteDeviceConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
