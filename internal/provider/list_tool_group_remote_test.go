package provider

import (
	"context"
	"testing"
)

// TestToolGroupListResource_List_Happy exercises ToolGroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestToolGroupListResource_List_Happy(t *testing.T) {
	r := &ToolGroupListResource{client: newMockClientStatus(t, 200, "{\"inlineToolGroups\":[]}")}
	m := ToolGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestToolGroupListResource_List_NilClient exercises ToolGroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolGroupListResource_List_NilClient(t *testing.T) {
	r := &ToolGroupListResource{}
	m := ToolGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestToolGroupListResource_List_BuildError exercises ToolGroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestToolGroupListResource_List_BuildError(t *testing.T) {
	r := &ToolGroupListResource{client: newMalformedBaseURLClient(t)}
	m := ToolGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestToolGroupListResource_List_SendError exercises ToolGroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestToolGroupListResource_List_SendError(t *testing.T) {
	r := &ToolGroupListResource{client: newTransportErrorClient(t)}
	m := ToolGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestToolGroupListResource_List_InvalidJSON exercises ToolGroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestToolGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &ToolGroupListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ToolGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
