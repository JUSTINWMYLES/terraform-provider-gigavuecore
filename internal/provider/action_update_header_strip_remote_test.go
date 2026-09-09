package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateHeaderStripAction_Invoke_Happy exercises UpdateHeaderStripAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateHeaderStripAction_Invoke_Happy(t *testing.T) {
	r := &UpdateHeaderStripAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateHeaderStripAction_Invoke_NilClient exercises UpdateHeaderStripAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateHeaderStripAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateHeaderStripAction{}
	m := UpdateHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateHeaderStripAction_Invoke_BuildError exercises UpdateHeaderStripAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateHeaderStripAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateHeaderStripAction{client: newMalformedBaseURLClient(t)}
	m := UpdateHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateHeaderStripAction_Invoke_SendError exercises UpdateHeaderStripAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateHeaderStripAction_Invoke_SendError(t *testing.T) {
	r := &UpdateHeaderStripAction{client: newTransportErrorClient(t)}
	m := UpdateHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateHeaderStripAction_Invoke_APIError exercises UpdateHeaderStripAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateHeaderStripAction_Invoke_APIError(t *testing.T) {
	r := &UpdateHeaderStripAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_header_strip")
}

// TestUpdateHeaderStripAction_Invoke_APIErrorReadBody exercises UpdateHeaderStripAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateHeaderStripAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateHeaderStripAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateHeaderStripActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
