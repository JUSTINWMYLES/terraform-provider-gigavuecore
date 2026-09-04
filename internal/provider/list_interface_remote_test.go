package provider

import (
	"context"
	"testing"
)

// TestInterfaceListResource_List_Happy exercises InterfaceListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestInterfaceListResource_List_Happy(t *testing.T) {
	r := &InterfaceListResource{client: newMockClientStatus(t, 200, "{\"ipInterfaces\":[]}")}
	m := InterfaceListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestInterfaceListResource_List_NilClient exercises InterfaceListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInterfaceListResource_List_NilClient(t *testing.T) {
	r := &InterfaceListResource{}
	m := InterfaceListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestInterfaceListResource_List_BuildError exercises InterfaceListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestInterfaceListResource_List_BuildError(t *testing.T) {
	r := &InterfaceListResource{client: newMalformedBaseURLClient(t)}
	m := InterfaceListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestInterfaceListResource_List_SendError exercises InterfaceListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestInterfaceListResource_List_SendError(t *testing.T) {
	r := &InterfaceListResource{client: newTransportErrorClient(t)}
	m := InterfaceListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestInterfaceListResource_List_InvalidJSON exercises InterfaceListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestInterfaceListResource_List_InvalidJSON(t *testing.T) {
	r := &InterfaceListResource{client: newMockClientStatus(t, 200, "{{")}
	m := InterfaceListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
