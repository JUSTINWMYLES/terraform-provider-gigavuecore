package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestImageServerResource_Create_Happy exercises ImageServerResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestImageServerResource_Create_Happy(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ImageServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestImageServerResource_Create_NilClient exercises ImageServerResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestImageServerResource_Create_NilClient(t *testing.T) {
	r := &ImageServerResource{}
	m := ImageServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestImageServerResource_Create_BuildError exercises ImageServerResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestImageServerResource_Create_BuildError(t *testing.T) {
	r := &ImageServerResource{client: newMalformedBaseURLClient(t)}
	m := ImageServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestImageServerResource_Create_SendError exercises ImageServerResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestImageServerResource_Create_SendError(t *testing.T) {
	r := &ImageServerResource{client: newTransportErrorClient(t)}
	m := ImageServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestImageServerResource_Create_APIError exercises ImageServerResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestImageServerResource_Create_APIError(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ImageServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_image_server")
}

// TestImageServerResource_Create_APIErrorReadBody exercises ImageServerResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestImageServerResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ImageServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ImageServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestImageServerResource_Create_InvalidJSON exercises ImageServerResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestImageServerResource_Create_InvalidJSON(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 201, "{{")}
	m := ImageServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestImageServerResource_Create_MapError exercises ImageServerResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestImageServerResource_Create_MapError(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ImageServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestImageServerResource_Create_MissingID exercises ImageServerResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestImageServerResource_Create_MissingID(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 201, "{}")}
	m := ImageServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestImageServerResource_Create_LocationFallback exercises ImageServerResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestImageServerResource_Create_LocationFallback(t *testing.T) {
	r := &ImageServerResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ImageServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestImageServerResource_Read_Happy exercises ImageServerResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestImageServerResource_Read_Happy(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := ImageServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestImageServerResource_Read_NilClient exercises ImageServerResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestImageServerResource_Read_NilClient(t *testing.T) {
	r := &ImageServerResource{}
	m := ImageServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestImageServerResource_Read_BuildError exercises ImageServerResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestImageServerResource_Read_BuildError(t *testing.T) {
	r := &ImageServerResource{client: newMalformedBaseURLClient(t)}
	m := ImageServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestImageServerResource_Read_SendError exercises ImageServerResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestImageServerResource_Read_SendError(t *testing.T) {
	r := &ImageServerResource{client: newTransportErrorClient(t)}
	m := ImageServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestImageServerResource_Read_NotFound exercises ImageServerResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestImageServerResource_Read_NotFound(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 404, "")}
	m := ImageServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestImageServerResource_Read_APIError exercises ImageServerResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestImageServerResource_Read_APIError(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ImageServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_image_server")
}

// TestImageServerResource_Read_APIErrorReadBody exercises ImageServerResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestImageServerResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ImageServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ImageServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestImageServerResource_Read_InvalidJSON exercises ImageServerResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestImageServerResource_Read_InvalidJSON(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := ImageServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestImageServerResource_Read_MapError exercises ImageServerResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestImageServerResource_Read_MapError(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ImageServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestImageServerResource_Update_Happy exercises ImageServerResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestImageServerResource_Update_Happy(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := ImageServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestImageServerResource_Update_NilClient exercises ImageServerResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestImageServerResource_Update_NilClient(t *testing.T) {
	r := &ImageServerResource{}
	m := ImageServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestImageServerResource_Update_BuildError exercises ImageServerResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestImageServerResource_Update_BuildError(t *testing.T) {
	r := &ImageServerResource{client: newMalformedBaseURLClient(t)}
	m := ImageServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestImageServerResource_Update_SendError exercises ImageServerResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestImageServerResource_Update_SendError(t *testing.T) {
	r := &ImageServerResource{client: newTransportErrorClient(t)}
	m := ImageServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestImageServerResource_Update_APIError exercises ImageServerResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestImageServerResource_Update_APIError(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ImageServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_image_server")
}

// TestImageServerResource_Update_APIErrorReadBody exercises ImageServerResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestImageServerResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ImageServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ImageServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestImageServerResource_Update_InvalidJSON exercises ImageServerResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestImageServerResource_Update_InvalidJSON(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := ImageServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestImageServerResource_Update_MapError exercises ImageServerResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestImageServerResource_Update_MapError(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ImageServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestImageServerResource_Delete_Happy exercises ImageServerResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestImageServerResource_Delete_Happy(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 204, "")}
	m := ImageServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestImageServerResource_Delete_NilClient exercises ImageServerResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestImageServerResource_Delete_NilClient(t *testing.T) {
	r := &ImageServerResource{}
	m := ImageServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestImageServerResource_Delete_BuildError exercises ImageServerResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestImageServerResource_Delete_BuildError(t *testing.T) {
	r := &ImageServerResource{client: newMalformedBaseURLClient(t)}
	m := ImageServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestImageServerResource_Delete_SendError exercises ImageServerResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestImageServerResource_Delete_SendError(t *testing.T) {
	r := &ImageServerResource{client: newTransportErrorClient(t)}
	m := ImageServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestImageServerResource_Delete_NotFoundSuccess exercises ImageServerResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestImageServerResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 404, "")}
	m := ImageServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestImageServerResource_Delete_APIError exercises ImageServerResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestImageServerResource_Delete_APIError(t *testing.T) {
	r := &ImageServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ImageServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_image_server")
}

// TestImageServerResource_Delete_APIErrorReadBody exercises ImageServerResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestImageServerResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ImageServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ImageServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
