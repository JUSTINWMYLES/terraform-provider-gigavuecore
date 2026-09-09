package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestGetHierarchialDataAction_Invoke_Happy exercises GetHierarchialDataAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestGetHierarchialDataAction_Invoke_Happy(t *testing.T) {
	r := &GetHierarchialDataAction{client: newMockClientStatus(t, 200, "{}")}
	m := GetHierarchialDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetHierarchialDataAction_Invoke_NilClient exercises GetHierarchialDataAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetHierarchialDataAction_Invoke_NilClient(t *testing.T) {
	r := &GetHierarchialDataAction{}
	m := GetHierarchialDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetHierarchialDataAction_Invoke_BuildError exercises GetHierarchialDataAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetHierarchialDataAction_Invoke_BuildError(t *testing.T) {
	r := &GetHierarchialDataAction{client: newMalformedBaseURLClient(t)}
	m := GetHierarchialDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetHierarchialDataAction_Invoke_SendError exercises GetHierarchialDataAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetHierarchialDataAction_Invoke_SendError(t *testing.T) {
	r := &GetHierarchialDataAction{client: newTransportErrorClient(t)}
	m := GetHierarchialDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetHierarchialDataAction_Invoke_APIError exercises GetHierarchialDataAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetHierarchialDataAction_Invoke_APIError(t *testing.T) {
	r := &GetHierarchialDataAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetHierarchialDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_get_hierarchial_data")
}

// TestGetHierarchialDataAction_Invoke_APIErrorReadBody exercises GetHierarchialDataAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetHierarchialDataAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &GetHierarchialDataAction{client: newMockClientReadErrorBody(t, 501)}
	m := GetHierarchialDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
