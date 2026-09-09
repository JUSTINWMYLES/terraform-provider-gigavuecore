package provider

import (
	"context"
	"testing"
)

// TestFabricMapListResource_List_Happy exercises FabricMapListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestFabricMapListResource_List_Happy(t *testing.T) {
	r := &FabricMapListResource{client: newMockClientStatus(t, 200, "{\"fabricMaps\":[]}")}
	m := FabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestFabricMapListResource_List_NilClient exercises FabricMapListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFabricMapListResource_List_NilClient(t *testing.T) {
	r := &FabricMapListResource{}
	m := FabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestFabricMapListResource_List_BuildError exercises FabricMapListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestFabricMapListResource_List_BuildError(t *testing.T) {
	r := &FabricMapListResource{client: newMalformedBaseURLClient(t)}
	m := FabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestFabricMapListResource_List_SendError exercises FabricMapListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestFabricMapListResource_List_SendError(t *testing.T) {
	r := &FabricMapListResource{client: newTransportErrorClient(t)}
	m := FabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestFabricMapListResource_List_InvalidJSON exercises FabricMapListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestFabricMapListResource_List_InvalidJSON(t *testing.T) {
	r := &FabricMapListResource{client: newMockClientStatus(t, 200, "{{")}
	m := FabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
