package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConfigureEmailServerAction_Invoke_Happy exercises ConfigureEmailServerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConfigureEmailServerAction_Invoke_Happy(t *testing.T) {
	r := &ConfigureEmailServerAction{client: newMockClientStatus(t, 200, "{}")}
	m := ConfigureEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConfigureEmailServerAction_Invoke_NilClient exercises ConfigureEmailServerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConfigureEmailServerAction_Invoke_NilClient(t *testing.T) {
	r := &ConfigureEmailServerAction{}
	m := ConfigureEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConfigureEmailServerAction_Invoke_BuildError exercises ConfigureEmailServerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConfigureEmailServerAction_Invoke_BuildError(t *testing.T) {
	r := &ConfigureEmailServerAction{client: newMalformedBaseURLClient(t)}
	m := ConfigureEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConfigureEmailServerAction_Invoke_SendError exercises ConfigureEmailServerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConfigureEmailServerAction_Invoke_SendError(t *testing.T) {
	r := &ConfigureEmailServerAction{client: newTransportErrorClient(t)}
	m := ConfigureEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConfigureEmailServerAction_Invoke_APIError exercises ConfigureEmailServerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConfigureEmailServerAction_Invoke_APIError(t *testing.T) {
	r := &ConfigureEmailServerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ConfigureEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_configure_email_server")
}

// TestConfigureEmailServerAction_Invoke_APIErrorReadBody exercises ConfigureEmailServerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConfigureEmailServerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConfigureEmailServerAction{client: newMockClientReadErrorBody(t, 501)}
	m := ConfigureEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
