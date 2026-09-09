package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestFetchInlineSslProfileListAction_Invoke_Happy exercises FetchInlineSslProfileListAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestFetchInlineSslProfileListAction_Invoke_Happy(t *testing.T) {
	r := &FetchInlineSslProfileListAction{client: newMockClientStatus(t, 201, "{}")}
	m := FetchInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFetchInlineSslProfileListAction_Invoke_NilClient exercises FetchInlineSslProfileListAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFetchInlineSslProfileListAction_Invoke_NilClient(t *testing.T) {
	r := &FetchInlineSslProfileListAction{}
	m := FetchInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFetchInlineSslProfileListAction_Invoke_BuildError exercises FetchInlineSslProfileListAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFetchInlineSslProfileListAction_Invoke_BuildError(t *testing.T) {
	r := &FetchInlineSslProfileListAction{client: newMalformedBaseURLClient(t)}
	m := FetchInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFetchInlineSslProfileListAction_Invoke_SendError exercises FetchInlineSslProfileListAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestFetchInlineSslProfileListAction_Invoke_SendError(t *testing.T) {
	r := &FetchInlineSslProfileListAction{client: newTransportErrorClient(t)}
	m := FetchInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFetchInlineSslProfileListAction_Invoke_APIError exercises FetchInlineSslProfileListAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFetchInlineSslProfileListAction_Invoke_APIError(t *testing.T) {
	r := &FetchInlineSslProfileListAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FetchInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_fetch_inline_ssl_profile_list")
}

// TestFetchInlineSslProfileListAction_Invoke_APIErrorReadBody exercises FetchInlineSslProfileListAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFetchInlineSslProfileListAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &FetchInlineSslProfileListAction{client: newMockClientReadErrorBody(t, 501)}
	m := FetchInlineSslProfileListActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
