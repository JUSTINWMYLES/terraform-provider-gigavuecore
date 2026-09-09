package provider

import (
	"context"
	"testing"
)

// TestCircuitTunnelListResource_List_Happy exercises CircuitTunnelListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestCircuitTunnelListResource_List_Happy(t *testing.T) {
	r := &CircuitTunnelListResource{client: newMockClientStatus(t, 200, "{\"circuitTunnels\":[]}")}
	m := CircuitTunnelListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestCircuitTunnelListResource_List_NilClient exercises CircuitTunnelListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCircuitTunnelListResource_List_NilClient(t *testing.T) {
	r := &CircuitTunnelListResource{}
	m := CircuitTunnelListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestCircuitTunnelListResource_List_BuildError exercises CircuitTunnelListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestCircuitTunnelListResource_List_BuildError(t *testing.T) {
	r := &CircuitTunnelListResource{client: newMalformedBaseURLClient(t)}
	m := CircuitTunnelListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestCircuitTunnelListResource_List_SendError exercises CircuitTunnelListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestCircuitTunnelListResource_List_SendError(t *testing.T) {
	r := &CircuitTunnelListResource{client: newTransportErrorClient(t)}
	m := CircuitTunnelListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestCircuitTunnelListResource_List_InvalidJSON exercises CircuitTunnelListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestCircuitTunnelListResource_List_InvalidJSON(t *testing.T) {
	r := &CircuitTunnelListResource{client: newMockClientStatus(t, 200, "{{")}
	m := CircuitTunnelListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
