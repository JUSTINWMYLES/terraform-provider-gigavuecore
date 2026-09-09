package provider

import (
	"context"
	"testing"
)

// TestGetAllPtpParentListResource_List_Happy exercises GetAllPtpParentListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllPtpParentListResource_List_Happy(t *testing.T) {
	r := &GetAllPtpParentListResource{client: newMockClientStatus(t, 200, "{\"parents\":[]}")}
	m := GetAllPtpParentListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllPtpParentListResource_List_NilClient exercises GetAllPtpParentListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllPtpParentListResource_List_NilClient(t *testing.T) {
	r := &GetAllPtpParentListResource{}
	m := GetAllPtpParentListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllPtpParentListResource_List_BuildError exercises GetAllPtpParentListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllPtpParentListResource_List_BuildError(t *testing.T) {
	r := &GetAllPtpParentListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllPtpParentListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllPtpParentListResource_List_SendError exercises GetAllPtpParentListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllPtpParentListResource_List_SendError(t *testing.T) {
	r := &GetAllPtpParentListResource{client: newTransportErrorClient(t)}
	m := GetAllPtpParentListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllPtpParentListResource_List_InvalidJSON exercises GetAllPtpParentListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllPtpParentListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllPtpParentListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllPtpParentListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
