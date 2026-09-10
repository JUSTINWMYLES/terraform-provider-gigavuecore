package provider

import (
	"context"
	"testing"
)

// TestLoadIpDestinationStatusesListResource_List_Happy exercises LoadIpDestinationStatusesListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadIpDestinationStatusesListResource_List_Happy(t *testing.T) {
	r := &LoadIpDestinationStatusesListResource{client: newMockClientStatus(t, 200, "{\"ipInterfaces\":[]}")}
	m := LoadIpDestinationStatusesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadIpDestinationStatusesListResource_List_NilClient exercises LoadIpDestinationStatusesListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadIpDestinationStatusesListResource_List_NilClient(t *testing.T) {
	r := &LoadIpDestinationStatusesListResource{}
	m := LoadIpDestinationStatusesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadIpDestinationStatusesListResource_List_BuildError exercises LoadIpDestinationStatusesListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadIpDestinationStatusesListResource_List_BuildError(t *testing.T) {
	r := &LoadIpDestinationStatusesListResource{client: newMalformedBaseURLClient(t)}
	m := LoadIpDestinationStatusesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadIpDestinationStatusesListResource_List_SendError exercises LoadIpDestinationStatusesListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadIpDestinationStatusesListResource_List_SendError(t *testing.T) {
	r := &LoadIpDestinationStatusesListResource{client: newTransportErrorClient(t)}
	m := LoadIpDestinationStatusesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadIpDestinationStatusesListResource_List_InvalidJSON exercises LoadIpDestinationStatusesListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadIpDestinationStatusesListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadIpDestinationStatusesListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadIpDestinationStatusesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
