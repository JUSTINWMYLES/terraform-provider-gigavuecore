package provider

import (
	"context"
	"testing"
)

// TestNegativeHbProfileListResource_List_Happy exercises NegativeHbProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestNegativeHbProfileListResource_List_Happy(t *testing.T) {
	r := &NegativeHbProfileListResource{client: newMockClientStatus(t, 200, "{\"negativeHeartbeatProfiles\":[]}")}
	m := NegativeHbProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestNegativeHbProfileListResource_List_NilClient exercises NegativeHbProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNegativeHbProfileListResource_List_NilClient(t *testing.T) {
	r := &NegativeHbProfileListResource{}
	m := NegativeHbProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestNegativeHbProfileListResource_List_BuildError exercises NegativeHbProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestNegativeHbProfileListResource_List_BuildError(t *testing.T) {
	r := &NegativeHbProfileListResource{client: newMalformedBaseURLClient(t)}
	m := NegativeHbProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestNegativeHbProfileListResource_List_SendError exercises NegativeHbProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestNegativeHbProfileListResource_List_SendError(t *testing.T) {
	r := &NegativeHbProfileListResource{client: newTransportErrorClient(t)}
	m := NegativeHbProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestNegativeHbProfileListResource_List_InvalidJSON exercises NegativeHbProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestNegativeHbProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &NegativeHbProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	m := NegativeHbProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
