package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupNetflowParamsAction_Invoke_Happy exercises RedefineGsGroupNetflowParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupNetflowParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupNetflowParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupNetflowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupNetflowParamsAction_Invoke_NilClient exercises RedefineGsGroupNetflowParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupNetflowParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupNetflowParamsAction{}
	m := RedefineGsGroupNetflowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupNetflowParamsAction_Invoke_BuildError exercises RedefineGsGroupNetflowParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupNetflowParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupNetflowParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupNetflowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupNetflowParamsAction_Invoke_SendError exercises RedefineGsGroupNetflowParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupNetflowParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupNetflowParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupNetflowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupNetflowParamsAction_Invoke_APIError exercises RedefineGsGroupNetflowParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupNetflowParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupNetflowParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupNetflowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_netflow_params")
}

// TestRedefineGsGroupNetflowParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupNetflowParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupNetflowParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupNetflowParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupNetflowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
