package provider

import (
	"context"
	"testing"
)

// TestPortFilterListResource_List_Happy exercises PortFilterListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestPortFilterListResource_List_Happy(t *testing.T) {
	r := &PortFilterListResource{client: newMockClientStatus(t, 200, "{\"portFilters\":[]}")}
	m := PortFilterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestPortFilterListResource_List_NilClient exercises PortFilterListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortFilterListResource_List_NilClient(t *testing.T) {
	r := &PortFilterListResource{}
	m := PortFilterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestPortFilterListResource_List_BuildError exercises PortFilterListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestPortFilterListResource_List_BuildError(t *testing.T) {
	r := &PortFilterListResource{client: newMalformedBaseURLClient(t)}
	m := PortFilterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPortFilterListResource_List_SendError exercises PortFilterListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestPortFilterListResource_List_SendError(t *testing.T) {
	r := &PortFilterListResource{client: newTransportErrorClient(t)}
	m := PortFilterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPortFilterListResource_List_InvalidJSON exercises PortFilterListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestPortFilterListResource_List_InvalidJSON(t *testing.T) {
	r := &PortFilterListResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortFilterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
