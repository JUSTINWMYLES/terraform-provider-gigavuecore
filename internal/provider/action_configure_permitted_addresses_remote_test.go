package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConfigurePermittedAddressesAction_Invoke_Happy exercises ConfigurePermittedAddressesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConfigurePermittedAddressesAction_Invoke_Happy(t *testing.T) {
	r := &ConfigurePermittedAddressesAction{client: newMockClientStatus(t, 200, "{}")}
	m := ConfigurePermittedAddressesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConfigurePermittedAddressesAction_Invoke_NilClient exercises ConfigurePermittedAddressesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConfigurePermittedAddressesAction_Invoke_NilClient(t *testing.T) {
	r := &ConfigurePermittedAddressesAction{}
	m := ConfigurePermittedAddressesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConfigurePermittedAddressesAction_Invoke_BuildError exercises ConfigurePermittedAddressesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConfigurePermittedAddressesAction_Invoke_BuildError(t *testing.T) {
	r := &ConfigurePermittedAddressesAction{client: newMalformedBaseURLClient(t)}
	m := ConfigurePermittedAddressesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConfigurePermittedAddressesAction_Invoke_SendError exercises ConfigurePermittedAddressesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConfigurePermittedAddressesAction_Invoke_SendError(t *testing.T) {
	r := &ConfigurePermittedAddressesAction{client: newTransportErrorClient(t)}
	m := ConfigurePermittedAddressesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConfigurePermittedAddressesAction_Invoke_APIError exercises ConfigurePermittedAddressesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConfigurePermittedAddressesAction_Invoke_APIError(t *testing.T) {
	r := &ConfigurePermittedAddressesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ConfigurePermittedAddressesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_configure_permitted_addresses")
}

// TestConfigurePermittedAddressesAction_Invoke_APIErrorReadBody exercises ConfigurePermittedAddressesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConfigurePermittedAddressesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConfigurePermittedAddressesAction{client: newMockClientReadErrorBody(t, 501)}
	m := ConfigurePermittedAddressesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
