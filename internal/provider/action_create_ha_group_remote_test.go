package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateHaGroupAction_Invoke_Happy exercises CreateHaGroupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateHaGroupAction_Invoke_Happy(t *testing.T) {
	r := &CreateHaGroupAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateHaGroupAction_Invoke_NilClient exercises CreateHaGroupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateHaGroupAction_Invoke_NilClient(t *testing.T) {
	r := &CreateHaGroupAction{}
	m := CreateHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateHaGroupAction_Invoke_BuildError exercises CreateHaGroupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateHaGroupAction_Invoke_BuildError(t *testing.T) {
	r := &CreateHaGroupAction{client: newMalformedBaseURLClient(t)}
	m := CreateHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateHaGroupAction_Invoke_SendError exercises CreateHaGroupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateHaGroupAction_Invoke_SendError(t *testing.T) {
	r := &CreateHaGroupAction{client: newTransportErrorClient(t)}
	m := CreateHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateHaGroupAction_Invoke_APIError exercises CreateHaGroupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateHaGroupAction_Invoke_APIError(t *testing.T) {
	r := &CreateHaGroupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_ha_group")
}

// TestCreateHaGroupAction_Invoke_APIErrorReadBody exercises CreateHaGroupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateHaGroupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateHaGroupAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateHaGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
