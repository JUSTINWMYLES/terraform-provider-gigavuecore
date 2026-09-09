package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestGetSankeyDataAction_Invoke_Happy exercises GetSankeyDataAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestGetSankeyDataAction_Invoke_Happy(t *testing.T) {
	r := &GetSankeyDataAction{client: newMockClientStatus(t, 200, "{}")}
	m := GetSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSankeyDataAction_Invoke_NilClient exercises GetSankeyDataAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSankeyDataAction_Invoke_NilClient(t *testing.T) {
	r := &GetSankeyDataAction{}
	m := GetSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSankeyDataAction_Invoke_BuildError exercises GetSankeyDataAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSankeyDataAction_Invoke_BuildError(t *testing.T) {
	r := &GetSankeyDataAction{client: newMalformedBaseURLClient(t)}
	m := GetSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSankeyDataAction_Invoke_SendError exercises GetSankeyDataAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSankeyDataAction_Invoke_SendError(t *testing.T) {
	r := &GetSankeyDataAction{client: newTransportErrorClient(t)}
	m := GetSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSankeyDataAction_Invoke_APIError exercises GetSankeyDataAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSankeyDataAction_Invoke_APIError(t *testing.T) {
	r := &GetSankeyDataAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_get_sankey_data")
}

// TestGetSankeyDataAction_Invoke_APIErrorReadBody exercises GetSankeyDataAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSankeyDataAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &GetSankeyDataAction{client: newMockClientReadErrorBody(t, 501)}
	m := GetSankeyDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
