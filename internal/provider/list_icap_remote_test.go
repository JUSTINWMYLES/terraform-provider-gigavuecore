package provider

import (
	"context"
	"testing"
)

// TestIcapListResource_List_Happy exercises IcapListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestIcapListResource_List_Happy(t *testing.T) {
	r := &IcapListResource{client: newMockClientStatus(t, 200, "{\"gigaIcapClients\":[]}")}
	m := IcapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestIcapListResource_List_NilClient exercises IcapListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIcapListResource_List_NilClient(t *testing.T) {
	r := &IcapListResource{}
	m := IcapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestIcapListResource_List_BuildError exercises IcapListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestIcapListResource_List_BuildError(t *testing.T) {
	r := &IcapListResource{client: newMalformedBaseURLClient(t)}
	m := IcapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestIcapListResource_List_SendError exercises IcapListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestIcapListResource_List_SendError(t *testing.T) {
	r := &IcapListResource{client: newTransportErrorClient(t)}
	m := IcapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestIcapListResource_List_InvalidJSON exercises IcapListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestIcapListResource_List_InvalidJSON(t *testing.T) {
	r := &IcapListResource{client: newMockClientStatus(t, 200, "{{")}
	m := IcapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
