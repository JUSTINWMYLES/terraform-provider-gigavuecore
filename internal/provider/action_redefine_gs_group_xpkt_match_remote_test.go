package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupXpktMatchAction_Invoke_Happy exercises RedefineGsGroupXpktMatchAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupXpktMatchAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupXpktMatchAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupXpktMatchActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupXpktMatchAction_Invoke_NilClient exercises RedefineGsGroupXpktMatchAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupXpktMatchAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupXpktMatchAction{}
	m := RedefineGsGroupXpktMatchActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupXpktMatchAction_Invoke_BuildError exercises RedefineGsGroupXpktMatchAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupXpktMatchAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupXpktMatchAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupXpktMatchActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupXpktMatchAction_Invoke_SendError exercises RedefineGsGroupXpktMatchAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupXpktMatchAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupXpktMatchAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupXpktMatchActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupXpktMatchAction_Invoke_APIError exercises RedefineGsGroupXpktMatchAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupXpktMatchAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupXpktMatchAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupXpktMatchActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_xpkt_match")
}

// TestRedefineGsGroupXpktMatchAction_Invoke_APIErrorReadBody exercises RedefineGsGroupXpktMatchAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupXpktMatchAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupXpktMatchAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupXpktMatchActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
