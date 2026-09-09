package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestRedundancyProfileResource_Create_Happy exercises RedundancyProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestRedundancyProfileResource_Create_Happy(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedundancyProfileResource_Create_NilClient exercises RedundancyProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedundancyProfileResource_Create_NilClient(t *testing.T) {
	r := &RedundancyProfileResource{}
	m := RedundancyProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedundancyProfileResource_Create_BuildError exercises RedundancyProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedundancyProfileResource_Create_BuildError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMalformedBaseURLClient(t)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedundancyProfileResource_Create_SendError exercises RedundancyProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedundancyProfileResource_Create_SendError(t *testing.T) {
	r := &RedundancyProfileResource{client: newTransportErrorClient(t)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedundancyProfileResource_Create_APIError exercises RedundancyProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedundancyProfileResource_Create_APIError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_redundancy_profile")
}

// TestRedundancyProfileResource_Create_APIErrorReadBody exercises RedundancyProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedundancyProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRedundancyProfileResource_Create_InvalidJSON exercises RedundancyProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRedundancyProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRedundancyProfileResource_Create_MapError exercises RedundancyProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRedundancyProfileResource_Create_MapError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRedundancyProfileResource_Create_MissingID exercises RedundancyProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestRedundancyProfileResource_Create_MissingID(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestRedundancyProfileResource_Create_LocationFallback exercises RedundancyProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestRedundancyProfileResource_Create_LocationFallback(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestRedundancyProfileResource_Read_Happy exercises RedundancyProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestRedundancyProfileResource_Read_Happy(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedundancyProfileResource_Read_NilClient exercises RedundancyProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedundancyProfileResource_Read_NilClient(t *testing.T) {
	r := &RedundancyProfileResource{}
	m := RedundancyProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedundancyProfileResource_Read_BuildError exercises RedundancyProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedundancyProfileResource_Read_BuildError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMalformedBaseURLClient(t)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedundancyProfileResource_Read_SendError exercises RedundancyProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedundancyProfileResource_Read_SendError(t *testing.T) {
	r := &RedundancyProfileResource{client: newTransportErrorClient(t)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedundancyProfileResource_Read_NotFound exercises RedundancyProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestRedundancyProfileResource_Read_NotFound(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 404, "")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedundancyProfileResource_Read_APIError exercises RedundancyProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedundancyProfileResource_Read_APIError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_redundancy_profile")
}

// TestRedundancyProfileResource_Read_APIErrorReadBody exercises RedundancyProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedundancyProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRedundancyProfileResource_Read_InvalidJSON exercises RedundancyProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRedundancyProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRedundancyProfileResource_Read_MapError exercises RedundancyProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRedundancyProfileResource_Read_MapError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRedundancyProfileResource_Update_Happy exercises RedundancyProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestRedundancyProfileResource_Update_Happy(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedundancyProfileResource_Update_NilClient exercises RedundancyProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedundancyProfileResource_Update_NilClient(t *testing.T) {
	r := &RedundancyProfileResource{}
	m := RedundancyProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedundancyProfileResource_Update_BuildError exercises RedundancyProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedundancyProfileResource_Update_BuildError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMalformedBaseURLClient(t)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedundancyProfileResource_Update_SendError exercises RedundancyProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedundancyProfileResource_Update_SendError(t *testing.T) {
	r := &RedundancyProfileResource{client: newTransportErrorClient(t)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedundancyProfileResource_Update_APIError exercises RedundancyProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedundancyProfileResource_Update_APIError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_redundancy_profile")
}

// TestRedundancyProfileResource_Update_APIErrorReadBody exercises RedundancyProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedundancyProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRedundancyProfileResource_Update_InvalidJSON exercises RedundancyProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRedundancyProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRedundancyProfileResource_Update_MapError exercises RedundancyProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRedundancyProfileResource_Update_MapError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRedundancyProfileResource_Delete_Happy exercises RedundancyProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestRedundancyProfileResource_Delete_Happy(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 204, "")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedundancyProfileResource_Delete_NilClient exercises RedundancyProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedundancyProfileResource_Delete_NilClient(t *testing.T) {
	r := &RedundancyProfileResource{}
	m := RedundancyProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedundancyProfileResource_Delete_BuildError exercises RedundancyProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedundancyProfileResource_Delete_BuildError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMalformedBaseURLClient(t)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedundancyProfileResource_Delete_SendError exercises RedundancyProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedundancyProfileResource_Delete_SendError(t *testing.T) {
	r := &RedundancyProfileResource{client: newTransportErrorClient(t)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedundancyProfileResource_Delete_NotFoundSuccess exercises RedundancyProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestRedundancyProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 404, "")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedundancyProfileResource_Delete_APIError exercises RedundancyProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedundancyProfileResource_Delete_APIError(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedundancyProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_redundancy_profile")
}

// TestRedundancyProfileResource_Delete_APIErrorReadBody exercises RedundancyProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedundancyProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &RedundancyProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := RedundancyProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
