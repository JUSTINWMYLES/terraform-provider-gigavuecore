package provider

import (
	"context"
	"testing"
)

// TestSerialToolGroupListResource_List_Happy exercises SerialToolGroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestSerialToolGroupListResource_List_Happy(t *testing.T) {
	r := &SerialToolGroupListResource{client: newMockClientStatus(t, 200, "{\"inlineSerialToolGroups\":[]}")}
	m := SerialToolGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestSerialToolGroupListResource_List_NilClient exercises SerialToolGroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSerialToolGroupListResource_List_NilClient(t *testing.T) {
	r := &SerialToolGroupListResource{}
	m := SerialToolGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestSerialToolGroupListResource_List_BuildError exercises SerialToolGroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestSerialToolGroupListResource_List_BuildError(t *testing.T) {
	r := &SerialToolGroupListResource{client: newMalformedBaseURLClient(t)}
	m := SerialToolGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSerialToolGroupListResource_List_SendError exercises SerialToolGroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestSerialToolGroupListResource_List_SendError(t *testing.T) {
	r := &SerialToolGroupListResource{client: newTransportErrorClient(t)}
	m := SerialToolGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSerialToolGroupListResource_List_InvalidJSON exercises SerialToolGroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestSerialToolGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &SerialToolGroupListResource{client: newMockClientStatus(t, 200, "{{")}
	m := SerialToolGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
