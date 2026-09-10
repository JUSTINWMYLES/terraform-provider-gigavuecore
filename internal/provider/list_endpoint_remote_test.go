package provider

import (
	"context"
	"testing"
)

// TestEndpointListResource_List_Happy exercises EndpointListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestEndpointListResource_List_Happy(t *testing.T) {
	r := &EndpointListResource{client: newMockClientStatus(t, 200, "{\"endpoints\":[]}")}
	m := EndpointListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestEndpointListResource_List_NilClient exercises EndpointListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEndpointListResource_List_NilClient(t *testing.T) {
	r := &EndpointListResource{}
	m := EndpointListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestEndpointListResource_List_BuildError exercises EndpointListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestEndpointListResource_List_BuildError(t *testing.T) {
	r := &EndpointListResource{client: newMalformedBaseURLClient(t)}
	m := EndpointListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestEndpointListResource_List_SendError exercises EndpointListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestEndpointListResource_List_SendError(t *testing.T) {
	r := &EndpointListResource{client: newTransportErrorClient(t)}
	m := EndpointListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestEndpointListResource_List_InvalidJSON exercises EndpointListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestEndpointListResource_List_InvalidJSON(t *testing.T) {
	r := &EndpointListResource{client: newMockClientStatus(t, 200, "{{")}
	m := EndpointListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
