package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMapGroupResource_Create_Happy exercises MapGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestMapGroupResource_Create_Happy(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := MapGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapGroupResource_Create_NilClient exercises MapGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapGroupResource_Create_NilClient(t *testing.T) {
	r := &MapGroupResource{}
	m := MapGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapGroupResource_Create_BuildError exercises MapGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapGroupResource_Create_BuildError(t *testing.T) {
	r := &MapGroupResource{client: newMalformedBaseURLClient(t)}
	m := MapGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapGroupResource_Create_SendError exercises MapGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapGroupResource_Create_SendError(t *testing.T) {
	r := &MapGroupResource{client: newTransportErrorClient(t)}
	m := MapGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapGroupResource_Create_APIError exercises MapGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapGroupResource_Create_APIError(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_map_group")
}

// TestMapGroupResource_Create_APIErrorReadBody exercises MapGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &MapGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMapGroupResource_Create_InvalidJSON exercises MapGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMapGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := MapGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMapGroupResource_Create_MapError exercises MapGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMapGroupResource_Create_MapError(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := MapGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMapGroupResource_Create_MissingID exercises MapGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestMapGroupResource_Create_MissingID(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := MapGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestMapGroupResource_Create_LocationFallback exercises MapGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestMapGroupResource_Create_LocationFallback(t *testing.T) {
	r := &MapGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := MapGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestMapGroupResource_Read_Happy exercises MapGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestMapGroupResource_Read_Happy(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := MapGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapGroupResource_Read_NilClient exercises MapGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapGroupResource_Read_NilClient(t *testing.T) {
	r := &MapGroupResource{}
	m := MapGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapGroupResource_Read_BuildError exercises MapGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapGroupResource_Read_BuildError(t *testing.T) {
	r := &MapGroupResource{client: newMalformedBaseURLClient(t)}
	m := MapGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapGroupResource_Read_SendError exercises MapGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapGroupResource_Read_SendError(t *testing.T) {
	r := &MapGroupResource{client: newTransportErrorClient(t)}
	m := MapGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapGroupResource_Read_NotFound exercises MapGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestMapGroupResource_Read_NotFound(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 404, "")}
	m := MapGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapGroupResource_Read_APIError exercises MapGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapGroupResource_Read_APIError(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_map_group")
}

// TestMapGroupResource_Read_APIErrorReadBody exercises MapGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &MapGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMapGroupResource_Read_InvalidJSON exercises MapGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMapGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := MapGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMapGroupResource_Read_MapError exercises MapGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMapGroupResource_Read_MapError(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := MapGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMapGroupResource_Update_Happy exercises MapGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestMapGroupResource_Update_Happy(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := MapGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapGroupResource_Update_NilClient exercises MapGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapGroupResource_Update_NilClient(t *testing.T) {
	r := &MapGroupResource{}
	m := MapGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapGroupResource_Update_BuildError exercises MapGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapGroupResource_Update_BuildError(t *testing.T) {
	r := &MapGroupResource{client: newMalformedBaseURLClient(t)}
	m := MapGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapGroupResource_Update_SendError exercises MapGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapGroupResource_Update_SendError(t *testing.T) {
	r := &MapGroupResource{client: newTransportErrorClient(t)}
	m := MapGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapGroupResource_Update_APIError exercises MapGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapGroupResource_Update_APIError(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_map_group")
}

// TestMapGroupResource_Update_APIErrorReadBody exercises MapGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &MapGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMapGroupResource_Update_InvalidJSON exercises MapGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMapGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := MapGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMapGroupResource_Update_MapError exercises MapGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMapGroupResource_Update_MapError(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := MapGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMapGroupResource_Delete_Happy exercises MapGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestMapGroupResource_Delete_Happy(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 204, "")}
	m := MapGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapGroupResource_Delete_NilClient exercises MapGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapGroupResource_Delete_NilClient(t *testing.T) {
	r := &MapGroupResource{}
	m := MapGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapGroupResource_Delete_BuildError exercises MapGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapGroupResource_Delete_BuildError(t *testing.T) {
	r := &MapGroupResource{client: newMalformedBaseURLClient(t)}
	m := MapGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapGroupResource_Delete_SendError exercises MapGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapGroupResource_Delete_SendError(t *testing.T) {
	r := &MapGroupResource{client: newTransportErrorClient(t)}
	m := MapGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapGroupResource_Delete_NotFoundSuccess exercises MapGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestMapGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 404, "")}
	m := MapGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapGroupResource_Delete_APIError exercises MapGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapGroupResource_Delete_APIError(t *testing.T) {
	r := &MapGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_map_group")
}

// TestMapGroupResource_Delete_APIErrorReadBody exercises MapGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &MapGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
