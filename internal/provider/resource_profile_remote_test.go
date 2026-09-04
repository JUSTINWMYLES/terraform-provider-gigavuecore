package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestProfileResource_Create_Happy exercises ProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestProfileResource_Create_Happy(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestProfileResource_Create_NilClient exercises ProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestProfileResource_Create_NilClient(t *testing.T) {
	r := &ProfileResource{}
	m := ProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestProfileResource_Create_BuildError exercises ProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestProfileResource_Create_BuildError(t *testing.T) {
	r := &ProfileResource{client: newMalformedBaseURLClient(t)}
	m := ProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestProfileResource_Create_SendError exercises ProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestProfileResource_Create_SendError(t *testing.T) {
	r := &ProfileResource{client: newTransportErrorClient(t)}
	m := ProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestProfileResource_Create_APIError exercises ProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestProfileResource_Create_APIError(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_profile")
}

// TestProfileResource_Create_APIErrorReadBody exercises ProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestProfileResource_Create_InvalidJSON exercises ProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := ProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestProfileResource_Create_MapError exercises ProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestProfileResource_Create_MapError(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestProfileResource_Create_MissingID exercises ProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestProfileResource_Create_MissingID(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := ProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestProfileResource_Create_LocationFallback exercises ProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestProfileResource_Create_LocationFallback(t *testing.T) {
	r := &ProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestProfileResource_Read_Happy exercises ProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestProfileResource_Read_Happy(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := ProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestProfileResource_Read_NilClient exercises ProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestProfileResource_Read_NilClient(t *testing.T) {
	r := &ProfileResource{}
	m := ProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestProfileResource_Read_BuildError exercises ProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestProfileResource_Read_BuildError(t *testing.T) {
	r := &ProfileResource{client: newMalformedBaseURLClient(t)}
	m := ProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestProfileResource_Read_SendError exercises ProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestProfileResource_Read_SendError(t *testing.T) {
	r := &ProfileResource{client: newTransportErrorClient(t)}
	m := ProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestProfileResource_Read_NotFound exercises ProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestProfileResource_Read_NotFound(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 404, "")}
	m := ProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestProfileResource_Read_APIError exercises ProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestProfileResource_Read_APIError(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_profile")
}

// TestProfileResource_Read_APIErrorReadBody exercises ProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestProfileResource_Read_InvalidJSON exercises ProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := ProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestProfileResource_Read_MapError exercises ProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestProfileResource_Read_MapError(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestProfileResource_Update_Happy exercises ProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestProfileResource_Update_Happy(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := ProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestProfileResource_Update_NilClient exercises ProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestProfileResource_Update_NilClient(t *testing.T) {
	r := &ProfileResource{}
	m := ProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestProfileResource_Update_BuildError exercises ProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestProfileResource_Update_BuildError(t *testing.T) {
	r := &ProfileResource{client: newMalformedBaseURLClient(t)}
	m := ProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestProfileResource_Update_SendError exercises ProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestProfileResource_Update_SendError(t *testing.T) {
	r := &ProfileResource{client: newTransportErrorClient(t)}
	m := ProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestProfileResource_Update_APIError exercises ProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestProfileResource_Update_APIError(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_profile")
}

// TestProfileResource_Update_APIErrorReadBody exercises ProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestProfileResource_Update_InvalidJSON exercises ProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := ProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestProfileResource_Update_MapError exercises ProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestProfileResource_Update_MapError(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestProfileResource_Delete_Happy exercises ProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestProfileResource_Delete_Happy(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 204, "")}
	m := ProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestProfileResource_Delete_NilClient exercises ProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestProfileResource_Delete_NilClient(t *testing.T) {
	r := &ProfileResource{}
	m := ProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestProfileResource_Delete_BuildError exercises ProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestProfileResource_Delete_BuildError(t *testing.T) {
	r := &ProfileResource{client: newMalformedBaseURLClient(t)}
	m := ProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestProfileResource_Delete_SendError exercises ProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestProfileResource_Delete_SendError(t *testing.T) {
	r := &ProfileResource{client: newTransportErrorClient(t)}
	m := ProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestProfileResource_Delete_NotFoundSuccess exercises ProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 404, "")}
	m := ProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestProfileResource_Delete_APIError exercises ProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestProfileResource_Delete_APIError(t *testing.T) {
	r := &ProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_profile")
}

// TestProfileResource_Delete_APIErrorReadBody exercises ProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
