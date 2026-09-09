package provider

import (
	"context"
	"testing"
)

// TestMobilityListResource_List_Happy exercises MobilityListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestMobilityListResource_List_Happy(t *testing.T) {
	r := &MobilityListResource{client: newMockClientStatus(t, 200, "[]")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestMobilityListResource_List_NilClient exercises MobilityListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMobilityListResource_List_NilClient(t *testing.T) {
	r := &MobilityListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestMobilityListResource_List_BuildError exercises MobilityListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestMobilityListResource_List_BuildError(t *testing.T) {
	r := &MobilityListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMobilityListResource_List_SendError exercises MobilityListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestMobilityListResource_List_SendError(t *testing.T) {
	r := &MobilityListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMobilityListResource_List_InvalidJSON exercises MobilityListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestMobilityListResource_List_InvalidJSON(t *testing.T) {
	r := &MobilityListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
