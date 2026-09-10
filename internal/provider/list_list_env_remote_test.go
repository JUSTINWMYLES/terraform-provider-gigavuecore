package provider

import (
	"context"
	"testing"
)

// TestListEnvListResource_List_Happy exercises ListEnvListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestListEnvListResource_List_Happy(t *testing.T) {
	r := &ListEnvListResource{client: newMockClientStatus(t, 200, "{\"env\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestListEnvListResource_List_NilClient exercises ListEnvListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListEnvListResource_List_NilClient(t *testing.T) {
	r := &ListEnvListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestListEnvListResource_List_BuildError exercises ListEnvListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestListEnvListResource_List_BuildError(t *testing.T) {
	r := &ListEnvListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestListEnvListResource_List_SendError exercises ListEnvListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestListEnvListResource_List_SendError(t *testing.T) {
	r := &ListEnvListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestListEnvListResource_List_InvalidJSON exercises ListEnvListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestListEnvListResource_List_InvalidJSON(t *testing.T) {
	r := &ListEnvListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
