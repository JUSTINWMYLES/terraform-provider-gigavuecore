package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestTunnelApplicationResource_Create_Happy exercises TunnelApplicationResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestTunnelApplicationResource_Create_Happy(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelApplicationResource_Create_NilClient exercises TunnelApplicationResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTunnelApplicationResource_Create_NilClient(t *testing.T) {
	r := &TunnelApplicationResource{}
	m := TunnelApplicationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTunnelApplicationResource_Create_BuildError exercises TunnelApplicationResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTunnelApplicationResource_Create_BuildError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMalformedBaseURLClient(t)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTunnelApplicationResource_Create_SendError exercises TunnelApplicationResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestTunnelApplicationResource_Create_SendError(t *testing.T) {
	r := &TunnelApplicationResource{client: newTransportErrorClient(t)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTunnelApplicationResource_Create_APIError exercises TunnelApplicationResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTunnelApplicationResource_Create_APIError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_tunnel_application")
}

// TestTunnelApplicationResource_Create_APIErrorReadBody exercises TunnelApplicationResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTunnelApplicationResource_Create_APIErrorReadBody(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientReadErrorBody(t, 501)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTunnelApplicationResource_Create_InvalidJSON exercises TunnelApplicationResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTunnelApplicationResource_Create_InvalidJSON(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 201, "{{")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTunnelApplicationResource_Create_MapError exercises TunnelApplicationResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTunnelApplicationResource_Create_MapError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTunnelApplicationResource_Create_MissingID exercises TunnelApplicationResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestTunnelApplicationResource_Create_MissingID(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 201, "{}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestTunnelApplicationResource_Create_LocationFallback exercises TunnelApplicationResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestTunnelApplicationResource_Create_LocationFallback(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestTunnelApplicationResource_Read_Happy exercises TunnelApplicationResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestTunnelApplicationResource_Read_Happy(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 200, "{}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelApplicationResource_Read_NilClient exercises TunnelApplicationResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTunnelApplicationResource_Read_NilClient(t *testing.T) {
	r := &TunnelApplicationResource{}
	m := TunnelApplicationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTunnelApplicationResource_Read_BuildError exercises TunnelApplicationResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTunnelApplicationResource_Read_BuildError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMalformedBaseURLClient(t)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTunnelApplicationResource_Read_SendError exercises TunnelApplicationResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestTunnelApplicationResource_Read_SendError(t *testing.T) {
	r := &TunnelApplicationResource{client: newTransportErrorClient(t)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTunnelApplicationResource_Read_NotFound exercises TunnelApplicationResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestTunnelApplicationResource_Read_NotFound(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 404, "")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelApplicationResource_Read_APIError exercises TunnelApplicationResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTunnelApplicationResource_Read_APIError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_tunnel_application")
}

// TestTunnelApplicationResource_Read_APIErrorReadBody exercises TunnelApplicationResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTunnelApplicationResource_Read_APIErrorReadBody(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientReadErrorBody(t, 501)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTunnelApplicationResource_Read_InvalidJSON exercises TunnelApplicationResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTunnelApplicationResource_Read_InvalidJSON(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 200, "{{")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTunnelApplicationResource_Read_MapError exercises TunnelApplicationResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTunnelApplicationResource_Read_MapError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTunnelApplicationResource_Update_Happy exercises TunnelApplicationResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestTunnelApplicationResource_Update_Happy(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 200, "{}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelApplicationResource_Update_NilClient exercises TunnelApplicationResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTunnelApplicationResource_Update_NilClient(t *testing.T) {
	r := &TunnelApplicationResource{}
	m := TunnelApplicationResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTunnelApplicationResource_Update_BuildError exercises TunnelApplicationResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTunnelApplicationResource_Update_BuildError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMalformedBaseURLClient(t)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTunnelApplicationResource_Update_SendError exercises TunnelApplicationResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestTunnelApplicationResource_Update_SendError(t *testing.T) {
	r := &TunnelApplicationResource{client: newTransportErrorClient(t)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTunnelApplicationResource_Update_APIError exercises TunnelApplicationResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTunnelApplicationResource_Update_APIError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_tunnel_application")
}

// TestTunnelApplicationResource_Update_APIErrorReadBody exercises TunnelApplicationResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTunnelApplicationResource_Update_APIErrorReadBody(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientReadErrorBody(t, 501)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTunnelApplicationResource_Update_InvalidJSON exercises TunnelApplicationResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTunnelApplicationResource_Update_InvalidJSON(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 200, "{{")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTunnelApplicationResource_Update_MapError exercises TunnelApplicationResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTunnelApplicationResource_Update_MapError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTunnelApplicationResource_Delete_Happy exercises TunnelApplicationResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestTunnelApplicationResource_Delete_Happy(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 204, "")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelApplicationResource_Delete_NilClient exercises TunnelApplicationResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTunnelApplicationResource_Delete_NilClient(t *testing.T) {
	r := &TunnelApplicationResource{}
	m := TunnelApplicationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTunnelApplicationResource_Delete_BuildError exercises TunnelApplicationResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTunnelApplicationResource_Delete_BuildError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMalformedBaseURLClient(t)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTunnelApplicationResource_Delete_SendError exercises TunnelApplicationResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestTunnelApplicationResource_Delete_SendError(t *testing.T) {
	r := &TunnelApplicationResource{client: newTransportErrorClient(t)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTunnelApplicationResource_Delete_NotFoundSuccess exercises TunnelApplicationResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestTunnelApplicationResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 404, "")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelApplicationResource_Delete_APIError exercises TunnelApplicationResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTunnelApplicationResource_Delete_APIError(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TunnelApplicationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_tunnel_application")
}

// TestTunnelApplicationResource_Delete_APIErrorReadBody exercises TunnelApplicationResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTunnelApplicationResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &TunnelApplicationResource{client: newMockClientReadErrorBody(t, 501)}
	m := TunnelApplicationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
