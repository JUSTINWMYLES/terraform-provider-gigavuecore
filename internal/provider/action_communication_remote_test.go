package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCommunicationAction_Invoke_Happy exercises CommunicationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCommunicationAction_Invoke_Happy(t *testing.T) {
	r := &CommunicationAction{client: newMockClientStatus(t, 200, "{}")}
	m := CommunicationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCommunicationAction_Invoke_NilClient exercises CommunicationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCommunicationAction_Invoke_NilClient(t *testing.T) {
	r := &CommunicationAction{}
	m := CommunicationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCommunicationAction_Invoke_BuildError exercises CommunicationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCommunicationAction_Invoke_BuildError(t *testing.T) {
	r := &CommunicationAction{client: newMalformedBaseURLClient(t)}
	m := CommunicationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCommunicationAction_Invoke_SendError exercises CommunicationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCommunicationAction_Invoke_SendError(t *testing.T) {
	r := &CommunicationAction{client: newTransportErrorClient(t)}
	m := CommunicationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCommunicationAction_Invoke_APIError exercises CommunicationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCommunicationAction_Invoke_APIError(t *testing.T) {
	r := &CommunicationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CommunicationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_communication")
}

// TestCommunicationAction_Invoke_APIErrorReadBody exercises CommunicationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCommunicationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CommunicationAction{client: newMockClientReadErrorBody(t, 501)}
	m := CommunicationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
