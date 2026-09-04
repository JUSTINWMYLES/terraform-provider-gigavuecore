package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestIssueAcmeCertificateAction_Invoke_Happy exercises IssueAcmeCertificateAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestIssueAcmeCertificateAction_Invoke_Happy(t *testing.T) {
	r := &IssueAcmeCertificateAction{client: newMockClientStatus(t, 202, "{}")}
	m := IssueAcmeCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIssueAcmeCertificateAction_Invoke_NilClient exercises IssueAcmeCertificateAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIssueAcmeCertificateAction_Invoke_NilClient(t *testing.T) {
	r := &IssueAcmeCertificateAction{}
	m := IssueAcmeCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIssueAcmeCertificateAction_Invoke_BuildError exercises IssueAcmeCertificateAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIssueAcmeCertificateAction_Invoke_BuildError(t *testing.T) {
	r := &IssueAcmeCertificateAction{client: newMalformedBaseURLClient(t)}
	m := IssueAcmeCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIssueAcmeCertificateAction_Invoke_SendError exercises IssueAcmeCertificateAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestIssueAcmeCertificateAction_Invoke_SendError(t *testing.T) {
	r := &IssueAcmeCertificateAction{client: newTransportErrorClient(t)}
	m := IssueAcmeCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIssueAcmeCertificateAction_Invoke_APIError exercises IssueAcmeCertificateAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIssueAcmeCertificateAction_Invoke_APIError(t *testing.T) {
	r := &IssueAcmeCertificateAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := IssueAcmeCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_issue_acme_certificate")
}

// TestIssueAcmeCertificateAction_Invoke_APIErrorReadBody exercises IssueAcmeCertificateAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIssueAcmeCertificateAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &IssueAcmeCertificateAction{client: newMockClientReadErrorBody(t, 501)}
	m := IssueAcmeCertificateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
