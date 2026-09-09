package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_Happy exercises DeleteAcmecertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfDeviceAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAcmecertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_NilClient exercises DeleteAcmecertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfDeviceAction{}
	m := DeleteAcmecertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_BuildError exercises DeleteAcmecertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfDeviceAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAcmecertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_SendError exercises DeleteAcmecertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfDeviceAction{client: newTransportErrorClient(t)}
	m := DeleteAcmecertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_APIError exercises DeleteAcmecertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfDeviceAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAcmecertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_acmecertificate_details_of_device")
}

// TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_APIErrorReadBody exercises DeleteAcmecertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAcmecertificateDetailsOfDeviceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAcmecertificateDetailsOfDeviceAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAcmecertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
