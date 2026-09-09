package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpgradeClusterGsCardImageAction_Invoke_Happy exercises UpgradeClusterGsCardImageAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpgradeClusterGsCardImageAction_Invoke_Happy(t *testing.T) {
	r := &UpgradeClusterGsCardImageAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpgradeClusterGsCardImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpgradeClusterGsCardImageAction_Invoke_NilClient exercises UpgradeClusterGsCardImageAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpgradeClusterGsCardImageAction_Invoke_NilClient(t *testing.T) {
	r := &UpgradeClusterGsCardImageAction{}
	m := UpgradeClusterGsCardImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpgradeClusterGsCardImageAction_Invoke_BuildError exercises UpgradeClusterGsCardImageAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpgradeClusterGsCardImageAction_Invoke_BuildError(t *testing.T) {
	r := &UpgradeClusterGsCardImageAction{client: newMalformedBaseURLClient(t)}
	m := UpgradeClusterGsCardImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpgradeClusterGsCardImageAction_Invoke_SendError exercises UpgradeClusterGsCardImageAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpgradeClusterGsCardImageAction_Invoke_SendError(t *testing.T) {
	r := &UpgradeClusterGsCardImageAction{client: newTransportErrorClient(t)}
	m := UpgradeClusterGsCardImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpgradeClusterGsCardImageAction_Invoke_APIError exercises UpgradeClusterGsCardImageAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpgradeClusterGsCardImageAction_Invoke_APIError(t *testing.T) {
	r := &UpgradeClusterGsCardImageAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpgradeClusterGsCardImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_upgrade_cluster_gs_card_image")
}

// TestUpgradeClusterGsCardImageAction_Invoke_APIErrorReadBody exercises UpgradeClusterGsCardImageAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpgradeClusterGsCardImageAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpgradeClusterGsCardImageAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpgradeClusterGsCardImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
