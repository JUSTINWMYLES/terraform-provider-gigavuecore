package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestStackLinkResource_Create_Happy exercises StackLinkResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestStackLinkResource_Create_Happy(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := StackLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestStackLinkResource_Create_NilClient exercises StackLinkResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestStackLinkResource_Create_NilClient(t *testing.T) {
	r := &StackLinkResource{}
	m := StackLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestStackLinkResource_Create_BuildError exercises StackLinkResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestStackLinkResource_Create_BuildError(t *testing.T) {
	r := &StackLinkResource{client: newMalformedBaseURLClient(t)}
	m := StackLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestStackLinkResource_Create_SendError exercises StackLinkResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestStackLinkResource_Create_SendError(t *testing.T) {
	r := &StackLinkResource{client: newTransportErrorClient(t)}
	m := StackLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestStackLinkResource_Create_APIError exercises StackLinkResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestStackLinkResource_Create_APIError(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := StackLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_stack_link")
}

// TestStackLinkResource_Create_APIErrorReadBody exercises StackLinkResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestStackLinkResource_Create_APIErrorReadBody(t *testing.T) {
	r := &StackLinkResource{client: newMockClientReadErrorBody(t, 501)}
	m := StackLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestStackLinkResource_Create_InvalidJSON exercises StackLinkResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestStackLinkResource_Create_InvalidJSON(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 201, "{{")}
	m := StackLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestStackLinkResource_Create_MapError exercises StackLinkResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestStackLinkResource_Create_MapError(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := StackLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestStackLinkResource_Create_MissingID exercises StackLinkResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestStackLinkResource_Create_MissingID(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 201, "{}")}
	m := StackLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestStackLinkResource_Create_LocationFallback exercises StackLinkResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestStackLinkResource_Create_LocationFallback(t *testing.T) {
	r := &StackLinkResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := StackLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestStackLinkResource_Read_Happy exercises StackLinkResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestStackLinkResource_Read_Happy(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 200, "{}")}
	m := StackLinkResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestStackLinkResource_Read_NilClient exercises StackLinkResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestStackLinkResource_Read_NilClient(t *testing.T) {
	r := &StackLinkResource{}
	m := StackLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestStackLinkResource_Read_BuildError exercises StackLinkResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestStackLinkResource_Read_BuildError(t *testing.T) {
	r := &StackLinkResource{client: newMalformedBaseURLClient(t)}
	m := StackLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestStackLinkResource_Read_SendError exercises StackLinkResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestStackLinkResource_Read_SendError(t *testing.T) {
	r := &StackLinkResource{client: newTransportErrorClient(t)}
	m := StackLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestStackLinkResource_Read_NotFound exercises StackLinkResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestStackLinkResource_Read_NotFound(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 404, "")}
	m := StackLinkResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestStackLinkResource_Read_APIError exercises StackLinkResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestStackLinkResource_Read_APIError(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := StackLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_stack_link")
}

// TestStackLinkResource_Read_APIErrorReadBody exercises StackLinkResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestStackLinkResource_Read_APIErrorReadBody(t *testing.T) {
	r := &StackLinkResource{client: newMockClientReadErrorBody(t, 501)}
	m := StackLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestStackLinkResource_Read_InvalidJSON exercises StackLinkResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestStackLinkResource_Read_InvalidJSON(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 200, "{{")}
	m := StackLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestStackLinkResource_Read_MapError exercises StackLinkResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestStackLinkResource_Read_MapError(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := StackLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestStackLinkResource_Update_Happy exercises StackLinkResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestStackLinkResource_Update_Happy(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 200, "{}")}
	m := StackLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestStackLinkResource_Update_NilClient exercises StackLinkResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestStackLinkResource_Update_NilClient(t *testing.T) {
	r := &StackLinkResource{}
	m := StackLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestStackLinkResource_Update_BuildError exercises StackLinkResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestStackLinkResource_Update_BuildError(t *testing.T) {
	r := &StackLinkResource{client: newMalformedBaseURLClient(t)}
	m := StackLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestStackLinkResource_Update_SendError exercises StackLinkResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestStackLinkResource_Update_SendError(t *testing.T) {
	r := &StackLinkResource{client: newTransportErrorClient(t)}
	m := StackLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestStackLinkResource_Update_APIError exercises StackLinkResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestStackLinkResource_Update_APIError(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := StackLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_stack_link")
}

// TestStackLinkResource_Update_APIErrorReadBody exercises StackLinkResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestStackLinkResource_Update_APIErrorReadBody(t *testing.T) {
	r := &StackLinkResource{client: newMockClientReadErrorBody(t, 501)}
	m := StackLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestStackLinkResource_Update_InvalidJSON exercises StackLinkResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestStackLinkResource_Update_InvalidJSON(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 200, "{{")}
	m := StackLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestStackLinkResource_Update_MapError exercises StackLinkResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestStackLinkResource_Update_MapError(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := StackLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestStackLinkResource_Delete_Happy exercises StackLinkResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestStackLinkResource_Delete_Happy(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 204, "")}
	m := StackLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestStackLinkResource_Delete_NilClient exercises StackLinkResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestStackLinkResource_Delete_NilClient(t *testing.T) {
	r := &StackLinkResource{}
	m := StackLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestStackLinkResource_Delete_BuildError exercises StackLinkResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestStackLinkResource_Delete_BuildError(t *testing.T) {
	r := &StackLinkResource{client: newMalformedBaseURLClient(t)}
	m := StackLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestStackLinkResource_Delete_SendError exercises StackLinkResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestStackLinkResource_Delete_SendError(t *testing.T) {
	r := &StackLinkResource{client: newTransportErrorClient(t)}
	m := StackLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestStackLinkResource_Delete_NotFoundSuccess exercises StackLinkResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestStackLinkResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 404, "")}
	m := StackLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestStackLinkResource_Delete_APIError exercises StackLinkResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestStackLinkResource_Delete_APIError(t *testing.T) {
	r := &StackLinkResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := StackLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_stack_link")
}

// TestStackLinkResource_Delete_APIErrorReadBody exercises StackLinkResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestStackLinkResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &StackLinkResource{client: newMockClientReadErrorBody(t, 501)}
	m := StackLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
