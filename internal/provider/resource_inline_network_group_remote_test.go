package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestInlineNetworkGroupResource_Create_Happy exercises InlineNetworkGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestInlineNetworkGroupResource_Create_Happy(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineNetworkGroupResource_Create_NilClient exercises InlineNetworkGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineNetworkGroupResource_Create_NilClient(t *testing.T) {
	r := &InlineNetworkGroupResource{}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineNetworkGroupResource_Create_BuildError exercises InlineNetworkGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineNetworkGroupResource_Create_BuildError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMalformedBaseURLClient(t)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineNetworkGroupResource_Create_SendError exercises InlineNetworkGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineNetworkGroupResource_Create_SendError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newTransportErrorClient(t)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineNetworkGroupResource_Create_APIError exercises InlineNetworkGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineNetworkGroupResource_Create_APIError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_inline_network_group")
}

// TestInlineNetworkGroupResource_Create_APIErrorReadBody exercises InlineNetworkGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineNetworkGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInlineNetworkGroupResource_Create_InvalidJSON exercises InlineNetworkGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInlineNetworkGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInlineNetworkGroupResource_Create_MapError exercises InlineNetworkGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInlineNetworkGroupResource_Create_MapError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInlineNetworkGroupResource_Create_MissingID exercises InlineNetworkGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestInlineNetworkGroupResource_Create_MissingID(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestInlineNetworkGroupResource_Create_LocationFallback exercises InlineNetworkGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestInlineNetworkGroupResource_Create_LocationFallback(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestInlineNetworkGroupResource_Read_Happy exercises InlineNetworkGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestInlineNetworkGroupResource_Read_Happy(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineNetworkGroupResource_Read_NilClient exercises InlineNetworkGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineNetworkGroupResource_Read_NilClient(t *testing.T) {
	r := &InlineNetworkGroupResource{}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineNetworkGroupResource_Read_BuildError exercises InlineNetworkGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineNetworkGroupResource_Read_BuildError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMalformedBaseURLClient(t)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineNetworkGroupResource_Read_SendError exercises InlineNetworkGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineNetworkGroupResource_Read_SendError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newTransportErrorClient(t)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineNetworkGroupResource_Read_NotFound exercises InlineNetworkGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestInlineNetworkGroupResource_Read_NotFound(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 404, "")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineNetworkGroupResource_Read_APIError exercises InlineNetworkGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineNetworkGroupResource_Read_APIError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_inline_network_group")
}

// TestInlineNetworkGroupResource_Read_APIErrorReadBody exercises InlineNetworkGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineNetworkGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInlineNetworkGroupResource_Read_InvalidJSON exercises InlineNetworkGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInlineNetworkGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInlineNetworkGroupResource_Read_MapError exercises InlineNetworkGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInlineNetworkGroupResource_Read_MapError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInlineNetworkGroupResource_Update_Happy exercises InlineNetworkGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestInlineNetworkGroupResource_Update_Happy(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineNetworkGroupResource_Update_NilClient exercises InlineNetworkGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineNetworkGroupResource_Update_NilClient(t *testing.T) {
	r := &InlineNetworkGroupResource{}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineNetworkGroupResource_Update_BuildError exercises InlineNetworkGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineNetworkGroupResource_Update_BuildError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMalformedBaseURLClient(t)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineNetworkGroupResource_Update_SendError exercises InlineNetworkGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineNetworkGroupResource_Update_SendError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newTransportErrorClient(t)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineNetworkGroupResource_Update_APIError exercises InlineNetworkGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineNetworkGroupResource_Update_APIError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_inline_network_group")
}

// TestInlineNetworkGroupResource_Update_APIErrorReadBody exercises InlineNetworkGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineNetworkGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInlineNetworkGroupResource_Update_InvalidJSON exercises InlineNetworkGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInlineNetworkGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInlineNetworkGroupResource_Update_MapError exercises InlineNetworkGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInlineNetworkGroupResource_Update_MapError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInlineNetworkGroupResource_Delete_Happy exercises InlineNetworkGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestInlineNetworkGroupResource_Delete_Happy(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 204, "")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineNetworkGroupResource_Delete_NilClient exercises InlineNetworkGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineNetworkGroupResource_Delete_NilClient(t *testing.T) {
	r := &InlineNetworkGroupResource{}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineNetworkGroupResource_Delete_BuildError exercises InlineNetworkGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineNetworkGroupResource_Delete_BuildError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMalformedBaseURLClient(t)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineNetworkGroupResource_Delete_SendError exercises InlineNetworkGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineNetworkGroupResource_Delete_SendError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newTransportErrorClient(t)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineNetworkGroupResource_Delete_NotFoundSuccess exercises InlineNetworkGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestInlineNetworkGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 404, "")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineNetworkGroupResource_Delete_APIError exercises InlineNetworkGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineNetworkGroupResource_Delete_APIError(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_inline_network_group")
}

// TestInlineNetworkGroupResource_Delete_APIErrorReadBody exercises InlineNetworkGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineNetworkGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &InlineNetworkGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineNetworkGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
