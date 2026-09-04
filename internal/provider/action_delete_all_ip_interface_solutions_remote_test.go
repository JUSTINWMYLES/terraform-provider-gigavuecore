package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllIpInterfaceSolutionsAction_Invoke_Happy exercises DeleteAllIpInterfaceSolutionsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllIpInterfaceSolutionsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllIpInterfaceSolutionsAction{client: newMockClientStatus(t, 207, "{}")}
	m := DeleteAllIpInterfaceSolutionsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllIpInterfaceSolutionsAction_Invoke_NilClient exercises DeleteAllIpInterfaceSolutionsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllIpInterfaceSolutionsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllIpInterfaceSolutionsAction{}
	m := DeleteAllIpInterfaceSolutionsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllIpInterfaceSolutionsAction_Invoke_BuildError exercises DeleteAllIpInterfaceSolutionsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllIpInterfaceSolutionsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllIpInterfaceSolutionsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllIpInterfaceSolutionsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllIpInterfaceSolutionsAction_Invoke_SendError exercises DeleteAllIpInterfaceSolutionsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllIpInterfaceSolutionsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllIpInterfaceSolutionsAction{client: newTransportErrorClient(t)}
	m := DeleteAllIpInterfaceSolutionsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllIpInterfaceSolutionsAction_Invoke_APIError exercises DeleteAllIpInterfaceSolutionsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllIpInterfaceSolutionsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllIpInterfaceSolutionsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllIpInterfaceSolutionsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_ip_interface_solutions")
}

// TestDeleteAllIpInterfaceSolutionsAction_Invoke_APIErrorReadBody exercises DeleteAllIpInterfaceSolutionsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllIpInterfaceSolutionsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllIpInterfaceSolutionsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllIpInterfaceSolutionsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
