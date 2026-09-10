package provider

import (
	"context"
	"testing"
)

// TestLoadAllFilterResourcesListResource_List_Happy exercises LoadAllFilterResourcesListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllFilterResourcesListResource_List_Happy(t *testing.T) {
	r := &LoadAllFilterResourcesListResource{client: newMockClientStatus(t, 200, "{\"filterResources\":[]}")}
	m := LoadAllFilterResourcesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadAllFilterResourcesListResource_List_NilClient exercises LoadAllFilterResourcesListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllFilterResourcesListResource_List_NilClient(t *testing.T) {
	r := &LoadAllFilterResourcesListResource{}
	m := LoadAllFilterResourcesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadAllFilterResourcesListResource_List_BuildError exercises LoadAllFilterResourcesListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllFilterResourcesListResource_List_BuildError(t *testing.T) {
	r := &LoadAllFilterResourcesListResource{client: newMalformedBaseURLClient(t)}
	m := LoadAllFilterResourcesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllFilterResourcesListResource_List_SendError exercises LoadAllFilterResourcesListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllFilterResourcesListResource_List_SendError(t *testing.T) {
	r := &LoadAllFilterResourcesListResource{client: newTransportErrorClient(t)}
	m := LoadAllFilterResourcesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllFilterResourcesListResource_List_InvalidJSON exercises LoadAllFilterResourcesListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllFilterResourcesListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadAllFilterResourcesListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllFilterResourcesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
