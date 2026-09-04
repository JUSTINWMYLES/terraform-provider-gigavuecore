package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestHbPacketResource_Create_Happy exercises HbPacketResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestHbPacketResource_Create_Happy(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := HbPacketResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbPacketResource_Create_NilClient exercises HbPacketResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHbPacketResource_Create_NilClient(t *testing.T) {
	r := &HbPacketResource{}
	m := HbPacketResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHbPacketResource_Create_BuildError exercises HbPacketResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHbPacketResource_Create_BuildError(t *testing.T) {
	r := &HbPacketResource{client: newMalformedBaseURLClient(t)}
	m := HbPacketResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHbPacketResource_Create_SendError exercises HbPacketResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestHbPacketResource_Create_SendError(t *testing.T) {
	r := &HbPacketResource{client: newTransportErrorClient(t)}
	m := HbPacketResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHbPacketResource_Create_APIError exercises HbPacketResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHbPacketResource_Create_APIError(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HbPacketResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_hb_packet")
}

// TestHbPacketResource_Create_APIErrorReadBody exercises HbPacketResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHbPacketResource_Create_APIErrorReadBody(t *testing.T) {
	r := &HbPacketResource{client: newMockClientReadErrorBody(t, 501)}
	m := HbPacketResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHbPacketResource_Create_InvalidJSON exercises HbPacketResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHbPacketResource_Create_InvalidJSON(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 201, "{{")}
	m := HbPacketResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHbPacketResource_Create_MapError exercises HbPacketResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHbPacketResource_Create_MapError(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := HbPacketResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHbPacketResource_Create_MissingID exercises HbPacketResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestHbPacketResource_Create_MissingID(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 201, "{}")}
	m := HbPacketResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestHbPacketResource_Create_LocationFallback exercises HbPacketResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestHbPacketResource_Create_LocationFallback(t *testing.T) {
	r := &HbPacketResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := HbPacketResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestHbPacketResource_Read_Happy exercises HbPacketResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestHbPacketResource_Read_Happy(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 200, "{}")}
	m := HbPacketResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbPacketResource_Read_NilClient exercises HbPacketResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHbPacketResource_Read_NilClient(t *testing.T) {
	r := &HbPacketResource{}
	m := HbPacketResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHbPacketResource_Read_BuildError exercises HbPacketResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHbPacketResource_Read_BuildError(t *testing.T) {
	r := &HbPacketResource{client: newMalformedBaseURLClient(t)}
	m := HbPacketResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHbPacketResource_Read_SendError exercises HbPacketResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestHbPacketResource_Read_SendError(t *testing.T) {
	r := &HbPacketResource{client: newTransportErrorClient(t)}
	m := HbPacketResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHbPacketResource_Read_NotFound exercises HbPacketResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestHbPacketResource_Read_NotFound(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 404, "")}
	m := HbPacketResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbPacketResource_Read_APIError exercises HbPacketResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHbPacketResource_Read_APIError(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HbPacketResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_hb_packet")
}

// TestHbPacketResource_Read_APIErrorReadBody exercises HbPacketResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHbPacketResource_Read_APIErrorReadBody(t *testing.T) {
	r := &HbPacketResource{client: newMockClientReadErrorBody(t, 501)}
	m := HbPacketResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHbPacketResource_Read_InvalidJSON exercises HbPacketResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHbPacketResource_Read_InvalidJSON(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 200, "{{")}
	m := HbPacketResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHbPacketResource_Read_MapError exercises HbPacketResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHbPacketResource_Read_MapError(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := HbPacketResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHbPacketResource_Delete_Happy exercises HbPacketResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestHbPacketResource_Delete_Happy(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 204, "")}
	m := HbPacketResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbPacketResource_Delete_NilClient exercises HbPacketResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHbPacketResource_Delete_NilClient(t *testing.T) {
	r := &HbPacketResource{}
	m := HbPacketResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHbPacketResource_Delete_BuildError exercises HbPacketResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHbPacketResource_Delete_BuildError(t *testing.T) {
	r := &HbPacketResource{client: newMalformedBaseURLClient(t)}
	m := HbPacketResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHbPacketResource_Delete_SendError exercises HbPacketResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestHbPacketResource_Delete_SendError(t *testing.T) {
	r := &HbPacketResource{client: newTransportErrorClient(t)}
	m := HbPacketResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHbPacketResource_Delete_NotFoundSuccess exercises HbPacketResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestHbPacketResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 404, "")}
	m := HbPacketResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbPacketResource_Delete_APIError exercises HbPacketResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHbPacketResource_Delete_APIError(t *testing.T) {
	r := &HbPacketResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HbPacketResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_hb_packet")
}

// TestHbPacketResource_Delete_APIErrorReadBody exercises HbPacketResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHbPacketResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &HbPacketResource{client: newMockClientReadErrorBody(t, 501)}
	m := HbPacketResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
