package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteUpgradeJobsAction_Invoke_Happy exercises DeleteUpgradeJobsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteUpgradeJobsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteUpgradeJobsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteUpgradeJobsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteUpgradeJobsAction_Invoke_NilClient exercises DeleteUpgradeJobsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteUpgradeJobsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteUpgradeJobsAction{}
	m := DeleteUpgradeJobsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteUpgradeJobsAction_Invoke_BuildError exercises DeleteUpgradeJobsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteUpgradeJobsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteUpgradeJobsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteUpgradeJobsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteUpgradeJobsAction_Invoke_SendError exercises DeleteUpgradeJobsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteUpgradeJobsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteUpgradeJobsAction{client: newTransportErrorClient(t)}
	m := DeleteUpgradeJobsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteUpgradeJobsAction_Invoke_APIError exercises DeleteUpgradeJobsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteUpgradeJobsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteUpgradeJobsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteUpgradeJobsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_upgrade_jobs")
}

// TestDeleteUpgradeJobsAction_Invoke_APIErrorReadBody exercises DeleteUpgradeJobsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteUpgradeJobsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteUpgradeJobsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteUpgradeJobsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
