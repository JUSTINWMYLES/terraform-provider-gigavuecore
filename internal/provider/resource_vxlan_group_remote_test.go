package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestVxlanGroupResource_Create_Happy exercises VxlanGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestVxlanGroupResource_Create_Happy(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestVxlanGroupResource_Create_NilClient exercises VxlanGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestVxlanGroupResource_Create_NilClient(t *testing.T) {
	r := &VxlanGroupResource{}
	m := VxlanGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestVxlanGroupResource_Create_BuildError exercises VxlanGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestVxlanGroupResource_Create_BuildError(t *testing.T) {
	r := &VxlanGroupResource{client: newMalformedBaseURLClient(t)}
	m := VxlanGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestVxlanGroupResource_Create_SendError exercises VxlanGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestVxlanGroupResource_Create_SendError(t *testing.T) {
	r := &VxlanGroupResource{client: newTransportErrorClient(t)}
	m := VxlanGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestVxlanGroupResource_Create_APIError exercises VxlanGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestVxlanGroupResource_Create_APIError(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_vxlan_group")
}

// TestVxlanGroupResource_Create_APIErrorReadBody exercises VxlanGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestVxlanGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := VxlanGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestVxlanGroupResource_Create_InvalidJSON exercises VxlanGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestVxlanGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := VxlanGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestVxlanGroupResource_Create_MapError exercises VxlanGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestVxlanGroupResource_Create_MapError(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestVxlanGroupResource_Create_MissingID exercises VxlanGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestVxlanGroupResource_Create_MissingID(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestVxlanGroupResource_Create_LocationFallback exercises VxlanGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestVxlanGroupResource_Create_LocationFallback(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestVxlanGroupResource_Read_Happy exercises VxlanGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestVxlanGroupResource_Read_Happy(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestVxlanGroupResource_Read_NilClient exercises VxlanGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestVxlanGroupResource_Read_NilClient(t *testing.T) {
	r := &VxlanGroupResource{}
	m := VxlanGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestVxlanGroupResource_Read_BuildError exercises VxlanGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestVxlanGroupResource_Read_BuildError(t *testing.T) {
	r := &VxlanGroupResource{client: newMalformedBaseURLClient(t)}
	m := VxlanGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestVxlanGroupResource_Read_SendError exercises VxlanGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestVxlanGroupResource_Read_SendError(t *testing.T) {
	r := &VxlanGroupResource{client: newTransportErrorClient(t)}
	m := VxlanGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestVxlanGroupResource_Read_NotFound exercises VxlanGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestVxlanGroupResource_Read_NotFound(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 404, "")}
	m := VxlanGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestVxlanGroupResource_Read_APIError exercises VxlanGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestVxlanGroupResource_Read_APIError(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_vxlan_group")
}

// TestVxlanGroupResource_Read_APIErrorReadBody exercises VxlanGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestVxlanGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := VxlanGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestVxlanGroupResource_Read_InvalidJSON exercises VxlanGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestVxlanGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := VxlanGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestVxlanGroupResource_Read_MapError exercises VxlanGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestVxlanGroupResource_Read_MapError(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestVxlanGroupResource_Update_Happy exercises VxlanGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestVxlanGroupResource_Update_Happy(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestVxlanGroupResource_Update_NilClient exercises VxlanGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestVxlanGroupResource_Update_NilClient(t *testing.T) {
	r := &VxlanGroupResource{}
	m := VxlanGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestVxlanGroupResource_Update_BuildError exercises VxlanGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestVxlanGroupResource_Update_BuildError(t *testing.T) {
	r := &VxlanGroupResource{client: newMalformedBaseURLClient(t)}
	m := VxlanGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestVxlanGroupResource_Update_SendError exercises VxlanGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestVxlanGroupResource_Update_SendError(t *testing.T) {
	r := &VxlanGroupResource{client: newTransportErrorClient(t)}
	m := VxlanGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestVxlanGroupResource_Update_APIError exercises VxlanGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestVxlanGroupResource_Update_APIError(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_vxlan_group")
}

// TestVxlanGroupResource_Update_APIErrorReadBody exercises VxlanGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestVxlanGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := VxlanGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestVxlanGroupResource_Update_InvalidJSON exercises VxlanGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestVxlanGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := VxlanGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestVxlanGroupResource_Update_MapError exercises VxlanGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestVxlanGroupResource_Update_MapError(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestVxlanGroupResource_Delete_Happy exercises VxlanGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestVxlanGroupResource_Delete_Happy(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 204, "")}
	m := VxlanGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestVxlanGroupResource_Delete_NilClient exercises VxlanGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestVxlanGroupResource_Delete_NilClient(t *testing.T) {
	r := &VxlanGroupResource{}
	m := VxlanGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestVxlanGroupResource_Delete_BuildError exercises VxlanGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestVxlanGroupResource_Delete_BuildError(t *testing.T) {
	r := &VxlanGroupResource{client: newMalformedBaseURLClient(t)}
	m := VxlanGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestVxlanGroupResource_Delete_SendError exercises VxlanGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestVxlanGroupResource_Delete_SendError(t *testing.T) {
	r := &VxlanGroupResource{client: newTransportErrorClient(t)}
	m := VxlanGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestVxlanGroupResource_Delete_NotFoundSuccess exercises VxlanGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestVxlanGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 404, "")}
	m := VxlanGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestVxlanGroupResource_Delete_APIError exercises VxlanGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestVxlanGroupResource_Delete_APIError(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := VxlanGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_vxlan_group")
}

// TestVxlanGroupResource_Delete_APIErrorReadBody exercises VxlanGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestVxlanGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &VxlanGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := VxlanGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
