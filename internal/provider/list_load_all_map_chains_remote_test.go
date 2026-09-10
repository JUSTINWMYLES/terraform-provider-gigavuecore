package provider

import (
	"context"
	"testing"
)

// TestLoadAllMapChainsListResource_List_Happy exercises LoadAllMapChainsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllMapChainsListResource_List_Happy(t *testing.T) {
	r := &LoadAllMapChainsListResource{client: newMockClientStatus(t, 200, "{\"mapChains\":[]}")}
	m := LoadAllMapChainsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadAllMapChainsListResource_List_NilClient exercises LoadAllMapChainsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllMapChainsListResource_List_NilClient(t *testing.T) {
	r := &LoadAllMapChainsListResource{}
	m := LoadAllMapChainsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadAllMapChainsListResource_List_BuildError exercises LoadAllMapChainsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllMapChainsListResource_List_BuildError(t *testing.T) {
	r := &LoadAllMapChainsListResource{client: newMalformedBaseURLClient(t)}
	m := LoadAllMapChainsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllMapChainsListResource_List_SendError exercises LoadAllMapChainsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllMapChainsListResource_List_SendError(t *testing.T) {
	r := &LoadAllMapChainsListResource{client: newTransportErrorClient(t)}
	m := LoadAllMapChainsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllMapChainsListResource_List_InvalidJSON exercises LoadAllMapChainsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllMapChainsListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadAllMapChainsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllMapChainsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
