package provider

import (
	"context"
	"testing"
)

// TestNetworkLagListResource_List_Happy exercises NetworkLagListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestNetworkLagListResource_List_Happy(t *testing.T) {
	r := &NetworkLagListResource{client: newMockClientStatus(t, 200, "{\"inlineNetworkLags\":[]}")}
	m := NetworkLagListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestNetworkLagListResource_List_NilClient exercises NetworkLagListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkLagListResource_List_NilClient(t *testing.T) {
	r := &NetworkLagListResource{}
	m := NetworkLagListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestNetworkLagListResource_List_BuildError exercises NetworkLagListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestNetworkLagListResource_List_BuildError(t *testing.T) {
	r := &NetworkLagListResource{client: newMalformedBaseURLClient(t)}
	m := NetworkLagListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestNetworkLagListResource_List_SendError exercises NetworkLagListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestNetworkLagListResource_List_SendError(t *testing.T) {
	r := &NetworkLagListResource{client: newTransportErrorClient(t)}
	m := NetworkLagListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestNetworkLagListResource_List_InvalidJSON exercises NetworkLagListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestNetworkLagListResource_List_InvalidJSON(t *testing.T) {
	r := &NetworkLagListResource{client: newMockClientStatus(t, 200, "{{")}
	m := NetworkLagListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
