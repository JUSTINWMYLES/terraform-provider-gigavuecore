package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestGetSmartSankeyDataAction_Invoke_Happy exercises GetSmartSankeyDataAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestGetSmartSankeyDataAction_Invoke_Happy(t *testing.T) {
	r := &GetSmartSankeyDataAction{client: newMockClientStatus(t, 200, "{}")}
	m := GetSmartSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSmartSankeyDataAction_Invoke_NilClient exercises GetSmartSankeyDataAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSmartSankeyDataAction_Invoke_NilClient(t *testing.T) {
	r := &GetSmartSankeyDataAction{}
	m := GetSmartSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSmartSankeyDataAction_Invoke_BuildError exercises GetSmartSankeyDataAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSmartSankeyDataAction_Invoke_BuildError(t *testing.T) {
	r := &GetSmartSankeyDataAction{client: newMalformedBaseURLClient(t)}
	m := GetSmartSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSmartSankeyDataAction_Invoke_SendError exercises GetSmartSankeyDataAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSmartSankeyDataAction_Invoke_SendError(t *testing.T) {
	r := &GetSmartSankeyDataAction{client: newTransportErrorClient(t)}
	m := GetSmartSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSmartSankeyDataAction_Invoke_APIError exercises GetSmartSankeyDataAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSmartSankeyDataAction_Invoke_APIError(t *testing.T) {
	r := &GetSmartSankeyDataAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSmartSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_get_smart_sankey_data")
}

// TestGetSmartSankeyDataAction_Invoke_APIErrorReadBody exercises GetSmartSankeyDataAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSmartSankeyDataAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &GetSmartSankeyDataAction{client: newMockClientReadErrorBody(t, 501)}
	m := GetSmartSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
