package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAcmeServerDetailsAction_Invoke_Happy exercises DeleteAcmeServerDetailsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAcmeServerDetailsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAcmeServerDetailsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAcmeServerDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAcmeServerDetailsAction_Invoke_NilClient exercises DeleteAcmeServerDetailsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAcmeServerDetailsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAcmeServerDetailsAction{}
	m := DeleteAcmeServerDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAcmeServerDetailsAction_Invoke_BuildError exercises DeleteAcmeServerDetailsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAcmeServerDetailsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAcmeServerDetailsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAcmeServerDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAcmeServerDetailsAction_Invoke_SendError exercises DeleteAcmeServerDetailsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAcmeServerDetailsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAcmeServerDetailsAction{client: newTransportErrorClient(t)}
	m := DeleteAcmeServerDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAcmeServerDetailsAction_Invoke_APIError exercises DeleteAcmeServerDetailsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAcmeServerDetailsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAcmeServerDetailsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAcmeServerDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_acme_server_details")
}

// TestDeleteAcmeServerDetailsAction_Invoke_APIErrorReadBody exercises DeleteAcmeServerDetailsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAcmeServerDetailsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAcmeServerDetailsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAcmeServerDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
