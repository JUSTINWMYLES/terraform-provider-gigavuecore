package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestGtpWhitelistResource_Create_Happy exercises GtpWhitelistResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestGtpWhitelistResource_Create_Happy(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtpWhitelistResource_Create_NilClient exercises GtpWhitelistResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGtpWhitelistResource_Create_NilClient(t *testing.T) {
	r := &GtpWhitelistResource{}
	m := GtpWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGtpWhitelistResource_Create_BuildError exercises GtpWhitelistResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGtpWhitelistResource_Create_BuildError(t *testing.T) {
	r := &GtpWhitelistResource{client: newMalformedBaseURLClient(t)}
	m := GtpWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGtpWhitelistResource_Create_SendError exercises GtpWhitelistResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestGtpWhitelistResource_Create_SendError(t *testing.T) {
	r := &GtpWhitelistResource{client: newTransportErrorClient(t)}
	m := GtpWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGtpWhitelistResource_Create_APIError exercises GtpWhitelistResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGtpWhitelistResource_Create_APIError(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_gtp_whitelist")
}

// TestGtpWhitelistResource_Create_APIErrorReadBody exercises GtpWhitelistResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGtpWhitelistResource_Create_APIErrorReadBody(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientReadErrorBody(t, 501)}
	m := GtpWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGtpWhitelistResource_Create_InvalidJSON exercises GtpWhitelistResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGtpWhitelistResource_Create_InvalidJSON(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 201, "{{")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGtpWhitelistResource_Create_MapError exercises GtpWhitelistResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGtpWhitelistResource_Create_MapError(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGtpWhitelistResource_Create_MissingID exercises GtpWhitelistResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestGtpWhitelistResource_Create_MissingID(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 201, "{}")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestGtpWhitelistResource_Create_LocationFallback exercises GtpWhitelistResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestGtpWhitelistResource_Create_LocationFallback(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestGtpWhitelistResource_Read_Happy exercises GtpWhitelistResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestGtpWhitelistResource_Read_Happy(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 200, "{}")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtpWhitelistResource_Read_NilClient exercises GtpWhitelistResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGtpWhitelistResource_Read_NilClient(t *testing.T) {
	r := &GtpWhitelistResource{}
	m := GtpWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGtpWhitelistResource_Read_BuildError exercises GtpWhitelistResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGtpWhitelistResource_Read_BuildError(t *testing.T) {
	r := &GtpWhitelistResource{client: newMalformedBaseURLClient(t)}
	m := GtpWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGtpWhitelistResource_Read_SendError exercises GtpWhitelistResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGtpWhitelistResource_Read_SendError(t *testing.T) {
	r := &GtpWhitelistResource{client: newTransportErrorClient(t)}
	m := GtpWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGtpWhitelistResource_Read_NotFound exercises GtpWhitelistResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestGtpWhitelistResource_Read_NotFound(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 404, "")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtpWhitelistResource_Read_APIError exercises GtpWhitelistResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGtpWhitelistResource_Read_APIError(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_gtp_whitelist")
}

// TestGtpWhitelistResource_Read_APIErrorReadBody exercises GtpWhitelistResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGtpWhitelistResource_Read_APIErrorReadBody(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientReadErrorBody(t, 501)}
	m := GtpWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGtpWhitelistResource_Read_InvalidJSON exercises GtpWhitelistResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGtpWhitelistResource_Read_InvalidJSON(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 200, "{{")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGtpWhitelistResource_Read_MapError exercises GtpWhitelistResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGtpWhitelistResource_Read_MapError(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGtpWhitelistResource_Delete_Happy exercises GtpWhitelistResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestGtpWhitelistResource_Delete_Happy(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 204, "")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtpWhitelistResource_Delete_NilClient exercises GtpWhitelistResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGtpWhitelistResource_Delete_NilClient(t *testing.T) {
	r := &GtpWhitelistResource{}
	m := GtpWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGtpWhitelistResource_Delete_BuildError exercises GtpWhitelistResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGtpWhitelistResource_Delete_BuildError(t *testing.T) {
	r := &GtpWhitelistResource{client: newMalformedBaseURLClient(t)}
	m := GtpWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGtpWhitelistResource_Delete_SendError exercises GtpWhitelistResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestGtpWhitelistResource_Delete_SendError(t *testing.T) {
	r := &GtpWhitelistResource{client: newTransportErrorClient(t)}
	m := GtpWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGtpWhitelistResource_Delete_NotFoundSuccess exercises GtpWhitelistResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestGtpWhitelistResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 404, "")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtpWhitelistResource_Delete_APIError exercises GtpWhitelistResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGtpWhitelistResource_Delete_APIError(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GtpWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_gtp_whitelist")
}

// TestGtpWhitelistResource_Delete_APIErrorReadBody exercises GtpWhitelistResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGtpWhitelistResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &GtpWhitelistResource{client: newMockClientReadErrorBody(t, 501)}
	m := GtpWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
