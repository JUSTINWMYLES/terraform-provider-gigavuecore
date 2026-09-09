package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestPatchIpInterfaceSolutionAction_Invoke_Happy exercises PatchIpInterfaceSolutionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestPatchIpInterfaceSolutionAction_Invoke_Happy(t *testing.T) {
	r := &PatchIpInterfaceSolutionAction{client: newMockClientStatus(t, 207, "{}")}
	m := PatchIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPatchIpInterfaceSolutionAction_Invoke_NilClient exercises PatchIpInterfaceSolutionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPatchIpInterfaceSolutionAction_Invoke_NilClient(t *testing.T) {
	r := &PatchIpInterfaceSolutionAction{}
	m := PatchIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPatchIpInterfaceSolutionAction_Invoke_BuildError exercises PatchIpInterfaceSolutionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPatchIpInterfaceSolutionAction_Invoke_BuildError(t *testing.T) {
	r := &PatchIpInterfaceSolutionAction{client: newMalformedBaseURLClient(t)}
	m := PatchIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPatchIpInterfaceSolutionAction_Invoke_SendError exercises PatchIpInterfaceSolutionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestPatchIpInterfaceSolutionAction_Invoke_SendError(t *testing.T) {
	r := &PatchIpInterfaceSolutionAction{client: newTransportErrorClient(t)}
	m := PatchIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPatchIpInterfaceSolutionAction_Invoke_APIError exercises PatchIpInterfaceSolutionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPatchIpInterfaceSolutionAction_Invoke_APIError(t *testing.T) {
	r := &PatchIpInterfaceSolutionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PatchIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_patch_ip_interface_solution")
}

// TestPatchIpInterfaceSolutionAction_Invoke_APIErrorReadBody exercises PatchIpInterfaceSolutionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPatchIpInterfaceSolutionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &PatchIpInterfaceSolutionAction{client: newMockClientReadErrorBody(t, 501)}
	m := PatchIpInterfaceSolutionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
