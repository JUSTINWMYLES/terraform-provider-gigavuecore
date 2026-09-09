package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUnconfigureChassisCardAction_Invoke_Happy exercises UnconfigureChassisCardAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUnconfigureChassisCardAction_Invoke_Happy(t *testing.T) {
	r := &UnconfigureChassisCardAction{client: newMockClientStatus(t, 204, "{}")}
	m := UnconfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUnconfigureChassisCardAction_Invoke_NilClient exercises UnconfigureChassisCardAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUnconfigureChassisCardAction_Invoke_NilClient(t *testing.T) {
	r := &UnconfigureChassisCardAction{}
	m := UnconfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUnconfigureChassisCardAction_Invoke_BuildError exercises UnconfigureChassisCardAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUnconfigureChassisCardAction_Invoke_BuildError(t *testing.T) {
	r := &UnconfigureChassisCardAction{client: newMalformedBaseURLClient(t)}
	m := UnconfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUnconfigureChassisCardAction_Invoke_SendError exercises UnconfigureChassisCardAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUnconfigureChassisCardAction_Invoke_SendError(t *testing.T) {
	r := &UnconfigureChassisCardAction{client: newTransportErrorClient(t)}
	m := UnconfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUnconfigureChassisCardAction_Invoke_APIError exercises UnconfigureChassisCardAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUnconfigureChassisCardAction_Invoke_APIError(t *testing.T) {
	r := &UnconfigureChassisCardAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UnconfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_unconfigure_chassis_card")
}

// TestUnconfigureChassisCardAction_Invoke_APIErrorReadBody exercises UnconfigureChassisCardAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUnconfigureChassisCardAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UnconfigureChassisCardAction{client: newMockClientReadErrorBody(t, 501)}
	m := UnconfigureChassisCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
