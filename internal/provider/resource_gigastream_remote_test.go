package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestGigastreamResource_Create_Happy exercises GigastreamResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestGigastreamResource_Create_Happy(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := GigastreamResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGigastreamResource_Create_NilClient exercises GigastreamResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGigastreamResource_Create_NilClient(t *testing.T) {
	r := &GigastreamResource{}
	m := GigastreamResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGigastreamResource_Create_BuildError exercises GigastreamResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGigastreamResource_Create_BuildError(t *testing.T) {
	r := &GigastreamResource{client: newMalformedBaseURLClient(t)}
	m := GigastreamResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGigastreamResource_Create_SendError exercises GigastreamResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestGigastreamResource_Create_SendError(t *testing.T) {
	r := &GigastreamResource{client: newTransportErrorClient(t)}
	m := GigastreamResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGigastreamResource_Create_APIError exercises GigastreamResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGigastreamResource_Create_APIError(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GigastreamResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_gigastream")
}

// TestGigastreamResource_Create_APIErrorReadBody exercises GigastreamResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGigastreamResource_Create_APIErrorReadBody(t *testing.T) {
	r := &GigastreamResource{client: newMockClientReadErrorBody(t, 501)}
	m := GigastreamResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGigastreamResource_Create_InvalidJSON exercises GigastreamResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGigastreamResource_Create_InvalidJSON(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 201, "{{")}
	m := GigastreamResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGigastreamResource_Create_MapError exercises GigastreamResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGigastreamResource_Create_MapError(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := GigastreamResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGigastreamResource_Create_MissingID exercises GigastreamResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestGigastreamResource_Create_MissingID(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 201, "{}")}
	m := GigastreamResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestGigastreamResource_Create_LocationFallback exercises GigastreamResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestGigastreamResource_Create_LocationFallback(t *testing.T) {
	r := &GigastreamResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := GigastreamResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestGigastreamResource_Read_Happy exercises GigastreamResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestGigastreamResource_Read_Happy(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 200, "{}")}
	m := GigastreamResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGigastreamResource_Read_NilClient exercises GigastreamResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGigastreamResource_Read_NilClient(t *testing.T) {
	r := &GigastreamResource{}
	m := GigastreamResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGigastreamResource_Read_BuildError exercises GigastreamResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGigastreamResource_Read_BuildError(t *testing.T) {
	r := &GigastreamResource{client: newMalformedBaseURLClient(t)}
	m := GigastreamResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGigastreamResource_Read_SendError exercises GigastreamResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGigastreamResource_Read_SendError(t *testing.T) {
	r := &GigastreamResource{client: newTransportErrorClient(t)}
	m := GigastreamResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGigastreamResource_Read_NotFound exercises GigastreamResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestGigastreamResource_Read_NotFound(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 404, "")}
	m := GigastreamResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGigastreamResource_Read_APIError exercises GigastreamResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGigastreamResource_Read_APIError(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GigastreamResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_gigastream")
}

// TestGigastreamResource_Read_APIErrorReadBody exercises GigastreamResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGigastreamResource_Read_APIErrorReadBody(t *testing.T) {
	r := &GigastreamResource{client: newMockClientReadErrorBody(t, 501)}
	m := GigastreamResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGigastreamResource_Read_InvalidJSON exercises GigastreamResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGigastreamResource_Read_InvalidJSON(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 200, "{{")}
	m := GigastreamResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGigastreamResource_Read_MapError exercises GigastreamResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGigastreamResource_Read_MapError(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GigastreamResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGigastreamResource_Update_Happy exercises GigastreamResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestGigastreamResource_Update_Happy(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 200, "{}")}
	m := GigastreamResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGigastreamResource_Update_NilClient exercises GigastreamResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGigastreamResource_Update_NilClient(t *testing.T) {
	r := &GigastreamResource{}
	m := GigastreamResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGigastreamResource_Update_BuildError exercises GigastreamResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGigastreamResource_Update_BuildError(t *testing.T) {
	r := &GigastreamResource{client: newMalformedBaseURLClient(t)}
	m := GigastreamResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGigastreamResource_Update_SendError exercises GigastreamResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestGigastreamResource_Update_SendError(t *testing.T) {
	r := &GigastreamResource{client: newTransportErrorClient(t)}
	m := GigastreamResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGigastreamResource_Update_APIError exercises GigastreamResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGigastreamResource_Update_APIError(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GigastreamResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_gigastream")
}

// TestGigastreamResource_Update_APIErrorReadBody exercises GigastreamResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGigastreamResource_Update_APIErrorReadBody(t *testing.T) {
	r := &GigastreamResource{client: newMockClientReadErrorBody(t, 501)}
	m := GigastreamResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGigastreamResource_Update_InvalidJSON exercises GigastreamResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGigastreamResource_Update_InvalidJSON(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 200, "{{")}
	m := GigastreamResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGigastreamResource_Update_MapError exercises GigastreamResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGigastreamResource_Update_MapError(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GigastreamResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGigastreamResource_Delete_Happy exercises GigastreamResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestGigastreamResource_Delete_Happy(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 204, "")}
	m := GigastreamResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGigastreamResource_Delete_NilClient exercises GigastreamResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGigastreamResource_Delete_NilClient(t *testing.T) {
	r := &GigastreamResource{}
	m := GigastreamResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGigastreamResource_Delete_BuildError exercises GigastreamResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGigastreamResource_Delete_BuildError(t *testing.T) {
	r := &GigastreamResource{client: newMalformedBaseURLClient(t)}
	m := GigastreamResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGigastreamResource_Delete_SendError exercises GigastreamResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestGigastreamResource_Delete_SendError(t *testing.T) {
	r := &GigastreamResource{client: newTransportErrorClient(t)}
	m := GigastreamResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGigastreamResource_Delete_NotFoundSuccess exercises GigastreamResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestGigastreamResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 404, "")}
	m := GigastreamResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGigastreamResource_Delete_APIError exercises GigastreamResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGigastreamResource_Delete_APIError(t *testing.T) {
	r := &GigastreamResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GigastreamResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_gigastream")
}

// TestGigastreamResource_Delete_APIErrorReadBody exercises GigastreamResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGigastreamResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &GigastreamResource{client: newMockClientReadErrorBody(t, 501)}
	m := GigastreamResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
