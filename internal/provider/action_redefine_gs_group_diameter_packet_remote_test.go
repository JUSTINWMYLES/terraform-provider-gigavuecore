package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupDiameterPacketAction_Invoke_Happy exercises RedefineGsGroupDiameterPacketAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupDiameterPacketAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupDiameterPacketAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupDiameterPacketActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupDiameterPacketAction_Invoke_NilClient exercises RedefineGsGroupDiameterPacketAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupDiameterPacketAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupDiameterPacketAction{}
	m := RedefineGsGroupDiameterPacketActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupDiameterPacketAction_Invoke_BuildError exercises RedefineGsGroupDiameterPacketAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupDiameterPacketAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupDiameterPacketAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupDiameterPacketActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupDiameterPacketAction_Invoke_SendError exercises RedefineGsGroupDiameterPacketAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupDiameterPacketAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupDiameterPacketAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupDiameterPacketActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupDiameterPacketAction_Invoke_APIError exercises RedefineGsGroupDiameterPacketAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupDiameterPacketAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupDiameterPacketAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupDiameterPacketActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_diameter_packet")
}

// TestRedefineGsGroupDiameterPacketAction_Invoke_APIErrorReadBody exercises RedefineGsGroupDiameterPacketAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupDiameterPacketAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupDiameterPacketAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupDiameterPacketActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
