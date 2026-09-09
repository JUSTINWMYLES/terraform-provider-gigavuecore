package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFabricMapResource_Create_Happy exercises FabricMapResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestFabricMapResource_Create_Happy(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := FabricMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricMapResource_Create_NilClient exercises FabricMapResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFabricMapResource_Create_NilClient(t *testing.T) {
	r := &FabricMapResource{}
	m := FabricMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFabricMapResource_Create_BuildError exercises FabricMapResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFabricMapResource_Create_BuildError(t *testing.T) {
	r := &FabricMapResource{client: newMalformedBaseURLClient(t)}
	m := FabricMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFabricMapResource_Create_SendError exercises FabricMapResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestFabricMapResource_Create_SendError(t *testing.T) {
	r := &FabricMapResource{client: newTransportErrorClient(t)}
	m := FabricMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFabricMapResource_Create_APIError exercises FabricMapResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFabricMapResource_Create_APIError(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FabricMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_fabric_map")
}

// TestFabricMapResource_Create_APIErrorReadBody exercises FabricMapResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFabricMapResource_Create_APIErrorReadBody(t *testing.T) {
	r := &FabricMapResource{client: newMockClientReadErrorBody(t, 501)}
	m := FabricMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFabricMapResource_Create_InvalidJSON exercises FabricMapResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFabricMapResource_Create_InvalidJSON(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 201, "{{")}
	m := FabricMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFabricMapResource_Create_MapError exercises FabricMapResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFabricMapResource_Create_MapError(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := FabricMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFabricMapResource_Create_MissingID exercises FabricMapResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestFabricMapResource_Create_MissingID(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 201, "{}")}
	m := FabricMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestFabricMapResource_Create_LocationFallback exercises FabricMapResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestFabricMapResource_Create_LocationFallback(t *testing.T) {
	r := &FabricMapResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := FabricMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestFabricMapResource_Read_Happy exercises FabricMapResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestFabricMapResource_Read_Happy(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 200, "{}")}
	m := FabricMapResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricMapResource_Read_NilClient exercises FabricMapResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFabricMapResource_Read_NilClient(t *testing.T) {
	r := &FabricMapResource{}
	m := FabricMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFabricMapResource_Read_BuildError exercises FabricMapResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFabricMapResource_Read_BuildError(t *testing.T) {
	r := &FabricMapResource{client: newMalformedBaseURLClient(t)}
	m := FabricMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFabricMapResource_Read_SendError exercises FabricMapResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestFabricMapResource_Read_SendError(t *testing.T) {
	r := &FabricMapResource{client: newTransportErrorClient(t)}
	m := FabricMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFabricMapResource_Read_NotFound exercises FabricMapResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestFabricMapResource_Read_NotFound(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 404, "")}
	m := FabricMapResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricMapResource_Read_APIError exercises FabricMapResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFabricMapResource_Read_APIError(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FabricMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_fabric_map")
}

// TestFabricMapResource_Read_APIErrorReadBody exercises FabricMapResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFabricMapResource_Read_APIErrorReadBody(t *testing.T) {
	r := &FabricMapResource{client: newMockClientReadErrorBody(t, 501)}
	m := FabricMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFabricMapResource_Read_InvalidJSON exercises FabricMapResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFabricMapResource_Read_InvalidJSON(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 200, "{{")}
	m := FabricMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFabricMapResource_Read_MapError exercises FabricMapResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFabricMapResource_Read_MapError(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := FabricMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFabricMapResource_Update_Happy exercises FabricMapResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestFabricMapResource_Update_Happy(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 200, "{}")}
	m := FabricMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricMapResource_Update_NilClient exercises FabricMapResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFabricMapResource_Update_NilClient(t *testing.T) {
	r := &FabricMapResource{}
	m := FabricMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFabricMapResource_Update_BuildError exercises FabricMapResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFabricMapResource_Update_BuildError(t *testing.T) {
	r := &FabricMapResource{client: newMalformedBaseURLClient(t)}
	m := FabricMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFabricMapResource_Update_SendError exercises FabricMapResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestFabricMapResource_Update_SendError(t *testing.T) {
	r := &FabricMapResource{client: newTransportErrorClient(t)}
	m := FabricMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFabricMapResource_Update_APIError exercises FabricMapResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFabricMapResource_Update_APIError(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FabricMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_fabric_map")
}

// TestFabricMapResource_Update_APIErrorReadBody exercises FabricMapResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFabricMapResource_Update_APIErrorReadBody(t *testing.T) {
	r := &FabricMapResource{client: newMockClientReadErrorBody(t, 501)}
	m := FabricMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFabricMapResource_Update_InvalidJSON exercises FabricMapResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFabricMapResource_Update_InvalidJSON(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 200, "{{")}
	m := FabricMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFabricMapResource_Update_MapError exercises FabricMapResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFabricMapResource_Update_MapError(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := FabricMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFabricMapResource_Delete_Happy exercises FabricMapResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestFabricMapResource_Delete_Happy(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 204, "")}
	m := FabricMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricMapResource_Delete_NilClient exercises FabricMapResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFabricMapResource_Delete_NilClient(t *testing.T) {
	r := &FabricMapResource{}
	m := FabricMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFabricMapResource_Delete_BuildError exercises FabricMapResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFabricMapResource_Delete_BuildError(t *testing.T) {
	r := &FabricMapResource{client: newMalformedBaseURLClient(t)}
	m := FabricMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFabricMapResource_Delete_SendError exercises FabricMapResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestFabricMapResource_Delete_SendError(t *testing.T) {
	r := &FabricMapResource{client: newTransportErrorClient(t)}
	m := FabricMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFabricMapResource_Delete_NotFoundSuccess exercises FabricMapResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestFabricMapResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 404, "")}
	m := FabricMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricMapResource_Delete_APIError exercises FabricMapResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFabricMapResource_Delete_APIError(t *testing.T) {
	r := &FabricMapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FabricMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_fabric_map")
}

// TestFabricMapResource_Delete_APIErrorReadBody exercises FabricMapResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFabricMapResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &FabricMapResource{client: newMockClientReadErrorBody(t, 501)}
	m := FabricMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
