package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestLdapServerResource_Create_Happy exercises LdapServerResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestLdapServerResource_Create_Happy(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 201, "{\"server_address\":\"example-id\"}")}
	m := LdapServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLdapServerResource_Create_NilClient exercises LdapServerResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLdapServerResource_Create_NilClient(t *testing.T) {
	r := &LdapServerResource{}
	m := LdapServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLdapServerResource_Create_BuildError exercises LdapServerResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLdapServerResource_Create_BuildError(t *testing.T) {
	r := &LdapServerResource{client: newMalformedBaseURLClient(t)}
	m := LdapServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLdapServerResource_Create_SendError exercises LdapServerResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestLdapServerResource_Create_SendError(t *testing.T) {
	r := &LdapServerResource{client: newTransportErrorClient(t)}
	m := LdapServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLdapServerResource_Create_APIError exercises LdapServerResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLdapServerResource_Create_APIError(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LdapServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_ldap_server")
}

// TestLdapServerResource_Create_APIErrorReadBody exercises LdapServerResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLdapServerResource_Create_APIErrorReadBody(t *testing.T) {
	r := &LdapServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := LdapServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLdapServerResource_Create_InvalidJSON exercises LdapServerResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLdapServerResource_Create_InvalidJSON(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 201, "{{")}
	m := LdapServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestLdapServerResource_Create_MapError exercises LdapServerResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestLdapServerResource_Create_MapError(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 201, "{\"server_address\":12345}")}
	m := LdapServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestLdapServerResource_Create_MissingID exercises LdapServerResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestLdapServerResource_Create_MissingID(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 201, "{}")}
	m := LdapServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestLdapServerResource_Create_LocationFallback exercises LdapServerResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestLdapServerResource_Create_LocationFallback(t *testing.T) {
	r := &LdapServerResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := LdapServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.ServerAddress.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.ServerAddress.ValueString(), "example-id")
	}
}

// TestLdapServerResource_Read_Happy exercises LdapServerResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestLdapServerResource_Read_Happy(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := LdapServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestLdapServerResource_Read_NilClient exercises LdapServerResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLdapServerResource_Read_NilClient(t *testing.T) {
	r := &LdapServerResource{}
	m := LdapServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLdapServerResource_Read_BuildError exercises LdapServerResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLdapServerResource_Read_BuildError(t *testing.T) {
	r := &LdapServerResource{client: newMalformedBaseURLClient(t)}
	m := LdapServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLdapServerResource_Read_SendError exercises LdapServerResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLdapServerResource_Read_SendError(t *testing.T) {
	r := &LdapServerResource{client: newTransportErrorClient(t)}
	m := LdapServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLdapServerResource_Read_NotFound exercises LdapServerResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestLdapServerResource_Read_NotFound(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 404, "")}
	m := LdapServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestLdapServerResource_Read_APIError exercises LdapServerResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLdapServerResource_Read_APIError(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LdapServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_ldap_server")
}

// TestLdapServerResource_Read_APIErrorReadBody exercises LdapServerResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLdapServerResource_Read_APIErrorReadBody(t *testing.T) {
	r := &LdapServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := LdapServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLdapServerResource_Read_InvalidJSON exercises LdapServerResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLdapServerResource_Read_InvalidJSON(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := LdapServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestLdapServerResource_Read_MapError exercises LdapServerResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestLdapServerResource_Read_MapError(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 200, "{\"server_address\":12345}")}
	m := LdapServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestLdapServerResource_Delete_Happy exercises LdapServerResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestLdapServerResource_Delete_Happy(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 204, "")}
	m := LdapServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLdapServerResource_Delete_NilClient exercises LdapServerResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLdapServerResource_Delete_NilClient(t *testing.T) {
	r := &LdapServerResource{}
	m := LdapServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLdapServerResource_Delete_BuildError exercises LdapServerResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLdapServerResource_Delete_BuildError(t *testing.T) {
	r := &LdapServerResource{client: newMalformedBaseURLClient(t)}
	m := LdapServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLdapServerResource_Delete_SendError exercises LdapServerResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestLdapServerResource_Delete_SendError(t *testing.T) {
	r := &LdapServerResource{client: newTransportErrorClient(t)}
	m := LdapServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLdapServerResource_Delete_NotFoundSuccess exercises LdapServerResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestLdapServerResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 404, "")}
	m := LdapServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLdapServerResource_Delete_APIError exercises LdapServerResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLdapServerResource_Delete_APIError(t *testing.T) {
	r := &LdapServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LdapServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_ldap_server")
}

// TestLdapServerResource_Delete_APIErrorReadBody exercises LdapServerResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLdapServerResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &LdapServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := LdapServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
