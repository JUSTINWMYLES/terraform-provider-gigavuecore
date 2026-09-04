package provider

import (
	"context"
	"testing"
)

// TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource_List_Happy exercises GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource_List_Happy(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource{client: newMockClientStatus(t, 200, "{\"upgradeTaskStatus\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource_List_NilClient exercises GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource_List_NilClient(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource_List_BuildError exercises GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource_List_BuildError(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource_List_SendError exercises GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource_List_SendError(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource_List_InvalidJSON exercises GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
