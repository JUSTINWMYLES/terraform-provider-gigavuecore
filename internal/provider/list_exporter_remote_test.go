package provider

import (
	"context"
	"testing"
)

// TestExporterListResource_List_Happy exercises ExporterListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestExporterListResource_List_Happy(t *testing.T) {
	r := &ExporterListResource{client: newMockClientStatus(t, 200, "{\"appsExporters\":[]}")}
	m := ExporterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestExporterListResource_List_NilClient exercises ExporterListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExporterListResource_List_NilClient(t *testing.T) {
	r := &ExporterListResource{}
	m := ExporterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestExporterListResource_List_BuildError exercises ExporterListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestExporterListResource_List_BuildError(t *testing.T) {
	r := &ExporterListResource{client: newMalformedBaseURLClient(t)}
	m := ExporterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestExporterListResource_List_SendError exercises ExporterListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestExporterListResource_List_SendError(t *testing.T) {
	r := &ExporterListResource{client: newTransportErrorClient(t)}
	m := ExporterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestExporterListResource_List_InvalidJSON exercises ExporterListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestExporterListResource_List_InvalidJSON(t *testing.T) {
	r := &ExporterListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ExporterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
