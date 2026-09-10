package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestKeyResource_Create_Happy exercises KeyResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestKeyResource_Create_Happy(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := KeyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyResource_Create_NilClient exercises KeyResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestKeyResource_Create_NilClient(t *testing.T) {
	r := &KeyResource{}
	m := KeyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestKeyResource_Create_BuildError exercises KeyResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestKeyResource_Create_BuildError(t *testing.T) {
	r := &KeyResource{client: newMalformedBaseURLClient(t)}
	m := KeyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestKeyResource_Create_SendError exercises KeyResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestKeyResource_Create_SendError(t *testing.T) {
	r := &KeyResource{client: newTransportErrorClient(t)}
	m := KeyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestKeyResource_Create_APIError exercises KeyResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestKeyResource_Create_APIError(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := KeyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_key")
}

// TestKeyResource_Create_APIErrorReadBody exercises KeyResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestKeyResource_Create_APIErrorReadBody(t *testing.T) {
	r := &KeyResource{client: newMockClientReadErrorBody(t, 501)}
	m := KeyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestKeyResource_Create_InvalidJSON exercises KeyResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestKeyResource_Create_InvalidJSON(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 201, "{{")}
	m := KeyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestKeyResource_Create_MapError exercises KeyResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestKeyResource_Create_MapError(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := KeyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestKeyResource_Create_MissingID exercises KeyResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestKeyResource_Create_MissingID(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 201, "{}")}
	m := KeyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestKeyResource_Create_LocationFallback exercises KeyResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestKeyResource_Create_LocationFallback(t *testing.T) {
	r := &KeyResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := KeyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestKeyResource_Read_Happy exercises KeyResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestKeyResource_Read_Happy(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 200, "{}")}
	m := KeyResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyResource_Read_NilClient exercises KeyResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestKeyResource_Read_NilClient(t *testing.T) {
	r := &KeyResource{}
	m := KeyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestKeyResource_Read_BuildError exercises KeyResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestKeyResource_Read_BuildError(t *testing.T) {
	r := &KeyResource{client: newMalformedBaseURLClient(t)}
	m := KeyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestKeyResource_Read_SendError exercises KeyResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestKeyResource_Read_SendError(t *testing.T) {
	r := &KeyResource{client: newTransportErrorClient(t)}
	m := KeyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestKeyResource_Read_NotFound exercises KeyResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestKeyResource_Read_NotFound(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 404, "")}
	m := KeyResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyResource_Read_APIError exercises KeyResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestKeyResource_Read_APIError(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := KeyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_key")
}

// TestKeyResource_Read_APIErrorReadBody exercises KeyResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestKeyResource_Read_APIErrorReadBody(t *testing.T) {
	r := &KeyResource{client: newMockClientReadErrorBody(t, 501)}
	m := KeyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestKeyResource_Read_InvalidJSON exercises KeyResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestKeyResource_Read_InvalidJSON(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 200, "{{")}
	m := KeyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestKeyResource_Read_MapError exercises KeyResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestKeyResource_Read_MapError(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := KeyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestKeyResource_Update_Happy exercises KeyResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestKeyResource_Update_Happy(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 200, "{}")}
	m := KeyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyResource_Update_NilClient exercises KeyResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestKeyResource_Update_NilClient(t *testing.T) {
	r := &KeyResource{}
	m := KeyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestKeyResource_Update_BuildError exercises KeyResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestKeyResource_Update_BuildError(t *testing.T) {
	r := &KeyResource{client: newMalformedBaseURLClient(t)}
	m := KeyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestKeyResource_Update_SendError exercises KeyResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestKeyResource_Update_SendError(t *testing.T) {
	r := &KeyResource{client: newTransportErrorClient(t)}
	m := KeyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestKeyResource_Update_APIError exercises KeyResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestKeyResource_Update_APIError(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := KeyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_key")
}

// TestKeyResource_Update_APIErrorReadBody exercises KeyResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestKeyResource_Update_APIErrorReadBody(t *testing.T) {
	r := &KeyResource{client: newMockClientReadErrorBody(t, 501)}
	m := KeyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestKeyResource_Update_InvalidJSON exercises KeyResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestKeyResource_Update_InvalidJSON(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 200, "{{")}
	m := KeyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestKeyResource_Update_MapError exercises KeyResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestKeyResource_Update_MapError(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := KeyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestKeyResource_Delete_Happy exercises KeyResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestKeyResource_Delete_Happy(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 204, "")}
	m := KeyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyResource_Delete_NilClient exercises KeyResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestKeyResource_Delete_NilClient(t *testing.T) {
	r := &KeyResource{}
	m := KeyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestKeyResource_Delete_BuildError exercises KeyResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestKeyResource_Delete_BuildError(t *testing.T) {
	r := &KeyResource{client: newMalformedBaseURLClient(t)}
	m := KeyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestKeyResource_Delete_SendError exercises KeyResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestKeyResource_Delete_SendError(t *testing.T) {
	r := &KeyResource{client: newTransportErrorClient(t)}
	m := KeyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestKeyResource_Delete_NotFoundSuccess exercises KeyResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestKeyResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 404, "")}
	m := KeyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestKeyResource_Delete_APIError exercises KeyResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestKeyResource_Delete_APIError(t *testing.T) {
	r := &KeyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := KeyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_key")
}

// TestKeyResource_Delete_APIErrorReadBody exercises KeyResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestKeyResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &KeyResource{client: newMockClientReadErrorBody(t, 501)}
	m := KeyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
