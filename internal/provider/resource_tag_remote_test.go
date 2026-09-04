package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestTagResource_Create_Happy exercises TagResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestTagResource_Create_Happy(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 201, "{\"tag_key\":\"example-id\"}")}
	m := TagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTagResource_Create_NilClient exercises TagResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTagResource_Create_NilClient(t *testing.T) {
	r := &TagResource{}
	m := TagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTagResource_Create_BuildError exercises TagResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTagResource_Create_BuildError(t *testing.T) {
	r := &TagResource{client: newMalformedBaseURLClient(t)}
	m := TagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTagResource_Create_SendError exercises TagResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestTagResource_Create_SendError(t *testing.T) {
	r := &TagResource{client: newTransportErrorClient(t)}
	m := TagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTagResource_Create_APIError exercises TagResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTagResource_Create_APIError(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_tag")
}

// TestTagResource_Create_APIErrorReadBody exercises TagResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTagResource_Create_APIErrorReadBody(t *testing.T) {
	r := &TagResource{client: newMockClientReadErrorBody(t, 501)}
	m := TagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTagResource_Create_InvalidJSON exercises TagResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTagResource_Create_InvalidJSON(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 201, "{{")}
	m := TagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTagResource_Create_MapError exercises TagResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTagResource_Create_MapError(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 201, "{\"tag_key\":12345}")}
	m := TagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTagResource_Create_MissingID exercises TagResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestTagResource_Create_MissingID(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 201, "{}")}
	m := TagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestTagResource_Create_LocationFallback exercises TagResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestTagResource_Create_LocationFallback(t *testing.T) {
	r := &TagResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := TagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.TagKey.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.TagKey.ValueString(), "example-id")
	}
}

// TestTagResource_Read_Happy exercises TagResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestTagResource_Read_Happy(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 200, "{}")}
	m := TagResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTagResource_Read_NilClient exercises TagResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTagResource_Read_NilClient(t *testing.T) {
	r := &TagResource{}
	m := TagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTagResource_Read_BuildError exercises TagResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTagResource_Read_BuildError(t *testing.T) {
	r := &TagResource{client: newMalformedBaseURLClient(t)}
	m := TagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTagResource_Read_SendError exercises TagResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestTagResource_Read_SendError(t *testing.T) {
	r := &TagResource{client: newTransportErrorClient(t)}
	m := TagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTagResource_Read_NotFound exercises TagResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestTagResource_Read_NotFound(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 404, "")}
	m := TagResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTagResource_Read_APIError exercises TagResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTagResource_Read_APIError(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_tag")
}

// TestTagResource_Read_APIErrorReadBody exercises TagResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTagResource_Read_APIErrorReadBody(t *testing.T) {
	r := &TagResource{client: newMockClientReadErrorBody(t, 501)}
	m := TagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTagResource_Read_InvalidJSON exercises TagResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTagResource_Read_InvalidJSON(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 200, "{{")}
	m := TagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTagResource_Read_MapError exercises TagResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTagResource_Read_MapError(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 200, "{\"tag_key\":12345}")}
	m := TagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTagResource_Update_Happy exercises TagResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestTagResource_Update_Happy(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 200, "{}")}
	m := TagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTagResource_Update_NilClient exercises TagResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTagResource_Update_NilClient(t *testing.T) {
	r := &TagResource{}
	m := TagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTagResource_Update_BuildError exercises TagResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTagResource_Update_BuildError(t *testing.T) {
	r := &TagResource{client: newMalformedBaseURLClient(t)}
	m := TagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTagResource_Update_SendError exercises TagResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestTagResource_Update_SendError(t *testing.T) {
	r := &TagResource{client: newTransportErrorClient(t)}
	m := TagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTagResource_Update_APIError exercises TagResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTagResource_Update_APIError(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_tag")
}

// TestTagResource_Update_APIErrorReadBody exercises TagResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTagResource_Update_APIErrorReadBody(t *testing.T) {
	r := &TagResource{client: newMockClientReadErrorBody(t, 501)}
	m := TagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTagResource_Update_InvalidJSON exercises TagResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTagResource_Update_InvalidJSON(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 200, "{{")}
	m := TagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTagResource_Update_MapError exercises TagResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTagResource_Update_MapError(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 200, "{\"tag_key\":12345}")}
	m := TagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTagResource_Delete_Happy exercises TagResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestTagResource_Delete_Happy(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 204, "")}
	m := TagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTagResource_Delete_NilClient exercises TagResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTagResource_Delete_NilClient(t *testing.T) {
	r := &TagResource{}
	m := TagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTagResource_Delete_BuildError exercises TagResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTagResource_Delete_BuildError(t *testing.T) {
	r := &TagResource{client: newMalformedBaseURLClient(t)}
	m := TagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTagResource_Delete_SendError exercises TagResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestTagResource_Delete_SendError(t *testing.T) {
	r := &TagResource{client: newTransportErrorClient(t)}
	m := TagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTagResource_Delete_NotFoundSuccess exercises TagResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestTagResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 404, "")}
	m := TagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTagResource_Delete_APIError exercises TagResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTagResource_Delete_APIError(t *testing.T) {
	r := &TagResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_tag")
}

// TestTagResource_Delete_APIErrorReadBody exercises TagResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTagResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &TagResource{client: newMockClientReadErrorBody(t, 501)}
	m := TagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
