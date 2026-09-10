package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMapResource_Create_Happy exercises MapResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestMapResource_Create_Happy(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := MapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapResource_Create_NilClient exercises MapResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapResource_Create_NilClient(t *testing.T) {
	r := &MapResource{}
	m := MapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapResource_Create_BuildError exercises MapResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapResource_Create_BuildError(t *testing.T) {
	r := &MapResource{client: newMalformedBaseURLClient(t)}
	m := MapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapResource_Create_SendError exercises MapResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapResource_Create_SendError(t *testing.T) {
	r := &MapResource{client: newTransportErrorClient(t)}
	m := MapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapResource_Create_APIError exercises MapResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapResource_Create_APIError(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_map")
}

// TestMapResource_Create_APIErrorReadBody exercises MapResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapResource_Create_APIErrorReadBody(t *testing.T) {
	r := &MapResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMapResource_Create_InvalidJSON exercises MapResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMapResource_Create_InvalidJSON(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 201, "{{")}
	m := MapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMapResource_Create_MapError exercises MapResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMapResource_Create_MapError(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := MapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMapResource_Create_MissingID exercises MapResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestMapResource_Create_MissingID(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 201, "{}")}
	m := MapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestMapResource_Create_LocationFallback exercises MapResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestMapResource_Create_LocationFallback(t *testing.T) {
	r := &MapResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := MapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestMapResource_Read_Happy exercises MapResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestMapResource_Read_Happy(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 200, "{}")}
	m := MapResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapResource_Read_NilClient exercises MapResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapResource_Read_NilClient(t *testing.T) {
	r := &MapResource{}
	m := MapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapResource_Read_BuildError exercises MapResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapResource_Read_BuildError(t *testing.T) {
	r := &MapResource{client: newMalformedBaseURLClient(t)}
	m := MapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapResource_Read_SendError exercises MapResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapResource_Read_SendError(t *testing.T) {
	r := &MapResource{client: newTransportErrorClient(t)}
	m := MapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapResource_Read_NotFound exercises MapResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestMapResource_Read_NotFound(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 404, "")}
	m := MapResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapResource_Read_APIError exercises MapResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapResource_Read_APIError(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_map")
}

// TestMapResource_Read_APIErrorReadBody exercises MapResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapResource_Read_APIErrorReadBody(t *testing.T) {
	r := &MapResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMapResource_Read_InvalidJSON exercises MapResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMapResource_Read_InvalidJSON(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 200, "{{")}
	m := MapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMapResource_Read_MapError exercises MapResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMapResource_Read_MapError(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := MapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMapResource_Update_Happy exercises MapResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestMapResource_Update_Happy(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 200, "{}")}
	m := MapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapResource_Update_NilClient exercises MapResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapResource_Update_NilClient(t *testing.T) {
	r := &MapResource{}
	m := MapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapResource_Update_BuildError exercises MapResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapResource_Update_BuildError(t *testing.T) {
	r := &MapResource{client: newMalformedBaseURLClient(t)}
	m := MapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapResource_Update_SendError exercises MapResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapResource_Update_SendError(t *testing.T) {
	r := &MapResource{client: newTransportErrorClient(t)}
	m := MapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapResource_Update_APIError exercises MapResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapResource_Update_APIError(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_map")
}

// TestMapResource_Update_APIErrorReadBody exercises MapResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapResource_Update_APIErrorReadBody(t *testing.T) {
	r := &MapResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMapResource_Update_InvalidJSON exercises MapResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMapResource_Update_InvalidJSON(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 200, "{{")}
	m := MapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMapResource_Update_MapError exercises MapResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMapResource_Update_MapError(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := MapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMapResource_Delete_Happy exercises MapResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestMapResource_Delete_Happy(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 204, "")}
	m := MapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapResource_Delete_NilClient exercises MapResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapResource_Delete_NilClient(t *testing.T) {
	r := &MapResource{}
	m := MapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapResource_Delete_BuildError exercises MapResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapResource_Delete_BuildError(t *testing.T) {
	r := &MapResource{client: newMalformedBaseURLClient(t)}
	m := MapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapResource_Delete_SendError exercises MapResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapResource_Delete_SendError(t *testing.T) {
	r := &MapResource{client: newTransportErrorClient(t)}
	m := MapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapResource_Delete_NotFoundSuccess exercises MapResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestMapResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 404, "")}
	m := MapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapResource_Delete_APIError exercises MapResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapResource_Delete_APIError(t *testing.T) {
	r := &MapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_map")
}

// TestMapResource_Delete_APIErrorReadBody exercises MapResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &MapResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
