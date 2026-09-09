package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestBulkUpdateNameserversAction_Invoke_Happy exercises BulkUpdateNameserversAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestBulkUpdateNameserversAction_Invoke_Happy(t *testing.T) {
	r := &BulkUpdateNameserversAction{client: newMockClientStatus(t, 201, "{}")}
	m := BulkUpdateNameserversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestBulkUpdateNameserversAction_Invoke_NilClient exercises BulkUpdateNameserversAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestBulkUpdateNameserversAction_Invoke_NilClient(t *testing.T) {
	r := &BulkUpdateNameserversAction{}
	m := BulkUpdateNameserversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestBulkUpdateNameserversAction_Invoke_BuildError exercises BulkUpdateNameserversAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestBulkUpdateNameserversAction_Invoke_BuildError(t *testing.T) {
	r := &BulkUpdateNameserversAction{client: newMalformedBaseURLClient(t)}
	m := BulkUpdateNameserversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestBulkUpdateNameserversAction_Invoke_SendError exercises BulkUpdateNameserversAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestBulkUpdateNameserversAction_Invoke_SendError(t *testing.T) {
	r := &BulkUpdateNameserversAction{client: newTransportErrorClient(t)}
	m := BulkUpdateNameserversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestBulkUpdateNameserversAction_Invoke_APIError exercises BulkUpdateNameserversAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestBulkUpdateNameserversAction_Invoke_APIError(t *testing.T) {
	r := &BulkUpdateNameserversAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := BulkUpdateNameserversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_bulk_update_nameservers")
}

// TestBulkUpdateNameserversAction_Invoke_APIErrorReadBody exercises BulkUpdateNameserversAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestBulkUpdateNameserversAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &BulkUpdateNameserversAction{client: newMockClientReadErrorBody(t, 501)}
	m := BulkUpdateNameserversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
