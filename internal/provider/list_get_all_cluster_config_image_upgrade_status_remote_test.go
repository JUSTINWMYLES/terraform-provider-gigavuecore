package provider

import (
	"context"
	"testing"
)

// TestGetAllClusterConfigImageUpgradeStatusListResource_List_Happy exercises GetAllClusterConfigImageUpgradeStatusListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterConfigImageUpgradeStatusListResource_List_Happy(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusListResource{client: newMockClientStatus(t, 200, "{\"clustersStatus\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestGetAllClusterConfigImageUpgradeStatusListResource_List_NilClient exercises GetAllClusterConfigImageUpgradeStatusListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterConfigImageUpgradeStatusListResource_List_NilClient(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllClusterConfigImageUpgradeStatusListResource_List_BuildError exercises GetAllClusterConfigImageUpgradeStatusListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterConfigImageUpgradeStatusListResource_List_BuildError(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllClusterConfigImageUpgradeStatusListResource_List_SendError exercises GetAllClusterConfigImageUpgradeStatusListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterConfigImageUpgradeStatusListResource_List_SendError(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllClusterConfigImageUpgradeStatusListResource_List_InvalidJSON exercises GetAllClusterConfigImageUpgradeStatusListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterConfigImageUpgradeStatusListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
