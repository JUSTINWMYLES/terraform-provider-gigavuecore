package provider

import (
	"context"
	"testing"
)

// TestServerListResource_List_Happy exercises ServerListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestServerListResource_List_Happy(t *testing.T) {
	r := &ServerListResource{client: newMockClientStatus(t, 200, "{\"icapServers\":[]}")}
	m := ServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestServerListResource_List_NilClient exercises ServerListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServerListResource_List_NilClient(t *testing.T) {
	r := &ServerListResource{}
	m := ServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestServerListResource_List_BuildError exercises ServerListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestServerListResource_List_BuildError(t *testing.T) {
	r := &ServerListResource{client: newMalformedBaseURLClient(t)}
	m := ServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestServerListResource_List_SendError exercises ServerListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestServerListResource_List_SendError(t *testing.T) {
	r := &ServerListResource{client: newTransportErrorClient(t)}
	m := ServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestServerListResource_List_InvalidJSON exercises ServerListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestServerListResource_List_InvalidJSON(t *testing.T) {
	r := &ServerListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
