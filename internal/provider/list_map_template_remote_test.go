package provider

import (
	"context"
	"testing"
)

// TestMapTemplateListResource_List_Happy exercises MapTemplateListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestMapTemplateListResource_List_Happy(t *testing.T) {
	r := &MapTemplateListResource{client: newMockClientStatus(t, 200, "{\"mapTemplates\":[]}")}
	m := MapTemplateListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestMapTemplateListResource_List_NilClient exercises MapTemplateListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapTemplateListResource_List_NilClient(t *testing.T) {
	r := &MapTemplateListResource{}
	m := MapTemplateListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestMapTemplateListResource_List_BuildError exercises MapTemplateListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestMapTemplateListResource_List_BuildError(t *testing.T) {
	r := &MapTemplateListResource{client: newMalformedBaseURLClient(t)}
	m := MapTemplateListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMapTemplateListResource_List_SendError exercises MapTemplateListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestMapTemplateListResource_List_SendError(t *testing.T) {
	r := &MapTemplateListResource{client: newTransportErrorClient(t)}
	m := MapTemplateListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMapTemplateListResource_List_InvalidJSON exercises MapTemplateListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestMapTemplateListResource_List_InvalidJSON(t *testing.T) {
	r := &MapTemplateListResource{client: newMockClientStatus(t, 200, "{{")}
	m := MapTemplateListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
