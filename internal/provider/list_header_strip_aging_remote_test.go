package provider

import (
	"context"
	"testing"
)

// TestHeaderStripAgingListResource_List_Happy exercises HeaderStripAgingListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestHeaderStripAgingListResource_List_Happy(t *testing.T) {
	r := &HeaderStripAgingListResource{client: newMockClientStatus(t, 200, "{\"headerStripsAgingDef\":[]}")}
	m := HeaderStripAgingListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestHeaderStripAgingListResource_List_NilClient exercises HeaderStripAgingListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHeaderStripAgingListResource_List_NilClient(t *testing.T) {
	r := &HeaderStripAgingListResource{}
	m := HeaderStripAgingListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestHeaderStripAgingListResource_List_BuildError exercises HeaderStripAgingListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestHeaderStripAgingListResource_List_BuildError(t *testing.T) {
	r := &HeaderStripAgingListResource{client: newMalformedBaseURLClient(t)}
	m := HeaderStripAgingListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestHeaderStripAgingListResource_List_SendError exercises HeaderStripAgingListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestHeaderStripAgingListResource_List_SendError(t *testing.T) {
	r := &HeaderStripAgingListResource{client: newTransportErrorClient(t)}
	m := HeaderStripAgingListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestHeaderStripAgingListResource_List_InvalidJSON exercises HeaderStripAgingListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestHeaderStripAgingListResource_List_InvalidJSON(t *testing.T) {
	r := &HeaderStripAgingListResource{client: newMockClientStatus(t, 200, "{{")}
	m := HeaderStripAgingListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
