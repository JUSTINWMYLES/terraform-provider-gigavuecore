package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeletePcapFileAction_Invoke_Happy exercises DeletePcapFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeletePcapFileAction_Invoke_Happy(t *testing.T) {
	r := &DeletePcapFileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeletePcapFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeletePcapFileAction_Invoke_NilClient exercises DeletePcapFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeletePcapFileAction_Invoke_NilClient(t *testing.T) {
	r := &DeletePcapFileAction{}
	m := DeletePcapFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeletePcapFileAction_Invoke_BuildError exercises DeletePcapFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeletePcapFileAction_Invoke_BuildError(t *testing.T) {
	r := &DeletePcapFileAction{client: newMalformedBaseURLClient(t)}
	m := DeletePcapFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeletePcapFileAction_Invoke_SendError exercises DeletePcapFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeletePcapFileAction_Invoke_SendError(t *testing.T) {
	r := &DeletePcapFileAction{client: newTransportErrorClient(t)}
	m := DeletePcapFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeletePcapFileAction_Invoke_APIError exercises DeletePcapFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeletePcapFileAction_Invoke_APIError(t *testing.T) {
	r := &DeletePcapFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeletePcapFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_pcap_file")
}

// TestDeletePcapFileAction_Invoke_APIErrorReadBody exercises DeletePcapFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeletePcapFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeletePcapFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeletePcapFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
