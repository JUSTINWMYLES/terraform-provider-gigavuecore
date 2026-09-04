package provider

import (
	"context"
	"testing"
)

// TestLoadAllAlarmsListResource_List_Happy exercises LoadAllAlarmsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllAlarmsListResource_List_Happy(t *testing.T) {
	r := &LoadAllAlarmsListResource{client: newMockClientStatus(t, 200, "{\"alarms\":[]}")}
	m := LoadAllAlarmsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadAllAlarmsListResource_List_NilClient exercises LoadAllAlarmsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllAlarmsListResource_List_NilClient(t *testing.T) {
	r := &LoadAllAlarmsListResource{}
	m := LoadAllAlarmsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadAllAlarmsListResource_List_BuildError exercises LoadAllAlarmsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllAlarmsListResource_List_BuildError(t *testing.T) {
	r := &LoadAllAlarmsListResource{client: newMalformedBaseURLClient(t)}
	m := LoadAllAlarmsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllAlarmsListResource_List_SendError exercises LoadAllAlarmsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllAlarmsListResource_List_SendError(t *testing.T) {
	r := &LoadAllAlarmsListResource{client: newTransportErrorClient(t)}
	m := LoadAllAlarmsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllAlarmsListResource_List_InvalidJSON exercises LoadAllAlarmsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllAlarmsListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadAllAlarmsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllAlarmsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
