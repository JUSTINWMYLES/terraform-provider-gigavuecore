package provider

import (
	"context"
	"testing"
)

// TestGetAllPtpClockStatesListResource_List_Happy exercises GetAllPtpClockStatesListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllPtpClockStatesListResource_List_Happy(t *testing.T) {
	r := &GetAllPtpClockStatesListResource{client: newMockClientStatus(t, 200, "{\"clockStates\":[]}")}
	m := GetAllPtpClockStatesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllPtpClockStatesListResource_List_NilClient exercises GetAllPtpClockStatesListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllPtpClockStatesListResource_List_NilClient(t *testing.T) {
	r := &GetAllPtpClockStatesListResource{}
	m := GetAllPtpClockStatesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllPtpClockStatesListResource_List_BuildError exercises GetAllPtpClockStatesListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllPtpClockStatesListResource_List_BuildError(t *testing.T) {
	r := &GetAllPtpClockStatesListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllPtpClockStatesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllPtpClockStatesListResource_List_SendError exercises GetAllPtpClockStatesListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllPtpClockStatesListResource_List_SendError(t *testing.T) {
	r := &GetAllPtpClockStatesListResource{client: newTransportErrorClient(t)}
	m := GetAllPtpClockStatesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllPtpClockStatesListResource_List_InvalidJSON exercises GetAllPtpClockStatesListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllPtpClockStatesListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllPtpClockStatesListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllPtpClockStatesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
