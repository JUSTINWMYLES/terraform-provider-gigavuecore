package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestLoadAliasesMetaDataAction_Invoke_Happy exercises LoadAliasesMetaDataAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestLoadAliasesMetaDataAction_Invoke_Happy(t *testing.T) {
	r := &LoadAliasesMetaDataAction{client: newMockClientStatus(t, 200, "{}")}
	m := LoadAliasesMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAliasesMetaDataAction_Invoke_NilClient exercises LoadAliasesMetaDataAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAliasesMetaDataAction_Invoke_NilClient(t *testing.T) {
	r := &LoadAliasesMetaDataAction{}
	m := LoadAliasesMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAliasesMetaDataAction_Invoke_BuildError exercises LoadAliasesMetaDataAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadAliasesMetaDataAction_Invoke_BuildError(t *testing.T) {
	r := &LoadAliasesMetaDataAction{client: newMalformedBaseURLClient(t)}
	m := LoadAliasesMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadAliasesMetaDataAction_Invoke_SendError exercises LoadAliasesMetaDataAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadAliasesMetaDataAction_Invoke_SendError(t *testing.T) {
	r := &LoadAliasesMetaDataAction{client: newTransportErrorClient(t)}
	m := LoadAliasesMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadAliasesMetaDataAction_Invoke_APIError exercises LoadAliasesMetaDataAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadAliasesMetaDataAction_Invoke_APIError(t *testing.T) {
	r := &LoadAliasesMetaDataAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadAliasesMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_load_aliases_meta_data")
}

// TestLoadAliasesMetaDataAction_Invoke_APIErrorReadBody exercises LoadAliasesMetaDataAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadAliasesMetaDataAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &LoadAliasesMetaDataAction{client: newMockClientReadErrorBody(t, 501)}
	m := LoadAliasesMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
