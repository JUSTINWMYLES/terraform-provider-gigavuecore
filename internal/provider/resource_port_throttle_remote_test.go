package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPortThrottleResource_Create_Happy exercises PortThrottleResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestPortThrottleResource_Create_Happy(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := PortThrottleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortThrottleResource_Create_NilClient exercises PortThrottleResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortThrottleResource_Create_NilClient(t *testing.T) {
	r := &PortThrottleResource{}
	m := PortThrottleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortThrottleResource_Create_BuildError exercises PortThrottleResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortThrottleResource_Create_BuildError(t *testing.T) {
	r := &PortThrottleResource{client: newMalformedBaseURLClient(t)}
	m := PortThrottleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortThrottleResource_Create_SendError exercises PortThrottleResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortThrottleResource_Create_SendError(t *testing.T) {
	r := &PortThrottleResource{client: newTransportErrorClient(t)}
	m := PortThrottleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortThrottleResource_Create_APIError exercises PortThrottleResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortThrottleResource_Create_APIError(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortThrottleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_port_throttle")
}

// TestPortThrottleResource_Create_APIErrorReadBody exercises PortThrottleResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortThrottleResource_Create_APIErrorReadBody(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortThrottleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortThrottleResource_Create_InvalidJSON exercises PortThrottleResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortThrottleResource_Create_InvalidJSON(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 201, "{{")}
	m := PortThrottleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortThrottleResource_Create_MapError exercises PortThrottleResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortThrottleResource_Create_MapError(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := PortThrottleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortThrottleResource_Create_MissingID exercises PortThrottleResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestPortThrottleResource_Create_MissingID(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 201, "{}")}
	m := PortThrottleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestPortThrottleResource_Create_LocationFallback exercises PortThrottleResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestPortThrottleResource_Create_LocationFallback(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := PortThrottleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestPortThrottleResource_Read_Happy exercises PortThrottleResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestPortThrottleResource_Read_Happy(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 200, "{}")}
	m := PortThrottleResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortThrottleResource_Read_NilClient exercises PortThrottleResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortThrottleResource_Read_NilClient(t *testing.T) {
	r := &PortThrottleResource{}
	m := PortThrottleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortThrottleResource_Read_BuildError exercises PortThrottleResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortThrottleResource_Read_BuildError(t *testing.T) {
	r := &PortThrottleResource{client: newMalformedBaseURLClient(t)}
	m := PortThrottleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortThrottleResource_Read_SendError exercises PortThrottleResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortThrottleResource_Read_SendError(t *testing.T) {
	r := &PortThrottleResource{client: newTransportErrorClient(t)}
	m := PortThrottleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortThrottleResource_Read_NotFound exercises PortThrottleResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestPortThrottleResource_Read_NotFound(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 404, "")}
	m := PortThrottleResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortThrottleResource_Read_APIError exercises PortThrottleResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortThrottleResource_Read_APIError(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortThrottleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_port_throttle")
}

// TestPortThrottleResource_Read_APIErrorReadBody exercises PortThrottleResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortThrottleResource_Read_APIErrorReadBody(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortThrottleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortThrottleResource_Read_InvalidJSON exercises PortThrottleResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortThrottleResource_Read_InvalidJSON(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortThrottleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortThrottleResource_Read_MapError exercises PortThrottleResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortThrottleResource_Read_MapError(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := PortThrottleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortThrottleResource_Update_Happy exercises PortThrottleResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestPortThrottleResource_Update_Happy(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 200, "{}")}
	m := PortThrottleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortThrottleResource_Update_NilClient exercises PortThrottleResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortThrottleResource_Update_NilClient(t *testing.T) {
	r := &PortThrottleResource{}
	m := PortThrottleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortThrottleResource_Update_BuildError exercises PortThrottleResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortThrottleResource_Update_BuildError(t *testing.T) {
	r := &PortThrottleResource{client: newMalformedBaseURLClient(t)}
	m := PortThrottleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortThrottleResource_Update_SendError exercises PortThrottleResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortThrottleResource_Update_SendError(t *testing.T) {
	r := &PortThrottleResource{client: newTransportErrorClient(t)}
	m := PortThrottleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortThrottleResource_Update_APIError exercises PortThrottleResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortThrottleResource_Update_APIError(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortThrottleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_port_throttle")
}

// TestPortThrottleResource_Update_APIErrorReadBody exercises PortThrottleResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortThrottleResource_Update_APIErrorReadBody(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortThrottleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortThrottleResource_Update_InvalidJSON exercises PortThrottleResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortThrottleResource_Update_InvalidJSON(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortThrottleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortThrottleResource_Update_MapError exercises PortThrottleResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortThrottleResource_Update_MapError(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := PortThrottleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortThrottleResource_Delete_Happy exercises PortThrottleResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestPortThrottleResource_Delete_Happy(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 204, "")}
	m := PortThrottleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortThrottleResource_Delete_NilClient exercises PortThrottleResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortThrottleResource_Delete_NilClient(t *testing.T) {
	r := &PortThrottleResource{}
	m := PortThrottleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortThrottleResource_Delete_BuildError exercises PortThrottleResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortThrottleResource_Delete_BuildError(t *testing.T) {
	r := &PortThrottleResource{client: newMalformedBaseURLClient(t)}
	m := PortThrottleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortThrottleResource_Delete_SendError exercises PortThrottleResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortThrottleResource_Delete_SendError(t *testing.T) {
	r := &PortThrottleResource{client: newTransportErrorClient(t)}
	m := PortThrottleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortThrottleResource_Delete_NotFoundSuccess exercises PortThrottleResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestPortThrottleResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 404, "")}
	m := PortThrottleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortThrottleResource_Delete_APIError exercises PortThrottleResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortThrottleResource_Delete_APIError(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortThrottleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_port_throttle")
}

// TestPortThrottleResource_Delete_APIErrorReadBody exercises PortThrottleResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortThrottleResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &PortThrottleResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortThrottleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
