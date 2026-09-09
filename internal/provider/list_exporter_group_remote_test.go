package provider

import (
	"context"
	"testing"
)

// TestExporterGroupListResource_List_Happy exercises ExporterGroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestExporterGroupListResource_List_Happy(t *testing.T) {
	r := &ExporterGroupListResource{client: newMockClientStatus(t, 200, "{\"appsExporterGroups\":[]}")}
	m := ExporterGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestExporterGroupListResource_List_NilClient exercises ExporterGroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExporterGroupListResource_List_NilClient(t *testing.T) {
	r := &ExporterGroupListResource{}
	m := ExporterGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestExporterGroupListResource_List_BuildError exercises ExporterGroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestExporterGroupListResource_List_BuildError(t *testing.T) {
	r := &ExporterGroupListResource{client: newMalformedBaseURLClient(t)}
	m := ExporterGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestExporterGroupListResource_List_SendError exercises ExporterGroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestExporterGroupListResource_List_SendError(t *testing.T) {
	r := &ExporterGroupListResource{client: newTransportErrorClient(t)}
	m := ExporterGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestExporterGroupListResource_List_InvalidJSON exercises ExporterGroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestExporterGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &ExporterGroupListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ExporterGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
