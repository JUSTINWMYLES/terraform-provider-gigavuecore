package provider

import (
	"context"
	"testing"
)

// TestSolutionListResource_List_Happy exercises SolutionListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestSolutionListResource_List_Happy(t *testing.T) {
	r := &SolutionListResource{client: newMockClientStatus(t, 200, "{\"appVzbilitySolutions\":[]}")}
	m := SolutionListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestSolutionListResource_List_NilClient exercises SolutionListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSolutionListResource_List_NilClient(t *testing.T) {
	r := &SolutionListResource{}
	m := SolutionListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestSolutionListResource_List_BuildError exercises SolutionListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestSolutionListResource_List_BuildError(t *testing.T) {
	r := &SolutionListResource{client: newMalformedBaseURLClient(t)}
	m := SolutionListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSolutionListResource_List_SendError exercises SolutionListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestSolutionListResource_List_SendError(t *testing.T) {
	r := &SolutionListResource{client: newTransportErrorClient(t)}
	m := SolutionListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSolutionListResource_List_InvalidJSON exercises SolutionListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestSolutionListResource_List_InvalidJSON(t *testing.T) {
	r := &SolutionListResource{client: newMockClientStatus(t, 200, "{{")}
	m := SolutionListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
