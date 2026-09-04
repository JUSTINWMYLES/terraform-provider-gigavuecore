package provider

import (
	"context"
	"testing"
)

// TestLoadEventsListResource_List_Happy exercises LoadEventsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadEventsListResource_List_Happy(t *testing.T) {
	r := &LoadEventsListResource{client: newMockClientStatus(t, 200, "{\"events\":[]}")}
	m := LoadEventsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadEventsListResource_List_NilClient exercises LoadEventsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadEventsListResource_List_NilClient(t *testing.T) {
	r := &LoadEventsListResource{}
	m := LoadEventsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadEventsListResource_List_BuildError exercises LoadEventsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadEventsListResource_List_BuildError(t *testing.T) {
	r := &LoadEventsListResource{client: newMalformedBaseURLClient(t)}
	m := LoadEventsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadEventsListResource_List_SendError exercises LoadEventsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadEventsListResource_List_SendError(t *testing.T) {
	r := &LoadEventsListResource{client: newTransportErrorClient(t)}
	m := LoadEventsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadEventsListResource_List_InvalidJSON exercises LoadEventsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadEventsListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadEventsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadEventsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
