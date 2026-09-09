package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllHeaderStripAction_Invoke_Happy exercises DeleteAllHeaderStripAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllHeaderStripAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllHeaderStripAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllHeaderStripAction_Invoke_NilClient exercises DeleteAllHeaderStripAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllHeaderStripAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllHeaderStripAction{}
	m := DeleteAllHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllHeaderStripAction_Invoke_BuildError exercises DeleteAllHeaderStripAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllHeaderStripAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllHeaderStripAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllHeaderStripAction_Invoke_SendError exercises DeleteAllHeaderStripAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllHeaderStripAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllHeaderStripAction{client: newTransportErrorClient(t)}
	m := DeleteAllHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllHeaderStripAction_Invoke_APIError exercises DeleteAllHeaderStripAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllHeaderStripAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllHeaderStripAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_header_strip")
}

// TestDeleteAllHeaderStripAction_Invoke_APIErrorReadBody exercises DeleteAllHeaderStripAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllHeaderStripAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllHeaderStripAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
