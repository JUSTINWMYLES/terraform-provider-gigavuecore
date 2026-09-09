package provider

import (
	"context"
	"testing"
)

// TestLoadAllGtapPortGroupListResource_List_Happy exercises LoadAllGtapPortGroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllGtapPortGroupListResource_List_Happy(t *testing.T) {
	r := &LoadAllGtapPortGroupListResource{client: newMockClientStatus(t, 200, "{\"gtapPortGroups\":[]}")}
	m := LoadAllGtapPortGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadAllGtapPortGroupListResource_List_NilClient exercises LoadAllGtapPortGroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllGtapPortGroupListResource_List_NilClient(t *testing.T) {
	r := &LoadAllGtapPortGroupListResource{}
	m := LoadAllGtapPortGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadAllGtapPortGroupListResource_List_BuildError exercises LoadAllGtapPortGroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllGtapPortGroupListResource_List_BuildError(t *testing.T) {
	r := &LoadAllGtapPortGroupListResource{client: newMalformedBaseURLClient(t)}
	m := LoadAllGtapPortGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllGtapPortGroupListResource_List_SendError exercises LoadAllGtapPortGroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllGtapPortGroupListResource_List_SendError(t *testing.T) {
	r := &LoadAllGtapPortGroupListResource{client: newTransportErrorClient(t)}
	m := LoadAllGtapPortGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllGtapPortGroupListResource_List_InvalidJSON exercises LoadAllGtapPortGroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllGtapPortGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadAllGtapPortGroupListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllGtapPortGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
