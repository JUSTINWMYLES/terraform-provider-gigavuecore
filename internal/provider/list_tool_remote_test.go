package provider

import (
	"context"
	"testing"
)

// TestToolListResource_List_Happy exercises ToolListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestToolListResource_List_Happy(t *testing.T) {
	r := &ToolListResource{client: newMockClientStatus(t, 200, "{\"inlineTools\":[]}")}
	m := ToolListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestToolListResource_List_NilClient exercises ToolListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolListResource_List_NilClient(t *testing.T) {
	r := &ToolListResource{}
	m := ToolListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestToolListResource_List_BuildError exercises ToolListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestToolListResource_List_BuildError(t *testing.T) {
	r := &ToolListResource{client: newMalformedBaseURLClient(t)}
	m := ToolListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestToolListResource_List_SendError exercises ToolListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestToolListResource_List_SendError(t *testing.T) {
	r := &ToolListResource{client: newTransportErrorClient(t)}
	m := ToolListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestToolListResource_List_InvalidJSON exercises ToolListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestToolListResource_List_InvalidJSON(t *testing.T) {
	r := &ToolListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ToolListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
