package provider

import (
	"context"
	"testing"
)

// TestLoadAllPortConfigListResource_List_Happy exercises LoadAllPortConfigListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllPortConfigListResource_List_Happy(t *testing.T) {
	r := &LoadAllPortConfigListResource{client: newMockClientStatus(t, 200, "{\"portConfigs\":[]}")}
	m := LoadAllPortConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadAllPortConfigListResource_List_NilClient exercises LoadAllPortConfigListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllPortConfigListResource_List_NilClient(t *testing.T) {
	r := &LoadAllPortConfigListResource{}
	m := LoadAllPortConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadAllPortConfigListResource_List_BuildError exercises LoadAllPortConfigListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllPortConfigListResource_List_BuildError(t *testing.T) {
	r := &LoadAllPortConfigListResource{client: newMalformedBaseURLClient(t)}
	m := LoadAllPortConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllPortConfigListResource_List_SendError exercises LoadAllPortConfigListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllPortConfigListResource_List_SendError(t *testing.T) {
	r := &LoadAllPortConfigListResource{client: newTransportErrorClient(t)}
	m := LoadAllPortConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllPortConfigListResource_List_InvalidJSON exercises LoadAllPortConfigListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllPortConfigListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadAllPortConfigListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllPortConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
