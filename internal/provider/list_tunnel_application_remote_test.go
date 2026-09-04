package provider

import (
	"context"
	"testing"
)

// TestTunnelApplicationListResource_List_Happy exercises TunnelApplicationListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestTunnelApplicationListResource_List_Happy(t *testing.T) {
	r := &TunnelApplicationListResource{client: newMockClientStatus(t, 200, "{\"tunnelApps\":[]}")}
	m := TunnelApplicationListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestTunnelApplicationListResource_List_NilClient exercises TunnelApplicationListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTunnelApplicationListResource_List_NilClient(t *testing.T) {
	r := &TunnelApplicationListResource{}
	m := TunnelApplicationListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestTunnelApplicationListResource_List_BuildError exercises TunnelApplicationListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestTunnelApplicationListResource_List_BuildError(t *testing.T) {
	r := &TunnelApplicationListResource{client: newMalformedBaseURLClient(t)}
	m := TunnelApplicationListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestTunnelApplicationListResource_List_SendError exercises TunnelApplicationListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestTunnelApplicationListResource_List_SendError(t *testing.T) {
	r := &TunnelApplicationListResource{client: newTransportErrorClient(t)}
	m := TunnelApplicationListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestTunnelApplicationListResource_List_InvalidJSON exercises TunnelApplicationListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestTunnelApplicationListResource_List_InvalidJSON(t *testing.T) {
	r := &TunnelApplicationListResource{client: newMockClientStatus(t, 200, "{{")}
	m := TunnelApplicationListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
