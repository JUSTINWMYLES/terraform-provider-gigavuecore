package provider

import (
	"context"
	"testing"
)

// TestEngineListResource_List_Happy exercises EngineListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestEngineListResource_List_Happy(t *testing.T) {
	r := &EngineListResource{client: newMockClientStatus(t, 200, "{\"gsEngineInterfaces\":[]}")}
	m := EngineListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestEngineListResource_List_NilClient exercises EngineListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEngineListResource_List_NilClient(t *testing.T) {
	r := &EngineListResource{}
	m := EngineListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestEngineListResource_List_BuildError exercises EngineListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestEngineListResource_List_BuildError(t *testing.T) {
	r := &EngineListResource{client: newMalformedBaseURLClient(t)}
	m := EngineListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestEngineListResource_List_SendError exercises EngineListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestEngineListResource_List_SendError(t *testing.T) {
	r := &EngineListResource{client: newTransportErrorClient(t)}
	m := EngineListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestEngineListResource_List_InvalidJSON exercises EngineListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestEngineListResource_List_InvalidJSON(t *testing.T) {
	r := &EngineListResource{client: newMockClientStatus(t, 200, "{{")}
	m := EngineListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
