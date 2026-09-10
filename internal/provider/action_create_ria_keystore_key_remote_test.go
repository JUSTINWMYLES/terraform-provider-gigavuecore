package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateRiaKeystoreKeyAction_Invoke_Happy exercises CreateRiaKeystoreKeyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateRiaKeystoreKeyAction_Invoke_Happy(t *testing.T) {
	r := &CreateRiaKeystoreKeyAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateRiaKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateRiaKeystoreKeyAction_Invoke_NilClient exercises CreateRiaKeystoreKeyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateRiaKeystoreKeyAction_Invoke_NilClient(t *testing.T) {
	r := &CreateRiaKeystoreKeyAction{}
	m := CreateRiaKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateRiaKeystoreKeyAction_Invoke_BuildError exercises CreateRiaKeystoreKeyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateRiaKeystoreKeyAction_Invoke_BuildError(t *testing.T) {
	r := &CreateRiaKeystoreKeyAction{client: newMalformedBaseURLClient(t)}
	m := CreateRiaKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateRiaKeystoreKeyAction_Invoke_SendError exercises CreateRiaKeystoreKeyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateRiaKeystoreKeyAction_Invoke_SendError(t *testing.T) {
	r := &CreateRiaKeystoreKeyAction{client: newTransportErrorClient(t)}
	m := CreateRiaKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateRiaKeystoreKeyAction_Invoke_APIError exercises CreateRiaKeystoreKeyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateRiaKeystoreKeyAction_Invoke_APIError(t *testing.T) {
	r := &CreateRiaKeystoreKeyAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateRiaKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_ria_keystore_key")
}

// TestCreateRiaKeystoreKeyAction_Invoke_APIErrorReadBody exercises CreateRiaKeystoreKeyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateRiaKeystoreKeyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateRiaKeystoreKeyAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateRiaKeystoreKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
