package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestEnableFipsModeAction_Invoke_Happy exercises EnableFipsModeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestEnableFipsModeAction_Invoke_Happy(t *testing.T) {
	r := &EnableFipsModeAction{client: newMockClientStatus(t, 200, "{}")}
	m := EnableFipsModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnableFipsModeAction_Invoke_NilClient exercises EnableFipsModeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEnableFipsModeAction_Invoke_NilClient(t *testing.T) {
	r := &EnableFipsModeAction{}
	m := EnableFipsModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEnableFipsModeAction_Invoke_BuildError exercises EnableFipsModeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEnableFipsModeAction_Invoke_BuildError(t *testing.T) {
	r := &EnableFipsModeAction{client: newMalformedBaseURLClient(t)}
	m := EnableFipsModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEnableFipsModeAction_Invoke_SendError exercises EnableFipsModeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestEnableFipsModeAction_Invoke_SendError(t *testing.T) {
	r := &EnableFipsModeAction{client: newTransportErrorClient(t)}
	m := EnableFipsModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEnableFipsModeAction_Invoke_APIError exercises EnableFipsModeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEnableFipsModeAction_Invoke_APIError(t *testing.T) {
	r := &EnableFipsModeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EnableFipsModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_enable_fips_mode")
}

// TestEnableFipsModeAction_Invoke_APIErrorReadBody exercises EnableFipsModeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEnableFipsModeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &EnableFipsModeAction{client: newMockClientReadErrorBody(t, 501)}
	m := EnableFipsModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
