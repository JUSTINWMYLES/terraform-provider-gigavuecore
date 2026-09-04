package provider

import (
	"context"
	"testing"
)

// TestGetAllInternalFabricMapsListResource_List_Happy exercises GetAllInternalFabricMapsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllInternalFabricMapsListResource_List_Happy(t *testing.T) {
	r := &GetAllInternalFabricMapsListResource{client: newMockClientStatus(t, 200, "{\"fabricMaps\":[]}")}
	m := GetAllInternalFabricMapsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllInternalFabricMapsListResource_List_NilClient exercises GetAllInternalFabricMapsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllInternalFabricMapsListResource_List_NilClient(t *testing.T) {
	r := &GetAllInternalFabricMapsListResource{}
	m := GetAllInternalFabricMapsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllInternalFabricMapsListResource_List_BuildError exercises GetAllInternalFabricMapsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllInternalFabricMapsListResource_List_BuildError(t *testing.T) {
	r := &GetAllInternalFabricMapsListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllInternalFabricMapsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllInternalFabricMapsListResource_List_SendError exercises GetAllInternalFabricMapsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllInternalFabricMapsListResource_List_SendError(t *testing.T) {
	r := &GetAllInternalFabricMapsListResource{client: newTransportErrorClient(t)}
	m := GetAllInternalFabricMapsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllInternalFabricMapsListResource_List_InvalidJSON exercises GetAllInternalFabricMapsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllInternalFabricMapsListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllInternalFabricMapsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllInternalFabricMapsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
