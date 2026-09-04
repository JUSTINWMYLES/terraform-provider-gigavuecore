package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSerialToolGroupResource_Create_Happy exercises SerialToolGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestSerialToolGroupResource_Create_Happy(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSerialToolGroupResource_Create_NilClient exercises SerialToolGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSerialToolGroupResource_Create_NilClient(t *testing.T) {
	r := &SerialToolGroupResource{}
	m := SerialToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSerialToolGroupResource_Create_BuildError exercises SerialToolGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSerialToolGroupResource_Create_BuildError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMalformedBaseURLClient(t)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSerialToolGroupResource_Create_SendError exercises SerialToolGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestSerialToolGroupResource_Create_SendError(t *testing.T) {
	r := &SerialToolGroupResource{client: newTransportErrorClient(t)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSerialToolGroupResource_Create_APIError exercises SerialToolGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSerialToolGroupResource_Create_APIError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_serial_tool_group")
}

// TestSerialToolGroupResource_Create_APIErrorReadBody exercises SerialToolGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSerialToolGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSerialToolGroupResource_Create_InvalidJSON exercises SerialToolGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSerialToolGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSerialToolGroupResource_Create_MapError exercises SerialToolGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSerialToolGroupResource_Create_MapError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSerialToolGroupResource_Create_MissingID exercises SerialToolGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestSerialToolGroupResource_Create_MissingID(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestSerialToolGroupResource_Create_LocationFallback exercises SerialToolGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestSerialToolGroupResource_Create_LocationFallback(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestSerialToolGroupResource_Read_Happy exercises SerialToolGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestSerialToolGroupResource_Read_Happy(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSerialToolGroupResource_Read_NilClient exercises SerialToolGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSerialToolGroupResource_Read_NilClient(t *testing.T) {
	r := &SerialToolGroupResource{}
	m := SerialToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSerialToolGroupResource_Read_BuildError exercises SerialToolGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSerialToolGroupResource_Read_BuildError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMalformedBaseURLClient(t)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSerialToolGroupResource_Read_SendError exercises SerialToolGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestSerialToolGroupResource_Read_SendError(t *testing.T) {
	r := &SerialToolGroupResource{client: newTransportErrorClient(t)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSerialToolGroupResource_Read_NotFound exercises SerialToolGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestSerialToolGroupResource_Read_NotFound(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 404, "")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSerialToolGroupResource_Read_APIError exercises SerialToolGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSerialToolGroupResource_Read_APIError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_serial_tool_group")
}

// TestSerialToolGroupResource_Read_APIErrorReadBody exercises SerialToolGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSerialToolGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSerialToolGroupResource_Read_InvalidJSON exercises SerialToolGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSerialToolGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSerialToolGroupResource_Read_MapError exercises SerialToolGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSerialToolGroupResource_Read_MapError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSerialToolGroupResource_Update_Happy exercises SerialToolGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestSerialToolGroupResource_Update_Happy(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSerialToolGroupResource_Update_NilClient exercises SerialToolGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSerialToolGroupResource_Update_NilClient(t *testing.T) {
	r := &SerialToolGroupResource{}
	m := SerialToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSerialToolGroupResource_Update_BuildError exercises SerialToolGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSerialToolGroupResource_Update_BuildError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMalformedBaseURLClient(t)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSerialToolGroupResource_Update_SendError exercises SerialToolGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestSerialToolGroupResource_Update_SendError(t *testing.T) {
	r := &SerialToolGroupResource{client: newTransportErrorClient(t)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSerialToolGroupResource_Update_APIError exercises SerialToolGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSerialToolGroupResource_Update_APIError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_serial_tool_group")
}

// TestSerialToolGroupResource_Update_APIErrorReadBody exercises SerialToolGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSerialToolGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSerialToolGroupResource_Update_InvalidJSON exercises SerialToolGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSerialToolGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSerialToolGroupResource_Update_MapError exercises SerialToolGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSerialToolGroupResource_Update_MapError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSerialToolGroupResource_Delete_Happy exercises SerialToolGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestSerialToolGroupResource_Delete_Happy(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 204, "")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSerialToolGroupResource_Delete_NilClient exercises SerialToolGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSerialToolGroupResource_Delete_NilClient(t *testing.T) {
	r := &SerialToolGroupResource{}
	m := SerialToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSerialToolGroupResource_Delete_BuildError exercises SerialToolGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSerialToolGroupResource_Delete_BuildError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMalformedBaseURLClient(t)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSerialToolGroupResource_Delete_SendError exercises SerialToolGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestSerialToolGroupResource_Delete_SendError(t *testing.T) {
	r := &SerialToolGroupResource{client: newTransportErrorClient(t)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSerialToolGroupResource_Delete_NotFoundSuccess exercises SerialToolGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestSerialToolGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 404, "")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSerialToolGroupResource_Delete_APIError exercises SerialToolGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSerialToolGroupResource_Delete_APIError(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SerialToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_serial_tool_group")
}

// TestSerialToolGroupResource_Delete_APIErrorReadBody exercises SerialToolGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSerialToolGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &SerialToolGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := SerialToolGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
