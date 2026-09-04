package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestDiameterWhitelistResource_Create_Happy exercises DiameterWhitelistResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestDiameterWhitelistResource_Create_Happy(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDiameterWhitelistResource_Create_NilClient exercises DiameterWhitelistResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDiameterWhitelistResource_Create_NilClient(t *testing.T) {
	r := &DiameterWhitelistResource{}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDiameterWhitelistResource_Create_BuildError exercises DiameterWhitelistResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDiameterWhitelistResource_Create_BuildError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMalformedBaseURLClient(t)}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDiameterWhitelistResource_Create_SendError exercises DiameterWhitelistResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestDiameterWhitelistResource_Create_SendError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newTransportErrorClient(t)}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDiameterWhitelistResource_Create_APIError exercises DiameterWhitelistResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDiameterWhitelistResource_Create_APIError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_diameter_whitelist")
}

// TestDiameterWhitelistResource_Create_APIErrorReadBody exercises DiameterWhitelistResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDiameterWhitelistResource_Create_APIErrorReadBody(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientReadErrorBody(t, 501)}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestDiameterWhitelistResource_Create_InvalidJSON exercises DiameterWhitelistResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestDiameterWhitelistResource_Create_InvalidJSON(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 201, "{{")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestDiameterWhitelistResource_Create_MapError exercises DiameterWhitelistResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestDiameterWhitelistResource_Create_MapError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestDiameterWhitelistResource_Create_MissingID exercises DiameterWhitelistResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestDiameterWhitelistResource_Create_MissingID(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 201, "{}")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestDiameterWhitelistResource_Create_LocationFallback exercises DiameterWhitelistResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestDiameterWhitelistResource_Create_LocationFallback(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestDiameterWhitelistResource_Read_Happy exercises DiameterWhitelistResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestDiameterWhitelistResource_Read_Happy(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 200, "{}")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestDiameterWhitelistResource_Read_NilClient exercises DiameterWhitelistResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDiameterWhitelistResource_Read_NilClient(t *testing.T) {
	r := &DiameterWhitelistResource{}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDiameterWhitelistResource_Read_BuildError exercises DiameterWhitelistResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDiameterWhitelistResource_Read_BuildError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMalformedBaseURLClient(t)}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDiameterWhitelistResource_Read_SendError exercises DiameterWhitelistResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestDiameterWhitelistResource_Read_SendError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newTransportErrorClient(t)}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDiameterWhitelistResource_Read_NotFound exercises DiameterWhitelistResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestDiameterWhitelistResource_Read_NotFound(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 404, "")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestDiameterWhitelistResource_Read_APIError exercises DiameterWhitelistResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDiameterWhitelistResource_Read_APIError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_diameter_whitelist")
}

// TestDiameterWhitelistResource_Read_APIErrorReadBody exercises DiameterWhitelistResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDiameterWhitelistResource_Read_APIErrorReadBody(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientReadErrorBody(t, 501)}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestDiameterWhitelistResource_Read_InvalidJSON exercises DiameterWhitelistResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestDiameterWhitelistResource_Read_InvalidJSON(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 200, "{{")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestDiameterWhitelistResource_Read_MapError exercises DiameterWhitelistResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestDiameterWhitelistResource_Read_MapError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestDiameterWhitelistResource_Delete_Happy exercises DiameterWhitelistResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestDiameterWhitelistResource_Delete_Happy(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 204, "")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDiameterWhitelistResource_Delete_NilClient exercises DiameterWhitelistResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDiameterWhitelistResource_Delete_NilClient(t *testing.T) {
	r := &DiameterWhitelistResource{}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDiameterWhitelistResource_Delete_BuildError exercises DiameterWhitelistResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDiameterWhitelistResource_Delete_BuildError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMalformedBaseURLClient(t)}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDiameterWhitelistResource_Delete_SendError exercises DiameterWhitelistResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestDiameterWhitelistResource_Delete_SendError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newTransportErrorClient(t)}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDiameterWhitelistResource_Delete_NotFoundSuccess exercises DiameterWhitelistResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestDiameterWhitelistResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 404, "")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDiameterWhitelistResource_Delete_APIError exercises DiameterWhitelistResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDiameterWhitelistResource_Delete_APIError(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_diameter_whitelist")
}

// TestDiameterWhitelistResource_Delete_APIErrorReadBody exercises DiameterWhitelistResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDiameterWhitelistResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &DiameterWhitelistResource{client: newMockClientReadErrorBody(t, 501)}
	m := DiameterWhitelistResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
