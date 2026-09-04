package provider

import (
	"context"
	"testing"
)

// TestIbPathwayListResource_List_Happy exercises IbPathwayListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestIbPathwayListResource_List_Happy(t *testing.T) {
	r := &IbPathwayListResource{client: newMockClientStatus(t, 200, "{\"ibPathways\":[]}")}
	m := IbPathwayListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestIbPathwayListResource_List_NilClient exercises IbPathwayListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIbPathwayListResource_List_NilClient(t *testing.T) {
	r := &IbPathwayListResource{}
	m := IbPathwayListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestIbPathwayListResource_List_BuildError exercises IbPathwayListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestIbPathwayListResource_List_BuildError(t *testing.T) {
	r := &IbPathwayListResource{client: newMalformedBaseURLClient(t)}
	m := IbPathwayListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestIbPathwayListResource_List_SendError exercises IbPathwayListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestIbPathwayListResource_List_SendError(t *testing.T) {
	r := &IbPathwayListResource{client: newTransportErrorClient(t)}
	m := IbPathwayListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestIbPathwayListResource_List_InvalidJSON exercises IbPathwayListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestIbPathwayListResource_List_InvalidJSON(t *testing.T) {
	r := &IbPathwayListResource{client: newMockClientStatus(t, 200, "{{")}
	m := IbPathwayListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
