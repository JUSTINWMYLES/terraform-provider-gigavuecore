package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestCircuitTunnelResource_Create_Happy exercises CircuitTunnelResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestCircuitTunnelResource_Create_Happy(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCircuitTunnelResource_Create_NilClient exercises CircuitTunnelResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCircuitTunnelResource_Create_NilClient(t *testing.T) {
	r := &CircuitTunnelResource{}
	m := CircuitTunnelResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCircuitTunnelResource_Create_BuildError exercises CircuitTunnelResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCircuitTunnelResource_Create_BuildError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMalformedBaseURLClient(t)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCircuitTunnelResource_Create_SendError exercises CircuitTunnelResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestCircuitTunnelResource_Create_SendError(t *testing.T) {
	r := &CircuitTunnelResource{client: newTransportErrorClient(t)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCircuitTunnelResource_Create_APIError exercises CircuitTunnelResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCircuitTunnelResource_Create_APIError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_circuit_tunnel")
}

// TestCircuitTunnelResource_Create_APIErrorReadBody exercises CircuitTunnelResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCircuitTunnelResource_Create_APIErrorReadBody(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientReadErrorBody(t, 501)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestCircuitTunnelResource_Create_InvalidJSON exercises CircuitTunnelResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestCircuitTunnelResource_Create_InvalidJSON(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 201, "{{")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestCircuitTunnelResource_Create_MapError exercises CircuitTunnelResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestCircuitTunnelResource_Create_MapError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestCircuitTunnelResource_Create_MissingID exercises CircuitTunnelResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestCircuitTunnelResource_Create_MissingID(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 201, "{}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestCircuitTunnelResource_Create_LocationFallback exercises CircuitTunnelResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestCircuitTunnelResource_Create_LocationFallback(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestCircuitTunnelResource_Read_Happy exercises CircuitTunnelResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestCircuitTunnelResource_Read_Happy(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 200, "{}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestCircuitTunnelResource_Read_NilClient exercises CircuitTunnelResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCircuitTunnelResource_Read_NilClient(t *testing.T) {
	r := &CircuitTunnelResource{}
	m := CircuitTunnelResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCircuitTunnelResource_Read_BuildError exercises CircuitTunnelResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCircuitTunnelResource_Read_BuildError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMalformedBaseURLClient(t)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCircuitTunnelResource_Read_SendError exercises CircuitTunnelResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestCircuitTunnelResource_Read_SendError(t *testing.T) {
	r := &CircuitTunnelResource{client: newTransportErrorClient(t)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCircuitTunnelResource_Read_NotFound exercises CircuitTunnelResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestCircuitTunnelResource_Read_NotFound(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 404, "")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestCircuitTunnelResource_Read_APIError exercises CircuitTunnelResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCircuitTunnelResource_Read_APIError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_circuit_tunnel")
}

// TestCircuitTunnelResource_Read_APIErrorReadBody exercises CircuitTunnelResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCircuitTunnelResource_Read_APIErrorReadBody(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientReadErrorBody(t, 501)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestCircuitTunnelResource_Read_InvalidJSON exercises CircuitTunnelResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestCircuitTunnelResource_Read_InvalidJSON(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 200, "{{")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestCircuitTunnelResource_Read_MapError exercises CircuitTunnelResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestCircuitTunnelResource_Read_MapError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestCircuitTunnelResource_Update_Happy exercises CircuitTunnelResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestCircuitTunnelResource_Update_Happy(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 200, "{}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCircuitTunnelResource_Update_NilClient exercises CircuitTunnelResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCircuitTunnelResource_Update_NilClient(t *testing.T) {
	r := &CircuitTunnelResource{}
	m := CircuitTunnelResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCircuitTunnelResource_Update_BuildError exercises CircuitTunnelResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCircuitTunnelResource_Update_BuildError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMalformedBaseURLClient(t)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCircuitTunnelResource_Update_SendError exercises CircuitTunnelResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestCircuitTunnelResource_Update_SendError(t *testing.T) {
	r := &CircuitTunnelResource{client: newTransportErrorClient(t)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCircuitTunnelResource_Update_APIError exercises CircuitTunnelResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCircuitTunnelResource_Update_APIError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_circuit_tunnel")
}

// TestCircuitTunnelResource_Update_APIErrorReadBody exercises CircuitTunnelResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCircuitTunnelResource_Update_APIErrorReadBody(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientReadErrorBody(t, 501)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestCircuitTunnelResource_Update_InvalidJSON exercises CircuitTunnelResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestCircuitTunnelResource_Update_InvalidJSON(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 200, "{{")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestCircuitTunnelResource_Update_MapError exercises CircuitTunnelResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestCircuitTunnelResource_Update_MapError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestCircuitTunnelResource_Delete_Happy exercises CircuitTunnelResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestCircuitTunnelResource_Delete_Happy(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 204, "")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCircuitTunnelResource_Delete_NilClient exercises CircuitTunnelResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCircuitTunnelResource_Delete_NilClient(t *testing.T) {
	r := &CircuitTunnelResource{}
	m := CircuitTunnelResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCircuitTunnelResource_Delete_BuildError exercises CircuitTunnelResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCircuitTunnelResource_Delete_BuildError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMalformedBaseURLClient(t)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCircuitTunnelResource_Delete_SendError exercises CircuitTunnelResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestCircuitTunnelResource_Delete_SendError(t *testing.T) {
	r := &CircuitTunnelResource{client: newTransportErrorClient(t)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCircuitTunnelResource_Delete_NotFoundSuccess exercises CircuitTunnelResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestCircuitTunnelResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 404, "")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCircuitTunnelResource_Delete_APIError exercises CircuitTunnelResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCircuitTunnelResource_Delete_APIError(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CircuitTunnelResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_circuit_tunnel")
}

// TestCircuitTunnelResource_Delete_APIErrorReadBody exercises CircuitTunnelResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCircuitTunnelResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &CircuitTunnelResource{client: newMockClientReadErrorBody(t, 501)}
	m := CircuitTunnelResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
