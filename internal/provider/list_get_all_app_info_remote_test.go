package provider

import (
	"context"
	"testing"
)

// TestGetAllAppInfoListResource_List_Happy exercises GetAllAppInfoListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllAppInfoListResource_List_Happy(t *testing.T) {
	r := &GetAllAppInfoListResource{client: newMockClientStatus(t, 200, "{\"env\":[]}")}
	m := GetAllAppInfoListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllAppInfoListResource_List_NilClient exercises GetAllAppInfoListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllAppInfoListResource_List_NilClient(t *testing.T) {
	r := &GetAllAppInfoListResource{}
	m := GetAllAppInfoListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllAppInfoListResource_List_BuildError exercises GetAllAppInfoListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllAppInfoListResource_List_BuildError(t *testing.T) {
	r := &GetAllAppInfoListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllAppInfoListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllAppInfoListResource_List_SendError exercises GetAllAppInfoListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllAppInfoListResource_List_SendError(t *testing.T) {
	r := &GetAllAppInfoListResource{client: newTransportErrorClient(t)}
	m := GetAllAppInfoListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllAppInfoListResource_List_InvalidJSON exercises GetAllAppInfoListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllAppInfoListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllAppInfoListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllAppInfoListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
