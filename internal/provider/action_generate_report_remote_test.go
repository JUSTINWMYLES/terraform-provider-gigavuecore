package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestGenerateReportAction_Invoke_Happy exercises GenerateReportAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestGenerateReportAction_Invoke_Happy(t *testing.T) {
	r := &GenerateReportAction{client: newMockClientStatus(t, 201, "{}")}
	m := GenerateReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGenerateReportAction_Invoke_NilClient exercises GenerateReportAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGenerateReportAction_Invoke_NilClient(t *testing.T) {
	r := &GenerateReportAction{}
	m := GenerateReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGenerateReportAction_Invoke_BuildError exercises GenerateReportAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGenerateReportAction_Invoke_BuildError(t *testing.T) {
	r := &GenerateReportAction{client: newMalformedBaseURLClient(t)}
	m := GenerateReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGenerateReportAction_Invoke_SendError exercises GenerateReportAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestGenerateReportAction_Invoke_SendError(t *testing.T) {
	r := &GenerateReportAction{client: newTransportErrorClient(t)}
	m := GenerateReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGenerateReportAction_Invoke_APIError exercises GenerateReportAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGenerateReportAction_Invoke_APIError(t *testing.T) {
	r := &GenerateReportAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GenerateReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_generate_report")
}

// TestGenerateReportAction_Invoke_APIErrorReadBody exercises GenerateReportAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGenerateReportAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &GenerateReportAction{client: newMockClientReadErrorBody(t, 501)}
	m := GenerateReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
