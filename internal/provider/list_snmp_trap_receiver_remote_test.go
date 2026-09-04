package provider

import (
	"context"
	"testing"
)

// TestSnmpTrapReceiverListResource_List_Happy exercises SnmpTrapReceiverListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestSnmpTrapReceiverListResource_List_Happy(t *testing.T) {
	r := &SnmpTrapReceiverListResource{client: newMockClientStatus(t, 200, "{\"externalTrapReceivers\":[]}")}
	m := SnmpTrapReceiverListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestSnmpTrapReceiverListResource_List_NilClient exercises SnmpTrapReceiverListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSnmpTrapReceiverListResource_List_NilClient(t *testing.T) {
	r := &SnmpTrapReceiverListResource{}
	m := SnmpTrapReceiverListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestSnmpTrapReceiverListResource_List_BuildError exercises SnmpTrapReceiverListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestSnmpTrapReceiverListResource_List_BuildError(t *testing.T) {
	r := &SnmpTrapReceiverListResource{client: newMalformedBaseURLClient(t)}
	m := SnmpTrapReceiverListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSnmpTrapReceiverListResource_List_SendError exercises SnmpTrapReceiverListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestSnmpTrapReceiverListResource_List_SendError(t *testing.T) {
	r := &SnmpTrapReceiverListResource{client: newTransportErrorClient(t)}
	m := SnmpTrapReceiverListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSnmpTrapReceiverListResource_List_InvalidJSON exercises SnmpTrapReceiverListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestSnmpTrapReceiverListResource_List_InvalidJSON(t *testing.T) {
	r := &SnmpTrapReceiverListResource{client: newMockClientStatus(t, 200, "{{")}
	m := SnmpTrapReceiverListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
