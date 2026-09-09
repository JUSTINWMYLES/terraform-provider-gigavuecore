package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddKeyMapToHsmGroupAction_Invoke_Happy exercises AddKeyMapToHsmGroupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddKeyMapToHsmGroupAction_Invoke_Happy(t *testing.T) {
	r := &AddKeyMapToHsmGroupAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddKeyMapToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddKeyMapToHsmGroupAction_Invoke_NilClient exercises AddKeyMapToHsmGroupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddKeyMapToHsmGroupAction_Invoke_NilClient(t *testing.T) {
	r := &AddKeyMapToHsmGroupAction{}
	m := AddKeyMapToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddKeyMapToHsmGroupAction_Invoke_BuildError exercises AddKeyMapToHsmGroupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddKeyMapToHsmGroupAction_Invoke_BuildError(t *testing.T) {
	r := &AddKeyMapToHsmGroupAction{client: newMalformedBaseURLClient(t)}
	m := AddKeyMapToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddKeyMapToHsmGroupAction_Invoke_SendError exercises AddKeyMapToHsmGroupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddKeyMapToHsmGroupAction_Invoke_SendError(t *testing.T) {
	r := &AddKeyMapToHsmGroupAction{client: newTransportErrorClient(t)}
	m := AddKeyMapToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddKeyMapToHsmGroupAction_Invoke_APIError exercises AddKeyMapToHsmGroupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddKeyMapToHsmGroupAction_Invoke_APIError(t *testing.T) {
	r := &AddKeyMapToHsmGroupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddKeyMapToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_key_map_to_hsm_group")
}

// TestAddKeyMapToHsmGroupAction_Invoke_APIErrorReadBody exercises AddKeyMapToHsmGroupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddKeyMapToHsmGroupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddKeyMapToHsmGroupAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddKeyMapToHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
