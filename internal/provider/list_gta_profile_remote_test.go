package provider

import (
	"context"
	"testing"
)

// TestGtaProfileListResource_List_Happy exercises GtaProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGtaProfileListResource_List_Happy(t *testing.T) {
	r := &GtaProfileListResource{client: newMockClientStatus(t, 200, "{\"gtaProfiles\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestGtaProfileListResource_List_NilClient exercises GtaProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGtaProfileListResource_List_NilClient(t *testing.T) {
	r := &GtaProfileListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGtaProfileListResource_List_BuildError exercises GtaProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGtaProfileListResource_List_BuildError(t *testing.T) {
	r := &GtaProfileListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGtaProfileListResource_List_SendError exercises GtaProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGtaProfileListResource_List_SendError(t *testing.T) {
	r := &GtaProfileListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGtaProfileListResource_List_InvalidJSON exercises GtaProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGtaProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &GtaProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
