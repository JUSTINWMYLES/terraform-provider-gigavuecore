package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConvertToFabricMapAction_Invoke_Happy exercises ConvertToFabricMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConvertToFabricMapAction_Invoke_Happy(t *testing.T) {
	r := &ConvertToFabricMapAction{client: newMockClientStatus(t, 202, "{}")}
	m := ConvertToFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConvertToFabricMapAction_Invoke_NilClient exercises ConvertToFabricMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConvertToFabricMapAction_Invoke_NilClient(t *testing.T) {
	r := &ConvertToFabricMapAction{}
	m := ConvertToFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConvertToFabricMapAction_Invoke_BuildError exercises ConvertToFabricMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConvertToFabricMapAction_Invoke_BuildError(t *testing.T) {
	r := &ConvertToFabricMapAction{client: newMalformedBaseURLClient(t)}
	m := ConvertToFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConvertToFabricMapAction_Invoke_SendError exercises ConvertToFabricMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConvertToFabricMapAction_Invoke_SendError(t *testing.T) {
	r := &ConvertToFabricMapAction{client: newTransportErrorClient(t)}
	m := ConvertToFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConvertToFabricMapAction_Invoke_APIError exercises ConvertToFabricMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConvertToFabricMapAction_Invoke_APIError(t *testing.T) {
	r := &ConvertToFabricMapAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ConvertToFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_convert_to_fabric_map")
}

// TestConvertToFabricMapAction_Invoke_APIErrorReadBody exercises ConvertToFabricMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConvertToFabricMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConvertToFabricMapAction{client: newMockClientReadErrorBody(t, 500)}
	m := ConvertToFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
