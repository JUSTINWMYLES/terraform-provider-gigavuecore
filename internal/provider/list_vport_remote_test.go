package provider

import (
	"context"
	"testing"
)

// TestVportListResource_List_Happy exercises VportListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestVportListResource_List_Happy(t *testing.T) {
	r := &VportListResource{client: newMockClientStatus(t, 200, "{\"vports\":[]}")}
	m := VportListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestVportListResource_List_NilClient exercises VportListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestVportListResource_List_NilClient(t *testing.T) {
	r := &VportListResource{}
	m := VportListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestVportListResource_List_BuildError exercises VportListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestVportListResource_List_BuildError(t *testing.T) {
	r := &VportListResource{client: newMalformedBaseURLClient(t)}
	m := VportListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestVportListResource_List_SendError exercises VportListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestVportListResource_List_SendError(t *testing.T) {
	r := &VportListResource{client: newTransportErrorClient(t)}
	m := VportListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestVportListResource_List_InvalidJSON exercises VportListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestVportListResource_List_InvalidJSON(t *testing.T) {
	r := &VportListResource{client: newMockClientStatus(t, 200, "{{")}
	m := VportListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
