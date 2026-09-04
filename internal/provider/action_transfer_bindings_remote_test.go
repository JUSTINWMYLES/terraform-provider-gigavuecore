package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestTransferBindingsAction_Invoke_Happy exercises TransferBindingsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestTransferBindingsAction_Invoke_Happy(t *testing.T) {
	r := &TransferBindingsAction{client: newMockClientStatus(t, 200, "{}")}
	m := TransferBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTransferBindingsAction_Invoke_NilClient exercises TransferBindingsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTransferBindingsAction_Invoke_NilClient(t *testing.T) {
	r := &TransferBindingsAction{}
	m := TransferBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTransferBindingsAction_Invoke_BuildError exercises TransferBindingsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTransferBindingsAction_Invoke_BuildError(t *testing.T) {
	r := &TransferBindingsAction{client: newMalformedBaseURLClient(t)}
	m := TransferBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTransferBindingsAction_Invoke_SendError exercises TransferBindingsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestTransferBindingsAction_Invoke_SendError(t *testing.T) {
	r := &TransferBindingsAction{client: newTransportErrorClient(t)}
	m := TransferBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTransferBindingsAction_Invoke_APIError exercises TransferBindingsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTransferBindingsAction_Invoke_APIError(t *testing.T) {
	r := &TransferBindingsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TransferBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_transfer_bindings")
}

// TestTransferBindingsAction_Invoke_APIErrorReadBody exercises TransferBindingsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTransferBindingsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &TransferBindingsAction{client: newMockClientReadErrorBody(t, 501)}
	m := TransferBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
