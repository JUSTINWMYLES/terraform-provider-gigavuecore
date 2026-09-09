package provider

import (
	"context"
	"testing"
)

// TestPortThrottleListResource_List_Happy exercises PortThrottleListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestPortThrottleListResource_List_Happy(t *testing.T) {
	r := &PortThrottleListResource{client: newMockClientStatus(t, 200, "{\"portThrottles\":[]}")}
	m := PortThrottleListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestPortThrottleListResource_List_NilClient exercises PortThrottleListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortThrottleListResource_List_NilClient(t *testing.T) {
	r := &PortThrottleListResource{}
	m := PortThrottleListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestPortThrottleListResource_List_BuildError exercises PortThrottleListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestPortThrottleListResource_List_BuildError(t *testing.T) {
	r := &PortThrottleListResource{client: newMalformedBaseURLClient(t)}
	m := PortThrottleListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPortThrottleListResource_List_SendError exercises PortThrottleListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestPortThrottleListResource_List_SendError(t *testing.T) {
	r := &PortThrottleListResource{client: newTransportErrorClient(t)}
	m := PortThrottleListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPortThrottleListResource_List_InvalidJSON exercises PortThrottleListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestPortThrottleListResource_List_InvalidJSON(t *testing.T) {
	r := &PortThrottleListResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortThrottleListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
