package provider

import (
	"context"
	"testing"
)

// TestHbPacketListResource_List_Happy exercises HbPacketListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestHbPacketListResource_List_Happy(t *testing.T) {
	r := &HbPacketListResource{client: newMockClientStatus(t, 200, "{\"hbCustomPackets\":[]}")}
	m := HbPacketListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestHbPacketListResource_List_NilClient exercises HbPacketListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHbPacketListResource_List_NilClient(t *testing.T) {
	r := &HbPacketListResource{}
	m := HbPacketListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestHbPacketListResource_List_BuildError exercises HbPacketListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestHbPacketListResource_List_BuildError(t *testing.T) {
	r := &HbPacketListResource{client: newMalformedBaseURLClient(t)}
	m := HbPacketListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestHbPacketListResource_List_SendError exercises HbPacketListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestHbPacketListResource_List_SendError(t *testing.T) {
	r := &HbPacketListResource{client: newTransportErrorClient(t)}
	m := HbPacketListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestHbPacketListResource_List_InvalidJSON exercises HbPacketListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestHbPacketListResource_List_InvalidJSON(t *testing.T) {
	r := &HbPacketListResource{client: newMockClientStatus(t, 200, "{{")}
	m := HbPacketListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
