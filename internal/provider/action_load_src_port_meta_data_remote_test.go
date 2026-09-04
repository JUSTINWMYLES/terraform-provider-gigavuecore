package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestLoadSrcPortMetaDataAction_Invoke_Happy exercises LoadSrcPortMetaDataAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestLoadSrcPortMetaDataAction_Invoke_Happy(t *testing.T) {
	r := &LoadSrcPortMetaDataAction{client: newMockClientStatus(t, 200, "{}")}
	m := LoadSrcPortMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadSrcPortMetaDataAction_Invoke_NilClient exercises LoadSrcPortMetaDataAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadSrcPortMetaDataAction_Invoke_NilClient(t *testing.T) {
	r := &LoadSrcPortMetaDataAction{}
	m := LoadSrcPortMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadSrcPortMetaDataAction_Invoke_BuildError exercises LoadSrcPortMetaDataAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadSrcPortMetaDataAction_Invoke_BuildError(t *testing.T) {
	r := &LoadSrcPortMetaDataAction{client: newMalformedBaseURLClient(t)}
	m := LoadSrcPortMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadSrcPortMetaDataAction_Invoke_SendError exercises LoadSrcPortMetaDataAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadSrcPortMetaDataAction_Invoke_SendError(t *testing.T) {
	r := &LoadSrcPortMetaDataAction{client: newTransportErrorClient(t)}
	m := LoadSrcPortMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadSrcPortMetaDataAction_Invoke_APIError exercises LoadSrcPortMetaDataAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadSrcPortMetaDataAction_Invoke_APIError(t *testing.T) {
	r := &LoadSrcPortMetaDataAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadSrcPortMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_load_src_port_meta_data")
}

// TestLoadSrcPortMetaDataAction_Invoke_APIErrorReadBody exercises LoadSrcPortMetaDataAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadSrcPortMetaDataAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &LoadSrcPortMetaDataAction{client: newMockClientReadErrorBody(t, 501)}
	m := LoadSrcPortMetaDataActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
