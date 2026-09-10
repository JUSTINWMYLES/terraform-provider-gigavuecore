package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSpineLinkResource_Create_Happy exercises SpineLinkResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestSpineLinkResource_Create_Happy(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := SpineLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSpineLinkResource_Create_NilClient exercises SpineLinkResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSpineLinkResource_Create_NilClient(t *testing.T) {
	r := &SpineLinkResource{}
	m := SpineLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSpineLinkResource_Create_BuildError exercises SpineLinkResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSpineLinkResource_Create_BuildError(t *testing.T) {
	r := &SpineLinkResource{client: newMalformedBaseURLClient(t)}
	m := SpineLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSpineLinkResource_Create_SendError exercises SpineLinkResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestSpineLinkResource_Create_SendError(t *testing.T) {
	r := &SpineLinkResource{client: newTransportErrorClient(t)}
	m := SpineLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSpineLinkResource_Create_APIError exercises SpineLinkResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSpineLinkResource_Create_APIError(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SpineLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_spine_link")
}

// TestSpineLinkResource_Create_APIErrorReadBody exercises SpineLinkResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSpineLinkResource_Create_APIErrorReadBody(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientReadErrorBody(t, 501)}
	m := SpineLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSpineLinkResource_Create_InvalidJSON exercises SpineLinkResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSpineLinkResource_Create_InvalidJSON(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 201, "{{")}
	m := SpineLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSpineLinkResource_Create_MapError exercises SpineLinkResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSpineLinkResource_Create_MapError(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := SpineLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSpineLinkResource_Create_MissingID exercises SpineLinkResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestSpineLinkResource_Create_MissingID(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 201, "{}")}
	m := SpineLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestSpineLinkResource_Create_LocationFallback exercises SpineLinkResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestSpineLinkResource_Create_LocationFallback(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := SpineLinkResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestSpineLinkResource_Read_Happy exercises SpineLinkResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestSpineLinkResource_Read_Happy(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 200, "{}")}
	m := SpineLinkResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSpineLinkResource_Read_NilClient exercises SpineLinkResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSpineLinkResource_Read_NilClient(t *testing.T) {
	r := &SpineLinkResource{}
	m := SpineLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSpineLinkResource_Read_BuildError exercises SpineLinkResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSpineLinkResource_Read_BuildError(t *testing.T) {
	r := &SpineLinkResource{client: newMalformedBaseURLClient(t)}
	m := SpineLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSpineLinkResource_Read_SendError exercises SpineLinkResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestSpineLinkResource_Read_SendError(t *testing.T) {
	r := &SpineLinkResource{client: newTransportErrorClient(t)}
	m := SpineLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSpineLinkResource_Read_NotFound exercises SpineLinkResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestSpineLinkResource_Read_NotFound(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 404, "")}
	m := SpineLinkResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSpineLinkResource_Read_APIError exercises SpineLinkResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSpineLinkResource_Read_APIError(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SpineLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_spine_link")
}

// TestSpineLinkResource_Read_APIErrorReadBody exercises SpineLinkResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSpineLinkResource_Read_APIErrorReadBody(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientReadErrorBody(t, 501)}
	m := SpineLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSpineLinkResource_Read_InvalidJSON exercises SpineLinkResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSpineLinkResource_Read_InvalidJSON(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 200, "{{")}
	m := SpineLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSpineLinkResource_Read_MapError exercises SpineLinkResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSpineLinkResource_Read_MapError(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SpineLinkResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSpineLinkResource_Update_Happy exercises SpineLinkResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestSpineLinkResource_Update_Happy(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 200, "{}")}
	m := SpineLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSpineLinkResource_Update_NilClient exercises SpineLinkResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSpineLinkResource_Update_NilClient(t *testing.T) {
	r := &SpineLinkResource{}
	m := SpineLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSpineLinkResource_Update_BuildError exercises SpineLinkResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSpineLinkResource_Update_BuildError(t *testing.T) {
	r := &SpineLinkResource{client: newMalformedBaseURLClient(t)}
	m := SpineLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSpineLinkResource_Update_SendError exercises SpineLinkResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestSpineLinkResource_Update_SendError(t *testing.T) {
	r := &SpineLinkResource{client: newTransportErrorClient(t)}
	m := SpineLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSpineLinkResource_Update_APIError exercises SpineLinkResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSpineLinkResource_Update_APIError(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SpineLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_spine_link")
}

// TestSpineLinkResource_Update_APIErrorReadBody exercises SpineLinkResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSpineLinkResource_Update_APIErrorReadBody(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientReadErrorBody(t, 501)}
	m := SpineLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSpineLinkResource_Update_InvalidJSON exercises SpineLinkResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSpineLinkResource_Update_InvalidJSON(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 200, "{{")}
	m := SpineLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSpineLinkResource_Update_MapError exercises SpineLinkResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSpineLinkResource_Update_MapError(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SpineLinkResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSpineLinkResource_Delete_Happy exercises SpineLinkResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestSpineLinkResource_Delete_Happy(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 204, "")}
	m := SpineLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSpineLinkResource_Delete_NilClient exercises SpineLinkResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSpineLinkResource_Delete_NilClient(t *testing.T) {
	r := &SpineLinkResource{}
	m := SpineLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSpineLinkResource_Delete_BuildError exercises SpineLinkResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSpineLinkResource_Delete_BuildError(t *testing.T) {
	r := &SpineLinkResource{client: newMalformedBaseURLClient(t)}
	m := SpineLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSpineLinkResource_Delete_SendError exercises SpineLinkResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestSpineLinkResource_Delete_SendError(t *testing.T) {
	r := &SpineLinkResource{client: newTransportErrorClient(t)}
	m := SpineLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSpineLinkResource_Delete_NotFoundSuccess exercises SpineLinkResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestSpineLinkResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 404, "")}
	m := SpineLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSpineLinkResource_Delete_APIError exercises SpineLinkResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSpineLinkResource_Delete_APIError(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SpineLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_spine_link")
}

// TestSpineLinkResource_Delete_APIErrorReadBody exercises SpineLinkResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSpineLinkResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &SpineLinkResource{client: newMockClientReadErrorBody(t, 501)}
	m := SpineLinkResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
