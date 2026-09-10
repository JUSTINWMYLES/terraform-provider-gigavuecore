package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateKeystorePreferenceAction_Invoke_Happy exercises UpdateKeystorePreferenceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateKeystorePreferenceAction_Invoke_Happy(t *testing.T) {
	r := &UpdateKeystorePreferenceAction{client: newMockClientStatus(t, 201, "{}")}
	m := UpdateKeystorePreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateKeystorePreferenceAction_Invoke_NilClient exercises UpdateKeystorePreferenceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateKeystorePreferenceAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateKeystorePreferenceAction{}
	m := UpdateKeystorePreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateKeystorePreferenceAction_Invoke_BuildError exercises UpdateKeystorePreferenceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateKeystorePreferenceAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateKeystorePreferenceAction{client: newMalformedBaseURLClient(t)}
	m := UpdateKeystorePreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateKeystorePreferenceAction_Invoke_SendError exercises UpdateKeystorePreferenceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateKeystorePreferenceAction_Invoke_SendError(t *testing.T) {
	r := &UpdateKeystorePreferenceAction{client: newTransportErrorClient(t)}
	m := UpdateKeystorePreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateKeystorePreferenceAction_Invoke_APIError exercises UpdateKeystorePreferenceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateKeystorePreferenceAction_Invoke_APIError(t *testing.T) {
	r := &UpdateKeystorePreferenceAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := UpdateKeystorePreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_keystore_preference")
}

// TestUpdateKeystorePreferenceAction_Invoke_APIErrorReadBody exercises UpdateKeystorePreferenceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateKeystorePreferenceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateKeystorePreferenceAction{client: newMockClientReadErrorBody(t, 500)}
	m := UpdateKeystorePreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
