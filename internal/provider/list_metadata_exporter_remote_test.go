package provider

import (
	"context"
	"testing"
)

// TestMetadataExporterListResource_List_Happy exercises MetadataExporterListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestMetadataExporterListResource_List_Happy(t *testing.T) {
	r := &MetadataExporterListResource{client: newMockClientStatus(t, 200, "{\"metadataExporters\":[]}")}
	m := MetadataExporterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestMetadataExporterListResource_List_NilClient exercises MetadataExporterListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMetadataExporterListResource_List_NilClient(t *testing.T) {
	r := &MetadataExporterListResource{}
	m := MetadataExporterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestMetadataExporterListResource_List_BuildError exercises MetadataExporterListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestMetadataExporterListResource_List_BuildError(t *testing.T) {
	r := &MetadataExporterListResource{client: newMalformedBaseURLClient(t)}
	m := MetadataExporterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMetadataExporterListResource_List_SendError exercises MetadataExporterListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestMetadataExporterListResource_List_SendError(t *testing.T) {
	r := &MetadataExporterListResource{client: newTransportErrorClient(t)}
	m := MetadataExporterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMetadataExporterListResource_List_InvalidJSON exercises MetadataExporterListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestMetadataExporterListResource_List_InvalidJSON(t *testing.T) {
	r := &MetadataExporterListResource{client: newMockClientStatus(t, 200, "{{")}
	m := MetadataExporterListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
