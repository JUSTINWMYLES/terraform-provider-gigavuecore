package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestLocalUserResource_Create_Happy exercises LocalUserResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestLocalUserResource_Create_Happy(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 201, "{\"username\":\"example-id\"}")}
	m := LocalUserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLocalUserResource_Create_NilClient exercises LocalUserResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLocalUserResource_Create_NilClient(t *testing.T) {
	r := &LocalUserResource{}
	m := LocalUserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLocalUserResource_Create_BuildError exercises LocalUserResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLocalUserResource_Create_BuildError(t *testing.T) {
	r := &LocalUserResource{client: newMalformedBaseURLClient(t)}
	m := LocalUserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLocalUserResource_Create_SendError exercises LocalUserResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestLocalUserResource_Create_SendError(t *testing.T) {
	r := &LocalUserResource{client: newTransportErrorClient(t)}
	m := LocalUserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLocalUserResource_Create_APIError exercises LocalUserResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLocalUserResource_Create_APIError(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LocalUserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_local_user")
}

// TestLocalUserResource_Create_APIErrorReadBody exercises LocalUserResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLocalUserResource_Create_APIErrorReadBody(t *testing.T) {
	r := &LocalUserResource{client: newMockClientReadErrorBody(t, 501)}
	m := LocalUserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLocalUserResource_Create_InvalidJSON exercises LocalUserResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLocalUserResource_Create_InvalidJSON(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 201, "{{")}
	m := LocalUserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestLocalUserResource_Create_MapError exercises LocalUserResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestLocalUserResource_Create_MapError(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 201, "{\"username\":12345}")}
	m := LocalUserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestLocalUserResource_Create_MissingID exercises LocalUserResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestLocalUserResource_Create_MissingID(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 201, "{}")}
	m := LocalUserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestLocalUserResource_Create_LocationFallback exercises LocalUserResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestLocalUserResource_Create_LocationFallback(t *testing.T) {
	r := &LocalUserResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := LocalUserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Username.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Username.ValueString(), "example-id")
	}
}

// TestLocalUserResource_Read_Happy exercises LocalUserResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestLocalUserResource_Read_Happy(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 200, "{}")}
	m := LocalUserResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestLocalUserResource_Read_NilClient exercises LocalUserResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLocalUserResource_Read_NilClient(t *testing.T) {
	r := &LocalUserResource{}
	m := LocalUserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLocalUserResource_Read_BuildError exercises LocalUserResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLocalUserResource_Read_BuildError(t *testing.T) {
	r := &LocalUserResource{client: newMalformedBaseURLClient(t)}
	m := LocalUserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLocalUserResource_Read_SendError exercises LocalUserResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLocalUserResource_Read_SendError(t *testing.T) {
	r := &LocalUserResource{client: newTransportErrorClient(t)}
	m := LocalUserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLocalUserResource_Read_NotFound exercises LocalUserResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestLocalUserResource_Read_NotFound(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 404, "")}
	m := LocalUserResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestLocalUserResource_Read_APIError exercises LocalUserResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLocalUserResource_Read_APIError(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LocalUserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_local_user")
}

// TestLocalUserResource_Read_APIErrorReadBody exercises LocalUserResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLocalUserResource_Read_APIErrorReadBody(t *testing.T) {
	r := &LocalUserResource{client: newMockClientReadErrorBody(t, 501)}
	m := LocalUserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLocalUserResource_Read_InvalidJSON exercises LocalUserResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLocalUserResource_Read_InvalidJSON(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 200, "{{")}
	m := LocalUserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestLocalUserResource_Read_MapError exercises LocalUserResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestLocalUserResource_Read_MapError(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 200, "{\"username\":12345}")}
	m := LocalUserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestLocalUserResource_Update_Happy exercises LocalUserResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestLocalUserResource_Update_Happy(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 200, "{}")}
	m := LocalUserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLocalUserResource_Update_NilClient exercises LocalUserResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLocalUserResource_Update_NilClient(t *testing.T) {
	r := &LocalUserResource{}
	m := LocalUserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLocalUserResource_Update_BuildError exercises LocalUserResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLocalUserResource_Update_BuildError(t *testing.T) {
	r := &LocalUserResource{client: newMalformedBaseURLClient(t)}
	m := LocalUserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLocalUserResource_Update_SendError exercises LocalUserResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestLocalUserResource_Update_SendError(t *testing.T) {
	r := &LocalUserResource{client: newTransportErrorClient(t)}
	m := LocalUserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLocalUserResource_Update_APIError exercises LocalUserResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLocalUserResource_Update_APIError(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LocalUserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_local_user")
}

// TestLocalUserResource_Update_APIErrorReadBody exercises LocalUserResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLocalUserResource_Update_APIErrorReadBody(t *testing.T) {
	r := &LocalUserResource{client: newMockClientReadErrorBody(t, 501)}
	m := LocalUserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLocalUserResource_Update_InvalidJSON exercises LocalUserResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLocalUserResource_Update_InvalidJSON(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 200, "{{")}
	m := LocalUserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestLocalUserResource_Update_MapError exercises LocalUserResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestLocalUserResource_Update_MapError(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 200, "{\"username\":12345}")}
	m := LocalUserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestLocalUserResource_Delete_Happy exercises LocalUserResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestLocalUserResource_Delete_Happy(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 204, "")}
	m := LocalUserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLocalUserResource_Delete_NilClient exercises LocalUserResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLocalUserResource_Delete_NilClient(t *testing.T) {
	r := &LocalUserResource{}
	m := LocalUserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLocalUserResource_Delete_BuildError exercises LocalUserResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLocalUserResource_Delete_BuildError(t *testing.T) {
	r := &LocalUserResource{client: newMalformedBaseURLClient(t)}
	m := LocalUserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLocalUserResource_Delete_SendError exercises LocalUserResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestLocalUserResource_Delete_SendError(t *testing.T) {
	r := &LocalUserResource{client: newTransportErrorClient(t)}
	m := LocalUserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLocalUserResource_Delete_NotFoundSuccess exercises LocalUserResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestLocalUserResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 404, "")}
	m := LocalUserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLocalUserResource_Delete_APIError exercises LocalUserResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLocalUserResource_Delete_APIError(t *testing.T) {
	r := &LocalUserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LocalUserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_local_user")
}

// TestLocalUserResource_Delete_APIErrorReadBody exercises LocalUserResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLocalUserResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &LocalUserResource{client: newMockClientReadErrorBody(t, 501)}
	m := LocalUserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
