package provider

import (
	"context"
	"testing"
)

// TestGetAllAppFilterRscListResource_List_Happy exercises GetAllAppFilterRscListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllAppFilterRscListResource_List_Happy(t *testing.T) {
	r := &GetAllAppFilterRscListResource{client: newMockClientStatus(t, 200, "{\"appFilterRscs\":[]}")}
	m := GetAllAppFilterRscListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAllAppFilterRscListResource_List_NilClient exercises GetAllAppFilterRscListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllAppFilterRscListResource_List_NilClient(t *testing.T) {
	r := &GetAllAppFilterRscListResource{}
	m := GetAllAppFilterRscListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAllAppFilterRscListResource_List_BuildError exercises GetAllAppFilterRscListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllAppFilterRscListResource_List_BuildError(t *testing.T) {
	r := &GetAllAppFilterRscListResource{client: newMalformedBaseURLClient(t)}
	m := GetAllAppFilterRscListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllAppFilterRscListResource_List_SendError exercises GetAllAppFilterRscListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllAppFilterRscListResource_List_SendError(t *testing.T) {
	r := &GetAllAppFilterRscListResource{client: newTransportErrorClient(t)}
	m := GetAllAppFilterRscListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAllAppFilterRscListResource_List_InvalidJSON exercises GetAllAppFilterRscListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllAppFilterRscListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAllAppFilterRscListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllAppFilterRscListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
