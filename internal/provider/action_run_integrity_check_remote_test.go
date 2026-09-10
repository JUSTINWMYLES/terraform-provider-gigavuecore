package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRunIntegrityCheckAction_Invoke_Happy exercises RunIntegrityCheckAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRunIntegrityCheckAction_Invoke_Happy(t *testing.T) {
	r := &RunIntegrityCheckAction{client: newMockClientStatus(t, 200, "{}")}
	m := RunIntegrityCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRunIntegrityCheckAction_Invoke_NilClient exercises RunIntegrityCheckAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRunIntegrityCheckAction_Invoke_NilClient(t *testing.T) {
	r := &RunIntegrityCheckAction{}
	m := RunIntegrityCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRunIntegrityCheckAction_Invoke_BuildError exercises RunIntegrityCheckAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRunIntegrityCheckAction_Invoke_BuildError(t *testing.T) {
	r := &RunIntegrityCheckAction{client: newMalformedBaseURLClient(t)}
	m := RunIntegrityCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRunIntegrityCheckAction_Invoke_SendError exercises RunIntegrityCheckAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRunIntegrityCheckAction_Invoke_SendError(t *testing.T) {
	r := &RunIntegrityCheckAction{client: newTransportErrorClient(t)}
	m := RunIntegrityCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRunIntegrityCheckAction_Invoke_APIError exercises RunIntegrityCheckAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRunIntegrityCheckAction_Invoke_APIError(t *testing.T) {
	r := &RunIntegrityCheckAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RunIntegrityCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_run_integrity_check")
}

// TestRunIntegrityCheckAction_Invoke_APIErrorReadBody exercises RunIntegrityCheckAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRunIntegrityCheckAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RunIntegrityCheckAction{client: newMockClientReadErrorBody(t, 501)}
	m := RunIntegrityCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
