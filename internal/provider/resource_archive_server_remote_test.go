package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestArchiveServerResource_Create_Happy exercises ArchiveServerResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestArchiveServerResource_Create_Happy(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestArchiveServerResource_Create_NilClient exercises ArchiveServerResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestArchiveServerResource_Create_NilClient(t *testing.T) {
	r := &ArchiveServerResource{}
	m := ArchiveServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestArchiveServerResource_Create_BuildError exercises ArchiveServerResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestArchiveServerResource_Create_BuildError(t *testing.T) {
	r := &ArchiveServerResource{client: newMalformedBaseURLClient(t)}
	m := ArchiveServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestArchiveServerResource_Create_SendError exercises ArchiveServerResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestArchiveServerResource_Create_SendError(t *testing.T) {
	r := &ArchiveServerResource{client: newTransportErrorClient(t)}
	m := ArchiveServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestArchiveServerResource_Create_APIError exercises ArchiveServerResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestArchiveServerResource_Create_APIError(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_archive_server")
}

// TestArchiveServerResource_Create_APIErrorReadBody exercises ArchiveServerResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestArchiveServerResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ArchiveServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestArchiveServerResource_Create_InvalidJSON exercises ArchiveServerResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestArchiveServerResource_Create_InvalidJSON(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 201, "{{")}
	m := ArchiveServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestArchiveServerResource_Create_MapError exercises ArchiveServerResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestArchiveServerResource_Create_MapError(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestArchiveServerResource_Create_MissingID exercises ArchiveServerResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestArchiveServerResource_Create_MissingID(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 201, "{}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestArchiveServerResource_Create_LocationFallback exercises ArchiveServerResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestArchiveServerResource_Create_LocationFallback(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestArchiveServerResource_Read_Happy exercises ArchiveServerResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestArchiveServerResource_Read_Happy(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestArchiveServerResource_Read_NilClient exercises ArchiveServerResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestArchiveServerResource_Read_NilClient(t *testing.T) {
	r := &ArchiveServerResource{}
	m := ArchiveServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestArchiveServerResource_Read_BuildError exercises ArchiveServerResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestArchiveServerResource_Read_BuildError(t *testing.T) {
	r := &ArchiveServerResource{client: newMalformedBaseURLClient(t)}
	m := ArchiveServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestArchiveServerResource_Read_SendError exercises ArchiveServerResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestArchiveServerResource_Read_SendError(t *testing.T) {
	r := &ArchiveServerResource{client: newTransportErrorClient(t)}
	m := ArchiveServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestArchiveServerResource_Read_NotFound exercises ArchiveServerResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestArchiveServerResource_Read_NotFound(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 404, "")}
	m := ArchiveServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestArchiveServerResource_Read_APIError exercises ArchiveServerResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestArchiveServerResource_Read_APIError(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_archive_server")
}

// TestArchiveServerResource_Read_APIErrorReadBody exercises ArchiveServerResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestArchiveServerResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ArchiveServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestArchiveServerResource_Read_InvalidJSON exercises ArchiveServerResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestArchiveServerResource_Read_InvalidJSON(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := ArchiveServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestArchiveServerResource_Read_MapError exercises ArchiveServerResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestArchiveServerResource_Read_MapError(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestArchiveServerResource_Update_Happy exercises ArchiveServerResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestArchiveServerResource_Update_Happy(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestArchiveServerResource_Update_NilClient exercises ArchiveServerResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestArchiveServerResource_Update_NilClient(t *testing.T) {
	r := &ArchiveServerResource{}
	m := ArchiveServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestArchiveServerResource_Update_BuildError exercises ArchiveServerResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestArchiveServerResource_Update_BuildError(t *testing.T) {
	r := &ArchiveServerResource{client: newMalformedBaseURLClient(t)}
	m := ArchiveServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestArchiveServerResource_Update_SendError exercises ArchiveServerResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestArchiveServerResource_Update_SendError(t *testing.T) {
	r := &ArchiveServerResource{client: newTransportErrorClient(t)}
	m := ArchiveServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestArchiveServerResource_Update_APIError exercises ArchiveServerResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestArchiveServerResource_Update_APIError(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_archive_server")
}

// TestArchiveServerResource_Update_APIErrorReadBody exercises ArchiveServerResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestArchiveServerResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ArchiveServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestArchiveServerResource_Update_InvalidJSON exercises ArchiveServerResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestArchiveServerResource_Update_InvalidJSON(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := ArchiveServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestArchiveServerResource_Update_MapError exercises ArchiveServerResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestArchiveServerResource_Update_MapError(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestArchiveServerResource_Delete_Happy exercises ArchiveServerResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestArchiveServerResource_Delete_Happy(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 204, "")}
	m := ArchiveServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestArchiveServerResource_Delete_NilClient exercises ArchiveServerResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestArchiveServerResource_Delete_NilClient(t *testing.T) {
	r := &ArchiveServerResource{}
	m := ArchiveServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestArchiveServerResource_Delete_BuildError exercises ArchiveServerResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestArchiveServerResource_Delete_BuildError(t *testing.T) {
	r := &ArchiveServerResource{client: newMalformedBaseURLClient(t)}
	m := ArchiveServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestArchiveServerResource_Delete_SendError exercises ArchiveServerResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestArchiveServerResource_Delete_SendError(t *testing.T) {
	r := &ArchiveServerResource{client: newTransportErrorClient(t)}
	m := ArchiveServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestArchiveServerResource_Delete_NotFoundSuccess exercises ArchiveServerResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestArchiveServerResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 404, "")}
	m := ArchiveServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestArchiveServerResource_Delete_APIError exercises ArchiveServerResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestArchiveServerResource_Delete_APIError(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ArchiveServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_archive_server")
}

// TestArchiveServerResource_Delete_APIErrorReadBody exercises ArchiveServerResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestArchiveServerResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ArchiveServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ArchiveServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
