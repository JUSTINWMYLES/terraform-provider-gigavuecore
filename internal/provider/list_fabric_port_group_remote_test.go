package provider

import (
	"context"
	"testing"
)

// TestFabricPortGroupListResource_List_Happy exercises FabricPortGroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestFabricPortGroupListResource_List_Happy(t *testing.T) {
	r := &FabricPortGroupListResource{client: newMockClientStatus(t, 200, "{\"gigaFabricPortGroups\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestFabricPortGroupListResource_List_NilClient exercises FabricPortGroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFabricPortGroupListResource_List_NilClient(t *testing.T) {
	r := &FabricPortGroupListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestFabricPortGroupListResource_List_BuildError exercises FabricPortGroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestFabricPortGroupListResource_List_BuildError(t *testing.T) {
	r := &FabricPortGroupListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestFabricPortGroupListResource_List_SendError exercises FabricPortGroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestFabricPortGroupListResource_List_SendError(t *testing.T) {
	r := &FabricPortGroupListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestFabricPortGroupListResource_List_InvalidJSON exercises FabricPortGroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestFabricPortGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &FabricPortGroupListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
