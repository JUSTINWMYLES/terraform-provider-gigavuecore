package provider

import (
	"context"
	"testing"
)

// TestGetAllPtpPortsCountersListResource_List_Happy exercises GetAllPtpPortsCountersListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllPtpPortsCountersListResource_List_Happy(t *testing.T) {
	r := &GetAllPtpPortsCountersListResource{client: newMockClientStatus(t, 200, "{\"ports\":[]}")}
	m := GetAllPtpPortsCountersListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllPtpPortsCountersListResource_List_NilClient exercises GetAllPtpPortsCountersListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllPtpPortsCountersListResource_List_NilClient(t *testing.T) {
	r := &GetAllPtpPortsCountersListResource{}
	m := GetAllPtpPortsCountersListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllPtpPortsCountersListResource_List_BuildError exercises GetAllPtpPortsCountersListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllPtpPortsCountersListResource_List_BuildError(t *testing.T) {
	r := &GetAllPtpPortsCountersListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllPtpPortsCountersListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllPtpPortsCountersListResource_List_SendError exercises GetAllPtpPortsCountersListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllPtpPortsCountersListResource_List_SendError(t *testing.T) {
	r := &GetAllPtpPortsCountersListResource{client: newTransportErrorClient(t)}
	m := GetAllPtpPortsCountersListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllPtpPortsCountersListResource_List_InvalidJSON exercises GetAllPtpPortsCountersListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllPtpPortsCountersListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllPtpPortsCountersListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllPtpPortsCountersListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
