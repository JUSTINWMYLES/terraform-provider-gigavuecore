package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestEnhancedSlicingResource_Create_Happy exercises EnhancedSlicingResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestEnhancedSlicingResource_Create_Happy(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnhancedSlicingResource_Create_NilClient exercises EnhancedSlicingResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEnhancedSlicingResource_Create_NilClient(t *testing.T) {
	r := &EnhancedSlicingResource{}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEnhancedSlicingResource_Create_BuildError exercises EnhancedSlicingResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEnhancedSlicingResource_Create_BuildError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMalformedBaseURLClient(t)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEnhancedSlicingResource_Create_SendError exercises EnhancedSlicingResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestEnhancedSlicingResource_Create_SendError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newTransportErrorClient(t)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEnhancedSlicingResource_Create_APIError exercises EnhancedSlicingResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEnhancedSlicingResource_Create_APIError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_enhanced_slicing")
}

// TestEnhancedSlicingResource_Create_APIErrorReadBody exercises EnhancedSlicingResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEnhancedSlicingResource_Create_APIErrorReadBody(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientReadErrorBody(t, 501)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestEnhancedSlicingResource_Create_InvalidJSON exercises EnhancedSlicingResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestEnhancedSlicingResource_Create_InvalidJSON(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 201, "{{")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestEnhancedSlicingResource_Create_MapError exercises EnhancedSlicingResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestEnhancedSlicingResource_Create_MapError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestEnhancedSlicingResource_Create_MissingID exercises EnhancedSlicingResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestEnhancedSlicingResource_Create_MissingID(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 201, "{}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestEnhancedSlicingResource_Create_LocationFallback exercises EnhancedSlicingResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestEnhancedSlicingResource_Create_LocationFallback(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestEnhancedSlicingResource_Read_Happy exercises EnhancedSlicingResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestEnhancedSlicingResource_Read_Happy(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 200, "{}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnhancedSlicingResource_Read_NilClient exercises EnhancedSlicingResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEnhancedSlicingResource_Read_NilClient(t *testing.T) {
	r := &EnhancedSlicingResource{}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEnhancedSlicingResource_Read_BuildError exercises EnhancedSlicingResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEnhancedSlicingResource_Read_BuildError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMalformedBaseURLClient(t)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEnhancedSlicingResource_Read_SendError exercises EnhancedSlicingResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestEnhancedSlicingResource_Read_SendError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newTransportErrorClient(t)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEnhancedSlicingResource_Read_NotFound exercises EnhancedSlicingResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestEnhancedSlicingResource_Read_NotFound(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 404, "")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnhancedSlicingResource_Read_APIError exercises EnhancedSlicingResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEnhancedSlicingResource_Read_APIError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_enhanced_slicing")
}

// TestEnhancedSlicingResource_Read_APIErrorReadBody exercises EnhancedSlicingResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEnhancedSlicingResource_Read_APIErrorReadBody(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientReadErrorBody(t, 501)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestEnhancedSlicingResource_Read_InvalidJSON exercises EnhancedSlicingResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestEnhancedSlicingResource_Read_InvalidJSON(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 200, "{{")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestEnhancedSlicingResource_Read_MapError exercises EnhancedSlicingResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestEnhancedSlicingResource_Read_MapError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestEnhancedSlicingResource_Update_Happy exercises EnhancedSlicingResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestEnhancedSlicingResource_Update_Happy(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 200, "{}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnhancedSlicingResource_Update_NilClient exercises EnhancedSlicingResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEnhancedSlicingResource_Update_NilClient(t *testing.T) {
	r := &EnhancedSlicingResource{}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEnhancedSlicingResource_Update_BuildError exercises EnhancedSlicingResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEnhancedSlicingResource_Update_BuildError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMalformedBaseURLClient(t)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEnhancedSlicingResource_Update_SendError exercises EnhancedSlicingResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestEnhancedSlicingResource_Update_SendError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newTransportErrorClient(t)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEnhancedSlicingResource_Update_APIError exercises EnhancedSlicingResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEnhancedSlicingResource_Update_APIError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_enhanced_slicing")
}

// TestEnhancedSlicingResource_Update_APIErrorReadBody exercises EnhancedSlicingResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEnhancedSlicingResource_Update_APIErrorReadBody(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientReadErrorBody(t, 501)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestEnhancedSlicingResource_Update_InvalidJSON exercises EnhancedSlicingResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestEnhancedSlicingResource_Update_InvalidJSON(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 200, "{{")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestEnhancedSlicingResource_Update_MapError exercises EnhancedSlicingResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestEnhancedSlicingResource_Update_MapError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestEnhancedSlicingResource_Delete_Happy exercises EnhancedSlicingResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestEnhancedSlicingResource_Delete_Happy(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 204, "")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnhancedSlicingResource_Delete_NilClient exercises EnhancedSlicingResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEnhancedSlicingResource_Delete_NilClient(t *testing.T) {
	r := &EnhancedSlicingResource{}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEnhancedSlicingResource_Delete_BuildError exercises EnhancedSlicingResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEnhancedSlicingResource_Delete_BuildError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMalformedBaseURLClient(t)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEnhancedSlicingResource_Delete_SendError exercises EnhancedSlicingResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestEnhancedSlicingResource_Delete_SendError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newTransportErrorClient(t)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEnhancedSlicingResource_Delete_NotFoundSuccess exercises EnhancedSlicingResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestEnhancedSlicingResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 404, "")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnhancedSlicingResource_Delete_APIError exercises EnhancedSlicingResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEnhancedSlicingResource_Delete_APIError(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_enhanced_slicing")
}

// TestEnhancedSlicingResource_Delete_APIErrorReadBody exercises EnhancedSlicingResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEnhancedSlicingResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &EnhancedSlicingResource{client: newMockClientReadErrorBody(t, 501)}
	m := EnhancedSlicingResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
