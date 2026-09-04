package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestToolResource_Create_Happy exercises ToolResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestToolResource_Create_Happy(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ToolResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolResource_Create_NilClient exercises ToolResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolResource_Create_NilClient(t *testing.T) {
	r := &ToolResource{}
	m := ToolResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolResource_Create_BuildError exercises ToolResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolResource_Create_BuildError(t *testing.T) {
	r := &ToolResource{client: newMalformedBaseURLClient(t)}
	m := ToolResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolResource_Create_SendError exercises ToolResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolResource_Create_SendError(t *testing.T) {
	r := &ToolResource{client: newTransportErrorClient(t)}
	m := ToolResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolResource_Create_APIError exercises ToolResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolResource_Create_APIError(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_tool")
}

// TestToolResource_Create_APIErrorReadBody exercises ToolResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ToolResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestToolResource_Create_InvalidJSON exercises ToolResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestToolResource_Create_InvalidJSON(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 201, "{{")}
	m := ToolResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestToolResource_Create_MapError exercises ToolResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestToolResource_Create_MapError(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ToolResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestToolResource_Create_MissingID exercises ToolResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestToolResource_Create_MissingID(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 201, "{}")}
	m := ToolResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestToolResource_Create_LocationFallback exercises ToolResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestToolResource_Create_LocationFallback(t *testing.T) {
	r := &ToolResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ToolResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestToolResource_Read_Happy exercises ToolResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestToolResource_Read_Happy(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 200, "{}")}
	m := ToolResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolResource_Read_NilClient exercises ToolResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolResource_Read_NilClient(t *testing.T) {
	r := &ToolResource{}
	m := ToolResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolResource_Read_BuildError exercises ToolResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolResource_Read_BuildError(t *testing.T) {
	r := &ToolResource{client: newMalformedBaseURLClient(t)}
	m := ToolResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolResource_Read_SendError exercises ToolResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolResource_Read_SendError(t *testing.T) {
	r := &ToolResource{client: newTransportErrorClient(t)}
	m := ToolResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolResource_Read_NotFound exercises ToolResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestToolResource_Read_NotFound(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 404, "")}
	m := ToolResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolResource_Read_APIError exercises ToolResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolResource_Read_APIError(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_tool")
}

// TestToolResource_Read_APIErrorReadBody exercises ToolResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ToolResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestToolResource_Read_InvalidJSON exercises ToolResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestToolResource_Read_InvalidJSON(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 200, "{{")}
	m := ToolResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestToolResource_Read_MapError exercises ToolResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestToolResource_Read_MapError(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ToolResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestToolResource_Update_Happy exercises ToolResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestToolResource_Update_Happy(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 200, "{}")}
	m := ToolResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolResource_Update_NilClient exercises ToolResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolResource_Update_NilClient(t *testing.T) {
	r := &ToolResource{}
	m := ToolResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolResource_Update_BuildError exercises ToolResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolResource_Update_BuildError(t *testing.T) {
	r := &ToolResource{client: newMalformedBaseURLClient(t)}
	m := ToolResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolResource_Update_SendError exercises ToolResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolResource_Update_SendError(t *testing.T) {
	r := &ToolResource{client: newTransportErrorClient(t)}
	m := ToolResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolResource_Update_APIError exercises ToolResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolResource_Update_APIError(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_tool")
}

// TestToolResource_Update_APIErrorReadBody exercises ToolResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ToolResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestToolResource_Update_InvalidJSON exercises ToolResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestToolResource_Update_InvalidJSON(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 200, "{{")}
	m := ToolResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestToolResource_Update_MapError exercises ToolResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestToolResource_Update_MapError(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ToolResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestToolResource_Delete_Happy exercises ToolResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestToolResource_Delete_Happy(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 204, "")}
	m := ToolResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolResource_Delete_NilClient exercises ToolResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestToolResource_Delete_NilClient(t *testing.T) {
	r := &ToolResource{}
	m := ToolResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestToolResource_Delete_BuildError exercises ToolResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestToolResource_Delete_BuildError(t *testing.T) {
	r := &ToolResource{client: newMalformedBaseURLClient(t)}
	m := ToolResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestToolResource_Delete_SendError exercises ToolResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestToolResource_Delete_SendError(t *testing.T) {
	r := &ToolResource{client: newTransportErrorClient(t)}
	m := ToolResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestToolResource_Delete_NotFoundSuccess exercises ToolResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestToolResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 404, "")}
	m := ToolResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestToolResource_Delete_APIError exercises ToolResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestToolResource_Delete_APIError(t *testing.T) {
	r := &ToolResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ToolResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_tool")
}

// TestToolResource_Delete_APIErrorReadBody exercises ToolResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestToolResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ToolResource{client: newMockClientReadErrorBody(t, 501)}
	m := ToolResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
