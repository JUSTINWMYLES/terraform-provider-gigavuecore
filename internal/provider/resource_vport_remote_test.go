package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestVportResource_Create_Happy exercises VportResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestVportResource_Create_Happy(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := VportResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestVportResource_Create_NilClient exercises VportResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestVportResource_Create_NilClient(t *testing.T) {
	r := &VportResource{}
	m := VportResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestVportResource_Create_BuildError exercises VportResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestVportResource_Create_BuildError(t *testing.T) {
	r := &VportResource{client: newMalformedBaseURLClient(t)}
	m := VportResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestVportResource_Create_SendError exercises VportResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestVportResource_Create_SendError(t *testing.T) {
	r := &VportResource{client: newTransportErrorClient(t)}
	m := VportResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestVportResource_Create_APIError exercises VportResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestVportResource_Create_APIError(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := VportResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_vport")
}

// TestVportResource_Create_APIErrorReadBody exercises VportResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestVportResource_Create_APIErrorReadBody(t *testing.T) {
	r := &VportResource{client: newMockClientReadErrorBody(t, 501)}
	m := VportResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestVportResource_Create_InvalidJSON exercises VportResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestVportResource_Create_InvalidJSON(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 201, "{{")}
	m := VportResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestVportResource_Create_MapError exercises VportResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestVportResource_Create_MapError(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := VportResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestVportResource_Create_MissingID exercises VportResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestVportResource_Create_MissingID(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 201, "{}")}
	m := VportResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestVportResource_Create_LocationFallback exercises VportResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestVportResource_Create_LocationFallback(t *testing.T) {
	r := &VportResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := VportResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestVportResource_Read_Happy exercises VportResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestVportResource_Read_Happy(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 200, "{}")}
	m := VportResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestVportResource_Read_NilClient exercises VportResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestVportResource_Read_NilClient(t *testing.T) {
	r := &VportResource{}
	m := VportResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestVportResource_Read_BuildError exercises VportResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestVportResource_Read_BuildError(t *testing.T) {
	r := &VportResource{client: newMalformedBaseURLClient(t)}
	m := VportResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestVportResource_Read_SendError exercises VportResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestVportResource_Read_SendError(t *testing.T) {
	r := &VportResource{client: newTransportErrorClient(t)}
	m := VportResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestVportResource_Read_NotFound exercises VportResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestVportResource_Read_NotFound(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 404, "")}
	m := VportResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestVportResource_Read_APIError exercises VportResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestVportResource_Read_APIError(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := VportResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_vport")
}

// TestVportResource_Read_APIErrorReadBody exercises VportResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestVportResource_Read_APIErrorReadBody(t *testing.T) {
	r := &VportResource{client: newMockClientReadErrorBody(t, 501)}
	m := VportResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestVportResource_Read_InvalidJSON exercises VportResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestVportResource_Read_InvalidJSON(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 200, "{{")}
	m := VportResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestVportResource_Read_MapError exercises VportResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestVportResource_Read_MapError(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := VportResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestVportResource_Update_Happy exercises VportResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestVportResource_Update_Happy(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 200, "{}")}
	m := VportResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestVportResource_Update_NilClient exercises VportResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestVportResource_Update_NilClient(t *testing.T) {
	r := &VportResource{}
	m := VportResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestVportResource_Update_BuildError exercises VportResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestVportResource_Update_BuildError(t *testing.T) {
	r := &VportResource{client: newMalformedBaseURLClient(t)}
	m := VportResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestVportResource_Update_SendError exercises VportResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestVportResource_Update_SendError(t *testing.T) {
	r := &VportResource{client: newTransportErrorClient(t)}
	m := VportResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestVportResource_Update_APIError exercises VportResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestVportResource_Update_APIError(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := VportResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_vport")
}

// TestVportResource_Update_APIErrorReadBody exercises VportResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestVportResource_Update_APIErrorReadBody(t *testing.T) {
	r := &VportResource{client: newMockClientReadErrorBody(t, 501)}
	m := VportResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestVportResource_Update_InvalidJSON exercises VportResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestVportResource_Update_InvalidJSON(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 200, "{{")}
	m := VportResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestVportResource_Update_MapError exercises VportResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestVportResource_Update_MapError(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := VportResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestVportResource_Delete_Happy exercises VportResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestVportResource_Delete_Happy(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 204, "")}
	m := VportResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestVportResource_Delete_NilClient exercises VportResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestVportResource_Delete_NilClient(t *testing.T) {
	r := &VportResource{}
	m := VportResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestVportResource_Delete_BuildError exercises VportResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestVportResource_Delete_BuildError(t *testing.T) {
	r := &VportResource{client: newMalformedBaseURLClient(t)}
	m := VportResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestVportResource_Delete_SendError exercises VportResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestVportResource_Delete_SendError(t *testing.T) {
	r := &VportResource{client: newTransportErrorClient(t)}
	m := VportResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestVportResource_Delete_NotFoundSuccess exercises VportResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestVportResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 404, "")}
	m := VportResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestVportResource_Delete_APIError exercises VportResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestVportResource_Delete_APIError(t *testing.T) {
	r := &VportResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := VportResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_vport")
}

// TestVportResource_Delete_APIErrorReadBody exercises VportResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestVportResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &VportResource{client: newMockClientReadErrorBody(t, 501)}
	m := VportResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
