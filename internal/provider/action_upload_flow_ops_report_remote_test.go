package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUploadFlowOpsReportAction_Invoke_Happy exercises UploadFlowOpsReportAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUploadFlowOpsReportAction_Invoke_Happy(t *testing.T) {
	r := &UploadFlowOpsReportAction{client: newMockClientStatus(t, 201, "{}")}
	m := UploadFlowOpsReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUploadFlowOpsReportAction_Invoke_NilClient exercises UploadFlowOpsReportAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUploadFlowOpsReportAction_Invoke_NilClient(t *testing.T) {
	r := &UploadFlowOpsReportAction{}
	m := UploadFlowOpsReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUploadFlowOpsReportAction_Invoke_BuildError exercises UploadFlowOpsReportAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUploadFlowOpsReportAction_Invoke_BuildError(t *testing.T) {
	r := &UploadFlowOpsReportAction{client: newMalformedBaseURLClient(t)}
	m := UploadFlowOpsReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUploadFlowOpsReportAction_Invoke_SendError exercises UploadFlowOpsReportAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUploadFlowOpsReportAction_Invoke_SendError(t *testing.T) {
	r := &UploadFlowOpsReportAction{client: newTransportErrorClient(t)}
	m := UploadFlowOpsReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUploadFlowOpsReportAction_Invoke_APIError exercises UploadFlowOpsReportAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUploadFlowOpsReportAction_Invoke_APIError(t *testing.T) {
	r := &UploadFlowOpsReportAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UploadFlowOpsReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_upload_flow_ops_report")
}

// TestUploadFlowOpsReportAction_Invoke_APIErrorReadBody exercises UploadFlowOpsReportAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUploadFlowOpsReportAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UploadFlowOpsReportAction{client: newMockClientReadErrorBody(t, 501)}
	m := UploadFlowOpsReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
