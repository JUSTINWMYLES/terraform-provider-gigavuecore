package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestUserRoleResource_Create_Happy exercises UserRoleResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestUserRoleResource_Create_Happy(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 201, "{\"name\":\"example-id\"}")}
	m := UserRoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserRoleResource_Create_NilClient exercises UserRoleResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUserRoleResource_Create_NilClient(t *testing.T) {
	r := &UserRoleResource{}
	m := UserRoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUserRoleResource_Create_BuildError exercises UserRoleResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUserRoleResource_Create_BuildError(t *testing.T) {
	r := &UserRoleResource{client: newMalformedBaseURLClient(t)}
	m := UserRoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUserRoleResource_Create_SendError exercises UserRoleResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestUserRoleResource_Create_SendError(t *testing.T) {
	r := &UserRoleResource{client: newTransportErrorClient(t)}
	m := UserRoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUserRoleResource_Create_APIError exercises UserRoleResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUserRoleResource_Create_APIError(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UserRoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_user_role")
}

// TestUserRoleResource_Create_APIErrorReadBody exercises UserRoleResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUserRoleResource_Create_APIErrorReadBody(t *testing.T) {
	r := &UserRoleResource{client: newMockClientReadErrorBody(t, 501)}
	m := UserRoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestUserRoleResource_Create_InvalidJSON exercises UserRoleResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestUserRoleResource_Create_InvalidJSON(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 201, "{{")}
	m := UserRoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestUserRoleResource_Create_MapError exercises UserRoleResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestUserRoleResource_Create_MapError(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 201, "{\"name\":12345}")}
	m := UserRoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestUserRoleResource_Create_MissingID exercises UserRoleResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestUserRoleResource_Create_MissingID(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 201, "{}")}
	m := UserRoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestUserRoleResource_Create_LocationFallback exercises UserRoleResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestUserRoleResource_Create_LocationFallback(t *testing.T) {
	r := &UserRoleResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := UserRoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Name.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Name.ValueString(), "example-id")
	}
}

// TestUserRoleResource_Read_Happy exercises UserRoleResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestUserRoleResource_Read_Happy(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 200, "{}")}
	m := UserRoleResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserRoleResource_Read_NilClient exercises UserRoleResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUserRoleResource_Read_NilClient(t *testing.T) {
	r := &UserRoleResource{}
	m := UserRoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUserRoleResource_Read_BuildError exercises UserRoleResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUserRoleResource_Read_BuildError(t *testing.T) {
	r := &UserRoleResource{client: newMalformedBaseURLClient(t)}
	m := UserRoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUserRoleResource_Read_SendError exercises UserRoleResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestUserRoleResource_Read_SendError(t *testing.T) {
	r := &UserRoleResource{client: newTransportErrorClient(t)}
	m := UserRoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUserRoleResource_Read_NotFound exercises UserRoleResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestUserRoleResource_Read_NotFound(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 404, "")}
	m := UserRoleResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserRoleResource_Read_APIError exercises UserRoleResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUserRoleResource_Read_APIError(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UserRoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_user_role")
}

// TestUserRoleResource_Read_APIErrorReadBody exercises UserRoleResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUserRoleResource_Read_APIErrorReadBody(t *testing.T) {
	r := &UserRoleResource{client: newMockClientReadErrorBody(t, 501)}
	m := UserRoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestUserRoleResource_Read_InvalidJSON exercises UserRoleResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestUserRoleResource_Read_InvalidJSON(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 200, "{{")}
	m := UserRoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestUserRoleResource_Read_MapError exercises UserRoleResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestUserRoleResource_Read_MapError(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 200, "{\"name\":12345}")}
	m := UserRoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestUserRoleResource_Update_Happy exercises UserRoleResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestUserRoleResource_Update_Happy(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 200, "{}")}
	m := UserRoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserRoleResource_Update_NilClient exercises UserRoleResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUserRoleResource_Update_NilClient(t *testing.T) {
	r := &UserRoleResource{}
	m := UserRoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUserRoleResource_Update_BuildError exercises UserRoleResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUserRoleResource_Update_BuildError(t *testing.T) {
	r := &UserRoleResource{client: newMalformedBaseURLClient(t)}
	m := UserRoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUserRoleResource_Update_SendError exercises UserRoleResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestUserRoleResource_Update_SendError(t *testing.T) {
	r := &UserRoleResource{client: newTransportErrorClient(t)}
	m := UserRoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUserRoleResource_Update_APIError exercises UserRoleResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUserRoleResource_Update_APIError(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UserRoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_user_role")
}

// TestUserRoleResource_Update_APIErrorReadBody exercises UserRoleResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUserRoleResource_Update_APIErrorReadBody(t *testing.T) {
	r := &UserRoleResource{client: newMockClientReadErrorBody(t, 501)}
	m := UserRoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestUserRoleResource_Update_InvalidJSON exercises UserRoleResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestUserRoleResource_Update_InvalidJSON(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 200, "{{")}
	m := UserRoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestUserRoleResource_Update_MapError exercises UserRoleResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestUserRoleResource_Update_MapError(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 200, "{\"name\":12345}")}
	m := UserRoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestUserRoleResource_Delete_Happy exercises UserRoleResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestUserRoleResource_Delete_Happy(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 204, "")}
	m := UserRoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserRoleResource_Delete_NilClient exercises UserRoleResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUserRoleResource_Delete_NilClient(t *testing.T) {
	r := &UserRoleResource{}
	m := UserRoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUserRoleResource_Delete_BuildError exercises UserRoleResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUserRoleResource_Delete_BuildError(t *testing.T) {
	r := &UserRoleResource{client: newMalformedBaseURLClient(t)}
	m := UserRoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUserRoleResource_Delete_SendError exercises UserRoleResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestUserRoleResource_Delete_SendError(t *testing.T) {
	r := &UserRoleResource{client: newTransportErrorClient(t)}
	m := UserRoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUserRoleResource_Delete_NotFoundSuccess exercises UserRoleResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestUserRoleResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 404, "")}
	m := UserRoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUserRoleResource_Delete_APIError exercises UserRoleResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUserRoleResource_Delete_APIError(t *testing.T) {
	r := &UserRoleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UserRoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_user_role")
}

// TestUserRoleResource_Delete_APIErrorReadBody exercises UserRoleResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUserRoleResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &UserRoleResource{client: newMockClientReadErrorBody(t, 501)}
	m := UserRoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
