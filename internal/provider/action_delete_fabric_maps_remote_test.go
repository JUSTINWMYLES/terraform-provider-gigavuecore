package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteFabricMapsAction_Invoke_Happy exercises DeleteFabricMapsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteFabricMapsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteFabricMapsAction{client: newMockClientStatus(t, 202, "{}")}
	m := DeleteFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteFabricMapsAction_Invoke_NilClient exercises DeleteFabricMapsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteFabricMapsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteFabricMapsAction{}
	m := DeleteFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteFabricMapsAction_Invoke_BuildError exercises DeleteFabricMapsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteFabricMapsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteFabricMapsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteFabricMapsAction_Invoke_SendError exercises DeleteFabricMapsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteFabricMapsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteFabricMapsAction{client: newTransportErrorClient(t)}
	m := DeleteFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteFabricMapsAction_Invoke_APIError exercises DeleteFabricMapsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteFabricMapsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteFabricMapsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_fabric_maps")
}

// TestDeleteFabricMapsAction_Invoke_APIErrorReadBody exercises DeleteFabricMapsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteFabricMapsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteFabricMapsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
