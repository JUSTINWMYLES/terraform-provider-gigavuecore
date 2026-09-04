package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNodeCredentialsResource_Create_Happy exercises NodeCredentialsResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestNodeCredentialsResource_Create_Happy(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 201, "{\"device_address\":\"example-id\"}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeCredentialsResource_Create_NilClient exercises NodeCredentialsResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNodeCredentialsResource_Create_NilClient(t *testing.T) {
	r := &NodeCredentialsResource{}
	m := NodeCredentialsResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNodeCredentialsResource_Create_BuildError exercises NodeCredentialsResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNodeCredentialsResource_Create_BuildError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMalformedBaseURLClient(t)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNodeCredentialsResource_Create_SendError exercises NodeCredentialsResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestNodeCredentialsResource_Create_SendError(t *testing.T) {
	r := &NodeCredentialsResource{client: newTransportErrorClient(t)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNodeCredentialsResource_Create_APIError exercises NodeCredentialsResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNodeCredentialsResource_Create_APIError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_node_credentials")
}

// TestNodeCredentialsResource_Create_APIErrorReadBody exercises NodeCredentialsResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNodeCredentialsResource_Create_APIErrorReadBody(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientReadErrorBody(t, 501)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNodeCredentialsResource_Create_InvalidJSON exercises NodeCredentialsResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNodeCredentialsResource_Create_InvalidJSON(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 201, "{{")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNodeCredentialsResource_Create_MapError exercises NodeCredentialsResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNodeCredentialsResource_Create_MapError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 201, "{\"device_address\":12345}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNodeCredentialsResource_Create_MissingID exercises NodeCredentialsResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestNodeCredentialsResource_Create_MissingID(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 201, "{}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestNodeCredentialsResource_Create_LocationFallback exercises NodeCredentialsResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestNodeCredentialsResource_Create_LocationFallback(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.DeviceAddress.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.DeviceAddress.ValueString(), "example-id")
	}
}

// TestNodeCredentialsResource_Read_Happy exercises NodeCredentialsResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestNodeCredentialsResource_Read_Happy(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 200, "{}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeCredentialsResource_Read_NilClient exercises NodeCredentialsResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNodeCredentialsResource_Read_NilClient(t *testing.T) {
	r := &NodeCredentialsResource{}
	m := NodeCredentialsResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNodeCredentialsResource_Read_BuildError exercises NodeCredentialsResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNodeCredentialsResource_Read_BuildError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMalformedBaseURLClient(t)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNodeCredentialsResource_Read_SendError exercises NodeCredentialsResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestNodeCredentialsResource_Read_SendError(t *testing.T) {
	r := &NodeCredentialsResource{client: newTransportErrorClient(t)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNodeCredentialsResource_Read_NotFound exercises NodeCredentialsResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestNodeCredentialsResource_Read_NotFound(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 404, "")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeCredentialsResource_Read_APIError exercises NodeCredentialsResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNodeCredentialsResource_Read_APIError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_node_credentials")
}

// TestNodeCredentialsResource_Read_APIErrorReadBody exercises NodeCredentialsResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNodeCredentialsResource_Read_APIErrorReadBody(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientReadErrorBody(t, 501)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNodeCredentialsResource_Read_InvalidJSON exercises NodeCredentialsResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNodeCredentialsResource_Read_InvalidJSON(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 200, "{{")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNodeCredentialsResource_Read_MapError exercises NodeCredentialsResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNodeCredentialsResource_Read_MapError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 200, "{\"device_address\":12345}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNodeCredentialsResource_Update_Happy exercises NodeCredentialsResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestNodeCredentialsResource_Update_Happy(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 200, "{}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeCredentialsResource_Update_NilClient exercises NodeCredentialsResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNodeCredentialsResource_Update_NilClient(t *testing.T) {
	r := &NodeCredentialsResource{}
	m := NodeCredentialsResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNodeCredentialsResource_Update_BuildError exercises NodeCredentialsResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNodeCredentialsResource_Update_BuildError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMalformedBaseURLClient(t)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNodeCredentialsResource_Update_SendError exercises NodeCredentialsResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestNodeCredentialsResource_Update_SendError(t *testing.T) {
	r := &NodeCredentialsResource{client: newTransportErrorClient(t)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNodeCredentialsResource_Update_APIError exercises NodeCredentialsResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNodeCredentialsResource_Update_APIError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_node_credentials")
}

// TestNodeCredentialsResource_Update_APIErrorReadBody exercises NodeCredentialsResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNodeCredentialsResource_Update_APIErrorReadBody(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientReadErrorBody(t, 501)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNodeCredentialsResource_Update_InvalidJSON exercises NodeCredentialsResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNodeCredentialsResource_Update_InvalidJSON(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 200, "{{")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNodeCredentialsResource_Update_MapError exercises NodeCredentialsResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNodeCredentialsResource_Update_MapError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 200, "{\"device_address\":12345}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNodeCredentialsResource_Delete_Happy exercises NodeCredentialsResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestNodeCredentialsResource_Delete_Happy(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 204, "")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeCredentialsResource_Delete_NilClient exercises NodeCredentialsResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNodeCredentialsResource_Delete_NilClient(t *testing.T) {
	r := &NodeCredentialsResource{}
	m := NodeCredentialsResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNodeCredentialsResource_Delete_BuildError exercises NodeCredentialsResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNodeCredentialsResource_Delete_BuildError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMalformedBaseURLClient(t)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNodeCredentialsResource_Delete_SendError exercises NodeCredentialsResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestNodeCredentialsResource_Delete_SendError(t *testing.T) {
	r := &NodeCredentialsResource{client: newTransportErrorClient(t)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNodeCredentialsResource_Delete_NotFoundSuccess exercises NodeCredentialsResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestNodeCredentialsResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 404, "")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNodeCredentialsResource_Delete_APIError exercises NodeCredentialsResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNodeCredentialsResource_Delete_APIError(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NodeCredentialsResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_node_credentials")
}

// TestNodeCredentialsResource_Delete_APIErrorReadBody exercises NodeCredentialsResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNodeCredentialsResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &NodeCredentialsResource{client: newMockClientReadErrorBody(t, 501)}
	m := NodeCredentialsResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
