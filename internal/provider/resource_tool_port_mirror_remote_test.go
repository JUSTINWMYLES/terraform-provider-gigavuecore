package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestToolPortMirrorResource_Create_Happy exercises ToolPortMirrorResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestToolPortMirrorResource_Create_Happy(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolPortMirrorResource_Create_NilClient exercises ToolPortMirrorResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolPortMirrorResource_Create_NilClient(t *testing.T) {
	r := &ToolPortMirrorResource{}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolPortMirrorResource_Create_BuildError exercises ToolPortMirrorResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolPortMirrorResource_Create_BuildError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMalformedBaseURLClient(t)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolPortMirrorResource_Create_SendError exercises ToolPortMirrorResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolPortMirrorResource_Create_SendError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newTransportErrorClient(t)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolPortMirrorResource_Create_APIError exercises ToolPortMirrorResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolPortMirrorResource_Create_APIError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_tool_port_mirror")
}

// TestToolPortMirrorResource_Create_APIErrorReadBody exercises ToolPortMirrorResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolPortMirrorResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestToolPortMirrorResource_Create_InvalidJSON exercises ToolPortMirrorResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestToolPortMirrorResource_Create_InvalidJSON(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 201, "{{")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestToolPortMirrorResource_Create_MapError exercises ToolPortMirrorResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestToolPortMirrorResource_Create_MapError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestToolPortMirrorResource_Create_MissingID exercises ToolPortMirrorResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestToolPortMirrorResource_Create_MissingID(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 201, "{}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestToolPortMirrorResource_Create_LocationFallback exercises ToolPortMirrorResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestToolPortMirrorResource_Create_LocationFallback(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestToolPortMirrorResource_Read_Happy exercises ToolPortMirrorResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestToolPortMirrorResource_Read_Happy(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 200, "{}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolPortMirrorResource_Read_NilClient exercises ToolPortMirrorResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolPortMirrorResource_Read_NilClient(t *testing.T) {
	r := &ToolPortMirrorResource{}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolPortMirrorResource_Read_BuildError exercises ToolPortMirrorResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolPortMirrorResource_Read_BuildError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMalformedBaseURLClient(t)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolPortMirrorResource_Read_SendError exercises ToolPortMirrorResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolPortMirrorResource_Read_SendError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newTransportErrorClient(t)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolPortMirrorResource_Read_NotFound exercises ToolPortMirrorResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestToolPortMirrorResource_Read_NotFound(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 404, "")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolPortMirrorResource_Read_APIError exercises ToolPortMirrorResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolPortMirrorResource_Read_APIError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_tool_port_mirror")
}

// TestToolPortMirrorResource_Read_APIErrorReadBody exercises ToolPortMirrorResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolPortMirrorResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestToolPortMirrorResource_Read_InvalidJSON exercises ToolPortMirrorResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestToolPortMirrorResource_Read_InvalidJSON(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 200, "{{")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestToolPortMirrorResource_Read_MapError exercises ToolPortMirrorResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestToolPortMirrorResource_Read_MapError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestToolPortMirrorResource_Update_Happy exercises ToolPortMirrorResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestToolPortMirrorResource_Update_Happy(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 200, "{}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolPortMirrorResource_Update_NilClient exercises ToolPortMirrorResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolPortMirrorResource_Update_NilClient(t *testing.T) {
	r := &ToolPortMirrorResource{}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolPortMirrorResource_Update_BuildError exercises ToolPortMirrorResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolPortMirrorResource_Update_BuildError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMalformedBaseURLClient(t)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolPortMirrorResource_Update_SendError exercises ToolPortMirrorResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolPortMirrorResource_Update_SendError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newTransportErrorClient(t)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolPortMirrorResource_Update_APIError exercises ToolPortMirrorResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolPortMirrorResource_Update_APIError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_tool_port_mirror")
}

// TestToolPortMirrorResource_Update_APIErrorReadBody exercises ToolPortMirrorResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolPortMirrorResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestToolPortMirrorResource_Update_InvalidJSON exercises ToolPortMirrorResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestToolPortMirrorResource_Update_InvalidJSON(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 200, "{{")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestToolPortMirrorResource_Update_MapError exercises ToolPortMirrorResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestToolPortMirrorResource_Update_MapError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestToolPortMirrorResource_Delete_Happy exercises ToolPortMirrorResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestToolPortMirrorResource_Delete_Happy(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 204, "")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolPortMirrorResource_Delete_NilClient exercises ToolPortMirrorResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolPortMirrorResource_Delete_NilClient(t *testing.T) {
	r := &ToolPortMirrorResource{}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolPortMirrorResource_Delete_BuildError exercises ToolPortMirrorResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolPortMirrorResource_Delete_BuildError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMalformedBaseURLClient(t)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolPortMirrorResource_Delete_SendError exercises ToolPortMirrorResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolPortMirrorResource_Delete_SendError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newTransportErrorClient(t)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolPortMirrorResource_Delete_NotFoundSuccess exercises ToolPortMirrorResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestToolPortMirrorResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 404, "")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolPortMirrorResource_Delete_APIError exercises ToolPortMirrorResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolPortMirrorResource_Delete_APIError(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_tool_port_mirror")
}

// TestToolPortMirrorResource_Delete_APIErrorReadBody exercises ToolPortMirrorResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolPortMirrorResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ToolPortMirrorResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolPortMirrorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
