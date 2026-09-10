package provider

import (
	"context"
	"testing"
)

// TestSslProfileListResource_List_Happy exercises SslProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestSslProfileListResource_List_Happy(t *testing.T) {
	r := &SslProfileListResource{client: newMockClientStatus(t, 200, "{\"appsSslProfiles\":[]}")}
	m := SslProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestSslProfileListResource_List_NilClient exercises SslProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSslProfileListResource_List_NilClient(t *testing.T) {
	r := &SslProfileListResource{}
	m := SslProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestSslProfileListResource_List_BuildError exercises SslProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestSslProfileListResource_List_BuildError(t *testing.T) {
	r := &SslProfileListResource{client: newMalformedBaseURLClient(t)}
	m := SslProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSslProfileListResource_List_SendError exercises SslProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestSslProfileListResource_List_SendError(t *testing.T) {
	r := &SslProfileListResource{client: newTransportErrorClient(t)}
	m := SslProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSslProfileListResource_List_InvalidJSON exercises SslProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestSslProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &SslProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	m := SslProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
