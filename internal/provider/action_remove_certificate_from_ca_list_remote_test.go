package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveCertificateFromCaListAction_Invoke_Happy exercises RemoveCertificateFromCaListAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveCertificateFromCaListAction_Invoke_Happy(t *testing.T) {
	r := &RemoveCertificateFromCaListAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveCertificateFromCaListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveCertificateFromCaListAction_Invoke_NilClient exercises RemoveCertificateFromCaListAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveCertificateFromCaListAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveCertificateFromCaListAction{}
	m := RemoveCertificateFromCaListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveCertificateFromCaListAction_Invoke_BuildError exercises RemoveCertificateFromCaListAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveCertificateFromCaListAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveCertificateFromCaListAction{client: newMalformedBaseURLClient(t)}
	m := RemoveCertificateFromCaListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveCertificateFromCaListAction_Invoke_SendError exercises RemoveCertificateFromCaListAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveCertificateFromCaListAction_Invoke_SendError(t *testing.T) {
	r := &RemoveCertificateFromCaListAction{client: newTransportErrorClient(t)}
	m := RemoveCertificateFromCaListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveCertificateFromCaListAction_Invoke_APIError exercises RemoveCertificateFromCaListAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveCertificateFromCaListAction_Invoke_APIError(t *testing.T) {
	r := &RemoveCertificateFromCaListAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveCertificateFromCaListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_certificate_from_ca_list")
}

// TestRemoveCertificateFromCaListAction_Invoke_APIErrorReadBody exercises RemoveCertificateFromCaListAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveCertificateFromCaListAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveCertificateFromCaListAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveCertificateFromCaListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
