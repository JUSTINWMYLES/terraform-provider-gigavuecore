package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteNetflowExporterFilterAction_Invoke_Happy exercises DeleteNetflowExporterFilterAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteNetflowExporterFilterAction_Invoke_Happy(t *testing.T) {
	r := &DeleteNetflowExporterFilterAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteNetflowExporterFilterAction_Invoke_NilClient exercises DeleteNetflowExporterFilterAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteNetflowExporterFilterAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteNetflowExporterFilterAction{}
	m := DeleteNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteNetflowExporterFilterAction_Invoke_BuildError exercises DeleteNetflowExporterFilterAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteNetflowExporterFilterAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteNetflowExporterFilterAction{client: newMalformedBaseURLClient(t)}
	m := DeleteNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteNetflowExporterFilterAction_Invoke_SendError exercises DeleteNetflowExporterFilterAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteNetflowExporterFilterAction_Invoke_SendError(t *testing.T) {
	r := &DeleteNetflowExporterFilterAction{client: newTransportErrorClient(t)}
	m := DeleteNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteNetflowExporterFilterAction_Invoke_APIError exercises DeleteNetflowExporterFilterAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteNetflowExporterFilterAction_Invoke_APIError(t *testing.T) {
	r := &DeleteNetflowExporterFilterAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_netflow_exporter_filter")
}

// TestDeleteNetflowExporterFilterAction_Invoke_APIErrorReadBody exercises DeleteNetflowExporterFilterAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteNetflowExporterFilterAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteNetflowExporterFilterAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
