package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUnconfigureAllChassisCardsAction_Invoke_Happy exercises UnconfigureAllChassisCardsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUnconfigureAllChassisCardsAction_Invoke_Happy(t *testing.T) {
	r := &UnconfigureAllChassisCardsAction{client: newMockClientStatus(t, 204, "{}")}
	m := UnconfigureAllChassisCardsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUnconfigureAllChassisCardsAction_Invoke_NilClient exercises UnconfigureAllChassisCardsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUnconfigureAllChassisCardsAction_Invoke_NilClient(t *testing.T) {
	r := &UnconfigureAllChassisCardsAction{}
	m := UnconfigureAllChassisCardsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUnconfigureAllChassisCardsAction_Invoke_BuildError exercises UnconfigureAllChassisCardsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUnconfigureAllChassisCardsAction_Invoke_BuildError(t *testing.T) {
	r := &UnconfigureAllChassisCardsAction{client: newMalformedBaseURLClient(t)}
	m := UnconfigureAllChassisCardsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUnconfigureAllChassisCardsAction_Invoke_SendError exercises UnconfigureAllChassisCardsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUnconfigureAllChassisCardsAction_Invoke_SendError(t *testing.T) {
	r := &UnconfigureAllChassisCardsAction{client: newTransportErrorClient(t)}
	m := UnconfigureAllChassisCardsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUnconfigureAllChassisCardsAction_Invoke_APIError exercises UnconfigureAllChassisCardsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUnconfigureAllChassisCardsAction_Invoke_APIError(t *testing.T) {
	r := &UnconfigureAllChassisCardsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UnconfigureAllChassisCardsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_unconfigure_all_chassis_cards")
}

// TestUnconfigureAllChassisCardsAction_Invoke_APIErrorReadBody exercises UnconfigureAllChassisCardsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUnconfigureAllChassisCardsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UnconfigureAllChassisCardsAction{client: newMockClientReadErrorBody(t, 501)}
	m := UnconfigureAllChassisCardsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
