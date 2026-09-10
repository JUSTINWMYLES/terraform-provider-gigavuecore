package provider

import (
	"context"
	"testing"
)

// TestLoadCardsDetailsListResource_List_Happy exercises LoadCardsDetailsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadCardsDetailsListResource_List_Happy(t *testing.T) {
	r := &LoadCardsDetailsListResource{client: newMockClientStatus(t, 200, "{\"cards\":[]}")}
	m := LoadCardsDetailsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadCardsDetailsListResource_List_NilClient exercises LoadCardsDetailsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadCardsDetailsListResource_List_NilClient(t *testing.T) {
	r := &LoadCardsDetailsListResource{}
	m := LoadCardsDetailsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadCardsDetailsListResource_List_BuildError exercises LoadCardsDetailsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadCardsDetailsListResource_List_BuildError(t *testing.T) {
	r := &LoadCardsDetailsListResource{client: newMalformedBaseURLClient(t)}
	m := LoadCardsDetailsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadCardsDetailsListResource_List_SendError exercises LoadCardsDetailsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadCardsDetailsListResource_List_SendError(t *testing.T) {
	r := &LoadCardsDetailsListResource{client: newTransportErrorClient(t)}
	m := LoadCardsDetailsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadCardsDetailsListResource_List_InvalidJSON exercises LoadCardsDetailsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadCardsDetailsListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadCardsDetailsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadCardsDetailsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
