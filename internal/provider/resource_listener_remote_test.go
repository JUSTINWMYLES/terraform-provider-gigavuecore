package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestListenerResource_Create_Happy exercises ListenerResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestListenerResource_Create_Happy(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ListenerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListenerResource_Create_NilClient exercises ListenerResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListenerResource_Create_NilClient(t *testing.T) {
	r := &ListenerResource{}
	m := ListenerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListenerResource_Create_BuildError exercises ListenerResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestListenerResource_Create_BuildError(t *testing.T) {
	r := &ListenerResource{client: newMalformedBaseURLClient(t)}
	m := ListenerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestListenerResource_Create_SendError exercises ListenerResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestListenerResource_Create_SendError(t *testing.T) {
	r := &ListenerResource{client: newTransportErrorClient(t)}
	m := ListenerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestListenerResource_Create_APIError exercises ListenerResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestListenerResource_Create_APIError(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ListenerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_listener")
}

// TestListenerResource_Create_APIErrorReadBody exercises ListenerResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestListenerResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ListenerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ListenerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestListenerResource_Create_InvalidJSON exercises ListenerResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestListenerResource_Create_InvalidJSON(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 201, "{{")}
	m := ListenerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestListenerResource_Create_MapError exercises ListenerResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestListenerResource_Create_MapError(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ListenerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestListenerResource_Create_MissingID exercises ListenerResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestListenerResource_Create_MissingID(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 201, "{}")}
	m := ListenerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestListenerResource_Create_LocationFallback exercises ListenerResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestListenerResource_Create_LocationFallback(t *testing.T) {
	r := &ListenerResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ListenerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestListenerResource_Read_Happy exercises ListenerResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestListenerResource_Read_Happy(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 200, "{}")}
	m := ListenerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestListenerResource_Read_NilClient exercises ListenerResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListenerResource_Read_NilClient(t *testing.T) {
	r := &ListenerResource{}
	m := ListenerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListenerResource_Read_BuildError exercises ListenerResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestListenerResource_Read_BuildError(t *testing.T) {
	r := &ListenerResource{client: newMalformedBaseURLClient(t)}
	m := ListenerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestListenerResource_Read_SendError exercises ListenerResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestListenerResource_Read_SendError(t *testing.T) {
	r := &ListenerResource{client: newTransportErrorClient(t)}
	m := ListenerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestListenerResource_Read_NotFound exercises ListenerResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestListenerResource_Read_NotFound(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 404, "")}
	m := ListenerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestListenerResource_Read_APIError exercises ListenerResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestListenerResource_Read_APIError(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ListenerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_listener")
}

// TestListenerResource_Read_APIErrorReadBody exercises ListenerResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestListenerResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ListenerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ListenerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestListenerResource_Read_InvalidJSON exercises ListenerResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestListenerResource_Read_InvalidJSON(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 200, "{{")}
	m := ListenerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestListenerResource_Read_MapError exercises ListenerResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestListenerResource_Read_MapError(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ListenerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestListenerResource_Update_Happy exercises ListenerResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestListenerResource_Update_Happy(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 200, "{}")}
	m := ListenerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListenerResource_Update_NilClient exercises ListenerResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListenerResource_Update_NilClient(t *testing.T) {
	r := &ListenerResource{}
	m := ListenerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListenerResource_Update_BuildError exercises ListenerResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestListenerResource_Update_BuildError(t *testing.T) {
	r := &ListenerResource{client: newMalformedBaseURLClient(t)}
	m := ListenerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestListenerResource_Update_SendError exercises ListenerResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestListenerResource_Update_SendError(t *testing.T) {
	r := &ListenerResource{client: newTransportErrorClient(t)}
	m := ListenerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestListenerResource_Update_APIError exercises ListenerResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestListenerResource_Update_APIError(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ListenerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_listener")
}

// TestListenerResource_Update_APIErrorReadBody exercises ListenerResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestListenerResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ListenerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ListenerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestListenerResource_Update_InvalidJSON exercises ListenerResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestListenerResource_Update_InvalidJSON(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 200, "{{")}
	m := ListenerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestListenerResource_Update_MapError exercises ListenerResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestListenerResource_Update_MapError(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ListenerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestListenerResource_Delete_Happy exercises ListenerResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestListenerResource_Delete_Happy(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 204, "")}
	m := ListenerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListenerResource_Delete_NilClient exercises ListenerResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListenerResource_Delete_NilClient(t *testing.T) {
	r := &ListenerResource{}
	m := ListenerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListenerResource_Delete_BuildError exercises ListenerResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestListenerResource_Delete_BuildError(t *testing.T) {
	r := &ListenerResource{client: newMalformedBaseURLClient(t)}
	m := ListenerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestListenerResource_Delete_SendError exercises ListenerResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestListenerResource_Delete_SendError(t *testing.T) {
	r := &ListenerResource{client: newTransportErrorClient(t)}
	m := ListenerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestListenerResource_Delete_NotFoundSuccess exercises ListenerResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestListenerResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 404, "")}
	m := ListenerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListenerResource_Delete_APIError exercises ListenerResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestListenerResource_Delete_APIError(t *testing.T) {
	r := &ListenerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ListenerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_listener")
}

// TestListenerResource_Delete_APIErrorReadBody exercises ListenerResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestListenerResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ListenerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ListenerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
