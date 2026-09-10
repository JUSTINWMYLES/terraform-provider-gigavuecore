package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateEmailNotifConfigSpecAction_Invoke_Happy exercises UpdateEmailNotifConfigSpecAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateEmailNotifConfigSpecAction_Invoke_Happy(t *testing.T) {
	r := &UpdateEmailNotifConfigSpecAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateEmailNotifConfigSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateEmailNotifConfigSpecAction_Invoke_NilClient exercises UpdateEmailNotifConfigSpecAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateEmailNotifConfigSpecAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateEmailNotifConfigSpecAction{}
	m := UpdateEmailNotifConfigSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateEmailNotifConfigSpecAction_Invoke_BuildError exercises UpdateEmailNotifConfigSpecAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateEmailNotifConfigSpecAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateEmailNotifConfigSpecAction{client: newMalformedBaseURLClient(t)}
	m := UpdateEmailNotifConfigSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateEmailNotifConfigSpecAction_Invoke_SendError exercises UpdateEmailNotifConfigSpecAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateEmailNotifConfigSpecAction_Invoke_SendError(t *testing.T) {
	r := &UpdateEmailNotifConfigSpecAction{client: newTransportErrorClient(t)}
	m := UpdateEmailNotifConfigSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateEmailNotifConfigSpecAction_Invoke_APIError exercises UpdateEmailNotifConfigSpecAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateEmailNotifConfigSpecAction_Invoke_APIError(t *testing.T) {
	r := &UpdateEmailNotifConfigSpecAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateEmailNotifConfigSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_email_notif_config_spec")
}

// TestUpdateEmailNotifConfigSpecAction_Invoke_APIErrorReadBody exercises UpdateEmailNotifConfigSpecAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateEmailNotifConfigSpecAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateEmailNotifConfigSpecAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateEmailNotifConfigSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
