package provider

import (
	"context"
	"testing"
)

// TestInlineNetworkGroupListResource_List_Happy exercises InlineNetworkGroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestInlineNetworkGroupListResource_List_Happy(t *testing.T) {
	r := &InlineNetworkGroupListResource{client: newMockClientStatus(t, 200, "{\"inlineNetworkGroups\":[]}")}
	m := InlineNetworkGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestInlineNetworkGroupListResource_List_NilClient exercises InlineNetworkGroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineNetworkGroupListResource_List_NilClient(t *testing.T) {
	r := &InlineNetworkGroupListResource{}
	m := InlineNetworkGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestInlineNetworkGroupListResource_List_BuildError exercises InlineNetworkGroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestInlineNetworkGroupListResource_List_BuildError(t *testing.T) {
	r := &InlineNetworkGroupListResource{client: newMalformedBaseURLClient(t)}
	m := InlineNetworkGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestInlineNetworkGroupListResource_List_SendError exercises InlineNetworkGroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestInlineNetworkGroupListResource_List_SendError(t *testing.T) {
	r := &InlineNetworkGroupListResource{client: newTransportErrorClient(t)}
	m := InlineNetworkGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestInlineNetworkGroupListResource_List_InvalidJSON exercises InlineNetworkGroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestInlineNetworkGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &InlineNetworkGroupListResource{client: newMockClientStatus(t, 200, "{{")}
	m := InlineNetworkGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
