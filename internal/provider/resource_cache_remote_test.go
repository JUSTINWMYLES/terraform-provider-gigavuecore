package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestCacheResource_Create_Happy exercises CacheResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestCacheResource_Create_Happy(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := CacheResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCacheResource_Create_NilClient exercises CacheResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCacheResource_Create_NilClient(t *testing.T) {
	r := &CacheResource{}
	m := CacheResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCacheResource_Create_BuildError exercises CacheResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCacheResource_Create_BuildError(t *testing.T) {
	r := &CacheResource{client: newMalformedBaseURLClient(t)}
	m := CacheResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCacheResource_Create_SendError exercises CacheResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestCacheResource_Create_SendError(t *testing.T) {
	r := &CacheResource{client: newTransportErrorClient(t)}
	m := CacheResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCacheResource_Create_APIError exercises CacheResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCacheResource_Create_APIError(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CacheResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_cache")
}

// TestCacheResource_Create_APIErrorReadBody exercises CacheResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCacheResource_Create_APIErrorReadBody(t *testing.T) {
	r := &CacheResource{client: newMockClientReadErrorBody(t, 501)}
	m := CacheResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestCacheResource_Create_InvalidJSON exercises CacheResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestCacheResource_Create_InvalidJSON(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 201, "{{")}
	m := CacheResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestCacheResource_Create_MapError exercises CacheResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestCacheResource_Create_MapError(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := CacheResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestCacheResource_Create_MissingID exercises CacheResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestCacheResource_Create_MissingID(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 201, "{}")}
	m := CacheResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestCacheResource_Create_LocationFallback exercises CacheResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestCacheResource_Create_LocationFallback(t *testing.T) {
	r := &CacheResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := CacheResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestCacheResource_Read_Happy exercises CacheResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestCacheResource_Read_Happy(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 200, "{}")}
	m := CacheResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestCacheResource_Read_NilClient exercises CacheResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCacheResource_Read_NilClient(t *testing.T) {
	r := &CacheResource{}
	m := CacheResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCacheResource_Read_BuildError exercises CacheResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCacheResource_Read_BuildError(t *testing.T) {
	r := &CacheResource{client: newMalformedBaseURLClient(t)}
	m := CacheResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCacheResource_Read_SendError exercises CacheResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestCacheResource_Read_SendError(t *testing.T) {
	r := &CacheResource{client: newTransportErrorClient(t)}
	m := CacheResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCacheResource_Read_NotFound exercises CacheResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestCacheResource_Read_NotFound(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 404, "")}
	m := CacheResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestCacheResource_Read_APIError exercises CacheResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCacheResource_Read_APIError(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CacheResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_cache")
}

// TestCacheResource_Read_APIErrorReadBody exercises CacheResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCacheResource_Read_APIErrorReadBody(t *testing.T) {
	r := &CacheResource{client: newMockClientReadErrorBody(t, 501)}
	m := CacheResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestCacheResource_Read_InvalidJSON exercises CacheResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestCacheResource_Read_InvalidJSON(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 200, "{{")}
	m := CacheResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestCacheResource_Read_MapError exercises CacheResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestCacheResource_Read_MapError(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := CacheResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestCacheResource_Update_Happy exercises CacheResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestCacheResource_Update_Happy(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 200, "{}")}
	m := CacheResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCacheResource_Update_NilClient exercises CacheResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCacheResource_Update_NilClient(t *testing.T) {
	r := &CacheResource{}
	m := CacheResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCacheResource_Update_BuildError exercises CacheResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCacheResource_Update_BuildError(t *testing.T) {
	r := &CacheResource{client: newMalformedBaseURLClient(t)}
	m := CacheResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCacheResource_Update_SendError exercises CacheResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestCacheResource_Update_SendError(t *testing.T) {
	r := &CacheResource{client: newTransportErrorClient(t)}
	m := CacheResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCacheResource_Update_APIError exercises CacheResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCacheResource_Update_APIError(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CacheResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_cache")
}

// TestCacheResource_Update_APIErrorReadBody exercises CacheResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCacheResource_Update_APIErrorReadBody(t *testing.T) {
	r := &CacheResource{client: newMockClientReadErrorBody(t, 501)}
	m := CacheResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestCacheResource_Update_InvalidJSON exercises CacheResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestCacheResource_Update_InvalidJSON(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 200, "{{")}
	m := CacheResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestCacheResource_Update_MapError exercises CacheResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestCacheResource_Update_MapError(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := CacheResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestCacheResource_Delete_Happy exercises CacheResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestCacheResource_Delete_Happy(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 204, "")}
	m := CacheResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCacheResource_Delete_NilClient exercises CacheResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCacheResource_Delete_NilClient(t *testing.T) {
	r := &CacheResource{}
	m := CacheResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCacheResource_Delete_BuildError exercises CacheResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCacheResource_Delete_BuildError(t *testing.T) {
	r := &CacheResource{client: newMalformedBaseURLClient(t)}
	m := CacheResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCacheResource_Delete_SendError exercises CacheResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestCacheResource_Delete_SendError(t *testing.T) {
	r := &CacheResource{client: newTransportErrorClient(t)}
	m := CacheResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCacheResource_Delete_NotFoundSuccess exercises CacheResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestCacheResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 404, "")}
	m := CacheResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCacheResource_Delete_APIError exercises CacheResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCacheResource_Delete_APIError(t *testing.T) {
	r := &CacheResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CacheResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_cache")
}

// TestCacheResource_Delete_APIErrorReadBody exercises CacheResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCacheResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &CacheResource{client: newMockClientReadErrorBody(t, 501)}
	m := CacheResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
