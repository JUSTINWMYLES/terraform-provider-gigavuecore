package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestPurgeStatsAction_Invoke_Happy exercises PurgeStatsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestPurgeStatsAction_Invoke_Happy(t *testing.T) {
	r := &PurgeStatsAction{client: newMockClientStatus(t, 204, "{}")}
	m := PurgeStatsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPurgeStatsAction_Invoke_NilClient exercises PurgeStatsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPurgeStatsAction_Invoke_NilClient(t *testing.T) {
	r := &PurgeStatsAction{}
	m := PurgeStatsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPurgeStatsAction_Invoke_BuildError exercises PurgeStatsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPurgeStatsAction_Invoke_BuildError(t *testing.T) {
	r := &PurgeStatsAction{client: newMalformedBaseURLClient(t)}
	m := PurgeStatsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPurgeStatsAction_Invoke_SendError exercises PurgeStatsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestPurgeStatsAction_Invoke_SendError(t *testing.T) {
	r := &PurgeStatsAction{client: newTransportErrorClient(t)}
	m := PurgeStatsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPurgeStatsAction_Invoke_APIError exercises PurgeStatsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPurgeStatsAction_Invoke_APIError(t *testing.T) {
	r := &PurgeStatsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PurgeStatsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_purge_stats")
}

// TestPurgeStatsAction_Invoke_APIErrorReadBody exercises PurgeStatsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPurgeStatsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &PurgeStatsAction{client: newMockClientReadErrorBody(t, 501)}
	m := PurgeStatsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
