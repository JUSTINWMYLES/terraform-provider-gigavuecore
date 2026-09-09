package provider

import (
	"context"
	"testing"
)

// TestGpfcpProfileListResource_List_Happy exercises GpfcpProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGpfcpProfileListResource_List_Happy(t *testing.T) {
	r := &GpfcpProfileListResource{client: newMockClientStatus(t, 200, "{\"gpfcpProfiles\":[]}")}
	m := GpfcpProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGpfcpProfileListResource_List_NilClient exercises GpfcpProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGpfcpProfileListResource_List_NilClient(t *testing.T) {
	r := &GpfcpProfileListResource{}
	m := GpfcpProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGpfcpProfileListResource_List_BuildError exercises GpfcpProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGpfcpProfileListResource_List_BuildError(t *testing.T) {
	r := &GpfcpProfileListResource{client: newMalformedBaseURLClient(t)}
	m := GpfcpProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGpfcpProfileListResource_List_SendError exercises GpfcpProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGpfcpProfileListResource_List_SendError(t *testing.T) {
	r := &GpfcpProfileListResource{client: newTransportErrorClient(t)}
	m := GpfcpProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGpfcpProfileListResource_List_InvalidJSON exercises GpfcpProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGpfcpProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &GpfcpProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GpfcpProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
