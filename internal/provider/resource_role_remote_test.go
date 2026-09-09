package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestRoleResource_Create_Happy exercises RoleResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestRoleResource_Create_Happy(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 201, "{\"name\":\"example-id\"}")}
	m := RoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRoleResource_Create_NilClient exercises RoleResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRoleResource_Create_NilClient(t *testing.T) {
	r := &RoleResource{}
	m := RoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRoleResource_Create_BuildError exercises RoleResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRoleResource_Create_BuildError(t *testing.T) {
	r := &RoleResource{client: newMalformedBaseURLClient(t)}
	m := RoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRoleResource_Create_SendError exercises RoleResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestRoleResource_Create_SendError(t *testing.T) {
	r := &RoleResource{client: newTransportErrorClient(t)}
	m := RoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRoleResource_Create_APIError exercises RoleResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRoleResource_Create_APIError(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_role")
}

// TestRoleResource_Create_APIErrorReadBody exercises RoleResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRoleResource_Create_APIErrorReadBody(t *testing.T) {
	r := &RoleResource{client: newMockClientReadErrorBody(t, 501)}
	m := RoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRoleResource_Create_InvalidJSON exercises RoleResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRoleResource_Create_InvalidJSON(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 201, "{{")}
	m := RoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRoleResource_Create_MapError exercises RoleResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRoleResource_Create_MapError(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 201, "{\"name\":12345}")}
	m := RoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRoleResource_Create_MissingID exercises RoleResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestRoleResource_Create_MissingID(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 201, "{}")}
	m := RoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestRoleResource_Create_LocationFallback exercises RoleResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestRoleResource_Create_LocationFallback(t *testing.T) {
	r := &RoleResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := RoleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Name.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Name.ValueString(), "example-id")
	}
}

// TestRoleResource_Read_Happy exercises RoleResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestRoleResource_Read_Happy(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 200, "{}")}
	m := RoleResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestRoleResource_Read_NilClient exercises RoleResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRoleResource_Read_NilClient(t *testing.T) {
	r := &RoleResource{}
	m := RoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRoleResource_Read_BuildError exercises RoleResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRoleResource_Read_BuildError(t *testing.T) {
	r := &RoleResource{client: newMalformedBaseURLClient(t)}
	m := RoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRoleResource_Read_SendError exercises RoleResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestRoleResource_Read_SendError(t *testing.T) {
	r := &RoleResource{client: newTransportErrorClient(t)}
	m := RoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRoleResource_Read_NotFound exercises RoleResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestRoleResource_Read_NotFound(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 404, "")}
	m := RoleResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestRoleResource_Read_APIError exercises RoleResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRoleResource_Read_APIError(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_role")
}

// TestRoleResource_Read_APIErrorReadBody exercises RoleResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRoleResource_Read_APIErrorReadBody(t *testing.T) {
	r := &RoleResource{client: newMockClientReadErrorBody(t, 501)}
	m := RoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRoleResource_Read_InvalidJSON exercises RoleResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRoleResource_Read_InvalidJSON(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 200, "{{")}
	m := RoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRoleResource_Read_MapError exercises RoleResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRoleResource_Read_MapError(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 200, "{\"name\":12345}")}
	m := RoleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRoleResource_Update_Happy exercises RoleResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestRoleResource_Update_Happy(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 200, "{}")}
	m := RoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRoleResource_Update_NilClient exercises RoleResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRoleResource_Update_NilClient(t *testing.T) {
	r := &RoleResource{}
	m := RoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRoleResource_Update_BuildError exercises RoleResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRoleResource_Update_BuildError(t *testing.T) {
	r := &RoleResource{client: newMalformedBaseURLClient(t)}
	m := RoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRoleResource_Update_SendError exercises RoleResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestRoleResource_Update_SendError(t *testing.T) {
	r := &RoleResource{client: newTransportErrorClient(t)}
	m := RoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRoleResource_Update_APIError exercises RoleResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRoleResource_Update_APIError(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_role")
}

// TestRoleResource_Update_APIErrorReadBody exercises RoleResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRoleResource_Update_APIErrorReadBody(t *testing.T) {
	r := &RoleResource{client: newMockClientReadErrorBody(t, 501)}
	m := RoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRoleResource_Update_InvalidJSON exercises RoleResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRoleResource_Update_InvalidJSON(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 200, "{{")}
	m := RoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRoleResource_Update_MapError exercises RoleResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRoleResource_Update_MapError(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 200, "{\"name\":12345}")}
	m := RoleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRoleResource_Delete_Happy exercises RoleResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestRoleResource_Delete_Happy(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 204, "")}
	m := RoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRoleResource_Delete_NilClient exercises RoleResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRoleResource_Delete_NilClient(t *testing.T) {
	r := &RoleResource{}
	m := RoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRoleResource_Delete_BuildError exercises RoleResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRoleResource_Delete_BuildError(t *testing.T) {
	r := &RoleResource{client: newMalformedBaseURLClient(t)}
	m := RoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRoleResource_Delete_SendError exercises RoleResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestRoleResource_Delete_SendError(t *testing.T) {
	r := &RoleResource{client: newTransportErrorClient(t)}
	m := RoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRoleResource_Delete_NotFoundSuccess exercises RoleResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestRoleResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 404, "")}
	m := RoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRoleResource_Delete_APIError exercises RoleResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRoleResource_Delete_APIError(t *testing.T) {
	r := &RoleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_role")
}

// TestRoleResource_Delete_APIErrorReadBody exercises RoleResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRoleResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &RoleResource{client: newMockClientReadErrorBody(t, 501)}
	m := RoleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
