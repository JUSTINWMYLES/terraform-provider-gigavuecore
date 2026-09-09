package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_Happy exercises CollectHeartbeatfromGigaInsightNodeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_Happy(t *testing.T) {
	r := &CollectHeartbeatfromGigaInsightNodeAction{client: newMockClientStatus(t, 200, "{}")}
	m := CollectHeartbeatfromGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_NilClient exercises CollectHeartbeatfromGigaInsightNodeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_NilClient(t *testing.T) {
	r := &CollectHeartbeatfromGigaInsightNodeAction{}
	m := CollectHeartbeatfromGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_BuildError exercises CollectHeartbeatfromGigaInsightNodeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_BuildError(t *testing.T) {
	r := &CollectHeartbeatfromGigaInsightNodeAction{client: newMalformedBaseURLClient(t)}
	m := CollectHeartbeatfromGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_SendError exercises CollectHeartbeatfromGigaInsightNodeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_SendError(t *testing.T) {
	r := &CollectHeartbeatfromGigaInsightNodeAction{client: newTransportErrorClient(t)}
	m := CollectHeartbeatfromGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_APIError exercises CollectHeartbeatfromGigaInsightNodeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_APIError(t *testing.T) {
	r := &CollectHeartbeatfromGigaInsightNodeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CollectHeartbeatfromGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node")
}

// TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_APIErrorReadBody exercises CollectHeartbeatfromGigaInsightNodeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCollectHeartbeatfromGigaInsightNodeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CollectHeartbeatfromGigaInsightNodeAction{client: newMockClientReadErrorBody(t, 501)}
	m := CollectHeartbeatfromGigaInsightNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
