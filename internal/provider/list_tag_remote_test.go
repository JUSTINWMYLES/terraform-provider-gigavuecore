package provider

import (
	"context"
	"testing"
)

// TestTagListResource_List_Happy exercises TagListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestTagListResource_List_Happy(t *testing.T) {
	r := &TagListResource{client: newMockClientStatus(t, 200, "{\"tags\":[]}")}
	m := TagListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestTagListResource_List_NilClient exercises TagListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTagListResource_List_NilClient(t *testing.T) {
	r := &TagListResource{}
	m := TagListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestTagListResource_List_BuildError exercises TagListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestTagListResource_List_BuildError(t *testing.T) {
	r := &TagListResource{client: newMalformedBaseURLClient(t)}
	m := TagListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestTagListResource_List_SendError exercises TagListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestTagListResource_List_SendError(t *testing.T) {
	r := &TagListResource{client: newTransportErrorClient(t)}
	m := TagListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestTagListResource_List_InvalidJSON exercises TagListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestTagListResource_List_InvalidJSON(t *testing.T) {
	r := &TagListResource{client: newMockClientStatus(t, 200, "{{")}
	m := TagListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
