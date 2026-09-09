package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRegisterWithEmsAction_Invoke_Happy exercises RegisterWithEmsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRegisterWithEmsAction_Invoke_Happy(t *testing.T) {
	r := &RegisterWithEmsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RegisterWithEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRegisterWithEmsAction_Invoke_NilClient exercises RegisterWithEmsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRegisterWithEmsAction_Invoke_NilClient(t *testing.T) {
	r := &RegisterWithEmsAction{}
	m := RegisterWithEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRegisterWithEmsAction_Invoke_BuildError exercises RegisterWithEmsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRegisterWithEmsAction_Invoke_BuildError(t *testing.T) {
	r := &RegisterWithEmsAction{client: newMalformedBaseURLClient(t)}
	m := RegisterWithEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRegisterWithEmsAction_Invoke_SendError exercises RegisterWithEmsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRegisterWithEmsAction_Invoke_SendError(t *testing.T) {
	r := &RegisterWithEmsAction{client: newTransportErrorClient(t)}
	m := RegisterWithEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRegisterWithEmsAction_Invoke_APIError exercises RegisterWithEmsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRegisterWithEmsAction_Invoke_APIError(t *testing.T) {
	r := &RegisterWithEmsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RegisterWithEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_register_with_ems")
}

// TestRegisterWithEmsAction_Invoke_APIErrorReadBody exercises RegisterWithEmsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRegisterWithEmsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RegisterWithEmsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RegisterWithEmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
