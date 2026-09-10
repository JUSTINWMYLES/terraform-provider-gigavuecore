package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSffpProfileResource_Create_Happy exercises SffpProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestSffpProfileResource_Create_Happy(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := SffpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSffpProfileResource_Create_NilClient exercises SffpProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSffpProfileResource_Create_NilClient(t *testing.T) {
	r := &SffpProfileResource{}
	m := SffpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSffpProfileResource_Create_BuildError exercises SffpProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSffpProfileResource_Create_BuildError(t *testing.T) {
	r := &SffpProfileResource{client: newMalformedBaseURLClient(t)}
	m := SffpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSffpProfileResource_Create_SendError exercises SffpProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestSffpProfileResource_Create_SendError(t *testing.T) {
	r := &SffpProfileResource{client: newTransportErrorClient(t)}
	m := SffpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSffpProfileResource_Create_APIError exercises SffpProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSffpProfileResource_Create_APIError(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SffpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_sffp_profile")
}

// TestSffpProfileResource_Create_APIErrorReadBody exercises SffpProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSffpProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SffpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSffpProfileResource_Create_InvalidJSON exercises SffpProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSffpProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := SffpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSffpProfileResource_Create_MapError exercises SffpProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSffpProfileResource_Create_MapError(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := SffpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSffpProfileResource_Create_MissingID exercises SffpProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestSffpProfileResource_Create_MissingID(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := SffpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestSffpProfileResource_Create_LocationFallback exercises SffpProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestSffpProfileResource_Create_LocationFallback(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := SffpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestSffpProfileResource_Read_Happy exercises SffpProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestSffpProfileResource_Read_Happy(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := SffpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSffpProfileResource_Read_NilClient exercises SffpProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSffpProfileResource_Read_NilClient(t *testing.T) {
	r := &SffpProfileResource{}
	m := SffpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSffpProfileResource_Read_BuildError exercises SffpProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSffpProfileResource_Read_BuildError(t *testing.T) {
	r := &SffpProfileResource{client: newMalformedBaseURLClient(t)}
	m := SffpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSffpProfileResource_Read_SendError exercises SffpProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestSffpProfileResource_Read_SendError(t *testing.T) {
	r := &SffpProfileResource{client: newTransportErrorClient(t)}
	m := SffpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSffpProfileResource_Read_NotFound exercises SffpProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestSffpProfileResource_Read_NotFound(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 404, "")}
	m := SffpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSffpProfileResource_Read_APIError exercises SffpProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSffpProfileResource_Read_APIError(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SffpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_sffp_profile")
}

// TestSffpProfileResource_Read_APIErrorReadBody exercises SffpProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSffpProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SffpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSffpProfileResource_Read_InvalidJSON exercises SffpProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSffpProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := SffpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSffpProfileResource_Read_MapError exercises SffpProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSffpProfileResource_Read_MapError(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SffpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSffpProfileResource_Update_Happy exercises SffpProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestSffpProfileResource_Update_Happy(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := SffpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSffpProfileResource_Update_NilClient exercises SffpProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSffpProfileResource_Update_NilClient(t *testing.T) {
	r := &SffpProfileResource{}
	m := SffpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSffpProfileResource_Update_BuildError exercises SffpProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSffpProfileResource_Update_BuildError(t *testing.T) {
	r := &SffpProfileResource{client: newMalformedBaseURLClient(t)}
	m := SffpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSffpProfileResource_Update_SendError exercises SffpProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestSffpProfileResource_Update_SendError(t *testing.T) {
	r := &SffpProfileResource{client: newTransportErrorClient(t)}
	m := SffpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSffpProfileResource_Update_APIError exercises SffpProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSffpProfileResource_Update_APIError(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SffpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_sffp_profile")
}

// TestSffpProfileResource_Update_APIErrorReadBody exercises SffpProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSffpProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SffpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSffpProfileResource_Update_InvalidJSON exercises SffpProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSffpProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := SffpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSffpProfileResource_Update_MapError exercises SffpProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSffpProfileResource_Update_MapError(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SffpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSffpProfileResource_Delete_Happy exercises SffpProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestSffpProfileResource_Delete_Happy(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 204, "")}
	m := SffpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSffpProfileResource_Delete_NilClient exercises SffpProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSffpProfileResource_Delete_NilClient(t *testing.T) {
	r := &SffpProfileResource{}
	m := SffpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSffpProfileResource_Delete_BuildError exercises SffpProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSffpProfileResource_Delete_BuildError(t *testing.T) {
	r := &SffpProfileResource{client: newMalformedBaseURLClient(t)}
	m := SffpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSffpProfileResource_Delete_SendError exercises SffpProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestSffpProfileResource_Delete_SendError(t *testing.T) {
	r := &SffpProfileResource{client: newTransportErrorClient(t)}
	m := SffpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSffpProfileResource_Delete_NotFoundSuccess exercises SffpProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestSffpProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 404, "")}
	m := SffpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSffpProfileResource_Delete_APIError exercises SffpProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSffpProfileResource_Delete_APIError(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SffpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_sffp_profile")
}

// TestSffpProfileResource_Delete_APIErrorReadBody exercises SffpProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSffpProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &SffpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SffpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
