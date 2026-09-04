package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRegisterOrDeRegisterNrtAction_Invoke_Happy exercises RegisterOrDeRegisterNrtAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRegisterOrDeRegisterNrtAction_Invoke_Happy(t *testing.T) {
	r := &RegisterOrDeRegisterNrtAction{client: newMockClientStatus(t, 200, "{}")}
	m := RegisterOrDeRegisterNrtActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRegisterOrDeRegisterNrtAction_Invoke_NilClient exercises RegisterOrDeRegisterNrtAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRegisterOrDeRegisterNrtAction_Invoke_NilClient(t *testing.T) {
	r := &RegisterOrDeRegisterNrtAction{}
	m := RegisterOrDeRegisterNrtActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRegisterOrDeRegisterNrtAction_Invoke_BuildError exercises RegisterOrDeRegisterNrtAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRegisterOrDeRegisterNrtAction_Invoke_BuildError(t *testing.T) {
	r := &RegisterOrDeRegisterNrtAction{client: newMalformedBaseURLClient(t)}
	m := RegisterOrDeRegisterNrtActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRegisterOrDeRegisterNrtAction_Invoke_SendError exercises RegisterOrDeRegisterNrtAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRegisterOrDeRegisterNrtAction_Invoke_SendError(t *testing.T) {
	r := &RegisterOrDeRegisterNrtAction{client: newTransportErrorClient(t)}
	m := RegisterOrDeRegisterNrtActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRegisterOrDeRegisterNrtAction_Invoke_APIError exercises RegisterOrDeRegisterNrtAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRegisterOrDeRegisterNrtAction_Invoke_APIError(t *testing.T) {
	r := &RegisterOrDeRegisterNrtAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RegisterOrDeRegisterNrtActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_register_or_de_register_nrt")
}

// TestRegisterOrDeRegisterNrtAction_Invoke_APIErrorReadBody exercises RegisterOrDeRegisterNrtAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRegisterOrDeRegisterNrtAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RegisterOrDeRegisterNrtAction{client: newMockClientReadErrorBody(t, 501)}
	m := RegisterOrDeRegisterNrtActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
