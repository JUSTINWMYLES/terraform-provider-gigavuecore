package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllAppsExporterGroupAction_Invoke_Happy exercises DeleteAllAppsExporterGroupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllAppsExporterGroupAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllAppsExporterGroupAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllAppsExporterGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllAppsExporterGroupAction_Invoke_NilClient exercises DeleteAllAppsExporterGroupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllAppsExporterGroupAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllAppsExporterGroupAction{}
	m := DeleteAllAppsExporterGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllAppsExporterGroupAction_Invoke_BuildError exercises DeleteAllAppsExporterGroupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllAppsExporterGroupAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllAppsExporterGroupAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllAppsExporterGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllAppsExporterGroupAction_Invoke_SendError exercises DeleteAllAppsExporterGroupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllAppsExporterGroupAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllAppsExporterGroupAction{client: newTransportErrorClient(t)}
	m := DeleteAllAppsExporterGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllAppsExporterGroupAction_Invoke_APIError exercises DeleteAllAppsExporterGroupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllAppsExporterGroupAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllAppsExporterGroupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllAppsExporterGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_apps_exporter_group")
}

// TestDeleteAllAppsExporterGroupAction_Invoke_APIErrorReadBody exercises DeleteAllAppsExporterGroupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllAppsExporterGroupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllAppsExporterGroupAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllAppsExporterGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
