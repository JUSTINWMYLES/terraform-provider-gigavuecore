package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestToolGroupResource_Create_Happy exercises ToolGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestToolGroupResource_Create_Happy(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolGroupResource_Create_NilClient exercises ToolGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolGroupResource_Create_NilClient(t *testing.T) {
	r := &ToolGroupResource{}
	m := ToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolGroupResource_Create_BuildError exercises ToolGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolGroupResource_Create_BuildError(t *testing.T) {
	r := &ToolGroupResource{client: newMalformedBaseURLClient(t)}
	m := ToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolGroupResource_Create_SendError exercises ToolGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolGroupResource_Create_SendError(t *testing.T) {
	r := &ToolGroupResource{client: newTransportErrorClient(t)}
	m := ToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolGroupResource_Create_APIError exercises ToolGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolGroupResource_Create_APIError(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_tool_group")
}

// TestToolGroupResource_Create_APIErrorReadBody exercises ToolGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestToolGroupResource_Create_InvalidJSON exercises ToolGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestToolGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := ToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestToolGroupResource_Create_MapError exercises ToolGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestToolGroupResource_Create_MapError(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestToolGroupResource_Create_MissingID exercises ToolGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestToolGroupResource_Create_MissingID(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := ToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestToolGroupResource_Create_LocationFallback exercises ToolGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestToolGroupResource_Create_LocationFallback(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestToolGroupResource_Read_Happy exercises ToolGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestToolGroupResource_Read_Happy(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := ToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolGroupResource_Read_NilClient exercises ToolGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolGroupResource_Read_NilClient(t *testing.T) {
	r := &ToolGroupResource{}
	m := ToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolGroupResource_Read_BuildError exercises ToolGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolGroupResource_Read_BuildError(t *testing.T) {
	r := &ToolGroupResource{client: newMalformedBaseURLClient(t)}
	m := ToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolGroupResource_Read_SendError exercises ToolGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolGroupResource_Read_SendError(t *testing.T) {
	r := &ToolGroupResource{client: newTransportErrorClient(t)}
	m := ToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolGroupResource_Read_NotFound exercises ToolGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestToolGroupResource_Read_NotFound(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 404, "")}
	m := ToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolGroupResource_Read_APIError exercises ToolGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolGroupResource_Read_APIError(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_tool_group")
}

// TestToolGroupResource_Read_APIErrorReadBody exercises ToolGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestToolGroupResource_Read_InvalidJSON exercises ToolGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestToolGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := ToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestToolGroupResource_Read_MapError exercises ToolGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestToolGroupResource_Read_MapError(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestToolGroupResource_Update_Happy exercises ToolGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestToolGroupResource_Update_Happy(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := ToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolGroupResource_Update_NilClient exercises ToolGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolGroupResource_Update_NilClient(t *testing.T) {
	r := &ToolGroupResource{}
	m := ToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolGroupResource_Update_BuildError exercises ToolGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolGroupResource_Update_BuildError(t *testing.T) {
	r := &ToolGroupResource{client: newMalformedBaseURLClient(t)}
	m := ToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolGroupResource_Update_SendError exercises ToolGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolGroupResource_Update_SendError(t *testing.T) {
	r := &ToolGroupResource{client: newTransportErrorClient(t)}
	m := ToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolGroupResource_Update_APIError exercises ToolGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolGroupResource_Update_APIError(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_tool_group")
}

// TestToolGroupResource_Update_APIErrorReadBody exercises ToolGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestToolGroupResource_Update_InvalidJSON exercises ToolGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestToolGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := ToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestToolGroupResource_Update_MapError exercises ToolGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestToolGroupResource_Update_MapError(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestToolGroupResource_Delete_Happy exercises ToolGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestToolGroupResource_Delete_Happy(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 204, "")}
	m := ToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolGroupResource_Delete_NilClient exercises ToolGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolGroupResource_Delete_NilClient(t *testing.T) {
	r := &ToolGroupResource{}
	m := ToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolGroupResource_Delete_BuildError exercises ToolGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolGroupResource_Delete_BuildError(t *testing.T) {
	r := &ToolGroupResource{client: newMalformedBaseURLClient(t)}
	m := ToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolGroupResource_Delete_SendError exercises ToolGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolGroupResource_Delete_SendError(t *testing.T) {
	r := &ToolGroupResource{client: newTransportErrorClient(t)}
	m := ToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolGroupResource_Delete_NotFoundSuccess exercises ToolGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestToolGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 404, "")}
	m := ToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolGroupResource_Delete_APIError exercises ToolGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolGroupResource_Delete_APIError(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_tool_group")
}

// TestToolGroupResource_Delete_APIErrorReadBody exercises ToolGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ToolGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
