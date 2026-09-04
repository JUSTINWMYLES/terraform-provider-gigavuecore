package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateManagementInterfaceAction_Invoke_Happy exercises UpdateManagementInterfaceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateManagementInterfaceAction_Invoke_Happy(t *testing.T) {
	r := &UpdateManagementInterfaceAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateManagementInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateManagementInterfaceAction_Invoke_NilClient exercises UpdateManagementInterfaceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateManagementInterfaceAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateManagementInterfaceAction{}
	m := UpdateManagementInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateManagementInterfaceAction_Invoke_BuildError exercises UpdateManagementInterfaceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateManagementInterfaceAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateManagementInterfaceAction{client: newMalformedBaseURLClient(t)}
	m := UpdateManagementInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateManagementInterfaceAction_Invoke_SendError exercises UpdateManagementInterfaceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateManagementInterfaceAction_Invoke_SendError(t *testing.T) {
	r := &UpdateManagementInterfaceAction{client: newTransportErrorClient(t)}
	m := UpdateManagementInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateManagementInterfaceAction_Invoke_APIError exercises UpdateManagementInterfaceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateManagementInterfaceAction_Invoke_APIError(t *testing.T) {
	r := &UpdateManagementInterfaceAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateManagementInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_management_interface")
}

// TestUpdateManagementInterfaceAction_Invoke_APIErrorReadBody exercises UpdateManagementInterfaceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateManagementInterfaceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateManagementInterfaceAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateManagementInterfaceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
