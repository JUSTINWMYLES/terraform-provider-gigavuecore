package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddTrafficFlowsDraftAction_Invoke_Happy exercises AddTrafficFlowsDraftAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddTrafficFlowsDraftAction_Invoke_Happy(t *testing.T) {
	r := &AddTrafficFlowsDraftAction{client: newMockClientStatus(t, 200, "{}")}
	m := AddTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddTrafficFlowsDraftAction_Invoke_NilClient exercises AddTrafficFlowsDraftAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddTrafficFlowsDraftAction_Invoke_NilClient(t *testing.T) {
	r := &AddTrafficFlowsDraftAction{}
	m := AddTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddTrafficFlowsDraftAction_Invoke_BuildError exercises AddTrafficFlowsDraftAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddTrafficFlowsDraftAction_Invoke_BuildError(t *testing.T) {
	r := &AddTrafficFlowsDraftAction{client: newMalformedBaseURLClient(t)}
	m := AddTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddTrafficFlowsDraftAction_Invoke_SendError exercises AddTrafficFlowsDraftAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddTrafficFlowsDraftAction_Invoke_SendError(t *testing.T) {
	r := &AddTrafficFlowsDraftAction{client: newTransportErrorClient(t)}
	m := AddTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddTrafficFlowsDraftAction_Invoke_APIError exercises AddTrafficFlowsDraftAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddTrafficFlowsDraftAction_Invoke_APIError(t *testing.T) {
	r := &AddTrafficFlowsDraftAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_traffic_flows_draft")
}

// TestAddTrafficFlowsDraftAction_Invoke_APIErrorReadBody exercises AddTrafficFlowsDraftAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddTrafficFlowsDraftAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddTrafficFlowsDraftAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
