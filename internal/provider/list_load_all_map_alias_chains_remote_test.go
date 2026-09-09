package provider

import (
	"context"
	"testing"
)

// TestLoadAllMapAliasChainsListResource_List_Happy exercises LoadAllMapAliasChainsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllMapAliasChainsListResource_List_Happy(t *testing.T) {
	r := &LoadAllMapAliasChainsListResource{client: newMockClientStatus(t, 200, "{\"mapAliasChains\":[]}")}
	m := LoadAllMapAliasChainsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadAllMapAliasChainsListResource_List_NilClient exercises LoadAllMapAliasChainsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllMapAliasChainsListResource_List_NilClient(t *testing.T) {
	r := &LoadAllMapAliasChainsListResource{}
	m := LoadAllMapAliasChainsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadAllMapAliasChainsListResource_List_BuildError exercises LoadAllMapAliasChainsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllMapAliasChainsListResource_List_BuildError(t *testing.T) {
	r := &LoadAllMapAliasChainsListResource{client: newMalformedBaseURLClient(t)}
	m := LoadAllMapAliasChainsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllMapAliasChainsListResource_List_SendError exercises LoadAllMapAliasChainsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllMapAliasChainsListResource_List_SendError(t *testing.T) {
	r := &LoadAllMapAliasChainsListResource{client: newTransportErrorClient(t)}
	m := LoadAllMapAliasChainsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllMapAliasChainsListResource_List_InvalidJSON exercises LoadAllMapAliasChainsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllMapAliasChainsListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadAllMapAliasChainsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllMapAliasChainsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
