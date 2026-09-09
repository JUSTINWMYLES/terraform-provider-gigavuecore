package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_Happy exercises RedefineGsGroupParamsGtpPersistenceParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupParamsGtpPersistenceParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupParamsGtpPersistenceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_NilClient exercises RedefineGsGroupParamsGtpPersistenceParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupParamsGtpPersistenceParamsAction{}
	m := RedefineGsGroupParamsGtpPersistenceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_BuildError exercises RedefineGsGroupParamsGtpPersistenceParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupParamsGtpPersistenceParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupParamsGtpPersistenceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_SendError exercises RedefineGsGroupParamsGtpPersistenceParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupParamsGtpPersistenceParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupParamsGtpPersistenceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_APIError exercises RedefineGsGroupParamsGtpPersistenceParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupParamsGtpPersistenceParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupParamsGtpPersistenceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_params_gtp_persistence_params")
}

// TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupParamsGtpPersistenceParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupParamsGtpPersistenceParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupParamsGtpPersistenceParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupParamsGtpPersistenceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
