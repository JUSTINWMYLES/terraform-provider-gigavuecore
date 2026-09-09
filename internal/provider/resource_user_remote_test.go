package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestUserResource_Create_Happy exercises UserResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestUserResource_Create_Happy(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 201, "{\"username\":\"example-id\"}")}
	m := UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserResource_Create_NilClient exercises UserResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUserResource_Create_NilClient(t *testing.T) {
	r := &UserResource{}
	m := UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUserResource_Create_BuildError exercises UserResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUserResource_Create_BuildError(t *testing.T) {
	r := &UserResource{client: newMalformedBaseURLClient(t)}
	m := UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUserResource_Create_SendError exercises UserResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestUserResource_Create_SendError(t *testing.T) {
	r := &UserResource{client: newTransportErrorClient(t)}
	m := UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUserResource_Create_APIError exercises UserResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUserResource_Create_APIError(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_user")
}

// TestUserResource_Create_APIErrorReadBody exercises UserResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUserResource_Create_APIErrorReadBody(t *testing.T) {
	r := &UserResource{client: newMockClientReadErrorBody(t, 501)}
	m := UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestUserResource_Create_InvalidJSON exercises UserResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestUserResource_Create_InvalidJSON(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 201, "{{")}
	m := UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestUserResource_Create_MapError exercises UserResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestUserResource_Create_MapError(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 201, "{\"username\":12345}")}
	m := UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestUserResource_Create_MissingID exercises UserResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestUserResource_Create_MissingID(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 201, "{}")}
	m := UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestUserResource_Create_LocationFallback exercises UserResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestUserResource_Create_LocationFallback(t *testing.T) {
	r := &UserResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Username.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Username.ValueString(), "example-id")
	}
}

// TestUserResource_Read_Happy exercises UserResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestUserResource_Read_Happy(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 200, "{}")}
	m := UserResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserResource_Read_NilClient exercises UserResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUserResource_Read_NilClient(t *testing.T) {
	r := &UserResource{}
	m := UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUserResource_Read_BuildError exercises UserResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUserResource_Read_BuildError(t *testing.T) {
	r := &UserResource{client: newMalformedBaseURLClient(t)}
	m := UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUserResource_Read_SendError exercises UserResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestUserResource_Read_SendError(t *testing.T) {
	r := &UserResource{client: newTransportErrorClient(t)}
	m := UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUserResource_Read_NotFound exercises UserResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestUserResource_Read_NotFound(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 404, "")}
	m := UserResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserResource_Read_APIError exercises UserResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUserResource_Read_APIError(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_user")
}

// TestUserResource_Read_APIErrorReadBody exercises UserResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUserResource_Read_APIErrorReadBody(t *testing.T) {
	r := &UserResource{client: newMockClientReadErrorBody(t, 501)}
	m := UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestUserResource_Read_InvalidJSON exercises UserResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestUserResource_Read_InvalidJSON(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 200, "{{")}
	m := UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestUserResource_Read_MapError exercises UserResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestUserResource_Read_MapError(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 200, "{\"username\":12345}")}
	m := UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestUserResource_Update_Happy exercises UserResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestUserResource_Update_Happy(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 200, "{}")}
	m := UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserResource_Update_NilClient exercises UserResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUserResource_Update_NilClient(t *testing.T) {
	r := &UserResource{}
	m := UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUserResource_Update_BuildError exercises UserResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUserResource_Update_BuildError(t *testing.T) {
	r := &UserResource{client: newMalformedBaseURLClient(t)}
	m := UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUserResource_Update_SendError exercises UserResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestUserResource_Update_SendError(t *testing.T) {
	r := &UserResource{client: newTransportErrorClient(t)}
	m := UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUserResource_Update_APIError exercises UserResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUserResource_Update_APIError(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_user")
}

// TestUserResource_Update_APIErrorReadBody exercises UserResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUserResource_Update_APIErrorReadBody(t *testing.T) {
	r := &UserResource{client: newMockClientReadErrorBody(t, 501)}
	m := UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestUserResource_Update_InvalidJSON exercises UserResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestUserResource_Update_InvalidJSON(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 200, "{{")}
	m := UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestUserResource_Update_MapError exercises UserResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestUserResource_Update_MapError(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 200, "{\"username\":12345}")}
	m := UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestUserResource_Delete_Happy exercises UserResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestUserResource_Delete_Happy(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 204, "")}
	m := UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserResource_Delete_NilClient exercises UserResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUserResource_Delete_NilClient(t *testing.T) {
	r := &UserResource{}
	m := UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUserResource_Delete_BuildError exercises UserResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUserResource_Delete_BuildError(t *testing.T) {
	r := &UserResource{client: newMalformedBaseURLClient(t)}
	m := UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUserResource_Delete_SendError exercises UserResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestUserResource_Delete_SendError(t *testing.T) {
	r := &UserResource{client: newTransportErrorClient(t)}
	m := UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUserResource_Delete_NotFoundSuccess exercises UserResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestUserResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 404, "")}
	m := UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserResource_Delete_APIError exercises UserResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUserResource_Delete_APIError(t *testing.T) {
	r := &UserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_user")
}

// TestUserResource_Delete_APIErrorReadBody exercises UserResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUserResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &UserResource{client: newMockClientReadErrorBody(t, 501)}
	m := UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
