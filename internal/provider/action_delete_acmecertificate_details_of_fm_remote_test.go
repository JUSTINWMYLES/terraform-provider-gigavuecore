package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAcmecertificateDetailsOfFmAction_Invoke_Happy exercises DeleteAcmecertificateDetailsOfFmAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAcmecertificateDetailsOfFmAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfFmAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAcmecertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAcmecertificateDetailsOfFmAction_Invoke_NilClient exercises DeleteAcmecertificateDetailsOfFmAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAcmecertificateDetailsOfFmAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfFmAction{}
	m := DeleteAcmecertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAcmecertificateDetailsOfFmAction_Invoke_BuildError exercises DeleteAcmecertificateDetailsOfFmAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAcmecertificateDetailsOfFmAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfFmAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAcmecertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAcmecertificateDetailsOfFmAction_Invoke_SendError exercises DeleteAcmecertificateDetailsOfFmAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAcmecertificateDetailsOfFmAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfFmAction{client: newTransportErrorClient(t)}
	m := DeleteAcmecertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAcmecertificateDetailsOfFmAction_Invoke_APIError exercises DeleteAcmecertificateDetailsOfFmAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAcmecertificateDetailsOfFmAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfFmAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAcmecertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_acmecertificate_details_of_fm")
}

// TestDeleteAcmecertificateDetailsOfFmAction_Invoke_APIErrorReadBody exercises DeleteAcmecertificateDetailsOfFmAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAcmecertificateDetailsOfFmAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfFmAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAcmecertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
