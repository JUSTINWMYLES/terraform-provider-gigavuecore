package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConfigureChassisCardAction_Invoke_Happy exercises ConfigureChassisCardAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConfigureChassisCardAction_Invoke_Happy(t *testing.T) {
	r := &ConfigureChassisCardAction{client: newMockClientStatus(t, 201, "{}")}
	m := ConfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConfigureChassisCardAction_Invoke_NilClient exercises ConfigureChassisCardAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConfigureChassisCardAction_Invoke_NilClient(t *testing.T) {
	r := &ConfigureChassisCardAction{}
	m := ConfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConfigureChassisCardAction_Invoke_BuildError exercises ConfigureChassisCardAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConfigureChassisCardAction_Invoke_BuildError(t *testing.T) {
	r := &ConfigureChassisCardAction{client: newMalformedBaseURLClient(t)}
	m := ConfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConfigureChassisCardAction_Invoke_SendError exercises ConfigureChassisCardAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConfigureChassisCardAction_Invoke_SendError(t *testing.T) {
	r := &ConfigureChassisCardAction{client: newTransportErrorClient(t)}
	m := ConfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConfigureChassisCardAction_Invoke_APIError exercises ConfigureChassisCardAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConfigureChassisCardAction_Invoke_APIError(t *testing.T) {
	r := &ConfigureChassisCardAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ConfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_configure_chassis_card")
}

// TestConfigureChassisCardAction_Invoke_APIErrorReadBody exercises ConfigureChassisCardAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConfigureChassisCardAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConfigureChassisCardAction{client: newMockClientReadErrorBody(t, 501)}
	m := ConfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
