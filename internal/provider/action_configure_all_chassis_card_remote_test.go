package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConfigureAllChassisCardAction_Invoke_Happy exercises ConfigureAllChassisCardAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConfigureAllChassisCardAction_Invoke_Happy(t *testing.T) {
	r := &ConfigureAllChassisCardAction{client: newMockClientStatus(t, 201, "{}")}
	m := ConfigureAllChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConfigureAllChassisCardAction_Invoke_NilClient exercises ConfigureAllChassisCardAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConfigureAllChassisCardAction_Invoke_NilClient(t *testing.T) {
	r := &ConfigureAllChassisCardAction{}
	m := ConfigureAllChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConfigureAllChassisCardAction_Invoke_BuildError exercises ConfigureAllChassisCardAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConfigureAllChassisCardAction_Invoke_BuildError(t *testing.T) {
	r := &ConfigureAllChassisCardAction{client: newMalformedBaseURLClient(t)}
	m := ConfigureAllChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConfigureAllChassisCardAction_Invoke_SendError exercises ConfigureAllChassisCardAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConfigureAllChassisCardAction_Invoke_SendError(t *testing.T) {
	r := &ConfigureAllChassisCardAction{client: newTransportErrorClient(t)}
	m := ConfigureAllChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConfigureAllChassisCardAction_Invoke_APIError exercises ConfigureAllChassisCardAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConfigureAllChassisCardAction_Invoke_APIError(t *testing.T) {
	r := &ConfigureAllChassisCardAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ConfigureAllChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_configure_all_chassis_card")
}

// TestConfigureAllChassisCardAction_Invoke_APIErrorReadBody exercises ConfigureAllChassisCardAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConfigureAllChassisCardAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConfigureAllChassisCardAction{client: newMockClientReadErrorBody(t, 501)}
	m := ConfigureAllChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
