package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRediscoverDomainClustersAction_Invoke_Happy exercises RediscoverDomainClustersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRediscoverDomainClustersAction_Invoke_Happy(t *testing.T) {
	r := &RediscoverDomainClustersAction{client: newMockClientStatus(t, 200, "{}")}
	m := RediscoverDomainClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRediscoverDomainClustersAction_Invoke_NilClient exercises RediscoverDomainClustersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRediscoverDomainClustersAction_Invoke_NilClient(t *testing.T) {
	r := &RediscoverDomainClustersAction{}
	m := RediscoverDomainClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRediscoverDomainClustersAction_Invoke_BuildError exercises RediscoverDomainClustersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRediscoverDomainClustersAction_Invoke_BuildError(t *testing.T) {
	r := &RediscoverDomainClustersAction{client: newMalformedBaseURLClient(t)}
	m := RediscoverDomainClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRediscoverDomainClustersAction_Invoke_SendError exercises RediscoverDomainClustersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRediscoverDomainClustersAction_Invoke_SendError(t *testing.T) {
	r := &RediscoverDomainClustersAction{client: newTransportErrorClient(t)}
	m := RediscoverDomainClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRediscoverDomainClustersAction_Invoke_APIError exercises RediscoverDomainClustersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRediscoverDomainClustersAction_Invoke_APIError(t *testing.T) {
	r := &RediscoverDomainClustersAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := RediscoverDomainClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_rediscover_domain_clusters")
}

// TestRediscoverDomainClustersAction_Invoke_APIErrorReadBody exercises RediscoverDomainClustersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRediscoverDomainClustersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RediscoverDomainClustersAction{client: newMockClientReadErrorBody(t, 500)}
	m := RediscoverDomainClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
