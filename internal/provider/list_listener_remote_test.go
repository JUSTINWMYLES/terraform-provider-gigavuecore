package provider

import (
	"context"
	"testing"
)

// TestListenerListResource_List_Happy exercises ListenerListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestListenerListResource_List_Happy(t *testing.T) {
	r := &ListenerListResource{client: newMockClientStatus(t, 200, "{\"appsListeners\":[]}")}
	m := ListenerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestListenerListResource_List_NilClient exercises ListenerListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListenerListResource_List_NilClient(t *testing.T) {
	r := &ListenerListResource{}
	m := ListenerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestListenerListResource_List_BuildError exercises ListenerListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestListenerListResource_List_BuildError(t *testing.T) {
	r := &ListenerListResource{client: newMalformedBaseURLClient(t)}
	m := ListenerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestListenerListResource_List_SendError exercises ListenerListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestListenerListResource_List_SendError(t *testing.T) {
	r := &ListenerListResource{client: newTransportErrorClient(t)}
	m := ListenerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestListenerListResource_List_InvalidJSON exercises ListenerListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestListenerListResource_List_InvalidJSON(t *testing.T) {
	r := &ListenerListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ListenerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
