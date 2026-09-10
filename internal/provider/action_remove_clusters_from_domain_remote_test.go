package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveClustersFromDomainAction_Invoke_Happy exercises RemoveClustersFromDomainAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveClustersFromDomainAction_Invoke_Happy(t *testing.T) {
	r := &RemoveClustersFromDomainAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveClustersFromDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveClustersFromDomainAction_Invoke_NilClient exercises RemoveClustersFromDomainAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveClustersFromDomainAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveClustersFromDomainAction{}
	m := RemoveClustersFromDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveClustersFromDomainAction_Invoke_BuildError exercises RemoveClustersFromDomainAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveClustersFromDomainAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveClustersFromDomainAction{client: newMalformedBaseURLClient(t)}
	m := RemoveClustersFromDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveClustersFromDomainAction_Invoke_SendError exercises RemoveClustersFromDomainAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveClustersFromDomainAction_Invoke_SendError(t *testing.T) {
	r := &RemoveClustersFromDomainAction{client: newTransportErrorClient(t)}
	m := RemoveClustersFromDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveClustersFromDomainAction_Invoke_APIError exercises RemoveClustersFromDomainAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveClustersFromDomainAction_Invoke_APIError(t *testing.T) {
	r := &RemoveClustersFromDomainAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveClustersFromDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_clusters_from_domain")
}

// TestRemoveClustersFromDomainAction_Invoke_APIErrorReadBody exercises RemoveClustersFromDomainAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveClustersFromDomainAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveClustersFromDomainAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveClustersFromDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
