package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateFlexInlineConfigAction_Invoke_Happy exercises CreateFlexInlineConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateFlexInlineConfigAction_Invoke_Happy(t *testing.T) {
	r := &CreateFlexInlineConfigAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateFlexInlineConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateFlexInlineConfigAction_Invoke_NilClient exercises CreateFlexInlineConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateFlexInlineConfigAction_Invoke_NilClient(t *testing.T) {
	r := &CreateFlexInlineConfigAction{}
	m := CreateFlexInlineConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateFlexInlineConfigAction_Invoke_BuildError exercises CreateFlexInlineConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateFlexInlineConfigAction_Invoke_BuildError(t *testing.T) {
	r := &CreateFlexInlineConfigAction{client: newMalformedBaseURLClient(t)}
	m := CreateFlexInlineConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateFlexInlineConfigAction_Invoke_SendError exercises CreateFlexInlineConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateFlexInlineConfigAction_Invoke_SendError(t *testing.T) {
	r := &CreateFlexInlineConfigAction{client: newTransportErrorClient(t)}
	m := CreateFlexInlineConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateFlexInlineConfigAction_Invoke_APIError exercises CreateFlexInlineConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateFlexInlineConfigAction_Invoke_APIError(t *testing.T) {
	r := &CreateFlexInlineConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateFlexInlineConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_flex_inline_config")
}

// TestCreateFlexInlineConfigAction_Invoke_APIErrorReadBody exercises CreateFlexInlineConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateFlexInlineConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateFlexInlineConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateFlexInlineConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
