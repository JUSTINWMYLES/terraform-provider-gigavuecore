package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_Happy exercises CreateApiTokenFmSystemApiRateLimitingAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_Happy(t *testing.T) {
	r := &CreateApiTokenFmSystemApiRateLimitingAction{client: newMockClientStatus(t, 200, "{}")}
	m := CreateApiTokenFmSystemApiRateLimitingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_NilClient exercises CreateApiTokenFmSystemApiRateLimitingAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_NilClient(t *testing.T) {
	r := &CreateApiTokenFmSystemApiRateLimitingAction{}
	m := CreateApiTokenFmSystemApiRateLimitingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_BuildError exercises CreateApiTokenFmSystemApiRateLimitingAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_BuildError(t *testing.T) {
	r := &CreateApiTokenFmSystemApiRateLimitingAction{client: newMalformedBaseURLClient(t)}
	m := CreateApiTokenFmSystemApiRateLimitingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_SendError exercises CreateApiTokenFmSystemApiRateLimitingAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_SendError(t *testing.T) {
	r := &CreateApiTokenFmSystemApiRateLimitingAction{client: newTransportErrorClient(t)}
	m := CreateApiTokenFmSystemApiRateLimitingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_APIError exercises CreateApiTokenFmSystemApiRateLimitingAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_APIError(t *testing.T) {
	r := &CreateApiTokenFmSystemApiRateLimitingAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateApiTokenFmSystemApiRateLimitingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting")
}

// TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_APIErrorReadBody exercises CreateApiTokenFmSystemApiRateLimitingAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateApiTokenFmSystemApiRateLimitingAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateApiTokenFmSystemApiRateLimitingAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateApiTokenFmSystemApiRateLimitingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
