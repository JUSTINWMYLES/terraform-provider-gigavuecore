package provider

import (
	"context"
	"testing"
)

// TestLoadAllFmNotificationTargetConfigListResource_List_Happy exercises LoadAllFmNotificationTargetConfigListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllFmNotificationTargetConfigListResource_List_Happy(t *testing.T) {
	r := &LoadAllFmNotificationTargetConfigListResource{client: newMockClientStatus(t, 200, "{\"fmNotificationTargetConfigs\":[]}")}
	m := LoadAllFmNotificationTargetConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadAllFmNotificationTargetConfigListResource_List_NilClient exercises LoadAllFmNotificationTargetConfigListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllFmNotificationTargetConfigListResource_List_NilClient(t *testing.T) {
	r := &LoadAllFmNotificationTargetConfigListResource{}
	m := LoadAllFmNotificationTargetConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadAllFmNotificationTargetConfigListResource_List_BuildError exercises LoadAllFmNotificationTargetConfigListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllFmNotificationTargetConfigListResource_List_BuildError(t *testing.T) {
	r := &LoadAllFmNotificationTargetConfigListResource{client: newMalformedBaseURLClient(t)}
	m := LoadAllFmNotificationTargetConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllFmNotificationTargetConfigListResource_List_SendError exercises LoadAllFmNotificationTargetConfigListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllFmNotificationTargetConfigListResource_List_SendError(t *testing.T) {
	r := &LoadAllFmNotificationTargetConfigListResource{client: newTransportErrorClient(t)}
	m := LoadAllFmNotificationTargetConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllFmNotificationTargetConfigListResource_List_InvalidJSON exercises LoadAllFmNotificationTargetConfigListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllFmNotificationTargetConfigListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadAllFmNotificationTargetConfigListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllFmNotificationTargetConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
