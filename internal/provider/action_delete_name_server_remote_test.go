package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteNameServerAction_Invoke_Happy exercises DeleteNameServerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteNameServerAction_Invoke_Happy(t *testing.T) {
	r := &DeleteNameServerAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteNameServerAction_Invoke_NilClient exercises DeleteNameServerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteNameServerAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteNameServerAction{}
	m := DeleteNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteNameServerAction_Invoke_BuildError exercises DeleteNameServerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteNameServerAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteNameServerAction{client: newMalformedBaseURLClient(t)}
	m := DeleteNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteNameServerAction_Invoke_SendError exercises DeleteNameServerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteNameServerAction_Invoke_SendError(t *testing.T) {
	r := &DeleteNameServerAction{client: newTransportErrorClient(t)}
	m := DeleteNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteNameServerAction_Invoke_APIError exercises DeleteNameServerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteNameServerAction_Invoke_APIError(t *testing.T) {
	r := &DeleteNameServerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_name_server")
}

// TestDeleteNameServerAction_Invoke_APIErrorReadBody exercises DeleteNameServerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteNameServerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteNameServerAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteNameServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
