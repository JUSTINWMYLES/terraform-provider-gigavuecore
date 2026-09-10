package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateNetflowExporterFilterAction_Invoke_Happy exercises CreateNetflowExporterFilterAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateNetflowExporterFilterAction_Invoke_Happy(t *testing.T) {
	r := &CreateNetflowExporterFilterAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateNetflowExporterFilterAction_Invoke_NilClient exercises CreateNetflowExporterFilterAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateNetflowExporterFilterAction_Invoke_NilClient(t *testing.T) {
	r := &CreateNetflowExporterFilterAction{}
	m := CreateNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateNetflowExporterFilterAction_Invoke_BuildError exercises CreateNetflowExporterFilterAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateNetflowExporterFilterAction_Invoke_BuildError(t *testing.T) {
	r := &CreateNetflowExporterFilterAction{client: newMalformedBaseURLClient(t)}
	m := CreateNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateNetflowExporterFilterAction_Invoke_SendError exercises CreateNetflowExporterFilterAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateNetflowExporterFilterAction_Invoke_SendError(t *testing.T) {
	r := &CreateNetflowExporterFilterAction{client: newTransportErrorClient(t)}
	m := CreateNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateNetflowExporterFilterAction_Invoke_APIError exercises CreateNetflowExporterFilterAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateNetflowExporterFilterAction_Invoke_APIError(t *testing.T) {
	r := &CreateNetflowExporterFilterAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_netflow_exporter_filter")
}

// TestCreateNetflowExporterFilterAction_Invoke_APIErrorReadBody exercises CreateNetflowExporterFilterAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateNetflowExporterFilterAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateNetflowExporterFilterAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateNetflowExporterFilterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
