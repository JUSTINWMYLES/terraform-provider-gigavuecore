package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteHsmFromHsmGroupAction_Invoke_Happy exercises DeleteHsmFromHsmGroupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteHsmFromHsmGroupAction_Invoke_Happy(t *testing.T) {
	r := &DeleteHsmFromHsmGroupAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteHsmFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteHsmFromHsmGroupAction_Invoke_NilClient exercises DeleteHsmFromHsmGroupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteHsmFromHsmGroupAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteHsmFromHsmGroupAction{}
	m := DeleteHsmFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteHsmFromHsmGroupAction_Invoke_BuildError exercises DeleteHsmFromHsmGroupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteHsmFromHsmGroupAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteHsmFromHsmGroupAction{client: newMalformedBaseURLClient(t)}
	m := DeleteHsmFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteHsmFromHsmGroupAction_Invoke_SendError exercises DeleteHsmFromHsmGroupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteHsmFromHsmGroupAction_Invoke_SendError(t *testing.T) {
	r := &DeleteHsmFromHsmGroupAction{client: newTransportErrorClient(t)}
	m := DeleteHsmFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteHsmFromHsmGroupAction_Invoke_APIError exercises DeleteHsmFromHsmGroupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteHsmFromHsmGroupAction_Invoke_APIError(t *testing.T) {
	r := &DeleteHsmFromHsmGroupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteHsmFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_hsm_from_hsm_group")
}

// TestDeleteHsmFromHsmGroupAction_Invoke_APIErrorReadBody exercises DeleteHsmFromHsmGroupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteHsmFromHsmGroupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteHsmFromHsmGroupAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteHsmFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
