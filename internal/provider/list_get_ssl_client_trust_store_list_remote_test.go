package provider

import (
	"context"
	"testing"
)

// TestGetSslClientTrustStoreListListResource_List_Happy exercises GetSslClientTrustStoreListListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetSslClientTrustStoreListListResource_List_Happy(t *testing.T) {
	r := &GetSslClientTrustStoreListListResource{client: newMockClientStatus(t, 200, "{\"clientTrustStores\":[]}")}
	m := GetSslClientTrustStoreListListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetSslClientTrustStoreListListResource_List_NilClient exercises GetSslClientTrustStoreListListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSslClientTrustStoreListListResource_List_NilClient(t *testing.T) {
	r := &GetSslClientTrustStoreListListResource{}
	m := GetSslClientTrustStoreListListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetSslClientTrustStoreListListResource_List_BuildError exercises GetSslClientTrustStoreListListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetSslClientTrustStoreListListResource_List_BuildError(t *testing.T) {
	r := &GetSslClientTrustStoreListListResource{client: newMalformedBaseURLClient(t)}
	m := GetSslClientTrustStoreListListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetSslClientTrustStoreListListResource_List_SendError exercises GetSslClientTrustStoreListListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetSslClientTrustStoreListListResource_List_SendError(t *testing.T) {
	r := &GetSslClientTrustStoreListListResource{client: newTransportErrorClient(t)}
	m := GetSslClientTrustStoreListListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetSslClientTrustStoreListListResource_List_InvalidJSON exercises GetSslClientTrustStoreListListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetSslClientTrustStoreListListResource_List_InvalidJSON(t *testing.T) {
	r := &GetSslClientTrustStoreListListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSslClientTrustStoreListListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
