package provider

import (
	"context"
	"testing"
)

// TestLoadGigaPortsListResource_List_Happy exercises LoadGigaPortsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadGigaPortsListResource_List_Happy(t *testing.T) {
	r := &LoadGigaPortsListResource{client: newMockClientStatus(t, 200, "{\"ports\":[]}")}
	m := LoadGigaPortsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadGigaPortsListResource_List_NilClient exercises LoadGigaPortsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadGigaPortsListResource_List_NilClient(t *testing.T) {
	r := &LoadGigaPortsListResource{}
	m := LoadGigaPortsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadGigaPortsListResource_List_BuildError exercises LoadGigaPortsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadGigaPortsListResource_List_BuildError(t *testing.T) {
	r := &LoadGigaPortsListResource{client: newMalformedBaseURLClient(t)}
	m := LoadGigaPortsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadGigaPortsListResource_List_SendError exercises LoadGigaPortsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadGigaPortsListResource_List_SendError(t *testing.T) {
	r := &LoadGigaPortsListResource{client: newTransportErrorClient(t)}
	m := LoadGigaPortsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadGigaPortsListResource_List_InvalidJSON exercises LoadGigaPortsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadGigaPortsListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadGigaPortsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadGigaPortsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
