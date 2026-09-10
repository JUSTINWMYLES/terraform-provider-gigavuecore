package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFlexInlineResource_Create_Happy exercises FlexInlineResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestFlexInlineResource_Create_Happy(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := FlexInlineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFlexInlineResource_Create_NilClient exercises FlexInlineResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFlexInlineResource_Create_NilClient(t *testing.T) {
	r := &FlexInlineResource{}
	m := FlexInlineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFlexInlineResource_Create_BuildError exercises FlexInlineResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFlexInlineResource_Create_BuildError(t *testing.T) {
	r := &FlexInlineResource{client: newMalformedBaseURLClient(t)}
	m := FlexInlineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFlexInlineResource_Create_SendError exercises FlexInlineResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestFlexInlineResource_Create_SendError(t *testing.T) {
	r := &FlexInlineResource{client: newTransportErrorClient(t)}
	m := FlexInlineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFlexInlineResource_Create_APIError exercises FlexInlineResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFlexInlineResource_Create_APIError(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FlexInlineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_flex_inline")
}

// TestFlexInlineResource_Create_APIErrorReadBody exercises FlexInlineResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFlexInlineResource_Create_APIErrorReadBody(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientReadErrorBody(t, 501)}
	m := FlexInlineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFlexInlineResource_Create_InvalidJSON exercises FlexInlineResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFlexInlineResource_Create_InvalidJSON(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 201, "{{")}
	m := FlexInlineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFlexInlineResource_Create_MapError exercises FlexInlineResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFlexInlineResource_Create_MapError(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := FlexInlineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFlexInlineResource_Create_MissingID exercises FlexInlineResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestFlexInlineResource_Create_MissingID(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 201, "{}")}
	m := FlexInlineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestFlexInlineResource_Create_LocationFallback exercises FlexInlineResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestFlexInlineResource_Create_LocationFallback(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := FlexInlineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestFlexInlineResource_Read_Happy exercises FlexInlineResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestFlexInlineResource_Read_Happy(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 200, "{}")}
	m := FlexInlineResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFlexInlineResource_Read_NilClient exercises FlexInlineResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFlexInlineResource_Read_NilClient(t *testing.T) {
	r := &FlexInlineResource{}
	m := FlexInlineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFlexInlineResource_Read_BuildError exercises FlexInlineResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFlexInlineResource_Read_BuildError(t *testing.T) {
	r := &FlexInlineResource{client: newMalformedBaseURLClient(t)}
	m := FlexInlineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFlexInlineResource_Read_SendError exercises FlexInlineResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestFlexInlineResource_Read_SendError(t *testing.T) {
	r := &FlexInlineResource{client: newTransportErrorClient(t)}
	m := FlexInlineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFlexInlineResource_Read_NotFound exercises FlexInlineResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestFlexInlineResource_Read_NotFound(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 404, "")}
	m := FlexInlineResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFlexInlineResource_Read_APIError exercises FlexInlineResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFlexInlineResource_Read_APIError(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FlexInlineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_flex_inline")
}

// TestFlexInlineResource_Read_APIErrorReadBody exercises FlexInlineResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFlexInlineResource_Read_APIErrorReadBody(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientReadErrorBody(t, 501)}
	m := FlexInlineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFlexInlineResource_Read_InvalidJSON exercises FlexInlineResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFlexInlineResource_Read_InvalidJSON(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 200, "{{")}
	m := FlexInlineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFlexInlineResource_Read_MapError exercises FlexInlineResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFlexInlineResource_Read_MapError(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := FlexInlineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFlexInlineResource_Update_Happy exercises FlexInlineResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestFlexInlineResource_Update_Happy(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 200, "{}")}
	m := FlexInlineResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFlexInlineResource_Update_NilClient exercises FlexInlineResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFlexInlineResource_Update_NilClient(t *testing.T) {
	r := &FlexInlineResource{}
	m := FlexInlineResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFlexInlineResource_Update_BuildError exercises FlexInlineResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFlexInlineResource_Update_BuildError(t *testing.T) {
	r := &FlexInlineResource{client: newMalformedBaseURLClient(t)}
	m := FlexInlineResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFlexInlineResource_Update_SendError exercises FlexInlineResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestFlexInlineResource_Update_SendError(t *testing.T) {
	r := &FlexInlineResource{client: newTransportErrorClient(t)}
	m := FlexInlineResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFlexInlineResource_Update_APIError exercises FlexInlineResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFlexInlineResource_Update_APIError(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FlexInlineResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_flex_inline")
}

// TestFlexInlineResource_Update_APIErrorReadBody exercises FlexInlineResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFlexInlineResource_Update_APIErrorReadBody(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientReadErrorBody(t, 501)}
	m := FlexInlineResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFlexInlineResource_Update_InvalidJSON exercises FlexInlineResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFlexInlineResource_Update_InvalidJSON(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 200, "{{")}
	m := FlexInlineResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFlexInlineResource_Update_MapError exercises FlexInlineResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFlexInlineResource_Update_MapError(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := FlexInlineResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFlexInlineResource_Delete_Happy exercises FlexInlineResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestFlexInlineResource_Delete_Happy(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 204, "")}
	m := FlexInlineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFlexInlineResource_Delete_NilClient exercises FlexInlineResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFlexInlineResource_Delete_NilClient(t *testing.T) {
	r := &FlexInlineResource{}
	m := FlexInlineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFlexInlineResource_Delete_BuildError exercises FlexInlineResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFlexInlineResource_Delete_BuildError(t *testing.T) {
	r := &FlexInlineResource{client: newMalformedBaseURLClient(t)}
	m := FlexInlineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFlexInlineResource_Delete_SendError exercises FlexInlineResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestFlexInlineResource_Delete_SendError(t *testing.T) {
	r := &FlexInlineResource{client: newTransportErrorClient(t)}
	m := FlexInlineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFlexInlineResource_Delete_NotFoundSuccess exercises FlexInlineResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestFlexInlineResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 404, "")}
	m := FlexInlineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFlexInlineResource_Delete_APIError exercises FlexInlineResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFlexInlineResource_Delete_APIError(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FlexInlineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_flex_inline")
}

// TestFlexInlineResource_Delete_APIErrorReadBody exercises FlexInlineResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFlexInlineResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &FlexInlineResource{client: newMockClientReadErrorBody(t, 501)}
	m := FlexInlineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
