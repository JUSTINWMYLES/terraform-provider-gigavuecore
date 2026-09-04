package provider

import (
	"context"
	"testing"
)

// TestTacacsServerListResource_List_Happy exercises TacacsServerListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestTacacsServerListResource_List_Happy(t *testing.T) {
	r := &TacacsServerListResource{client: newMockClientStatus(t, 200, "{\"tacacsServers\":[]}")}
	m := TacacsServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestTacacsServerListResource_List_NilClient exercises TacacsServerListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTacacsServerListResource_List_NilClient(t *testing.T) {
	r := &TacacsServerListResource{}
	m := TacacsServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestTacacsServerListResource_List_BuildError exercises TacacsServerListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestTacacsServerListResource_List_BuildError(t *testing.T) {
	r := &TacacsServerListResource{client: newMalformedBaseURLClient(t)}
	m := TacacsServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestTacacsServerListResource_List_SendError exercises TacacsServerListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestTacacsServerListResource_List_SendError(t *testing.T) {
	r := &TacacsServerListResource{client: newTransportErrorClient(t)}
	m := TacacsServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestTacacsServerListResource_List_InvalidJSON exercises TacacsServerListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestTacacsServerListResource_List_InvalidJSON(t *testing.T) {
	r := &TacacsServerListResource{client: newMockClientStatus(t, 200, "{{")}
	m := TacacsServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
