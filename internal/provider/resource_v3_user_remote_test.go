package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestV3UserResource_Create_Happy exercises V3UserResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestV3UserResource_Create_Happy(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 201, "{\"username\":\"example-id\"}")}
	m := V3UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestV3UserResource_Create_NilClient exercises V3UserResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestV3UserResource_Create_NilClient(t *testing.T) {
	r := &V3UserResource{}
	m := V3UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestV3UserResource_Create_BuildError exercises V3UserResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestV3UserResource_Create_BuildError(t *testing.T) {
	r := &V3UserResource{client: newMalformedBaseURLClient(t)}
	m := V3UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestV3UserResource_Create_SendError exercises V3UserResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestV3UserResource_Create_SendError(t *testing.T) {
	r := &V3UserResource{client: newTransportErrorClient(t)}
	m := V3UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestV3UserResource_Create_APIError exercises V3UserResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestV3UserResource_Create_APIError(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := V3UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_v3_user")
}

// TestV3UserResource_Create_APIErrorReadBody exercises V3UserResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestV3UserResource_Create_APIErrorReadBody(t *testing.T) {
	r := &V3UserResource{client: newMockClientReadErrorBody(t, 501)}
	m := V3UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestV3UserResource_Create_InvalidJSON exercises V3UserResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestV3UserResource_Create_InvalidJSON(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 201, "{{")}
	m := V3UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestV3UserResource_Create_MapError exercises V3UserResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestV3UserResource_Create_MapError(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 201, "{\"username\":12345}")}
	m := V3UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestV3UserResource_Create_MissingID exercises V3UserResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestV3UserResource_Create_MissingID(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 201, "{}")}
	m := V3UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestV3UserResource_Create_LocationFallback exercises V3UserResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestV3UserResource_Create_LocationFallback(t *testing.T) {
	r := &V3UserResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := V3UserResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Username.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Username.ValueString(), "example-id")
	}
}

// TestV3UserResource_Read_Happy exercises V3UserResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestV3UserResource_Read_Happy(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 200, "{}")}
	m := V3UserResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestV3UserResource_Read_NilClient exercises V3UserResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestV3UserResource_Read_NilClient(t *testing.T) {
	r := &V3UserResource{}
	m := V3UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestV3UserResource_Read_BuildError exercises V3UserResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestV3UserResource_Read_BuildError(t *testing.T) {
	r := &V3UserResource{client: newMalformedBaseURLClient(t)}
	m := V3UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestV3UserResource_Read_SendError exercises V3UserResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestV3UserResource_Read_SendError(t *testing.T) {
	r := &V3UserResource{client: newTransportErrorClient(t)}
	m := V3UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestV3UserResource_Read_NotFound exercises V3UserResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestV3UserResource_Read_NotFound(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 404, "")}
	m := V3UserResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestV3UserResource_Read_APIError exercises V3UserResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestV3UserResource_Read_APIError(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := V3UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_v3_user")
}

// TestV3UserResource_Read_APIErrorReadBody exercises V3UserResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestV3UserResource_Read_APIErrorReadBody(t *testing.T) {
	r := &V3UserResource{client: newMockClientReadErrorBody(t, 501)}
	m := V3UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestV3UserResource_Read_InvalidJSON exercises V3UserResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestV3UserResource_Read_InvalidJSON(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 200, "{{")}
	m := V3UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestV3UserResource_Read_MapError exercises V3UserResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestV3UserResource_Read_MapError(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 200, "{\"username\":12345}")}
	m := V3UserResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestV3UserResource_Update_Happy exercises V3UserResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestV3UserResource_Update_Happy(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 200, "{}")}
	m := V3UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestV3UserResource_Update_NilClient exercises V3UserResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestV3UserResource_Update_NilClient(t *testing.T) {
	r := &V3UserResource{}
	m := V3UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestV3UserResource_Update_BuildError exercises V3UserResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestV3UserResource_Update_BuildError(t *testing.T) {
	r := &V3UserResource{client: newMalformedBaseURLClient(t)}
	m := V3UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestV3UserResource_Update_SendError exercises V3UserResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestV3UserResource_Update_SendError(t *testing.T) {
	r := &V3UserResource{client: newTransportErrorClient(t)}
	m := V3UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestV3UserResource_Update_APIError exercises V3UserResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestV3UserResource_Update_APIError(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := V3UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_v3_user")
}

// TestV3UserResource_Update_APIErrorReadBody exercises V3UserResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestV3UserResource_Update_APIErrorReadBody(t *testing.T) {
	r := &V3UserResource{client: newMockClientReadErrorBody(t, 501)}
	m := V3UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestV3UserResource_Update_InvalidJSON exercises V3UserResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestV3UserResource_Update_InvalidJSON(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 200, "{{")}
	m := V3UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestV3UserResource_Update_MapError exercises V3UserResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestV3UserResource_Update_MapError(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 200, "{\"username\":12345}")}
	m := V3UserResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestV3UserResource_Delete_Happy exercises V3UserResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestV3UserResource_Delete_Happy(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 204, "")}
	m := V3UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestV3UserResource_Delete_NilClient exercises V3UserResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestV3UserResource_Delete_NilClient(t *testing.T) {
	r := &V3UserResource{}
	m := V3UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestV3UserResource_Delete_BuildError exercises V3UserResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestV3UserResource_Delete_BuildError(t *testing.T) {
	r := &V3UserResource{client: newMalformedBaseURLClient(t)}
	m := V3UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestV3UserResource_Delete_SendError exercises V3UserResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestV3UserResource_Delete_SendError(t *testing.T) {
	r := &V3UserResource{client: newTransportErrorClient(t)}
	m := V3UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestV3UserResource_Delete_NotFoundSuccess exercises V3UserResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestV3UserResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 404, "")}
	m := V3UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestV3UserResource_Delete_APIError exercises V3UserResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestV3UserResource_Delete_APIError(t *testing.T) {
	r := &V3UserResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := V3UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_v3_user")
}

// TestV3UserResource_Delete_APIErrorReadBody exercises V3UserResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestV3UserResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &V3UserResource{client: newMockClientReadErrorBody(t, 501)}
	m := V3UserResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
