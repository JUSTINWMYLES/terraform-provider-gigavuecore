package provider

import (
	"context"
	"testing"
)

// TestSaApfProfileListResource_List_Happy exercises SaApfProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestSaApfProfileListResource_List_Happy(t *testing.T) {
	r := &SaApfProfileListResource{client: newMockClientStatus(t, 200, "{\"saApfProfiles\":[]}")}
	m := SaApfProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestSaApfProfileListResource_List_NilClient exercises SaApfProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSaApfProfileListResource_List_NilClient(t *testing.T) {
	r := &SaApfProfileListResource{}
	m := SaApfProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestSaApfProfileListResource_List_BuildError exercises SaApfProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestSaApfProfileListResource_List_BuildError(t *testing.T) {
	r := &SaApfProfileListResource{client: newMalformedBaseURLClient(t)}
	m := SaApfProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSaApfProfileListResource_List_SendError exercises SaApfProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestSaApfProfileListResource_List_SendError(t *testing.T) {
	r := &SaApfProfileListResource{client: newTransportErrorClient(t)}
	m := SaApfProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSaApfProfileListResource_List_InvalidJSON exercises SaApfProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestSaApfProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &SaApfProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	m := SaApfProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
