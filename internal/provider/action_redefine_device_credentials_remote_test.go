package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineDeviceCredentialsAction_Invoke_Happy exercises RedefineDeviceCredentialsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineDeviceCredentialsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineDeviceCredentialsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineDeviceCredentialsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineDeviceCredentialsAction_Invoke_NilClient exercises RedefineDeviceCredentialsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineDeviceCredentialsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineDeviceCredentialsAction{}
	m := RedefineDeviceCredentialsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineDeviceCredentialsAction_Invoke_BuildError exercises RedefineDeviceCredentialsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineDeviceCredentialsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineDeviceCredentialsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineDeviceCredentialsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineDeviceCredentialsAction_Invoke_SendError exercises RedefineDeviceCredentialsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineDeviceCredentialsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineDeviceCredentialsAction{client: newTransportErrorClient(t)}
	m := RedefineDeviceCredentialsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineDeviceCredentialsAction_Invoke_APIError exercises RedefineDeviceCredentialsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineDeviceCredentialsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineDeviceCredentialsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineDeviceCredentialsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_device_credentials")
}

// TestRedefineDeviceCredentialsAction_Invoke_APIErrorReadBody exercises RedefineDeviceCredentialsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineDeviceCredentialsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineDeviceCredentialsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineDeviceCredentialsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
