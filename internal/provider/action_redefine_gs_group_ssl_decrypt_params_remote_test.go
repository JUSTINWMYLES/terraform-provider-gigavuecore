package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupSslDecryptParamsAction_Invoke_Happy exercises RedefineGsGroupSslDecryptParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupSslDecryptParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupSslDecryptParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupSslDecryptParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupSslDecryptParamsAction_Invoke_NilClient exercises RedefineGsGroupSslDecryptParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupSslDecryptParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupSslDecryptParamsAction{}
	m := RedefineGsGroupSslDecryptParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupSslDecryptParamsAction_Invoke_BuildError exercises RedefineGsGroupSslDecryptParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupSslDecryptParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupSslDecryptParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupSslDecryptParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupSslDecryptParamsAction_Invoke_SendError exercises RedefineGsGroupSslDecryptParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupSslDecryptParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupSslDecryptParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupSslDecryptParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupSslDecryptParamsAction_Invoke_APIError exercises RedefineGsGroupSslDecryptParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupSslDecryptParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupSslDecryptParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupSslDecryptParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_ssl_decrypt_params")
}

// TestRedefineGsGroupSslDecryptParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupSslDecryptParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupSslDecryptParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupSslDecryptParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupSslDecryptParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
