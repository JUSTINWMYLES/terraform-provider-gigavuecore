package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteImageFileAction_Invoke_Happy exercises DeleteImageFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteImageFileAction_Invoke_Happy(t *testing.T) {
	r := &DeleteImageFileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteImageFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteImageFileAction_Invoke_NilClient exercises DeleteImageFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteImageFileAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteImageFileAction{}
	m := DeleteImageFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteImageFileAction_Invoke_BuildError exercises DeleteImageFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteImageFileAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteImageFileAction{client: newMalformedBaseURLClient(t)}
	m := DeleteImageFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteImageFileAction_Invoke_SendError exercises DeleteImageFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteImageFileAction_Invoke_SendError(t *testing.T) {
	r := &DeleteImageFileAction{client: newTransportErrorClient(t)}
	m := DeleteImageFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteImageFileAction_Invoke_APIError exercises DeleteImageFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteImageFileAction_Invoke_APIError(t *testing.T) {
	r := &DeleteImageFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteImageFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_image_file")
}

// TestDeleteImageFileAction_Invoke_APIErrorReadBody exercises DeleteImageFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteImageFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteImageFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteImageFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
