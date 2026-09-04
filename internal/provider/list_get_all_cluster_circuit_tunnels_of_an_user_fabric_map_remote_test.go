package provider

import (
	"context"
	"testing"
)

// TestGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource_List_Happy exercises GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource_List_Happy(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource{client: newMockClientStatus(t, 200, "{\"clusterMaps\":[]}")}
	m := GetAllClusterCircuitTunnelsOfAnUserFabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource_List_NilClient exercises GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource_List_NilClient(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource{}
	m := GetAllClusterCircuitTunnelsOfAnUserFabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource_List_BuildError exercises GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource_List_BuildError(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllClusterCircuitTunnelsOfAnUserFabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource_List_SendError exercises GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource_List_SendError(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource{client: newTransportErrorClient(t)}
	m := GetAllClusterCircuitTunnelsOfAnUserFabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource_List_InvalidJSON exercises GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllClusterCircuitTunnelsOfAnUserFabricMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
