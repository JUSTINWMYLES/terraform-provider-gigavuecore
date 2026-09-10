package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestGpfcpProfileResource_Create_Happy exercises GpfcpProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestGpfcpProfileResource_Create_Happy(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGpfcpProfileResource_Create_NilClient exercises GpfcpProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGpfcpProfileResource_Create_NilClient(t *testing.T) {
	r := &GpfcpProfileResource{}
	m := GpfcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGpfcpProfileResource_Create_BuildError exercises GpfcpProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGpfcpProfileResource_Create_BuildError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMalformedBaseURLClient(t)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGpfcpProfileResource_Create_SendError exercises GpfcpProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestGpfcpProfileResource_Create_SendError(t *testing.T) {
	r := &GpfcpProfileResource{client: newTransportErrorClient(t)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGpfcpProfileResource_Create_APIError exercises GpfcpProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGpfcpProfileResource_Create_APIError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_gpfcp_profile")
}

// TestGpfcpProfileResource_Create_APIErrorReadBody exercises GpfcpProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGpfcpProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGpfcpProfileResource_Create_InvalidJSON exercises GpfcpProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGpfcpProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGpfcpProfileResource_Create_MapError exercises GpfcpProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGpfcpProfileResource_Create_MapError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGpfcpProfileResource_Create_MissingID exercises GpfcpProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestGpfcpProfileResource_Create_MissingID(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestGpfcpProfileResource_Create_LocationFallback exercises GpfcpProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestGpfcpProfileResource_Create_LocationFallback(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestGpfcpProfileResource_Read_Happy exercises GpfcpProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestGpfcpProfileResource_Read_Happy(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGpfcpProfileResource_Read_NilClient exercises GpfcpProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGpfcpProfileResource_Read_NilClient(t *testing.T) {
	r := &GpfcpProfileResource{}
	m := GpfcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGpfcpProfileResource_Read_BuildError exercises GpfcpProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGpfcpProfileResource_Read_BuildError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMalformedBaseURLClient(t)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGpfcpProfileResource_Read_SendError exercises GpfcpProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGpfcpProfileResource_Read_SendError(t *testing.T) {
	r := &GpfcpProfileResource{client: newTransportErrorClient(t)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGpfcpProfileResource_Read_NotFound exercises GpfcpProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestGpfcpProfileResource_Read_NotFound(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 404, "")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGpfcpProfileResource_Read_APIError exercises GpfcpProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGpfcpProfileResource_Read_APIError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_gpfcp_profile")
}

// TestGpfcpProfileResource_Read_APIErrorReadBody exercises GpfcpProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGpfcpProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGpfcpProfileResource_Read_InvalidJSON exercises GpfcpProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGpfcpProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGpfcpProfileResource_Read_MapError exercises GpfcpProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGpfcpProfileResource_Read_MapError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGpfcpProfileResource_Update_Happy exercises GpfcpProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestGpfcpProfileResource_Update_Happy(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGpfcpProfileResource_Update_NilClient exercises GpfcpProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGpfcpProfileResource_Update_NilClient(t *testing.T) {
	r := &GpfcpProfileResource{}
	m := GpfcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGpfcpProfileResource_Update_BuildError exercises GpfcpProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGpfcpProfileResource_Update_BuildError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMalformedBaseURLClient(t)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGpfcpProfileResource_Update_SendError exercises GpfcpProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestGpfcpProfileResource_Update_SendError(t *testing.T) {
	r := &GpfcpProfileResource{client: newTransportErrorClient(t)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGpfcpProfileResource_Update_APIError exercises GpfcpProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGpfcpProfileResource_Update_APIError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_gpfcp_profile")
}

// TestGpfcpProfileResource_Update_APIErrorReadBody exercises GpfcpProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGpfcpProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGpfcpProfileResource_Update_InvalidJSON exercises GpfcpProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGpfcpProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGpfcpProfileResource_Update_MapError exercises GpfcpProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGpfcpProfileResource_Update_MapError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGpfcpProfileResource_Delete_Happy exercises GpfcpProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestGpfcpProfileResource_Delete_Happy(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 204, "")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGpfcpProfileResource_Delete_NilClient exercises GpfcpProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGpfcpProfileResource_Delete_NilClient(t *testing.T) {
	r := &GpfcpProfileResource{}
	m := GpfcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGpfcpProfileResource_Delete_BuildError exercises GpfcpProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGpfcpProfileResource_Delete_BuildError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMalformedBaseURLClient(t)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGpfcpProfileResource_Delete_SendError exercises GpfcpProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestGpfcpProfileResource_Delete_SendError(t *testing.T) {
	r := &GpfcpProfileResource{client: newTransportErrorClient(t)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGpfcpProfileResource_Delete_NotFoundSuccess exercises GpfcpProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestGpfcpProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 404, "")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGpfcpProfileResource_Delete_APIError exercises GpfcpProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGpfcpProfileResource_Delete_APIError(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GpfcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_gpfcp_profile")
}

// TestGpfcpProfileResource_Delete_APIErrorReadBody exercises GpfcpProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGpfcpProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &GpfcpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := GpfcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
