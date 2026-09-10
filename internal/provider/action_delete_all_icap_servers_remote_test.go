package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllIcapServersAction_Invoke_Happy exercises DeleteAllIcapServersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllIcapServersAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllIcapServersAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllIcapServersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllIcapServersAction_Invoke_NilClient exercises DeleteAllIcapServersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllIcapServersAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllIcapServersAction{}
	m := DeleteAllIcapServersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllIcapServersAction_Invoke_BuildError exercises DeleteAllIcapServersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllIcapServersAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllIcapServersAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllIcapServersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllIcapServersAction_Invoke_SendError exercises DeleteAllIcapServersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllIcapServersAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllIcapServersAction{client: newTransportErrorClient(t)}
	m := DeleteAllIcapServersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllIcapServersAction_Invoke_APIError exercises DeleteAllIcapServersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllIcapServersAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllIcapServersAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllIcapServersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_icap_servers")
}

// TestDeleteAllIcapServersAction_Invoke_APIErrorReadBody exercises DeleteAllIcapServersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllIcapServersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllIcapServersAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllIcapServersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
