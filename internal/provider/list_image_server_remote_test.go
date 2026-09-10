package provider

import (
	"context"
	"testing"
)

// TestImageServerListResource_List_Happy exercises ImageServerListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestImageServerListResource_List_Happy(t *testing.T) {
	r := &ImageServerListResource{client: newMockClientStatus(t, 200, "{\"imageServers\":[]}")}
	m := ImageServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestImageServerListResource_List_NilClient exercises ImageServerListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestImageServerListResource_List_NilClient(t *testing.T) {
	r := &ImageServerListResource{}
	m := ImageServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestImageServerListResource_List_BuildError exercises ImageServerListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestImageServerListResource_List_BuildError(t *testing.T) {
	r := &ImageServerListResource{client: newMalformedBaseURLClient(t)}
	m := ImageServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestImageServerListResource_List_SendError exercises ImageServerListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestImageServerListResource_List_SendError(t *testing.T) {
	r := &ImageServerListResource{client: newTransportErrorClient(t)}
	m := ImageServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestImageServerListResource_List_InvalidJSON exercises ImageServerListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestImageServerListResource_List_InvalidJSON(t *testing.T) {
	r := &ImageServerListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ImageServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
