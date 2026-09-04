package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestGsGroupResource_Create_Happy exercises GsGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestGsGroupResource_Create_Happy(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := GsGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsGroupResource_Create_NilClient exercises GsGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsGroupResource_Create_NilClient(t *testing.T) {
	r := &GsGroupResource{}
	m := GsGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGsGroupResource_Create_BuildError exercises GsGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGsGroupResource_Create_BuildError(t *testing.T) {
	r := &GsGroupResource{client: newMalformedBaseURLClient(t)}
	m := GsGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGsGroupResource_Create_SendError exercises GsGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestGsGroupResource_Create_SendError(t *testing.T) {
	r := &GsGroupResource{client: newTransportErrorClient(t)}
	m := GsGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGsGroupResource_Create_APIError exercises GsGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGsGroupResource_Create_APIError(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GsGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_gs_group")
}

// TestGsGroupResource_Create_APIErrorReadBody exercises GsGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGsGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &GsGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := GsGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGsGroupResource_Create_InvalidJSON exercises GsGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGsGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := GsGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGsGroupResource_Create_MapError exercises GsGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGsGroupResource_Create_MapError(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := GsGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGsGroupResource_Create_MissingID exercises GsGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestGsGroupResource_Create_MissingID(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := GsGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestGsGroupResource_Create_LocationFallback exercises GsGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestGsGroupResource_Create_LocationFallback(t *testing.T) {
	r := &GsGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := GsGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestGsGroupResource_Read_Happy exercises GsGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestGsGroupResource_Read_Happy(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := GsGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsGroupResource_Read_NilClient exercises GsGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsGroupResource_Read_NilClient(t *testing.T) {
	r := &GsGroupResource{}
	m := GsGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGsGroupResource_Read_BuildError exercises GsGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGsGroupResource_Read_BuildError(t *testing.T) {
	r := &GsGroupResource{client: newMalformedBaseURLClient(t)}
	m := GsGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGsGroupResource_Read_SendError exercises GsGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGsGroupResource_Read_SendError(t *testing.T) {
	r := &GsGroupResource{client: newTransportErrorClient(t)}
	m := GsGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGsGroupResource_Read_NotFound exercises GsGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestGsGroupResource_Read_NotFound(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 404, "")}
	m := GsGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsGroupResource_Read_APIError exercises GsGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGsGroupResource_Read_APIError(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GsGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_gs_group")
}

// TestGsGroupResource_Read_APIErrorReadBody exercises GsGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGsGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &GsGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := GsGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGsGroupResource_Read_InvalidJSON exercises GsGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGsGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := GsGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGsGroupResource_Read_MapError exercises GsGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGsGroupResource_Read_MapError(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GsGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGsGroupResource_Update_Happy exercises GsGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestGsGroupResource_Update_Happy(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := GsGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsGroupResource_Update_NilClient exercises GsGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsGroupResource_Update_NilClient(t *testing.T) {
	r := &GsGroupResource{}
	m := GsGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGsGroupResource_Update_BuildError exercises GsGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGsGroupResource_Update_BuildError(t *testing.T) {
	r := &GsGroupResource{client: newMalformedBaseURLClient(t)}
	m := GsGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGsGroupResource_Update_SendError exercises GsGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestGsGroupResource_Update_SendError(t *testing.T) {
	r := &GsGroupResource{client: newTransportErrorClient(t)}
	m := GsGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGsGroupResource_Update_APIError exercises GsGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGsGroupResource_Update_APIError(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GsGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_gs_group")
}

// TestGsGroupResource_Update_APIErrorReadBody exercises GsGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGsGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &GsGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := GsGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGsGroupResource_Update_InvalidJSON exercises GsGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGsGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := GsGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGsGroupResource_Update_MapError exercises GsGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGsGroupResource_Update_MapError(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GsGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGsGroupResource_Delete_Happy exercises GsGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestGsGroupResource_Delete_Happy(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 204, "")}
	m := GsGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsGroupResource_Delete_NilClient exercises GsGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsGroupResource_Delete_NilClient(t *testing.T) {
	r := &GsGroupResource{}
	m := GsGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGsGroupResource_Delete_BuildError exercises GsGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGsGroupResource_Delete_BuildError(t *testing.T) {
	r := &GsGroupResource{client: newMalformedBaseURLClient(t)}
	m := GsGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGsGroupResource_Delete_SendError exercises GsGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestGsGroupResource_Delete_SendError(t *testing.T) {
	r := &GsGroupResource{client: newTransportErrorClient(t)}
	m := GsGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGsGroupResource_Delete_NotFoundSuccess exercises GsGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestGsGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 404, "")}
	m := GsGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsGroupResource_Delete_APIError exercises GsGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGsGroupResource_Delete_APIError(t *testing.T) {
	r := &GsGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GsGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_gs_group")
}

// TestGsGroupResource_Delete_APIErrorReadBody exercises GsGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGsGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &GsGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := GsGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
