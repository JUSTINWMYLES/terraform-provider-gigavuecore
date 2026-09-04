package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDelKeyMapFromHsmGroupAction_Invoke_Happy exercises DelKeyMapFromHsmGroupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDelKeyMapFromHsmGroupAction_Invoke_Happy(t *testing.T) {
	r := &DelKeyMapFromHsmGroupAction{client: newMockClientStatus(t, 201, "{}")}
	m := DelKeyMapFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDelKeyMapFromHsmGroupAction_Invoke_NilClient exercises DelKeyMapFromHsmGroupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDelKeyMapFromHsmGroupAction_Invoke_NilClient(t *testing.T) {
	r := &DelKeyMapFromHsmGroupAction{}
	m := DelKeyMapFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDelKeyMapFromHsmGroupAction_Invoke_BuildError exercises DelKeyMapFromHsmGroupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDelKeyMapFromHsmGroupAction_Invoke_BuildError(t *testing.T) {
	r := &DelKeyMapFromHsmGroupAction{client: newMalformedBaseURLClient(t)}
	m := DelKeyMapFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDelKeyMapFromHsmGroupAction_Invoke_SendError exercises DelKeyMapFromHsmGroupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDelKeyMapFromHsmGroupAction_Invoke_SendError(t *testing.T) {
	r := &DelKeyMapFromHsmGroupAction{client: newTransportErrorClient(t)}
	m := DelKeyMapFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDelKeyMapFromHsmGroupAction_Invoke_APIError exercises DelKeyMapFromHsmGroupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDelKeyMapFromHsmGroupAction_Invoke_APIError(t *testing.T) {
	r := &DelKeyMapFromHsmGroupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DelKeyMapFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_del_key_map_from_hsm_group")
}

// TestDelKeyMapFromHsmGroupAction_Invoke_APIErrorReadBody exercises DelKeyMapFromHsmGroupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDelKeyMapFromHsmGroupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DelKeyMapFromHsmGroupAction{client: newMockClientReadErrorBody(t, 501)}
	m := DelKeyMapFromHsmGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
