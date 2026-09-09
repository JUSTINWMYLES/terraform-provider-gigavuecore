package provider

import (
	"context"
	"testing"
)

// TestAlertPolicyListResource_List_Happy exercises AlertPolicyListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestAlertPolicyListResource_List_Happy(t *testing.T) {
	r := &AlertPolicyListResource{client: newMockClientStatus(t, 200, "{\"alertPolicies\":[]}")}
	m := AlertPolicyListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestAlertPolicyListResource_List_NilClient exercises AlertPolicyListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAlertPolicyListResource_List_NilClient(t *testing.T) {
	r := &AlertPolicyListResource{}
	m := AlertPolicyListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestAlertPolicyListResource_List_BuildError exercises AlertPolicyListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestAlertPolicyListResource_List_BuildError(t *testing.T) {
	r := &AlertPolicyListResource{client: newMalformedBaseURLClient(t)}
	m := AlertPolicyListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestAlertPolicyListResource_List_SendError exercises AlertPolicyListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestAlertPolicyListResource_List_SendError(t *testing.T) {
	r := &AlertPolicyListResource{client: newTransportErrorClient(t)}
	m := AlertPolicyListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestAlertPolicyListResource_List_InvalidJSON exercises AlertPolicyListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestAlertPolicyListResource_List_InvalidJSON(t *testing.T) {
	r := &AlertPolicyListResource{client: newMockClientStatus(t, 200, "{{")}
	m := AlertPolicyListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
