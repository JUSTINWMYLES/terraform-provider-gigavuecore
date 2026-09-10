package provider

import (
	"context"
	"testing"
)

// TestSysdumpListResource_List_Happy exercises SysdumpListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestSysdumpListResource_List_Happy(t *testing.T) {
	r := &SysdumpListResource{client: newMockClientStatus(t, 200, "{\"sysdumpFiles\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestSysdumpListResource_List_NilClient exercises SysdumpListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSysdumpListResource_List_NilClient(t *testing.T) {
	r := &SysdumpListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestSysdumpListResource_List_BuildError exercises SysdumpListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestSysdumpListResource_List_BuildError(t *testing.T) {
	r := &SysdumpListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSysdumpListResource_List_SendError exercises SysdumpListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestSysdumpListResource_List_SendError(t *testing.T) {
	r := &SysdumpListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSysdumpListResource_List_InvalidJSON exercises SysdumpListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestSysdumpListResource_List_InvalidJSON(t *testing.T) {
	r := &SysdumpListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
