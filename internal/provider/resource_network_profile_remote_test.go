package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNetworkProfileResource_Create_Happy exercises NetworkProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestNetworkProfileResource_Create_Happy(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkProfileResource_Create_NilClient exercises NetworkProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkProfileResource_Create_NilClient(t *testing.T) {
	r := &NetworkProfileResource{}
	m := NetworkProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkProfileResource_Create_BuildError exercises NetworkProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkProfileResource_Create_BuildError(t *testing.T) {
	r := &NetworkProfileResource{client: newMalformedBaseURLClient(t)}
	m := NetworkProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkProfileResource_Create_SendError exercises NetworkProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkProfileResource_Create_SendError(t *testing.T) {
	r := &NetworkProfileResource{client: newTransportErrorClient(t)}
	m := NetworkProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkProfileResource_Create_APIError exercises NetworkProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkProfileResource_Create_APIError(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_network_profile")
}

// TestNetworkProfileResource_Create_APIErrorReadBody exercises NetworkProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetworkProfileResource_Create_InvalidJSON exercises NetworkProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetworkProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := NetworkProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetworkProfileResource_Create_MapError exercises NetworkProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetworkProfileResource_Create_MapError(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetworkProfileResource_Create_MissingID exercises NetworkProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestNetworkProfileResource_Create_MissingID(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestNetworkProfileResource_Create_LocationFallback exercises NetworkProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestNetworkProfileResource_Create_LocationFallback(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestNetworkProfileResource_Read_Happy exercises NetworkProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestNetworkProfileResource_Read_Happy(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkProfileResource_Read_NilClient exercises NetworkProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkProfileResource_Read_NilClient(t *testing.T) {
	r := &NetworkProfileResource{}
	m := NetworkProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkProfileResource_Read_BuildError exercises NetworkProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkProfileResource_Read_BuildError(t *testing.T) {
	r := &NetworkProfileResource{client: newMalformedBaseURLClient(t)}
	m := NetworkProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkProfileResource_Read_SendError exercises NetworkProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkProfileResource_Read_SendError(t *testing.T) {
	r := &NetworkProfileResource{client: newTransportErrorClient(t)}
	m := NetworkProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkProfileResource_Read_NotFound exercises NetworkProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestNetworkProfileResource_Read_NotFound(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 404, "")}
	m := NetworkProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkProfileResource_Read_APIError exercises NetworkProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkProfileResource_Read_APIError(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_network_profile")
}

// TestNetworkProfileResource_Read_APIErrorReadBody exercises NetworkProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetworkProfileResource_Read_InvalidJSON exercises NetworkProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetworkProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := NetworkProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetworkProfileResource_Read_MapError exercises NetworkProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetworkProfileResource_Read_MapError(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetworkProfileResource_Update_Happy exercises NetworkProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestNetworkProfileResource_Update_Happy(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkProfileResource_Update_NilClient exercises NetworkProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkProfileResource_Update_NilClient(t *testing.T) {
	r := &NetworkProfileResource{}
	m := NetworkProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkProfileResource_Update_BuildError exercises NetworkProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkProfileResource_Update_BuildError(t *testing.T) {
	r := &NetworkProfileResource{client: newMalformedBaseURLClient(t)}
	m := NetworkProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkProfileResource_Update_SendError exercises NetworkProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkProfileResource_Update_SendError(t *testing.T) {
	r := &NetworkProfileResource{client: newTransportErrorClient(t)}
	m := NetworkProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkProfileResource_Update_APIError exercises NetworkProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkProfileResource_Update_APIError(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_network_profile")
}

// TestNetworkProfileResource_Update_APIErrorReadBody exercises NetworkProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetworkProfileResource_Update_InvalidJSON exercises NetworkProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetworkProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := NetworkProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetworkProfileResource_Update_MapError exercises NetworkProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetworkProfileResource_Update_MapError(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetworkProfileResource_Delete_Happy exercises NetworkProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestNetworkProfileResource_Delete_Happy(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 204, "")}
	m := NetworkProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkProfileResource_Delete_NilClient exercises NetworkProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkProfileResource_Delete_NilClient(t *testing.T) {
	r := &NetworkProfileResource{}
	m := NetworkProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkProfileResource_Delete_BuildError exercises NetworkProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkProfileResource_Delete_BuildError(t *testing.T) {
	r := &NetworkProfileResource{client: newMalformedBaseURLClient(t)}
	m := NetworkProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkProfileResource_Delete_SendError exercises NetworkProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkProfileResource_Delete_SendError(t *testing.T) {
	r := &NetworkProfileResource{client: newTransportErrorClient(t)}
	m := NetworkProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkProfileResource_Delete_NotFoundSuccess exercises NetworkProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestNetworkProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 404, "")}
	m := NetworkProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkProfileResource_Delete_APIError exercises NetworkProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkProfileResource_Delete_APIError(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_network_profile")
}

// TestNetworkProfileResource_Delete_APIErrorReadBody exercises NetworkProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &NetworkProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
