package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestGenerateKeystoreKeyAction_Invoke_Happy exercises GenerateKeystoreKeyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestGenerateKeystoreKeyAction_Invoke_Happy(t *testing.T) {
	r := &GenerateKeystoreKeyAction{client: newMockClientStatus(t, 201, "{}")}
	m := GenerateKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGenerateKeystoreKeyAction_Invoke_NilClient exercises GenerateKeystoreKeyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGenerateKeystoreKeyAction_Invoke_NilClient(t *testing.T) {
	r := &GenerateKeystoreKeyAction{}
	m := GenerateKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGenerateKeystoreKeyAction_Invoke_BuildError exercises GenerateKeystoreKeyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGenerateKeystoreKeyAction_Invoke_BuildError(t *testing.T) {
	r := &GenerateKeystoreKeyAction{client: newMalformedBaseURLClient(t)}
	m := GenerateKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGenerateKeystoreKeyAction_Invoke_SendError exercises GenerateKeystoreKeyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestGenerateKeystoreKeyAction_Invoke_SendError(t *testing.T) {
	r := &GenerateKeystoreKeyAction{client: newTransportErrorClient(t)}
	m := GenerateKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGenerateKeystoreKeyAction_Invoke_APIError exercises GenerateKeystoreKeyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGenerateKeystoreKeyAction_Invoke_APIError(t *testing.T) {
	r := &GenerateKeystoreKeyAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GenerateKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_generate_keystore_key")
}

// TestGenerateKeystoreKeyAction_Invoke_APIErrorReadBody exercises GenerateKeystoreKeyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGenerateKeystoreKeyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &GenerateKeystoreKeyAction{client: newMockClientReadErrorBody(t, 501)}
	m := GenerateKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
