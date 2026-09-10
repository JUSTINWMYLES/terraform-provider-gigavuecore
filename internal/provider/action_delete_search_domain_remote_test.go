package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteSearchDomainAction_Invoke_Happy exercises DeleteSearchDomainAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteSearchDomainAction_Invoke_Happy(t *testing.T) {
	r := &DeleteSearchDomainAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteSearchDomainAction_Invoke_NilClient exercises DeleteSearchDomainAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteSearchDomainAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteSearchDomainAction{}
	m := DeleteSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteSearchDomainAction_Invoke_BuildError exercises DeleteSearchDomainAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteSearchDomainAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteSearchDomainAction{client: newMalformedBaseURLClient(t)}
	m := DeleteSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteSearchDomainAction_Invoke_SendError exercises DeleteSearchDomainAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteSearchDomainAction_Invoke_SendError(t *testing.T) {
	r := &DeleteSearchDomainAction{client: newTransportErrorClient(t)}
	m := DeleteSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteSearchDomainAction_Invoke_APIError exercises DeleteSearchDomainAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteSearchDomainAction_Invoke_APIError(t *testing.T) {
	r := &DeleteSearchDomainAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_search_domain")
}

// TestDeleteSearchDomainAction_Invoke_APIErrorReadBody exercises DeleteSearchDomainAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteSearchDomainAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteSearchDomainAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
