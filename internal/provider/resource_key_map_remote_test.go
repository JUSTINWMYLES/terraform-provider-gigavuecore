package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestKeyMapResource_Create_Happy exercises KeyMapResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestKeyMapResource_Create_Happy(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := KeyMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyMapResource_Create_NilClient exercises KeyMapResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestKeyMapResource_Create_NilClient(t *testing.T) {
	r := &KeyMapResource{}
	m := KeyMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestKeyMapResource_Create_BuildError exercises KeyMapResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestKeyMapResource_Create_BuildError(t *testing.T) {
	r := &KeyMapResource{client: newMalformedBaseURLClient(t)}
	m := KeyMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestKeyMapResource_Create_SendError exercises KeyMapResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestKeyMapResource_Create_SendError(t *testing.T) {
	r := &KeyMapResource{client: newTransportErrorClient(t)}
	m := KeyMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestKeyMapResource_Create_APIError exercises KeyMapResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestKeyMapResource_Create_APIError(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := KeyMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_key_map")
}

// TestKeyMapResource_Create_APIErrorReadBody exercises KeyMapResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestKeyMapResource_Create_APIErrorReadBody(t *testing.T) {
	r := &KeyMapResource{client: newMockClientReadErrorBody(t, 501)}
	m := KeyMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestKeyMapResource_Create_InvalidJSON exercises KeyMapResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestKeyMapResource_Create_InvalidJSON(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 201, "{{")}
	m := KeyMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestKeyMapResource_Create_MapError exercises KeyMapResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestKeyMapResource_Create_MapError(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := KeyMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestKeyMapResource_Create_MissingID exercises KeyMapResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestKeyMapResource_Create_MissingID(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 201, "{}")}
	m := KeyMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestKeyMapResource_Create_LocationFallback exercises KeyMapResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestKeyMapResource_Create_LocationFallback(t *testing.T) {
	r := &KeyMapResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := KeyMapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestKeyMapResource_Read_Happy exercises KeyMapResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestKeyMapResource_Read_Happy(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 200, "{}")}
	m := KeyMapResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyMapResource_Read_NilClient exercises KeyMapResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestKeyMapResource_Read_NilClient(t *testing.T) {
	r := &KeyMapResource{}
	m := KeyMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestKeyMapResource_Read_BuildError exercises KeyMapResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestKeyMapResource_Read_BuildError(t *testing.T) {
	r := &KeyMapResource{client: newMalformedBaseURLClient(t)}
	m := KeyMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestKeyMapResource_Read_SendError exercises KeyMapResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestKeyMapResource_Read_SendError(t *testing.T) {
	r := &KeyMapResource{client: newTransportErrorClient(t)}
	m := KeyMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestKeyMapResource_Read_NotFound exercises KeyMapResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestKeyMapResource_Read_NotFound(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 404, "")}
	m := KeyMapResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyMapResource_Read_APIError exercises KeyMapResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestKeyMapResource_Read_APIError(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := KeyMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_key_map")
}

// TestKeyMapResource_Read_APIErrorReadBody exercises KeyMapResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestKeyMapResource_Read_APIErrorReadBody(t *testing.T) {
	r := &KeyMapResource{client: newMockClientReadErrorBody(t, 501)}
	m := KeyMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestKeyMapResource_Read_InvalidJSON exercises KeyMapResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestKeyMapResource_Read_InvalidJSON(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 200, "{{")}
	m := KeyMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestKeyMapResource_Read_MapError exercises KeyMapResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestKeyMapResource_Read_MapError(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := KeyMapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestKeyMapResource_Update_Happy exercises KeyMapResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestKeyMapResource_Update_Happy(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 200, "{}")}
	m := KeyMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyMapResource_Update_NilClient exercises KeyMapResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestKeyMapResource_Update_NilClient(t *testing.T) {
	r := &KeyMapResource{}
	m := KeyMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestKeyMapResource_Update_BuildError exercises KeyMapResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestKeyMapResource_Update_BuildError(t *testing.T) {
	r := &KeyMapResource{client: newMalformedBaseURLClient(t)}
	m := KeyMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestKeyMapResource_Update_SendError exercises KeyMapResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestKeyMapResource_Update_SendError(t *testing.T) {
	r := &KeyMapResource{client: newTransportErrorClient(t)}
	m := KeyMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestKeyMapResource_Update_APIError exercises KeyMapResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestKeyMapResource_Update_APIError(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := KeyMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_key_map")
}

// TestKeyMapResource_Update_APIErrorReadBody exercises KeyMapResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestKeyMapResource_Update_APIErrorReadBody(t *testing.T) {
	r := &KeyMapResource{client: newMockClientReadErrorBody(t, 501)}
	m := KeyMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestKeyMapResource_Update_InvalidJSON exercises KeyMapResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestKeyMapResource_Update_InvalidJSON(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 200, "{{")}
	m := KeyMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestKeyMapResource_Update_MapError exercises KeyMapResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestKeyMapResource_Update_MapError(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := KeyMapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestKeyMapResource_Delete_Happy exercises KeyMapResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestKeyMapResource_Delete_Happy(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 204, "")}
	m := KeyMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyMapResource_Delete_NilClient exercises KeyMapResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestKeyMapResource_Delete_NilClient(t *testing.T) {
	r := &KeyMapResource{}
	m := KeyMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestKeyMapResource_Delete_BuildError exercises KeyMapResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestKeyMapResource_Delete_BuildError(t *testing.T) {
	r := &KeyMapResource{client: newMalformedBaseURLClient(t)}
	m := KeyMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestKeyMapResource_Delete_SendError exercises KeyMapResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestKeyMapResource_Delete_SendError(t *testing.T) {
	r := &KeyMapResource{client: newTransportErrorClient(t)}
	m := KeyMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestKeyMapResource_Delete_NotFoundSuccess exercises KeyMapResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestKeyMapResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 404, "")}
	m := KeyMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyMapResource_Delete_APIError exercises KeyMapResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestKeyMapResource_Delete_APIError(t *testing.T) {
	r := &KeyMapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := KeyMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_key_map")
}

// TestKeyMapResource_Delete_APIErrorReadBody exercises KeyMapResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestKeyMapResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &KeyMapResource{client: newMockClientReadErrorBody(t, 501)}
	m := KeyMapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
