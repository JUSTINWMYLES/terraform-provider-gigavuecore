package provider

import (
	"context"
	"testing"
)

// TestFilterTemplateListResource_List_Happy exercises FilterTemplateListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestFilterTemplateListResource_List_Happy(t *testing.T) {
	r := &FilterTemplateListResource{client: newMockClientStatus(t, 200, "{\"filterTemplates\":[]}")}
	m := FilterTemplateListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestFilterTemplateListResource_List_NilClient exercises FilterTemplateListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFilterTemplateListResource_List_NilClient(t *testing.T) {
	r := &FilterTemplateListResource{}
	m := FilterTemplateListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestFilterTemplateListResource_List_BuildError exercises FilterTemplateListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestFilterTemplateListResource_List_BuildError(t *testing.T) {
	r := &FilterTemplateListResource{client: newMalformedBaseURLClient(t)}
	m := FilterTemplateListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestFilterTemplateListResource_List_SendError exercises FilterTemplateListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestFilterTemplateListResource_List_SendError(t *testing.T) {
	r := &FilterTemplateListResource{client: newTransportErrorClient(t)}
	m := FilterTemplateListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestFilterTemplateListResource_List_InvalidJSON exercises FilterTemplateListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestFilterTemplateListResource_List_InvalidJSON(t *testing.T) {
	r := &FilterTemplateListResource{client: newMockClientStatus(t, 200, "{{")}
	m := FilterTemplateListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
