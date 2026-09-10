package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteTrafficConfigAction_Invoke_Happy exercises DeleteTrafficConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteTrafficConfigAction_Invoke_Happy(t *testing.T) {
	r := &DeleteTrafficConfigAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteTrafficConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteTrafficConfigAction_Invoke_NilClient exercises DeleteTrafficConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteTrafficConfigAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteTrafficConfigAction{}
	m := DeleteTrafficConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteTrafficConfigAction_Invoke_BuildError exercises DeleteTrafficConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteTrafficConfigAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteTrafficConfigAction{client: newMalformedBaseURLClient(t)}
	m := DeleteTrafficConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteTrafficConfigAction_Invoke_SendError exercises DeleteTrafficConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteTrafficConfigAction_Invoke_SendError(t *testing.T) {
	r := &DeleteTrafficConfigAction{client: newTransportErrorClient(t)}
	m := DeleteTrafficConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteTrafficConfigAction_Invoke_APIError exercises DeleteTrafficConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteTrafficConfigAction_Invoke_APIError(t *testing.T) {
	r := &DeleteTrafficConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteTrafficConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_traffic_config")
}

// TestDeleteTrafficConfigAction_Invoke_APIErrorReadBody exercises DeleteTrafficConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteTrafficConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteTrafficConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteTrafficConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
