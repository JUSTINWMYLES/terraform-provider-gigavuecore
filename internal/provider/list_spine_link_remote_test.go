package provider

import (
	"context"
	"testing"
)

// TestSpineLinkListResource_List_Happy exercises SpineLinkListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestSpineLinkListResource_List_Happy(t *testing.T) {
	r := &SpineLinkListResource{client: newMockClientStatus(t, 200, "{\"spineLinks\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestSpineLinkListResource_List_NilClient exercises SpineLinkListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSpineLinkListResource_List_NilClient(t *testing.T) {
	r := &SpineLinkListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestSpineLinkListResource_List_BuildError exercises SpineLinkListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestSpineLinkListResource_List_BuildError(t *testing.T) {
	r := &SpineLinkListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSpineLinkListResource_List_SendError exercises SpineLinkListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestSpineLinkListResource_List_SendError(t *testing.T) {
	r := &SpineLinkListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSpineLinkListResource_List_InvalidJSON exercises SpineLinkListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestSpineLinkListResource_List_InvalidJSON(t *testing.T) {
	r := &SpineLinkListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
