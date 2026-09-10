package provider

import (
	"context"
	"testing"
)

// TestStackLinkListResource_List_Happy exercises StackLinkListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestStackLinkListResource_List_Happy(t *testing.T) {
	r := &StackLinkListResource{client: newMockClientStatus(t, 200, "{\"stackLinks\":[]}")}
	m := StackLinkListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestStackLinkListResource_List_NilClient exercises StackLinkListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestStackLinkListResource_List_NilClient(t *testing.T) {
	r := &StackLinkListResource{}
	m := StackLinkListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestStackLinkListResource_List_BuildError exercises StackLinkListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestStackLinkListResource_List_BuildError(t *testing.T) {
	r := &StackLinkListResource{client: newMalformedBaseURLClient(t)}
	m := StackLinkListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestStackLinkListResource_List_SendError exercises StackLinkListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestStackLinkListResource_List_SendError(t *testing.T) {
	r := &StackLinkListResource{client: newTransportErrorClient(t)}
	m := StackLinkListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestStackLinkListResource_List_InvalidJSON exercises StackLinkListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestStackLinkListResource_List_InvalidJSON(t *testing.T) {
	r := &StackLinkListResource{client: newMockClientStatus(t, 200, "{{")}
	m := StackLinkListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
