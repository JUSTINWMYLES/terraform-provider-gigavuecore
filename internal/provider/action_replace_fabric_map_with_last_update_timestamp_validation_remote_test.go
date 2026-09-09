package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_Happy exercises ReplaceFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_Happy(t *testing.T) {
	r := &ReplaceFabricMapWithLastUpdateTimestampValidationAction{client: newMockClientStatus(t, 200, "{}")}
	m := ReplaceFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_NilClient exercises ReplaceFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_NilClient(t *testing.T) {
	r := &ReplaceFabricMapWithLastUpdateTimestampValidationAction{}
	m := ReplaceFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_BuildError exercises ReplaceFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_BuildError(t *testing.T) {
	r := &ReplaceFabricMapWithLastUpdateTimestampValidationAction{client: newMalformedBaseURLClient(t)}
	m := ReplaceFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_SendError exercises ReplaceFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_SendError(t *testing.T) {
	r := &ReplaceFabricMapWithLastUpdateTimestampValidationAction{client: newTransportErrorClient(t)}
	m := ReplaceFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIError exercises ReplaceFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIError(t *testing.T) {
	r := &ReplaceFabricMapWithLastUpdateTimestampValidationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReplaceFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_replace_fabric_map_with_last_update_timestamp_validation")
}

// TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIErrorReadBody exercises ReplaceFabricMapWithLastUpdateTimestampValidationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReplaceFabricMapWithLastUpdateTimestampValidationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReplaceFabricMapWithLastUpdateTimestampValidationAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReplaceFabricMapWithLastUpdateTimestampValidationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
