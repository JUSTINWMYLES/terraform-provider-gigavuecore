package provider

import (
	"context"
	"testing"
)

// TestCacheListResource_List_Happy exercises CacheListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestCacheListResource_List_Happy(t *testing.T) {
	r := &CacheListResource{client: newMockClientStatus(t, 200, "{\"metadataCaches\":[]}")}
	m := CacheListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestCacheListResource_List_NilClient exercises CacheListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCacheListResource_List_NilClient(t *testing.T) {
	r := &CacheListResource{}
	m := CacheListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestCacheListResource_List_BuildError exercises CacheListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestCacheListResource_List_BuildError(t *testing.T) {
	r := &CacheListResource{client: newMalformedBaseURLClient(t)}
	m := CacheListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestCacheListResource_List_SendError exercises CacheListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestCacheListResource_List_SendError(t *testing.T) {
	r := &CacheListResource{client: newTransportErrorClient(t)}
	m := CacheListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestCacheListResource_List_InvalidJSON exercises CacheListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestCacheListResource_List_InvalidJSON(t *testing.T) {
	r := &CacheListResource{client: newMockClientStatus(t, 200, "{{")}
	m := CacheListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
