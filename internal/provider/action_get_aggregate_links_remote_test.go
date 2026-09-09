package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestGetAggregateLinksAction_Invoke_Happy exercises GetAggregateLinksAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestGetAggregateLinksAction_Invoke_Happy(t *testing.T) {
	r := &GetAggregateLinksAction{client: newMockClientStatus(t, 200, "{}")}
	m := GetAggregateLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAggregateLinksAction_Invoke_NilClient exercises GetAggregateLinksAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAggregateLinksAction_Invoke_NilClient(t *testing.T) {
	r := &GetAggregateLinksAction{}
	m := GetAggregateLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAggregateLinksAction_Invoke_BuildError exercises GetAggregateLinksAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAggregateLinksAction_Invoke_BuildError(t *testing.T) {
	r := &GetAggregateLinksAction{client: newMalformedBaseURLClient(t)}
	m := GetAggregateLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAggregateLinksAction_Invoke_SendError exercises GetAggregateLinksAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAggregateLinksAction_Invoke_SendError(t *testing.T) {
	r := &GetAggregateLinksAction{client: newTransportErrorClient(t)}
	m := GetAggregateLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAggregateLinksAction_Invoke_APIError exercises GetAggregateLinksAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAggregateLinksAction_Invoke_APIError(t *testing.T) {
	r := &GetAggregateLinksAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAggregateLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_get_aggregate_links")
}

// TestGetAggregateLinksAction_Invoke_APIErrorReadBody exercises GetAggregateLinksAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAggregateLinksAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &GetAggregateLinksAction{client: newMockClientReadErrorBody(t, 501)}
	m := GetAggregateLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
