package provider

import (
	"context"
	"testing"
)

// TestGetAllClusterMapsOfAnUserFabricMapListResource_List_Happy exercises GetAllClusterMapsOfAnUserFabricMapListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterMapsOfAnUserFabricMapListResource_List_Happy(t *testing.T) {
	r := &GetAllClusterMapsOfAnUserFabricMapListResource{client: newMockClientStatus(t, 200, "{\"clusterMaps\":[]}")}
	m := GetAllClusterMapsOfAnUserFabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllClusterMapsOfAnUserFabricMapListResource_List_NilClient exercises GetAllClusterMapsOfAnUserFabricMapListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterMapsOfAnUserFabricMapListResource_List_NilClient(t *testing.T) {
	r := &GetAllClusterMapsOfAnUserFabricMapListResource{}
	m := GetAllClusterMapsOfAnUserFabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllClusterMapsOfAnUserFabricMapListResource_List_BuildError exercises GetAllClusterMapsOfAnUserFabricMapListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterMapsOfAnUserFabricMapListResource_List_BuildError(t *testing.T) {
	r := &GetAllClusterMapsOfAnUserFabricMapListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllClusterMapsOfAnUserFabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllClusterMapsOfAnUserFabricMapListResource_List_SendError exercises GetAllClusterMapsOfAnUserFabricMapListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterMapsOfAnUserFabricMapListResource_List_SendError(t *testing.T) {
	r := &GetAllClusterMapsOfAnUserFabricMapListResource{client: newTransportErrorClient(t)}
	m := GetAllClusterMapsOfAnUserFabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllClusterMapsOfAnUserFabricMapListResource_List_InvalidJSON exercises GetAllClusterMapsOfAnUserFabricMapListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterMapsOfAnUserFabricMapListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllClusterMapsOfAnUserFabricMapListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllClusterMapsOfAnUserFabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
