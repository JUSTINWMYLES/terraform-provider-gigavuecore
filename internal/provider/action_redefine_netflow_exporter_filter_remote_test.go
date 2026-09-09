package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineNetflowExporterFilterAction_Invoke_Happy exercises RedefineNetflowExporterFilterAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineNetflowExporterFilterAction_Invoke_Happy(t *testing.T) {
	r := &RedefineNetflowExporterFilterAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineNetflowExporterFilterAction_Invoke_NilClient exercises RedefineNetflowExporterFilterAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineNetflowExporterFilterAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineNetflowExporterFilterAction{}
	m := RedefineNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineNetflowExporterFilterAction_Invoke_BuildError exercises RedefineNetflowExporterFilterAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineNetflowExporterFilterAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineNetflowExporterFilterAction{client: newMalformedBaseURLClient(t)}
	m := RedefineNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineNetflowExporterFilterAction_Invoke_SendError exercises RedefineNetflowExporterFilterAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineNetflowExporterFilterAction_Invoke_SendError(t *testing.T) {
	r := &RedefineNetflowExporterFilterAction{client: newTransportErrorClient(t)}
	m := RedefineNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineNetflowExporterFilterAction_Invoke_APIError exercises RedefineNetflowExporterFilterAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineNetflowExporterFilterAction_Invoke_APIError(t *testing.T) {
	r := &RedefineNetflowExporterFilterAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_netflow_exporter_filter")
}

// TestRedefineNetflowExporterFilterAction_Invoke_APIErrorReadBody exercises RedefineNetflowExporterFilterAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineNetflowExporterFilterAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineNetflowExporterFilterAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
