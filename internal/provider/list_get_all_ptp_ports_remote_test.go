package provider

import (
	"context"
	"testing"
)

// TestGetAllPtpPortsListResource_List_Happy exercises GetAllPtpPortsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllPtpPortsListResource_List_Happy(t *testing.T) {
	r := &GetAllPtpPortsListResource{client: newMockClientStatus(t, 200, "{\"portStates\":[]}")}
	m := GetAllPtpPortsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllPtpPortsListResource_List_NilClient exercises GetAllPtpPortsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllPtpPortsListResource_List_NilClient(t *testing.T) {
	r := &GetAllPtpPortsListResource{}
	m := GetAllPtpPortsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllPtpPortsListResource_List_BuildError exercises GetAllPtpPortsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllPtpPortsListResource_List_BuildError(t *testing.T) {
	r := &GetAllPtpPortsListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllPtpPortsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllPtpPortsListResource_List_SendError exercises GetAllPtpPortsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllPtpPortsListResource_List_SendError(t *testing.T) {
	r := &GetAllPtpPortsListResource{client: newTransportErrorClient(t)}
	m := GetAllPtpPortsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllPtpPortsListResource_List_InvalidJSON exercises GetAllPtpPortsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllPtpPortsListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllPtpPortsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllPtpPortsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
