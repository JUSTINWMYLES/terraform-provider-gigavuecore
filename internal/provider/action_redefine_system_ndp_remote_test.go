package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineSystemNdpAction_Invoke_Happy exercises RedefineSystemNdpAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineSystemNdpAction_Invoke_Happy(t *testing.T) {
	r := &RedefineSystemNdpAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineSystemNdpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineSystemNdpAction_Invoke_NilClient exercises RedefineSystemNdpAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineSystemNdpAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineSystemNdpAction{}
	m := RedefineSystemNdpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineSystemNdpAction_Invoke_BuildError exercises RedefineSystemNdpAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineSystemNdpAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineSystemNdpAction{client: newMalformedBaseURLClient(t)}
	m := RedefineSystemNdpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineSystemNdpAction_Invoke_SendError exercises RedefineSystemNdpAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineSystemNdpAction_Invoke_SendError(t *testing.T) {
	r := &RedefineSystemNdpAction{client: newTransportErrorClient(t)}
	m := RedefineSystemNdpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineSystemNdpAction_Invoke_APIError exercises RedefineSystemNdpAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineSystemNdpAction_Invoke_APIError(t *testing.T) {
	r := &RedefineSystemNdpAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineSystemNdpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_system_ndp")
}

// TestRedefineSystemNdpAction_Invoke_APIErrorReadBody exercises RedefineSystemNdpAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineSystemNdpAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineSystemNdpAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineSystemNdpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
