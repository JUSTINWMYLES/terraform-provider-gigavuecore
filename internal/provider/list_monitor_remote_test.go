package provider

import (
	"context"
	"testing"
)

// TestMonitorListResource_List_Happy exercises MonitorListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestMonitorListResource_List_Happy(t *testing.T) {
	r := &MonitorListResource{client: newMockClientStatus(t, 200, "{\"nfMonitors\":[]}")}
	m := MonitorListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestMonitorListResource_List_NilClient exercises MonitorListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMonitorListResource_List_NilClient(t *testing.T) {
	r := &MonitorListResource{}
	m := MonitorListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestMonitorListResource_List_BuildError exercises MonitorListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestMonitorListResource_List_BuildError(t *testing.T) {
	r := &MonitorListResource{client: newMalformedBaseURLClient(t)}
	m := MonitorListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMonitorListResource_List_SendError exercises MonitorListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestMonitorListResource_List_SendError(t *testing.T) {
	r := &MonitorListResource{client: newTransportErrorClient(t)}
	m := MonitorListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMonitorListResource_List_InvalidJSON exercises MonitorListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestMonitorListResource_List_InvalidJSON(t *testing.T) {
	r := &MonitorListResource{client: newMockClientStatus(t, 200, "{{")}
	m := MonitorListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
