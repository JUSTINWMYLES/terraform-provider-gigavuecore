package provider

import (
	"context"
	"testing"
)

// TestGetAllTimePropertiesListResource_List_Happy exercises GetAllTimePropertiesListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllTimePropertiesListResource_List_Happy(t *testing.T) {
	r := &GetAllTimePropertiesListResource{client: newMockClientStatus(t, 200, "{\"timeProperties\":[]}")}
	m := GetAllTimePropertiesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllTimePropertiesListResource_List_NilClient exercises GetAllTimePropertiesListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllTimePropertiesListResource_List_NilClient(t *testing.T) {
	r := &GetAllTimePropertiesListResource{}
	m := GetAllTimePropertiesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllTimePropertiesListResource_List_BuildError exercises GetAllTimePropertiesListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllTimePropertiesListResource_List_BuildError(t *testing.T) {
	r := &GetAllTimePropertiesListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllTimePropertiesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllTimePropertiesListResource_List_SendError exercises GetAllTimePropertiesListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllTimePropertiesListResource_List_SendError(t *testing.T) {
	r := &GetAllTimePropertiesListResource{client: newTransportErrorClient(t)}
	m := GetAllTimePropertiesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllTimePropertiesListResource_List_InvalidJSON exercises GetAllTimePropertiesListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllTimePropertiesListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllTimePropertiesListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllTimePropertiesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
