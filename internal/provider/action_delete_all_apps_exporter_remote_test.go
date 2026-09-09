package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllAppsExporterAction_Invoke_Happy exercises DeleteAllAppsExporterAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllAppsExporterAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllAppsExporterAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllAppsExporterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllAppsExporterAction_Invoke_NilClient exercises DeleteAllAppsExporterAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllAppsExporterAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllAppsExporterAction{}
	m := DeleteAllAppsExporterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllAppsExporterAction_Invoke_BuildError exercises DeleteAllAppsExporterAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllAppsExporterAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllAppsExporterAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllAppsExporterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllAppsExporterAction_Invoke_SendError exercises DeleteAllAppsExporterAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllAppsExporterAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllAppsExporterAction{client: newTransportErrorClient(t)}
	m := DeleteAllAppsExporterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllAppsExporterAction_Invoke_APIError exercises DeleteAllAppsExporterAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllAppsExporterAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllAppsExporterAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllAppsExporterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_apps_exporter")
}

// TestDeleteAllAppsExporterAction_Invoke_APIErrorReadBody exercises DeleteAllAppsExporterAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllAppsExporterAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllAppsExporterAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllAppsExporterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
