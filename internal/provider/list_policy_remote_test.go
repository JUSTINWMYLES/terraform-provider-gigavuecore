package provider

import (
	"context"
	"testing"
)

// TestPolicyListResource_List_Happy exercises PolicyListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestPolicyListResource_List_Happy(t *testing.T) {
	r := &PolicyListResource{client: newMockClientStatus(t, 200, "{\"avPolicies\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestPolicyListResource_List_NilClient exercises PolicyListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPolicyListResource_List_NilClient(t *testing.T) {
	r := &PolicyListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestPolicyListResource_List_BuildError exercises PolicyListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestPolicyListResource_List_BuildError(t *testing.T) {
	r := &PolicyListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPolicyListResource_List_SendError exercises PolicyListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestPolicyListResource_List_SendError(t *testing.T) {
	r := &PolicyListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPolicyListResource_List_InvalidJSON exercises PolicyListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestPolicyListResource_List_InvalidJSON(t *testing.T) {
	r := &PolicyListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
