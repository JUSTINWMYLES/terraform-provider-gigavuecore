package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_Happy exercises UpdateFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_Happy(t *testing.T) {
	r := &UpdateFabricMapWithLastUpdateTimestampValidationAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_NilClient exercises UpdateFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateFabricMapWithLastUpdateTimestampValidationAction{}
	m := UpdateFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_BuildError exercises UpdateFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateFabricMapWithLastUpdateTimestampValidationAction{client: newMalformedBaseURLClient(t)}
	m := UpdateFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_SendError exercises UpdateFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_SendError(t *testing.T) {
	r := &UpdateFabricMapWithLastUpdateTimestampValidationAction{client: newTransportErrorClient(t)}
	m := UpdateFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIError exercises UpdateFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIError(t *testing.T) {
	r := &UpdateFabricMapWithLastUpdateTimestampValidationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation")
}

// TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIErrorReadBody exercises UpdateFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateFabricMapWithLastUpdateTimestampValidationAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
