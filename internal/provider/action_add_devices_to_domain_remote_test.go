package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddDevicesToDomainAction_Invoke_Happy exercises AddDevicesToDomainAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddDevicesToDomainAction_Invoke_Happy(t *testing.T) {
	r := &AddDevicesToDomainAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddDevicesToDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddDevicesToDomainAction_Invoke_NilClient exercises AddDevicesToDomainAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddDevicesToDomainAction_Invoke_NilClient(t *testing.T) {
	r := &AddDevicesToDomainAction{}
	m := AddDevicesToDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddDevicesToDomainAction_Invoke_BuildError exercises AddDevicesToDomainAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddDevicesToDomainAction_Invoke_BuildError(t *testing.T) {
	r := &AddDevicesToDomainAction{client: newMalformedBaseURLClient(t)}
	m := AddDevicesToDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddDevicesToDomainAction_Invoke_SendError exercises AddDevicesToDomainAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddDevicesToDomainAction_Invoke_SendError(t *testing.T) {
	r := &AddDevicesToDomainAction{client: newTransportErrorClient(t)}
	m := AddDevicesToDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddDevicesToDomainAction_Invoke_APIError exercises AddDevicesToDomainAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddDevicesToDomainAction_Invoke_APIError(t *testing.T) {
	r := &AddDevicesToDomainAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddDevicesToDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_devices_to_domain")
}

// TestAddDevicesToDomainAction_Invoke_APIErrorReadBody exercises AddDevicesToDomainAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddDevicesToDomainAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddDevicesToDomainAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddDevicesToDomainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
