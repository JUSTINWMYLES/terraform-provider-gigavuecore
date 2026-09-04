package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNodeResource_Create_Happy exercises NodeResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestNodeResource_Create_Happy(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 200, "{\"node_id\":\"example-id\"}")}
	m := NodeResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeResource_Create_NilClient exercises NodeResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNodeResource_Create_NilClient(t *testing.T) {
	r := &NodeResource{}
	m := NodeResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNodeResource_Create_BuildError exercises NodeResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNodeResource_Create_BuildError(t *testing.T) {
	r := &NodeResource{client: newMalformedBaseURLClient(t)}
	m := NodeResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNodeResource_Create_SendError exercises NodeResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestNodeResource_Create_SendError(t *testing.T) {
	r := &NodeResource{client: newTransportErrorClient(t)}
	m := NodeResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNodeResource_Create_APIError exercises NodeResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNodeResource_Create_APIError(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NodeResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_node")
}

// TestNodeResource_Create_APIErrorReadBody exercises NodeResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNodeResource_Create_APIErrorReadBody(t *testing.T) {
	r := &NodeResource{client: newMockClientReadErrorBody(t, 501)}
	m := NodeResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNodeResource_Create_InvalidJSON exercises NodeResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNodeResource_Create_InvalidJSON(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 200, "{{")}
	m := NodeResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNodeResource_Create_MapError exercises NodeResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNodeResource_Create_MapError(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 200, "{\"node_id\":12345}")}
	m := NodeResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNodeResource_Create_MissingID exercises NodeResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestNodeResource_Create_MissingID(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 200, "{}")}
	m := NodeResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestNodeResource_Create_LocationFallback exercises NodeResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestNodeResource_Create_LocationFallback(t *testing.T) {
	r := &NodeResource{client: newMockClientWithLocation(t, 200, "http://example.test/folders/example-id", "{}")}
	m := NodeResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.NodeId.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.NodeId.ValueString(), "example-id")
	}
}

// TestNodeResource_Read_Happy exercises NodeResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestNodeResource_Read_Happy(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 200, "{}")}
	m := NodeResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeResource_Read_NilClient exercises NodeResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNodeResource_Read_NilClient(t *testing.T) {
	r := &NodeResource{}
	m := NodeResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNodeResource_Read_BuildError exercises NodeResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNodeResource_Read_BuildError(t *testing.T) {
	r := &NodeResource{client: newMalformedBaseURLClient(t)}
	m := NodeResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNodeResource_Read_SendError exercises NodeResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestNodeResource_Read_SendError(t *testing.T) {
	r := &NodeResource{client: newTransportErrorClient(t)}
	m := NodeResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNodeResource_Read_NotFound exercises NodeResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestNodeResource_Read_NotFound(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 404, "")}
	m := NodeResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeResource_Read_APIError exercises NodeResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNodeResource_Read_APIError(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NodeResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_node")
}

// TestNodeResource_Read_APIErrorReadBody exercises NodeResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNodeResource_Read_APIErrorReadBody(t *testing.T) {
	r := &NodeResource{client: newMockClientReadErrorBody(t, 501)}
	m := NodeResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNodeResource_Read_InvalidJSON exercises NodeResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNodeResource_Read_InvalidJSON(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 200, "{{")}
	m := NodeResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNodeResource_Read_MapError exercises NodeResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNodeResource_Read_MapError(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 200, "{\"node_id\":12345}")}
	m := NodeResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNodeResource_Update_Happy exercises NodeResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestNodeResource_Update_Happy(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 200, "{}")}
	m := NodeResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeResource_Update_NilClient exercises NodeResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNodeResource_Update_NilClient(t *testing.T) {
	r := &NodeResource{}
	m := NodeResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNodeResource_Update_BuildError exercises NodeResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNodeResource_Update_BuildError(t *testing.T) {
	r := &NodeResource{client: newMalformedBaseURLClient(t)}
	m := NodeResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNodeResource_Update_SendError exercises NodeResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestNodeResource_Update_SendError(t *testing.T) {
	r := &NodeResource{client: newTransportErrorClient(t)}
	m := NodeResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNodeResource_Update_APIError exercises NodeResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNodeResource_Update_APIError(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NodeResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_node")
}

// TestNodeResource_Update_APIErrorReadBody exercises NodeResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNodeResource_Update_APIErrorReadBody(t *testing.T) {
	r := &NodeResource{client: newMockClientReadErrorBody(t, 501)}
	m := NodeResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNodeResource_Update_InvalidJSON exercises NodeResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNodeResource_Update_InvalidJSON(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 200, "{{")}
	m := NodeResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNodeResource_Update_MapError exercises NodeResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNodeResource_Update_MapError(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 200, "{\"node_id\":12345}")}
	m := NodeResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNodeResource_Delete_Happy exercises NodeResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestNodeResource_Delete_Happy(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 204, "")}
	m := NodeResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeResource_Delete_NilClient exercises NodeResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNodeResource_Delete_NilClient(t *testing.T) {
	r := &NodeResource{}
	m := NodeResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNodeResource_Delete_BuildError exercises NodeResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNodeResource_Delete_BuildError(t *testing.T) {
	r := &NodeResource{client: newMalformedBaseURLClient(t)}
	m := NodeResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNodeResource_Delete_SendError exercises NodeResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestNodeResource_Delete_SendError(t *testing.T) {
	r := &NodeResource{client: newTransportErrorClient(t)}
	m := NodeResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNodeResource_Delete_NotFoundSuccess exercises NodeResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestNodeResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 404, "")}
	m := NodeResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeResource_Delete_APIError exercises NodeResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNodeResource_Delete_APIError(t *testing.T) {
	r := &NodeResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NodeResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_node")
}

// TestNodeResource_Delete_APIErrorReadBody exercises NodeResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNodeResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &NodeResource{client: newMockClientReadErrorBody(t, 501)}
	m := NodeResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
