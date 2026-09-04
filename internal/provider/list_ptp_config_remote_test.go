package provider

import (
	"context"
	"testing"
)

// TestPtpConfigListResource_List_Happy exercises PtpConfigListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestPtpConfigListResource_List_Happy(t *testing.T) {
	r := &PtpConfigListResource{client: newMockClientStatus(t, 200, "{\"ptpConfigs\":[]}")}
	m := PtpConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestPtpConfigListResource_List_NilClient exercises PtpConfigListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPtpConfigListResource_List_NilClient(t *testing.T) {
	r := &PtpConfigListResource{}
	m := PtpConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestPtpConfigListResource_List_BuildError exercises PtpConfigListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestPtpConfigListResource_List_BuildError(t *testing.T) {
	r := &PtpConfigListResource{client: newMalformedBaseURLClient(t)}
	m := PtpConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPtpConfigListResource_List_SendError exercises PtpConfigListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestPtpConfigListResource_List_SendError(t *testing.T) {
	r := &PtpConfigListResource{client: newTransportErrorClient(t)}
	m := PtpConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestPtpConfigListResource_List_InvalidJSON exercises PtpConfigListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestPtpConfigListResource_List_InvalidJSON(t *testing.T) {
	r := &PtpConfigListResource{client: newMockClientStatus(t, 200, "{{")}
	m := PtpConfigListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
