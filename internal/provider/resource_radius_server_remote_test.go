package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestRadiusServerResource_Create_Happy exercises RadiusServerResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestRadiusServerResource_Create_Happy(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 201, "{\"server_address\":\"example-id\"}")}
	m := RadiusServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRadiusServerResource_Create_NilClient exercises RadiusServerResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRadiusServerResource_Create_NilClient(t *testing.T) {
	r := &RadiusServerResource{}
	m := RadiusServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRadiusServerResource_Create_BuildError exercises RadiusServerResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRadiusServerResource_Create_BuildError(t *testing.T) {
	r := &RadiusServerResource{client: newMalformedBaseURLClient(t)}
	m := RadiusServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRadiusServerResource_Create_SendError exercises RadiusServerResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestRadiusServerResource_Create_SendError(t *testing.T) {
	r := &RadiusServerResource{client: newTransportErrorClient(t)}
	m := RadiusServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRadiusServerResource_Create_APIError exercises RadiusServerResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRadiusServerResource_Create_APIError(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RadiusServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_radius_server")
}

// TestRadiusServerResource_Create_APIErrorReadBody exercises RadiusServerResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRadiusServerResource_Create_APIErrorReadBody(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := RadiusServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRadiusServerResource_Create_InvalidJSON exercises RadiusServerResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRadiusServerResource_Create_InvalidJSON(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 201, "{{")}
	m := RadiusServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRadiusServerResource_Create_MapError exercises RadiusServerResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRadiusServerResource_Create_MapError(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 201, "{\"server_address\":12345}")}
	m := RadiusServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRadiusServerResource_Create_MissingID exercises RadiusServerResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestRadiusServerResource_Create_MissingID(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 201, "{}")}
	m := RadiusServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestRadiusServerResource_Create_LocationFallback exercises RadiusServerResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestRadiusServerResource_Create_LocationFallback(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := RadiusServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.ServerAddress.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.ServerAddress.ValueString(), "example-id")
	}
}

// TestRadiusServerResource_Read_Happy exercises RadiusServerResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestRadiusServerResource_Read_Happy(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := RadiusServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestRadiusServerResource_Read_NilClient exercises RadiusServerResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRadiusServerResource_Read_NilClient(t *testing.T) {
	r := &RadiusServerResource{}
	m := RadiusServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRadiusServerResource_Read_BuildError exercises RadiusServerResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRadiusServerResource_Read_BuildError(t *testing.T) {
	r := &RadiusServerResource{client: newMalformedBaseURLClient(t)}
	m := RadiusServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRadiusServerResource_Read_SendError exercises RadiusServerResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestRadiusServerResource_Read_SendError(t *testing.T) {
	r := &RadiusServerResource{client: newTransportErrorClient(t)}
	m := RadiusServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRadiusServerResource_Read_NotFound exercises RadiusServerResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestRadiusServerResource_Read_NotFound(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 404, "")}
	m := RadiusServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestRadiusServerResource_Read_APIError exercises RadiusServerResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRadiusServerResource_Read_APIError(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RadiusServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_radius_server")
}

// TestRadiusServerResource_Read_APIErrorReadBody exercises RadiusServerResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRadiusServerResource_Read_APIErrorReadBody(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := RadiusServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRadiusServerResource_Read_InvalidJSON exercises RadiusServerResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRadiusServerResource_Read_InvalidJSON(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := RadiusServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRadiusServerResource_Read_MapError exercises RadiusServerResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRadiusServerResource_Read_MapError(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 200, "{\"server_address\":12345}")}
	m := RadiusServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRadiusServerResource_Update_Happy exercises RadiusServerResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestRadiusServerResource_Update_Happy(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := RadiusServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRadiusServerResource_Update_NilClient exercises RadiusServerResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRadiusServerResource_Update_NilClient(t *testing.T) {
	r := &RadiusServerResource{}
	m := RadiusServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRadiusServerResource_Update_BuildError exercises RadiusServerResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRadiusServerResource_Update_BuildError(t *testing.T) {
	r := &RadiusServerResource{client: newMalformedBaseURLClient(t)}
	m := RadiusServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRadiusServerResource_Update_SendError exercises RadiusServerResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestRadiusServerResource_Update_SendError(t *testing.T) {
	r := &RadiusServerResource{client: newTransportErrorClient(t)}
	m := RadiusServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRadiusServerResource_Update_APIError exercises RadiusServerResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRadiusServerResource_Update_APIError(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RadiusServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_radius_server")
}

// TestRadiusServerResource_Update_APIErrorReadBody exercises RadiusServerResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRadiusServerResource_Update_APIErrorReadBody(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := RadiusServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRadiusServerResource_Update_InvalidJSON exercises RadiusServerResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRadiusServerResource_Update_InvalidJSON(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := RadiusServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRadiusServerResource_Update_MapError exercises RadiusServerResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRadiusServerResource_Update_MapError(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 200, "{\"server_address\":12345}")}
	m := RadiusServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRadiusServerResource_Delete_Happy exercises RadiusServerResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestRadiusServerResource_Delete_Happy(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 204, "")}
	m := RadiusServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRadiusServerResource_Delete_NilClient exercises RadiusServerResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRadiusServerResource_Delete_NilClient(t *testing.T) {
	r := &RadiusServerResource{}
	m := RadiusServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRadiusServerResource_Delete_BuildError exercises RadiusServerResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRadiusServerResource_Delete_BuildError(t *testing.T) {
	r := &RadiusServerResource{client: newMalformedBaseURLClient(t)}
	m := RadiusServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRadiusServerResource_Delete_SendError exercises RadiusServerResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestRadiusServerResource_Delete_SendError(t *testing.T) {
	r := &RadiusServerResource{client: newTransportErrorClient(t)}
	m := RadiusServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRadiusServerResource_Delete_NotFoundSuccess exercises RadiusServerResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestRadiusServerResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 404, "")}
	m := RadiusServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRadiusServerResource_Delete_APIError exercises RadiusServerResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRadiusServerResource_Delete_APIError(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RadiusServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_radius_server")
}

// TestRadiusServerResource_Delete_APIErrorReadBody exercises RadiusServerResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRadiusServerResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &RadiusServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := RadiusServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
