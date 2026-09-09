package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestAppVisibilityResource_Create_Happy exercises AppVisibilityResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestAppVisibilityResource_Create_Happy(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 201, "{\"solution_alias\":\"example-id\"}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAppVisibilityResource_Create_NilClient exercises AppVisibilityResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAppVisibilityResource_Create_NilClient(t *testing.T) {
	r := &AppVisibilityResource{}
	m := AppVisibilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAppVisibilityResource_Create_BuildError exercises AppVisibilityResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAppVisibilityResource_Create_BuildError(t *testing.T) {
	r := &AppVisibilityResource{client: newMalformedBaseURLClient(t)}
	m := AppVisibilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAppVisibilityResource_Create_SendError exercises AppVisibilityResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestAppVisibilityResource_Create_SendError(t *testing.T) {
	r := &AppVisibilityResource{client: newTransportErrorClient(t)}
	m := AppVisibilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAppVisibilityResource_Create_APIError exercises AppVisibilityResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAppVisibilityResource_Create_APIError(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_app_visibility")
}

// TestAppVisibilityResource_Create_APIErrorReadBody exercises AppVisibilityResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAppVisibilityResource_Create_APIErrorReadBody(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientReadErrorBody(t, 501)}
	m := AppVisibilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestAppVisibilityResource_Create_InvalidJSON exercises AppVisibilityResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestAppVisibilityResource_Create_InvalidJSON(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 201, "{{")}
	m := AppVisibilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestAppVisibilityResource_Create_MapError exercises AppVisibilityResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestAppVisibilityResource_Create_MapError(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 201, "{\"solution_alias\":12345}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestAppVisibilityResource_Create_MissingID exercises AppVisibilityResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestAppVisibilityResource_Create_MissingID(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 201, "{}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestAppVisibilityResource_Create_LocationFallback exercises AppVisibilityResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestAppVisibilityResource_Create_LocationFallback(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.SolutionAlias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.SolutionAlias.ValueString(), "example-id")
	}
}

// TestAppVisibilityResource_Read_Happy exercises AppVisibilityResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestAppVisibilityResource_Read_Happy(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 200, "{}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestAppVisibilityResource_Read_NilClient exercises AppVisibilityResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAppVisibilityResource_Read_NilClient(t *testing.T) {
	r := &AppVisibilityResource{}
	m := AppVisibilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAppVisibilityResource_Read_BuildError exercises AppVisibilityResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAppVisibilityResource_Read_BuildError(t *testing.T) {
	r := &AppVisibilityResource{client: newMalformedBaseURLClient(t)}
	m := AppVisibilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAppVisibilityResource_Read_SendError exercises AppVisibilityResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestAppVisibilityResource_Read_SendError(t *testing.T) {
	r := &AppVisibilityResource{client: newTransportErrorClient(t)}
	m := AppVisibilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAppVisibilityResource_Read_NotFound exercises AppVisibilityResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestAppVisibilityResource_Read_NotFound(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 404, "")}
	m := AppVisibilityResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestAppVisibilityResource_Read_APIError exercises AppVisibilityResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAppVisibilityResource_Read_APIError(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_app_visibility")
}

// TestAppVisibilityResource_Read_APIErrorReadBody exercises AppVisibilityResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAppVisibilityResource_Read_APIErrorReadBody(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientReadErrorBody(t, 501)}
	m := AppVisibilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestAppVisibilityResource_Read_InvalidJSON exercises AppVisibilityResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestAppVisibilityResource_Read_InvalidJSON(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 200, "{{")}
	m := AppVisibilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestAppVisibilityResource_Read_MapError exercises AppVisibilityResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestAppVisibilityResource_Read_MapError(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 200, "{\"solution_alias\":12345}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestAppVisibilityResource_Update_Happy exercises AppVisibilityResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestAppVisibilityResource_Update_Happy(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 200, "{}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAppVisibilityResource_Update_NilClient exercises AppVisibilityResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAppVisibilityResource_Update_NilClient(t *testing.T) {
	r := &AppVisibilityResource{}
	m := AppVisibilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAppVisibilityResource_Update_BuildError exercises AppVisibilityResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAppVisibilityResource_Update_BuildError(t *testing.T) {
	r := &AppVisibilityResource{client: newMalformedBaseURLClient(t)}
	m := AppVisibilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAppVisibilityResource_Update_SendError exercises AppVisibilityResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestAppVisibilityResource_Update_SendError(t *testing.T) {
	r := &AppVisibilityResource{client: newTransportErrorClient(t)}
	m := AppVisibilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAppVisibilityResource_Update_APIError exercises AppVisibilityResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAppVisibilityResource_Update_APIError(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_app_visibility")
}

// TestAppVisibilityResource_Update_APIErrorReadBody exercises AppVisibilityResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAppVisibilityResource_Update_APIErrorReadBody(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientReadErrorBody(t, 501)}
	m := AppVisibilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestAppVisibilityResource_Update_InvalidJSON exercises AppVisibilityResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestAppVisibilityResource_Update_InvalidJSON(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 200, "{{")}
	m := AppVisibilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestAppVisibilityResource_Update_MapError exercises AppVisibilityResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestAppVisibilityResource_Update_MapError(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 200, "{\"solution_alias\":12345}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestAppVisibilityResource_Delete_Happy exercises AppVisibilityResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestAppVisibilityResource_Delete_Happy(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 204, "")}
	m := AppVisibilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAppVisibilityResource_Delete_NilClient exercises AppVisibilityResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAppVisibilityResource_Delete_NilClient(t *testing.T) {
	r := &AppVisibilityResource{}
	m := AppVisibilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAppVisibilityResource_Delete_BuildError exercises AppVisibilityResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAppVisibilityResource_Delete_BuildError(t *testing.T) {
	r := &AppVisibilityResource{client: newMalformedBaseURLClient(t)}
	m := AppVisibilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAppVisibilityResource_Delete_SendError exercises AppVisibilityResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestAppVisibilityResource_Delete_SendError(t *testing.T) {
	r := &AppVisibilityResource{client: newTransportErrorClient(t)}
	m := AppVisibilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAppVisibilityResource_Delete_NotFoundSuccess exercises AppVisibilityResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestAppVisibilityResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 404, "")}
	m := AppVisibilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAppVisibilityResource_Delete_APIError exercises AppVisibilityResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAppVisibilityResource_Delete_APIError(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AppVisibilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_app_visibility")
}

// TestAppVisibilityResource_Delete_APIErrorReadBody exercises AppVisibilityResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAppVisibilityResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &AppVisibilityResource{client: newMockClientReadErrorBody(t, 501)}
	m := AppVisibilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
