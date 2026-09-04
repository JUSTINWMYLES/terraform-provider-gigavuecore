package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestTunnelLbEndpointResource_Create_Happy exercises TunnelLbEndpointResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestTunnelLbEndpointResource_Create_Happy(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 201, "{\"te_id\":\"example-id\"}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelLbEndpointResource_Create_NilClient exercises TunnelLbEndpointResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTunnelLbEndpointResource_Create_NilClient(t *testing.T) {
	r := &TunnelLbEndpointResource{}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTunnelLbEndpointResource_Create_BuildError exercises TunnelLbEndpointResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTunnelLbEndpointResource_Create_BuildError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMalformedBaseURLClient(t)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTunnelLbEndpointResource_Create_SendError exercises TunnelLbEndpointResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestTunnelLbEndpointResource_Create_SendError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newTransportErrorClient(t)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTunnelLbEndpointResource_Create_APIError exercises TunnelLbEndpointResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTunnelLbEndpointResource_Create_APIError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_tunnel_lb_endpoint")
}

// TestTunnelLbEndpointResource_Create_APIErrorReadBody exercises TunnelLbEndpointResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTunnelLbEndpointResource_Create_APIErrorReadBody(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientReadErrorBody(t, 501)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTunnelLbEndpointResource_Create_InvalidJSON exercises TunnelLbEndpointResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTunnelLbEndpointResource_Create_InvalidJSON(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 201, "{{")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTunnelLbEndpointResource_Create_MapError exercises TunnelLbEndpointResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTunnelLbEndpointResource_Create_MapError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 201, "{\"te_id\":12345}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTunnelLbEndpointResource_Create_MissingID exercises TunnelLbEndpointResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestTunnelLbEndpointResource_Create_MissingID(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 201, "{}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestTunnelLbEndpointResource_Create_LocationFallback exercises TunnelLbEndpointResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestTunnelLbEndpointResource_Create_LocationFallback(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.TeId.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.TeId.ValueString(), "example-id")
	}
}

// TestTunnelLbEndpointResource_Read_Happy exercises TunnelLbEndpointResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestTunnelLbEndpointResource_Read_Happy(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 200, "{}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelLbEndpointResource_Read_NilClient exercises TunnelLbEndpointResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTunnelLbEndpointResource_Read_NilClient(t *testing.T) {
	r := &TunnelLbEndpointResource{}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTunnelLbEndpointResource_Read_BuildError exercises TunnelLbEndpointResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTunnelLbEndpointResource_Read_BuildError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMalformedBaseURLClient(t)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTunnelLbEndpointResource_Read_SendError exercises TunnelLbEndpointResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestTunnelLbEndpointResource_Read_SendError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newTransportErrorClient(t)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTunnelLbEndpointResource_Read_NotFound exercises TunnelLbEndpointResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestTunnelLbEndpointResource_Read_NotFound(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 404, "")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelLbEndpointResource_Read_APIError exercises TunnelLbEndpointResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTunnelLbEndpointResource_Read_APIError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_tunnel_lb_endpoint")
}

// TestTunnelLbEndpointResource_Read_APIErrorReadBody exercises TunnelLbEndpointResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTunnelLbEndpointResource_Read_APIErrorReadBody(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientReadErrorBody(t, 501)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTunnelLbEndpointResource_Read_InvalidJSON exercises TunnelLbEndpointResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTunnelLbEndpointResource_Read_InvalidJSON(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 200, "{{")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTunnelLbEndpointResource_Read_MapError exercises TunnelLbEndpointResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTunnelLbEndpointResource_Read_MapError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 200, "{\"te_id\":12345}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTunnelLbEndpointResource_Update_Happy exercises TunnelLbEndpointResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestTunnelLbEndpointResource_Update_Happy(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 200, "{}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelLbEndpointResource_Update_NilClient exercises TunnelLbEndpointResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTunnelLbEndpointResource_Update_NilClient(t *testing.T) {
	r := &TunnelLbEndpointResource{}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTunnelLbEndpointResource_Update_BuildError exercises TunnelLbEndpointResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTunnelLbEndpointResource_Update_BuildError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMalformedBaseURLClient(t)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTunnelLbEndpointResource_Update_SendError exercises TunnelLbEndpointResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestTunnelLbEndpointResource_Update_SendError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newTransportErrorClient(t)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTunnelLbEndpointResource_Update_APIError exercises TunnelLbEndpointResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTunnelLbEndpointResource_Update_APIError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_tunnel_lb_endpoint")
}

// TestTunnelLbEndpointResource_Update_APIErrorReadBody exercises TunnelLbEndpointResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTunnelLbEndpointResource_Update_APIErrorReadBody(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientReadErrorBody(t, 501)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTunnelLbEndpointResource_Update_InvalidJSON exercises TunnelLbEndpointResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTunnelLbEndpointResource_Update_InvalidJSON(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 200, "{{")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTunnelLbEndpointResource_Update_MapError exercises TunnelLbEndpointResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTunnelLbEndpointResource_Update_MapError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 200, "{\"te_id\":12345}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTunnelLbEndpointResource_Delete_Happy exercises TunnelLbEndpointResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestTunnelLbEndpointResource_Delete_Happy(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 204, "")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelLbEndpointResource_Delete_NilClient exercises TunnelLbEndpointResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTunnelLbEndpointResource_Delete_NilClient(t *testing.T) {
	r := &TunnelLbEndpointResource{}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTunnelLbEndpointResource_Delete_BuildError exercises TunnelLbEndpointResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTunnelLbEndpointResource_Delete_BuildError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMalformedBaseURLClient(t)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTunnelLbEndpointResource_Delete_SendError exercises TunnelLbEndpointResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestTunnelLbEndpointResource_Delete_SendError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newTransportErrorClient(t)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTunnelLbEndpointResource_Delete_NotFoundSuccess exercises TunnelLbEndpointResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestTunnelLbEndpointResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 404, "")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTunnelLbEndpointResource_Delete_APIError exercises TunnelLbEndpointResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTunnelLbEndpointResource_Delete_APIError(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_tunnel_lb_endpoint")
}

// TestTunnelLbEndpointResource_Delete_APIErrorReadBody exercises TunnelLbEndpointResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTunnelLbEndpointResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &TunnelLbEndpointResource{client: newMockClientReadErrorBody(t, 501)}
	m := TunnelLbEndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
