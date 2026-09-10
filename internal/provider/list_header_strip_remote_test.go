package provider

import (
	"context"
	"testing"
)

// TestHeaderStripListResource_List_Happy exercises HeaderStripListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestHeaderStripListResource_List_Happy(t *testing.T) {
	r := &HeaderStripListResource{client: newMockClientStatus(t, 200, "{\"headerStripsDef\":[]}")}
	m := HeaderStripListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestHeaderStripListResource_List_NilClient exercises HeaderStripListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHeaderStripListResource_List_NilClient(t *testing.T) {
	r := &HeaderStripListResource{}
	m := HeaderStripListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestHeaderStripListResource_List_BuildError exercises HeaderStripListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestHeaderStripListResource_List_BuildError(t *testing.T) {
	r := &HeaderStripListResource{client: newMalformedBaseURLClient(t)}
	m := HeaderStripListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestHeaderStripListResource_List_SendError exercises HeaderStripListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestHeaderStripListResource_List_SendError(t *testing.T) {
	r := &HeaderStripListResource{client: newTransportErrorClient(t)}
	m := HeaderStripListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestHeaderStripListResource_List_InvalidJSON exercises HeaderStripListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestHeaderStripListResource_List_InvalidJSON(t *testing.T) {
	r := &HeaderStripListResource{client: newMockClientStatus(t, 200, "{{")}
	m := HeaderStripListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
