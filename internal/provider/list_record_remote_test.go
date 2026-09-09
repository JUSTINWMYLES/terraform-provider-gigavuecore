package provider

import (
	"context"
	"testing"
)

// TestRecordListResource_List_Happy exercises RecordListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestRecordListResource_List_Happy(t *testing.T) {
	r := &RecordListResource{client: newMockClientStatus(t, 200, "{\"nfRecords\":[]}")}
	m := RecordListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestRecordListResource_List_NilClient exercises RecordListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRecordListResource_List_NilClient(t *testing.T) {
	r := &RecordListResource{}
	m := RecordListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestRecordListResource_List_BuildError exercises RecordListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestRecordListResource_List_BuildError(t *testing.T) {
	r := &RecordListResource{client: newMalformedBaseURLClient(t)}
	m := RecordListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestRecordListResource_List_SendError exercises RecordListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestRecordListResource_List_SendError(t *testing.T) {
	r := &RecordListResource{client: newTransportErrorClient(t)}
	m := RecordListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestRecordListResource_List_InvalidJSON exercises RecordListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestRecordListResource_List_InvalidJSON(t *testing.T) {
	r := &RecordListResource{client: newMockClientStatus(t, 200, "{{")}
	m := RecordListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
