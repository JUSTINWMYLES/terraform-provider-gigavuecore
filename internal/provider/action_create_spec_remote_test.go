package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateSpecAction_Invoke_Happy exercises CreateSpecAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateSpecAction_Invoke_Happy(t *testing.T) {
	r := &CreateSpecAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateSpecAction_Invoke_NilClient exercises CreateSpecAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateSpecAction_Invoke_NilClient(t *testing.T) {
	r := &CreateSpecAction{}
	m := CreateSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateSpecAction_Invoke_BuildError exercises CreateSpecAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateSpecAction_Invoke_BuildError(t *testing.T) {
	r := &CreateSpecAction{client: newMalformedBaseURLClient(t)}
	m := CreateSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateSpecAction_Invoke_SendError exercises CreateSpecAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateSpecAction_Invoke_SendError(t *testing.T) {
	r := &CreateSpecAction{client: newTransportErrorClient(t)}
	m := CreateSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateSpecAction_Invoke_APIError exercises CreateSpecAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateSpecAction_Invoke_APIError(t *testing.T) {
	r := &CreateSpecAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_spec")
}

// TestCreateSpecAction_Invoke_APIErrorReadBody exercises CreateSpecAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateSpecAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateSpecAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
