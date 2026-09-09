package provider

import (
	"context"
	"testing"
)

// TestToolPortMirrorListResource_List_Happy exercises ToolPortMirrorListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestToolPortMirrorListResource_List_Happy(t *testing.T) {
	r := &ToolPortMirrorListResource{client: newMockClientStatus(t, 200, "{\"toolPortMirrors\":[]}")}
	m := ToolPortMirrorListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestToolPortMirrorListResource_List_NilClient exercises ToolPortMirrorListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolPortMirrorListResource_List_NilClient(t *testing.T) {
	r := &ToolPortMirrorListResource{}
	m := ToolPortMirrorListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestToolPortMirrorListResource_List_BuildError exercises ToolPortMirrorListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestToolPortMirrorListResource_List_BuildError(t *testing.T) {
	r := &ToolPortMirrorListResource{client: newMalformedBaseURLClient(t)}
	m := ToolPortMirrorListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestToolPortMirrorListResource_List_SendError exercises ToolPortMirrorListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestToolPortMirrorListResource_List_SendError(t *testing.T) {
	r := &ToolPortMirrorListResource{client: newTransportErrorClient(t)}
	m := ToolPortMirrorListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestToolPortMirrorListResource_List_InvalidJSON exercises ToolPortMirrorListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestToolPortMirrorListResource_List_InvalidJSON(t *testing.T) {
	r := &ToolPortMirrorListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ToolPortMirrorListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
