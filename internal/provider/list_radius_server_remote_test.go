package provider

import (
	"context"
	"testing"
)

// TestRadiusServerListResource_List_Happy exercises RadiusServerListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestRadiusServerListResource_List_Happy(t *testing.T) {
	r := &RadiusServerListResource{client: newMockClientStatus(t, 200, "{\"radiusServers\":[]}")}
	m := RadiusServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestRadiusServerListResource_List_NilClient exercises RadiusServerListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRadiusServerListResource_List_NilClient(t *testing.T) {
	r := &RadiusServerListResource{}
	m := RadiusServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestRadiusServerListResource_List_BuildError exercises RadiusServerListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestRadiusServerListResource_List_BuildError(t *testing.T) {
	r := &RadiusServerListResource{client: newMalformedBaseURLClient(t)}
	m := RadiusServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestRadiusServerListResource_List_SendError exercises RadiusServerListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestRadiusServerListResource_List_SendError(t *testing.T) {
	r := &RadiusServerListResource{client: newTransportErrorClient(t)}
	m := RadiusServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestRadiusServerListResource_List_InvalidJSON exercises RadiusServerListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestRadiusServerListResource_List_InvalidJSON(t *testing.T) {
	r := &RadiusServerListResource{client: newMockClientStatus(t, 200, "{{")}
	m := RadiusServerListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
