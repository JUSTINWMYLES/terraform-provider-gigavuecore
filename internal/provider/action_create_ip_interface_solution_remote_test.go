package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateIpInterfaceSolutionAction_Invoke_Happy exercises CreateIpInterfaceSolutionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateIpInterfaceSolutionAction_Invoke_Happy(t *testing.T) {
	r := &CreateIpInterfaceSolutionAction{client: newMockClientStatus(t, 207, "{}")}
	m := CreateIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateIpInterfaceSolutionAction_Invoke_NilClient exercises CreateIpInterfaceSolutionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateIpInterfaceSolutionAction_Invoke_NilClient(t *testing.T) {
	r := &CreateIpInterfaceSolutionAction{}
	m := CreateIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateIpInterfaceSolutionAction_Invoke_BuildError exercises CreateIpInterfaceSolutionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateIpInterfaceSolutionAction_Invoke_BuildError(t *testing.T) {
	r := &CreateIpInterfaceSolutionAction{client: newMalformedBaseURLClient(t)}
	m := CreateIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateIpInterfaceSolutionAction_Invoke_SendError exercises CreateIpInterfaceSolutionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateIpInterfaceSolutionAction_Invoke_SendError(t *testing.T) {
	r := &CreateIpInterfaceSolutionAction{client: newTransportErrorClient(t)}
	m := CreateIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateIpInterfaceSolutionAction_Invoke_APIError exercises CreateIpInterfaceSolutionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateIpInterfaceSolutionAction_Invoke_APIError(t *testing.T) {
	r := &CreateIpInterfaceSolutionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_ip_interface_solution")
}

// TestCreateIpInterfaceSolutionAction_Invoke_APIErrorReadBody exercises CreateIpInterfaceSolutionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateIpInterfaceSolutionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateIpInterfaceSolutionAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
