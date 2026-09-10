package provider

import (
	"context"
	"testing"
)

// TestPortGroupListResource_List_Happy exercises PortGroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestPortGroupListResource_List_Happy(t *testing.T) {
	r := &PortGroupListResource{client: newMockClientStatus(t, 200, "{\"portGroups\":[]}")}
	m := PortGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestPortGroupListResource_List_NilClient exercises PortGroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortGroupListResource_List_NilClient(t *testing.T) {
	r := &PortGroupListResource{}
	m := PortGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestPortGroupListResource_List_BuildError exercises PortGroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestPortGroupListResource_List_BuildError(t *testing.T) {
	r := &PortGroupListResource{client: newMalformedBaseURLClient(t)}
	m := PortGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPortGroupListResource_List_SendError exercises PortGroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestPortGroupListResource_List_SendError(t *testing.T) {
	r := &PortGroupListResource{client: newTransportErrorClient(t)}
	m := PortGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPortGroupListResource_List_InvalidJSON exercises PortGroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestPortGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &PortGroupListResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
