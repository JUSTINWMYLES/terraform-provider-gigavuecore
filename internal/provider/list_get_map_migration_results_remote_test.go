package provider

import (
	"context"
	"testing"
)

// TestGetMapMigrationResultsListResource_List_Happy exercises GetMapMigrationResultsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetMapMigrationResultsListResource_List_Happy(t *testing.T) {
	r := &GetMapMigrationResultsListResource{client: newMockClientStatus(t, 200, "{\"mapMigrationResults\":[]}")}
	m := GetMapMigrationResultsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetMapMigrationResultsListResource_List_NilClient exercises GetMapMigrationResultsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMapMigrationResultsListResource_List_NilClient(t *testing.T) {
	r := &GetMapMigrationResultsListResource{}
	m := GetMapMigrationResultsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetMapMigrationResultsListResource_List_BuildError exercises GetMapMigrationResultsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetMapMigrationResultsListResource_List_BuildError(t *testing.T) {
	r := &GetMapMigrationResultsListResource{client: newMalformedBaseURLClient(t)}
	m := GetMapMigrationResultsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetMapMigrationResultsListResource_List_SendError exercises GetMapMigrationResultsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetMapMigrationResultsListResource_List_SendError(t *testing.T) {
	r := &GetMapMigrationResultsListResource{client: newTransportErrorClient(t)}
	m := GetMapMigrationResultsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetMapMigrationResultsListResource_List_InvalidJSON exercises GetMapMigrationResultsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetMapMigrationResultsListResource_List_InvalidJSON(t *testing.T) {
	r := &GetMapMigrationResultsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMapMigrationResultsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
