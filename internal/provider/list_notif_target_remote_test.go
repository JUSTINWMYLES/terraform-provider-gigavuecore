package provider

import (
	"context"
	"testing"
)

// TestNotifTargetListResource_List_Happy exercises NotifTargetListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestNotifTargetListResource_List_Happy(t *testing.T) {
	r := &NotifTargetListResource{client: newMockClientStatus(t, 200, "{\"notifTargets\":[]}")}
	m := NotifTargetListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestNotifTargetListResource_List_NilClient exercises NotifTargetListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNotifTargetListResource_List_NilClient(t *testing.T) {
	r := &NotifTargetListResource{}
	m := NotifTargetListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestNotifTargetListResource_List_BuildError exercises NotifTargetListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestNotifTargetListResource_List_BuildError(t *testing.T) {
	r := &NotifTargetListResource{client: newMalformedBaseURLClient(t)}
	m := NotifTargetListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestNotifTargetListResource_List_SendError exercises NotifTargetListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestNotifTargetListResource_List_SendError(t *testing.T) {
	r := &NotifTargetListResource{client: newTransportErrorClient(t)}
	m := NotifTargetListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestNotifTargetListResource_List_InvalidJSON exercises NotifTargetListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestNotifTargetListResource_List_InvalidJSON(t *testing.T) {
	r := &NotifTargetListResource{client: newMockClientStatus(t, 200, "{{")}
	m := NotifTargetListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
