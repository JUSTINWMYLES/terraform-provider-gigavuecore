package provider

import (
	"context"
	"testing"
)

// TestGetAllPtpCountersListResource_List_Happy exercises GetAllPtpCountersListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllPtpCountersListResource_List_Happy(t *testing.T) {
	r := &GetAllPtpCountersListResource{client: newMockClientStatus(t, 200, "{\"counters\":[]}")}
	m := GetAllPtpCountersListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllPtpCountersListResource_List_NilClient exercises GetAllPtpCountersListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllPtpCountersListResource_List_NilClient(t *testing.T) {
	r := &GetAllPtpCountersListResource{}
	m := GetAllPtpCountersListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllPtpCountersListResource_List_BuildError exercises GetAllPtpCountersListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllPtpCountersListResource_List_BuildError(t *testing.T) {
	r := &GetAllPtpCountersListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllPtpCountersListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllPtpCountersListResource_List_SendError exercises GetAllPtpCountersListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllPtpCountersListResource_List_SendError(t *testing.T) {
	r := &GetAllPtpCountersListResource{client: newTransportErrorClient(t)}
	m := GetAllPtpCountersListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllPtpCountersListResource_List_InvalidJSON exercises GetAllPtpCountersListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllPtpCountersListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllPtpCountersListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllPtpCountersListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
