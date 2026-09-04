package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteReportAction_Invoke_Happy exercises DeleteReportAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteReportAction_Invoke_Happy(t *testing.T) {
	r := &DeleteReportAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteReportAction_Invoke_NilClient exercises DeleteReportAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteReportAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteReportAction{}
	m := DeleteReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteReportAction_Invoke_BuildError exercises DeleteReportAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteReportAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteReportAction{client: newMalformedBaseURLClient(t)}
	m := DeleteReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteReportAction_Invoke_SendError exercises DeleteReportAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteReportAction_Invoke_SendError(t *testing.T) {
	r := &DeleteReportAction{client: newTransportErrorClient(t)}
	m := DeleteReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteReportAction_Invoke_APIError exercises DeleteReportAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteReportAction_Invoke_APIError(t *testing.T) {
	r := &DeleteReportAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_report")
}

// TestDeleteReportAction_Invoke_APIErrorReadBody exercises DeleteReportAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteReportAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteReportAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteReportActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
