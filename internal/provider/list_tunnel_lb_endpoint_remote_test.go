package provider

import (
	"context"
	"testing"
)

// TestTunnelLbEndpointListResource_List_Happy exercises TunnelLbEndpointListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestTunnelLbEndpointListResource_List_Happy(t *testing.T) {
	r := &TunnelLbEndpointListResource{client: newMockClientStatus(t, 200, "{\"tunnelLbEndpoints\":[]}")}
	m := TunnelLbEndpointListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestTunnelLbEndpointListResource_List_NilClient exercises TunnelLbEndpointListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTunnelLbEndpointListResource_List_NilClient(t *testing.T) {
	r := &TunnelLbEndpointListResource{}
	m := TunnelLbEndpointListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestTunnelLbEndpointListResource_List_BuildError exercises TunnelLbEndpointListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestTunnelLbEndpointListResource_List_BuildError(t *testing.T) {
	r := &TunnelLbEndpointListResource{client: newMalformedBaseURLClient(t)}
	m := TunnelLbEndpointListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestTunnelLbEndpointListResource_List_SendError exercises TunnelLbEndpointListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestTunnelLbEndpointListResource_List_SendError(t *testing.T) {
	r := &TunnelLbEndpointListResource{client: newTransportErrorClient(t)}
	m := TunnelLbEndpointListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestTunnelLbEndpointListResource_List_InvalidJSON exercises TunnelLbEndpointListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestTunnelLbEndpointListResource_List_InvalidJSON(t *testing.T) {
	r := &TunnelLbEndpointListResource{client: newMockClientStatus(t, 200, "{{")}
	m := TunnelLbEndpointListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
