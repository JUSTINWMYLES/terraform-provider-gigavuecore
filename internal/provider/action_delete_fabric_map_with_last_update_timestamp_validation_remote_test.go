package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_Happy exercises DeleteFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_Happy(t *testing.T) {
	r := &DeleteFabricMapWithLastUpdateTimestampValidationAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_NilClient exercises DeleteFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteFabricMapWithLastUpdateTimestampValidationAction{}
	m := DeleteFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_BuildError exercises DeleteFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteFabricMapWithLastUpdateTimestampValidationAction{client: newMalformedBaseURLClient(t)}
	m := DeleteFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_SendError exercises DeleteFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_SendError(t *testing.T) {
	r := &DeleteFabricMapWithLastUpdateTimestampValidationAction{client: newTransportErrorClient(t)}
	m := DeleteFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIError exercises DeleteFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIError(t *testing.T) {
	r := &DeleteFabricMapWithLastUpdateTimestampValidationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_fabric_map_with_last_update_timestamp_validation")
}

// TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIErrorReadBody exercises DeleteFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteFabricMapWithLastUpdateTimestampValidationAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
