package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddHsmToHsmGroupAction_Invoke_Happy exercises AddHsmToHsmGroupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddHsmToHsmGroupAction_Invoke_Happy(t *testing.T) {
	r := &AddHsmToHsmGroupAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddHsmToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddHsmToHsmGroupAction_Invoke_NilClient exercises AddHsmToHsmGroupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddHsmToHsmGroupAction_Invoke_NilClient(t *testing.T) {
	r := &AddHsmToHsmGroupAction{}
	m := AddHsmToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddHsmToHsmGroupAction_Invoke_BuildError exercises AddHsmToHsmGroupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddHsmToHsmGroupAction_Invoke_BuildError(t *testing.T) {
	r := &AddHsmToHsmGroupAction{client: newMalformedBaseURLClient(t)}
	m := AddHsmToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddHsmToHsmGroupAction_Invoke_SendError exercises AddHsmToHsmGroupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddHsmToHsmGroupAction_Invoke_SendError(t *testing.T) {
	r := &AddHsmToHsmGroupAction{client: newTransportErrorClient(t)}
	m := AddHsmToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddHsmToHsmGroupAction_Invoke_APIError exercises AddHsmToHsmGroupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddHsmToHsmGroupAction_Invoke_APIError(t *testing.T) {
	r := &AddHsmToHsmGroupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddHsmToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_hsm_to_hsm_group")
}

// TestAddHsmToHsmGroupAction_Invoke_APIErrorReadBody exercises AddHsmToHsmGroupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddHsmToHsmGroupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddHsmToHsmGroupAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddHsmToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
