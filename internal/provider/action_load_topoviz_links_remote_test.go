package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestLoadTopovizLinksAction_Invoke_Happy exercises LoadTopovizLinksAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestLoadTopovizLinksAction_Invoke_Happy(t *testing.T) {
	r := &LoadTopovizLinksAction{client: newMockClientStatus(t, 200, "{}")}
	m := LoadTopovizLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadTopovizLinksAction_Invoke_NilClient exercises LoadTopovizLinksAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadTopovizLinksAction_Invoke_NilClient(t *testing.T) {
	r := &LoadTopovizLinksAction{}
	m := LoadTopovizLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadTopovizLinksAction_Invoke_BuildError exercises LoadTopovizLinksAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadTopovizLinksAction_Invoke_BuildError(t *testing.T) {
	r := &LoadTopovizLinksAction{client: newMalformedBaseURLClient(t)}
	m := LoadTopovizLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadTopovizLinksAction_Invoke_SendError exercises LoadTopovizLinksAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadTopovizLinksAction_Invoke_SendError(t *testing.T) {
	r := &LoadTopovizLinksAction{client: newTransportErrorClient(t)}
	m := LoadTopovizLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadTopovizLinksAction_Invoke_APIError exercises LoadTopovizLinksAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadTopovizLinksAction_Invoke_APIError(t *testing.T) {
	r := &LoadTopovizLinksAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadTopovizLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_load_topoviz_links")
}

// TestLoadTopovizLinksAction_Invoke_APIErrorReadBody exercises LoadTopovizLinksAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadTopovizLinksAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &LoadTopovizLinksAction{client: newMockClientReadErrorBody(t, 501)}
	m := LoadTopovizLinksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
