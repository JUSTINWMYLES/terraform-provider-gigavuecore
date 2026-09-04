package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFabricPortGroupResource_Create_Happy exercises FabricPortGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestFabricPortGroupResource_Create_Happy(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricPortGroupResource_Create_NilClient exercises FabricPortGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFabricPortGroupResource_Create_NilClient(t *testing.T) {
	r := &FabricPortGroupResource{}
	m := FabricPortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFabricPortGroupResource_Create_BuildError exercises FabricPortGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFabricPortGroupResource_Create_BuildError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMalformedBaseURLClient(t)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFabricPortGroupResource_Create_SendError exercises FabricPortGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestFabricPortGroupResource_Create_SendError(t *testing.T) {
	r := &FabricPortGroupResource{client: newTransportErrorClient(t)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFabricPortGroupResource_Create_APIError exercises FabricPortGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFabricPortGroupResource_Create_APIError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_fabric_port_group")
}

// TestFabricPortGroupResource_Create_APIErrorReadBody exercises FabricPortGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFabricPortGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFabricPortGroupResource_Create_InvalidJSON exercises FabricPortGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFabricPortGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFabricPortGroupResource_Create_MapError exercises FabricPortGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFabricPortGroupResource_Create_MapError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFabricPortGroupResource_Create_MissingID exercises FabricPortGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestFabricPortGroupResource_Create_MissingID(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestFabricPortGroupResource_Create_LocationFallback exercises FabricPortGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestFabricPortGroupResource_Create_LocationFallback(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestFabricPortGroupResource_Read_Happy exercises FabricPortGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestFabricPortGroupResource_Read_Happy(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricPortGroupResource_Read_NilClient exercises FabricPortGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFabricPortGroupResource_Read_NilClient(t *testing.T) {
	r := &FabricPortGroupResource{}
	m := FabricPortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFabricPortGroupResource_Read_BuildError exercises FabricPortGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFabricPortGroupResource_Read_BuildError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMalformedBaseURLClient(t)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFabricPortGroupResource_Read_SendError exercises FabricPortGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestFabricPortGroupResource_Read_SendError(t *testing.T) {
	r := &FabricPortGroupResource{client: newTransportErrorClient(t)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFabricPortGroupResource_Read_NotFound exercises FabricPortGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestFabricPortGroupResource_Read_NotFound(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 404, "")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricPortGroupResource_Read_APIError exercises FabricPortGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFabricPortGroupResource_Read_APIError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_fabric_port_group")
}

// TestFabricPortGroupResource_Read_APIErrorReadBody exercises FabricPortGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFabricPortGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFabricPortGroupResource_Read_InvalidJSON exercises FabricPortGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFabricPortGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFabricPortGroupResource_Read_MapError exercises FabricPortGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFabricPortGroupResource_Read_MapError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFabricPortGroupResource_Update_Happy exercises FabricPortGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestFabricPortGroupResource_Update_Happy(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricPortGroupResource_Update_NilClient exercises FabricPortGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFabricPortGroupResource_Update_NilClient(t *testing.T) {
	r := &FabricPortGroupResource{}
	m := FabricPortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFabricPortGroupResource_Update_BuildError exercises FabricPortGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFabricPortGroupResource_Update_BuildError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMalformedBaseURLClient(t)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFabricPortGroupResource_Update_SendError exercises FabricPortGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestFabricPortGroupResource_Update_SendError(t *testing.T) {
	r := &FabricPortGroupResource{client: newTransportErrorClient(t)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFabricPortGroupResource_Update_APIError exercises FabricPortGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFabricPortGroupResource_Update_APIError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_fabric_port_group")
}

// TestFabricPortGroupResource_Update_APIErrorReadBody exercises FabricPortGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFabricPortGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFabricPortGroupResource_Update_InvalidJSON exercises FabricPortGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFabricPortGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFabricPortGroupResource_Update_MapError exercises FabricPortGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFabricPortGroupResource_Update_MapError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFabricPortGroupResource_Delete_Happy exercises FabricPortGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestFabricPortGroupResource_Delete_Happy(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 204, "")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricPortGroupResource_Delete_NilClient exercises FabricPortGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFabricPortGroupResource_Delete_NilClient(t *testing.T) {
	r := &FabricPortGroupResource{}
	m := FabricPortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFabricPortGroupResource_Delete_BuildError exercises FabricPortGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFabricPortGroupResource_Delete_BuildError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMalformedBaseURLClient(t)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFabricPortGroupResource_Delete_SendError exercises FabricPortGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestFabricPortGroupResource_Delete_SendError(t *testing.T) {
	r := &FabricPortGroupResource{client: newTransportErrorClient(t)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFabricPortGroupResource_Delete_NotFoundSuccess exercises FabricPortGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestFabricPortGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 404, "")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFabricPortGroupResource_Delete_APIError exercises FabricPortGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFabricPortGroupResource_Delete_APIError(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FabricPortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_fabric_port_group")
}

// TestFabricPortGroupResource_Delete_APIErrorReadBody exercises FabricPortGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFabricPortGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &FabricPortGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := FabricPortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
