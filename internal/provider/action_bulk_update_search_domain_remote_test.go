package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestBulkUpdateSearchDomainAction_Invoke_Happy exercises BulkUpdateSearchDomainAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestBulkUpdateSearchDomainAction_Invoke_Happy(t *testing.T) {
	r := &BulkUpdateSearchDomainAction{client: newMockClientStatus(t, 201, "{}")}
	m := BulkUpdateSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestBulkUpdateSearchDomainAction_Invoke_NilClient exercises BulkUpdateSearchDomainAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestBulkUpdateSearchDomainAction_Invoke_NilClient(t *testing.T) {
	r := &BulkUpdateSearchDomainAction{}
	m := BulkUpdateSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestBulkUpdateSearchDomainAction_Invoke_BuildError exercises BulkUpdateSearchDomainAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestBulkUpdateSearchDomainAction_Invoke_BuildError(t *testing.T) {
	r := &BulkUpdateSearchDomainAction{client: newMalformedBaseURLClient(t)}
	m := BulkUpdateSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestBulkUpdateSearchDomainAction_Invoke_SendError exercises BulkUpdateSearchDomainAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestBulkUpdateSearchDomainAction_Invoke_SendError(t *testing.T) {
	r := &BulkUpdateSearchDomainAction{client: newTransportErrorClient(t)}
	m := BulkUpdateSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestBulkUpdateSearchDomainAction_Invoke_APIError exercises BulkUpdateSearchDomainAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestBulkUpdateSearchDomainAction_Invoke_APIError(t *testing.T) {
	r := &BulkUpdateSearchDomainAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := BulkUpdateSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_bulk_update_search_domain")
}

// TestBulkUpdateSearchDomainAction_Invoke_APIErrorReadBody exercises BulkUpdateSearchDomainAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestBulkUpdateSearchDomainAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &BulkUpdateSearchDomainAction{client: newMockClientReadErrorBody(t, 501)}
	m := BulkUpdateSearchDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
