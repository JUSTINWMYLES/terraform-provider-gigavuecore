package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestEnableCryptoModeAction_Invoke_Happy exercises EnableCryptoModeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestEnableCryptoModeAction_Invoke_Happy(t *testing.T) {
	r := &EnableCryptoModeAction{client: newMockClientStatus(t, 200, "{}")}
	m := EnableCryptoModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnableCryptoModeAction_Invoke_NilClient exercises EnableCryptoModeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEnableCryptoModeAction_Invoke_NilClient(t *testing.T) {
	r := &EnableCryptoModeAction{}
	m := EnableCryptoModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEnableCryptoModeAction_Invoke_BuildError exercises EnableCryptoModeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEnableCryptoModeAction_Invoke_BuildError(t *testing.T) {
	r := &EnableCryptoModeAction{client: newMalformedBaseURLClient(t)}
	m := EnableCryptoModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEnableCryptoModeAction_Invoke_SendError exercises EnableCryptoModeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestEnableCryptoModeAction_Invoke_SendError(t *testing.T) {
	r := &EnableCryptoModeAction{client: newTransportErrorClient(t)}
	m := EnableCryptoModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEnableCryptoModeAction_Invoke_APIError exercises EnableCryptoModeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEnableCryptoModeAction_Invoke_APIError(t *testing.T) {
	r := &EnableCryptoModeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EnableCryptoModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_enable_crypto_mode")
}

// TestEnableCryptoModeAction_Invoke_APIErrorReadBody exercises EnableCryptoModeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEnableCryptoModeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &EnableCryptoModeAction{client: newMockClientReadErrorBody(t, 501)}
	m := EnableCryptoModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
