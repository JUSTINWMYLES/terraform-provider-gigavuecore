package provider

import (
	"context"
	"testing"
)

// TestFlexInlineListResource_List_Happy exercises FlexInlineListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestFlexInlineListResource_List_Happy(t *testing.T) {
	r := &FlexInlineListResource{client: newMockClientStatus(t, 200, "{\"flexInlineSolutions\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestFlexInlineListResource_List_NilClient exercises FlexInlineListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFlexInlineListResource_List_NilClient(t *testing.T) {
	r := &FlexInlineListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestFlexInlineListResource_List_BuildError exercises FlexInlineListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestFlexInlineListResource_List_BuildError(t *testing.T) {
	r := &FlexInlineListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestFlexInlineListResource_List_SendError exercises FlexInlineListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestFlexInlineListResource_List_SendError(t *testing.T) {
	r := &FlexInlineListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestFlexInlineListResource_List_InvalidJSON exercises FlexInlineListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestFlexInlineListResource_List_InvalidJSON(t *testing.T) {
	r := &FlexInlineListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
