package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSipWhitelistResource_Create_Happy exercises SipWhitelistResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestSipWhitelistResource_Create_Happy(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := SipWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSipWhitelistResource_Create_NilClient exercises SipWhitelistResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSipWhitelistResource_Create_NilClient(t *testing.T) {
	r := &SipWhitelistResource{}
	m := SipWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSipWhitelistResource_Create_BuildError exercises SipWhitelistResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSipWhitelistResource_Create_BuildError(t *testing.T) {
	r := &SipWhitelistResource{client: newMalformedBaseURLClient(t)}
	m := SipWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSipWhitelistResource_Create_SendError exercises SipWhitelistResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestSipWhitelistResource_Create_SendError(t *testing.T) {
	r := &SipWhitelistResource{client: newTransportErrorClient(t)}
	m := SipWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSipWhitelistResource_Create_APIError exercises SipWhitelistResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSipWhitelistResource_Create_APIError(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SipWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_sip_whitelist")
}

// TestSipWhitelistResource_Create_APIErrorReadBody exercises SipWhitelistResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSipWhitelistResource_Create_APIErrorReadBody(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientReadErrorBody(t, 501)}
	m := SipWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSipWhitelistResource_Create_InvalidJSON exercises SipWhitelistResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSipWhitelistResource_Create_InvalidJSON(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 201, "{{")}
	m := SipWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSipWhitelistResource_Create_MapError exercises SipWhitelistResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSipWhitelistResource_Create_MapError(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := SipWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSipWhitelistResource_Create_MissingID exercises SipWhitelistResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestSipWhitelistResource_Create_MissingID(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 201, "{}")}
	m := SipWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestSipWhitelistResource_Create_LocationFallback exercises SipWhitelistResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestSipWhitelistResource_Create_LocationFallback(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := SipWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestSipWhitelistResource_Read_Happy exercises SipWhitelistResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestSipWhitelistResource_Read_Happy(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 200, "{}")}
	m := SipWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSipWhitelistResource_Read_NilClient exercises SipWhitelistResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSipWhitelistResource_Read_NilClient(t *testing.T) {
	r := &SipWhitelistResource{}
	m := SipWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSipWhitelistResource_Read_BuildError exercises SipWhitelistResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSipWhitelistResource_Read_BuildError(t *testing.T) {
	r := &SipWhitelistResource{client: newMalformedBaseURLClient(t)}
	m := SipWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSipWhitelistResource_Read_SendError exercises SipWhitelistResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestSipWhitelistResource_Read_SendError(t *testing.T) {
	r := &SipWhitelistResource{client: newTransportErrorClient(t)}
	m := SipWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSipWhitelistResource_Read_NotFound exercises SipWhitelistResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestSipWhitelistResource_Read_NotFound(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 404, "")}
	m := SipWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSipWhitelistResource_Read_APIError exercises SipWhitelistResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSipWhitelistResource_Read_APIError(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SipWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_sip_whitelist")
}

// TestSipWhitelistResource_Read_APIErrorReadBody exercises SipWhitelistResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSipWhitelistResource_Read_APIErrorReadBody(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientReadErrorBody(t, 501)}
	m := SipWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSipWhitelistResource_Read_InvalidJSON exercises SipWhitelistResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSipWhitelistResource_Read_InvalidJSON(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 200, "{{")}
	m := SipWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSipWhitelistResource_Read_MapError exercises SipWhitelistResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSipWhitelistResource_Read_MapError(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SipWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSipWhitelistResource_Delete_Happy exercises SipWhitelistResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestSipWhitelistResource_Delete_Happy(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 204, "")}
	m := SipWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSipWhitelistResource_Delete_NilClient exercises SipWhitelistResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSipWhitelistResource_Delete_NilClient(t *testing.T) {
	r := &SipWhitelistResource{}
	m := SipWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSipWhitelistResource_Delete_BuildError exercises SipWhitelistResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSipWhitelistResource_Delete_BuildError(t *testing.T) {
	r := &SipWhitelistResource{client: newMalformedBaseURLClient(t)}
	m := SipWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSipWhitelistResource_Delete_SendError exercises SipWhitelistResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestSipWhitelistResource_Delete_SendError(t *testing.T) {
	r := &SipWhitelistResource{client: newTransportErrorClient(t)}
	m := SipWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSipWhitelistResource_Delete_NotFoundSuccess exercises SipWhitelistResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestSipWhitelistResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 404, "")}
	m := SipWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSipWhitelistResource_Delete_APIError exercises SipWhitelistResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSipWhitelistResource_Delete_APIError(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SipWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_sip_whitelist")
}

// TestSipWhitelistResource_Delete_APIErrorReadBody exercises SipWhitelistResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSipWhitelistResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &SipWhitelistResource{client: newMockClientReadErrorBody(t, 501)}
	m := SipWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
