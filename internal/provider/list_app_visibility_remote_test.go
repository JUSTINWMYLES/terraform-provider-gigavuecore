package provider

import (
	"context"
	"testing"
)

// TestAppVisibilityListResource_List_Happy exercises AppVisibilityListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestAppVisibilityListResource_List_Happy(t *testing.T) {
	r := &AppVisibilityListResource{client: newMockClientStatus(t, 200, "{\"appsVisibilitySolutions\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestAppVisibilityListResource_List_NilClient exercises AppVisibilityListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAppVisibilityListResource_List_NilClient(t *testing.T) {
	r := &AppVisibilityListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestAppVisibilityListResource_List_BuildError exercises AppVisibilityListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestAppVisibilityListResource_List_BuildError(t *testing.T) {
	r := &AppVisibilityListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestAppVisibilityListResource_List_SendError exercises AppVisibilityListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestAppVisibilityListResource_List_SendError(t *testing.T) {
	r := &AppVisibilityListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestAppVisibilityListResource_List_InvalidJSON exercises AppVisibilityListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestAppVisibilityListResource_List_InvalidJSON(t *testing.T) {
	r := &AppVisibilityListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
