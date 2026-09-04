package provider

import (
	"context"
	"testing"
)

// TestEnhancedSlicingListResource_List_Happy exercises EnhancedSlicingListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestEnhancedSlicingListResource_List_Happy(t *testing.T) {
	r := &EnhancedSlicingListResource{client: newMockClientStatus(t, 200, "{\"enhancedSlicings\":[]}")}
	m := EnhancedSlicingListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestEnhancedSlicingListResource_List_NilClient exercises EnhancedSlicingListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEnhancedSlicingListResource_List_NilClient(t *testing.T) {
	r := &EnhancedSlicingListResource{}
	m := EnhancedSlicingListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestEnhancedSlicingListResource_List_BuildError exercises EnhancedSlicingListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestEnhancedSlicingListResource_List_BuildError(t *testing.T) {
	r := &EnhancedSlicingListResource{client: newMalformedBaseURLClient(t)}
	m := EnhancedSlicingListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestEnhancedSlicingListResource_List_SendError exercises EnhancedSlicingListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestEnhancedSlicingListResource_List_SendError(t *testing.T) {
	r := &EnhancedSlicingListResource{client: newTransportErrorClient(t)}
	m := EnhancedSlicingListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestEnhancedSlicingListResource_List_InvalidJSON exercises EnhancedSlicingListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestEnhancedSlicingListResource_List_InvalidJSON(t *testing.T) {
	r := &EnhancedSlicingListResource{client: newMockClientStatus(t, 200, "{{")}
	m := EnhancedSlicingListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
