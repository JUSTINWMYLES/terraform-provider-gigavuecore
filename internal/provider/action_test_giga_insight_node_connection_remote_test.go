package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestTestGigaInsightNodeConnectionAction_Invoke_Happy exercises TestGigaInsightNodeConnectionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestTestGigaInsightNodeConnectionAction_Invoke_Happy(t *testing.T) {
	r := &TestGigaInsightNodeConnectionAction{client: newMockClientStatus(t, 200, "{}")}
	m := TestGigaInsightNodeConnectionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTestGigaInsightNodeConnectionAction_Invoke_NilClient exercises TestGigaInsightNodeConnectionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTestGigaInsightNodeConnectionAction_Invoke_NilClient(t *testing.T) {
	r := &TestGigaInsightNodeConnectionAction{}
	m := TestGigaInsightNodeConnectionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTestGigaInsightNodeConnectionAction_Invoke_BuildError exercises TestGigaInsightNodeConnectionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTestGigaInsightNodeConnectionAction_Invoke_BuildError(t *testing.T) {
	r := &TestGigaInsightNodeConnectionAction{client: newMalformedBaseURLClient(t)}
	m := TestGigaInsightNodeConnectionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTestGigaInsightNodeConnectionAction_Invoke_SendError exercises TestGigaInsightNodeConnectionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestTestGigaInsightNodeConnectionAction_Invoke_SendError(t *testing.T) {
	r := &TestGigaInsightNodeConnectionAction{client: newTransportErrorClient(t)}
	m := TestGigaInsightNodeConnectionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTestGigaInsightNodeConnectionAction_Invoke_APIError exercises TestGigaInsightNodeConnectionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTestGigaInsightNodeConnectionAction_Invoke_APIError(t *testing.T) {
	r := &TestGigaInsightNodeConnectionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TestGigaInsightNodeConnectionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_test_giga_insight_node_connection")
}

// TestTestGigaInsightNodeConnectionAction_Invoke_APIErrorReadBody exercises TestGigaInsightNodeConnectionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTestGigaInsightNodeConnectionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &TestGigaInsightNodeConnectionAction{client: newMockClientReadErrorBody(t, 501)}
	m := TestGigaInsightNodeConnectionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
