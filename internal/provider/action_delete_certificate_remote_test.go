package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteCertificateAction_Invoke_Happy exercises DeleteCertificateAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteCertificateAction_Invoke_Happy(t *testing.T) {
	r := &DeleteCertificateAction{client: newMockClientStatus(t, 201, "{}")}
	m := DeleteCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteCertificateAction_Invoke_NilClient exercises DeleteCertificateAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteCertificateAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteCertificateAction{}
	m := DeleteCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteCertificateAction_Invoke_BuildError exercises DeleteCertificateAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteCertificateAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteCertificateAction{client: newMalformedBaseURLClient(t)}
	m := DeleteCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteCertificateAction_Invoke_SendError exercises DeleteCertificateAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteCertificateAction_Invoke_SendError(t *testing.T) {
	r := &DeleteCertificateAction{client: newTransportErrorClient(t)}
	m := DeleteCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteCertificateAction_Invoke_APIError exercises DeleteCertificateAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteCertificateAction_Invoke_APIError(t *testing.T) {
	r := &DeleteCertificateAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_certificate")
}

// TestDeleteCertificateAction_Invoke_APIErrorReadBody exercises DeleteCertificateAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteCertificateAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteCertificateAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
