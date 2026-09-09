package provider

import (
	"context"
	"testing"
)

// TestNetworkProfileListResource_List_Happy exercises NetworkProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestNetworkProfileListResource_List_Happy(t *testing.T) {
	r := &NetworkProfileListResource{client: newMockClientStatus(t, 200, "{\"networkProfiles\":[]}")}
	m := NetworkProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestNetworkProfileListResource_List_NilClient exercises NetworkProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkProfileListResource_List_NilClient(t *testing.T) {
	r := &NetworkProfileListResource{}
	m := NetworkProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestNetworkProfileListResource_List_BuildError exercises NetworkProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestNetworkProfileListResource_List_BuildError(t *testing.T) {
	r := &NetworkProfileListResource{client: newMalformedBaseURLClient(t)}
	m := NetworkProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestNetworkProfileListResource_List_SendError exercises NetworkProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestNetworkProfileListResource_List_SendError(t *testing.T) {
	r := &NetworkProfileListResource{client: newTransportErrorClient(t)}
	m := NetworkProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestNetworkProfileListResource_List_InvalidJSON exercises NetworkProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestNetworkProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &NetworkProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	m := NetworkProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
