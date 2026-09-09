package provider

import (
	"context"
	"testing"
)

// TestLoadAllSystemAcmeCertificateDetailsListResource_List_Happy exercises LoadAllSystemAcmeCertificateDetailsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllSystemAcmeCertificateDetailsListResource_List_Happy(t *testing.T) {
	r := &LoadAllSystemAcmeCertificateDetailsListResource{client: newMockClientStatus(t, 200, "{\"certificateInfo\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestLoadAllSystemAcmeCertificateDetailsListResource_List_NilClient exercises LoadAllSystemAcmeCertificateDetailsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllSystemAcmeCertificateDetailsListResource_List_NilClient(t *testing.T) {
	r := &LoadAllSystemAcmeCertificateDetailsListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadAllSystemAcmeCertificateDetailsListResource_List_BuildError exercises LoadAllSystemAcmeCertificateDetailsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllSystemAcmeCertificateDetailsListResource_List_BuildError(t *testing.T) {
	r := &LoadAllSystemAcmeCertificateDetailsListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllSystemAcmeCertificateDetailsListResource_List_SendError exercises LoadAllSystemAcmeCertificateDetailsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllSystemAcmeCertificateDetailsListResource_List_SendError(t *testing.T) {
	r := &LoadAllSystemAcmeCertificateDetailsListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAllSystemAcmeCertificateDetailsListResource_List_InvalidJSON exercises LoadAllSystemAcmeCertificateDetailsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllSystemAcmeCertificateDetailsListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadAllSystemAcmeCertificateDetailsListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
