package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddNameServerAction_Invoke_Happy exercises AddNameServerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddNameServerAction_Invoke_Happy(t *testing.T) {
	r := &AddNameServerAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddNameServerAction_Invoke_NilClient exercises AddNameServerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddNameServerAction_Invoke_NilClient(t *testing.T) {
	r := &AddNameServerAction{}
	m := AddNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddNameServerAction_Invoke_BuildError exercises AddNameServerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddNameServerAction_Invoke_BuildError(t *testing.T) {
	r := &AddNameServerAction{client: newMalformedBaseURLClient(t)}
	m := AddNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddNameServerAction_Invoke_SendError exercises AddNameServerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddNameServerAction_Invoke_SendError(t *testing.T) {
	r := &AddNameServerAction{client: newTransportErrorClient(t)}
	m := AddNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddNameServerAction_Invoke_APIError exercises AddNameServerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddNameServerAction_Invoke_APIError(t *testing.T) {
	r := &AddNameServerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_name_server")
}

// TestAddNameServerAction_Invoke_APIErrorReadBody exercises AddNameServerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddNameServerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddNameServerAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
