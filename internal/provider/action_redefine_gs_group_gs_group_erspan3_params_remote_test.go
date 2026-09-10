package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_Happy exercises RedefineGsGroupGsGroupErspan3ParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupGsGroupErspan3ParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupGsGroupErspan3ParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_NilClient exercises RedefineGsGroupGsGroupErspan3ParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupGsGroupErspan3ParamsAction{}
	m := RedefineGsGroupGsGroupErspan3ParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_BuildError exercises RedefineGsGroupGsGroupErspan3ParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupGsGroupErspan3ParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupGsGroupErspan3ParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_SendError exercises RedefineGsGroupGsGroupErspan3ParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupGsGroupErspan3ParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupGsGroupErspan3ParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_APIError exercises RedefineGsGroupGsGroupErspan3ParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupGsGroupErspan3ParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupGsGroupErspan3ParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_gs_group_erspan3_params")
}

// TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupGsGroupErspan3ParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupGsGroupErspan3ParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupGsGroupErspan3ParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupGsGroupErspan3ParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
