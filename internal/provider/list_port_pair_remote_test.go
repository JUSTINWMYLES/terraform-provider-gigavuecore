package provider

import (
	"context"
	"testing"
)

// TestPortPairListResource_List_Happy exercises PortPairListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestPortPairListResource_List_Happy(t *testing.T) {
	r := &PortPairListResource{client: newMockClientStatus(t, 200, "{\"portPairs\":[]}")}
	m := PortPairListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestPortPairListResource_List_NilClient exercises PortPairListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortPairListResource_List_NilClient(t *testing.T) {
	r := &PortPairListResource{}
	m := PortPairListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestPortPairListResource_List_BuildError exercises PortPairListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestPortPairListResource_List_BuildError(t *testing.T) {
	r := &PortPairListResource{client: newMalformedBaseURLClient(t)}
	m := PortPairListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPortPairListResource_List_SendError exercises PortPairListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestPortPairListResource_List_SendError(t *testing.T) {
	r := &PortPairListResource{client: newTransportErrorClient(t)}
	m := PortPairListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPortPairListResource_List_InvalidJSON exercises PortPairListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestPortPairListResource_List_InvalidJSON(t *testing.T) {
	r := &PortPairListResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortPairListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
