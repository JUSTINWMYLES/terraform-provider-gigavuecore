package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNetworkGroupResource_Create_Happy exercises NetworkGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestNetworkGroupResource_Create_Happy(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkGroupResource_Create_NilClient exercises NetworkGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkGroupResource_Create_NilClient(t *testing.T) {
	r := &NetworkGroupResource{}
	m := NetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkGroupResource_Create_BuildError exercises NetworkGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkGroupResource_Create_BuildError(t *testing.T) {
	r := &NetworkGroupResource{client: newMalformedBaseURLClient(t)}
	m := NetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkGroupResource_Create_SendError exercises NetworkGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkGroupResource_Create_SendError(t *testing.T) {
	r := &NetworkGroupResource{client: newTransportErrorClient(t)}
	m := NetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkGroupResource_Create_APIError exercises NetworkGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkGroupResource_Create_APIError(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_network_group")
}

// TestNetworkGroupResource_Create_APIErrorReadBody exercises NetworkGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetworkGroupResource_Create_InvalidJSON exercises NetworkGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetworkGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := NetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetworkGroupResource_Create_MapError exercises NetworkGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetworkGroupResource_Create_MapError(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetworkGroupResource_Create_MissingID exercises NetworkGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestNetworkGroupResource_Create_MissingID(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestNetworkGroupResource_Create_LocationFallback exercises NetworkGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestNetworkGroupResource_Create_LocationFallback(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestNetworkGroupResource_Read_Happy exercises NetworkGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestNetworkGroupResource_Read_Happy(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkGroupResource_Read_NilClient exercises NetworkGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkGroupResource_Read_NilClient(t *testing.T) {
	r := &NetworkGroupResource{}
	m := NetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkGroupResource_Read_BuildError exercises NetworkGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkGroupResource_Read_BuildError(t *testing.T) {
	r := &NetworkGroupResource{client: newMalformedBaseURLClient(t)}
	m := NetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkGroupResource_Read_SendError exercises NetworkGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkGroupResource_Read_SendError(t *testing.T) {
	r := &NetworkGroupResource{client: newTransportErrorClient(t)}
	m := NetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkGroupResource_Read_NotFound exercises NetworkGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestNetworkGroupResource_Read_NotFound(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 404, "")}
	m := NetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkGroupResource_Read_APIError exercises NetworkGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkGroupResource_Read_APIError(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_network_group")
}

// TestNetworkGroupResource_Read_APIErrorReadBody exercises NetworkGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetworkGroupResource_Read_InvalidJSON exercises NetworkGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetworkGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := NetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetworkGroupResource_Read_MapError exercises NetworkGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetworkGroupResource_Read_MapError(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetworkGroupResource_Update_Happy exercises NetworkGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestNetworkGroupResource_Update_Happy(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkGroupResource_Update_NilClient exercises NetworkGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkGroupResource_Update_NilClient(t *testing.T) {
	r := &NetworkGroupResource{}
	m := NetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkGroupResource_Update_BuildError exercises NetworkGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkGroupResource_Update_BuildError(t *testing.T) {
	r := &NetworkGroupResource{client: newMalformedBaseURLClient(t)}
	m := NetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkGroupResource_Update_SendError exercises NetworkGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkGroupResource_Update_SendError(t *testing.T) {
	r := &NetworkGroupResource{client: newTransportErrorClient(t)}
	m := NetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkGroupResource_Update_APIError exercises NetworkGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkGroupResource_Update_APIError(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_network_group")
}

// TestNetworkGroupResource_Update_APIErrorReadBody exercises NetworkGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetworkGroupResource_Update_InvalidJSON exercises NetworkGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetworkGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := NetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetworkGroupResource_Update_MapError exercises NetworkGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetworkGroupResource_Update_MapError(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetworkGroupResource_Delete_Happy exercises NetworkGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestNetworkGroupResource_Delete_Happy(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 204, "")}
	m := NetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkGroupResource_Delete_NilClient exercises NetworkGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkGroupResource_Delete_NilClient(t *testing.T) {
	r := &NetworkGroupResource{}
	m := NetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkGroupResource_Delete_BuildError exercises NetworkGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkGroupResource_Delete_BuildError(t *testing.T) {
	r := &NetworkGroupResource{client: newMalformedBaseURLClient(t)}
	m := NetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkGroupResource_Delete_SendError exercises NetworkGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkGroupResource_Delete_SendError(t *testing.T) {
	r := &NetworkGroupResource{client: newTransportErrorClient(t)}
	m := NetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkGroupResource_Delete_NotFoundSuccess exercises NetworkGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestNetworkGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 404, "")}
	m := NetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkGroupResource_Delete_APIError exercises NetworkGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkGroupResource_Delete_APIError(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_network_group")
}

// TestNetworkGroupResource_Delete_APIErrorReadBody exercises NetworkGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &NetworkGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
