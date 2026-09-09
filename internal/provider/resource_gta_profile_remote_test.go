package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestGtaProfileResource_Create_Happy exercises GtaProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestGtaProfileResource_Create_Happy(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := GtaProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtaProfileResource_Create_NilClient exercises GtaProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGtaProfileResource_Create_NilClient(t *testing.T) {
	r := &GtaProfileResource{}
	m := GtaProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGtaProfileResource_Create_BuildError exercises GtaProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGtaProfileResource_Create_BuildError(t *testing.T) {
	r := &GtaProfileResource{client: newMalformedBaseURLClient(t)}
	m := GtaProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGtaProfileResource_Create_SendError exercises GtaProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestGtaProfileResource_Create_SendError(t *testing.T) {
	r := &GtaProfileResource{client: newTransportErrorClient(t)}
	m := GtaProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGtaProfileResource_Create_APIError exercises GtaProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGtaProfileResource_Create_APIError(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GtaProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_gta_profile")
}

// TestGtaProfileResource_Create_APIErrorReadBody exercises GtaProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGtaProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := GtaProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGtaProfileResource_Create_InvalidJSON exercises GtaProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGtaProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := GtaProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGtaProfileResource_Create_MapError exercises GtaProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGtaProfileResource_Create_MapError(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := GtaProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGtaProfileResource_Create_MissingID exercises GtaProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestGtaProfileResource_Create_MissingID(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := GtaProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestGtaProfileResource_Create_LocationFallback exercises GtaProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestGtaProfileResource_Create_LocationFallback(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := GtaProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestGtaProfileResource_Read_Happy exercises GtaProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestGtaProfileResource_Read_Happy(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := GtaProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtaProfileResource_Read_NilClient exercises GtaProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGtaProfileResource_Read_NilClient(t *testing.T) {
	r := &GtaProfileResource{}
	m := GtaProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGtaProfileResource_Read_BuildError exercises GtaProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGtaProfileResource_Read_BuildError(t *testing.T) {
	r := &GtaProfileResource{client: newMalformedBaseURLClient(t)}
	m := GtaProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGtaProfileResource_Read_SendError exercises GtaProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGtaProfileResource_Read_SendError(t *testing.T) {
	r := &GtaProfileResource{client: newTransportErrorClient(t)}
	m := GtaProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGtaProfileResource_Read_NotFound exercises GtaProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestGtaProfileResource_Read_NotFound(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 404, "")}
	m := GtaProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtaProfileResource_Read_APIError exercises GtaProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGtaProfileResource_Read_APIError(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GtaProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_gta_profile")
}

// TestGtaProfileResource_Read_APIErrorReadBody exercises GtaProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGtaProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := GtaProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGtaProfileResource_Read_InvalidJSON exercises GtaProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGtaProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := GtaProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGtaProfileResource_Read_MapError exercises GtaProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGtaProfileResource_Read_MapError(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GtaProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGtaProfileResource_Update_Happy exercises GtaProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestGtaProfileResource_Update_Happy(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := GtaProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtaProfileResource_Update_NilClient exercises GtaProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGtaProfileResource_Update_NilClient(t *testing.T) {
	r := &GtaProfileResource{}
	m := GtaProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGtaProfileResource_Update_BuildError exercises GtaProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGtaProfileResource_Update_BuildError(t *testing.T) {
	r := &GtaProfileResource{client: newMalformedBaseURLClient(t)}
	m := GtaProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGtaProfileResource_Update_SendError exercises GtaProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestGtaProfileResource_Update_SendError(t *testing.T) {
	r := &GtaProfileResource{client: newTransportErrorClient(t)}
	m := GtaProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGtaProfileResource_Update_APIError exercises GtaProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGtaProfileResource_Update_APIError(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GtaProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_gta_profile")
}

// TestGtaProfileResource_Update_APIErrorReadBody exercises GtaProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGtaProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := GtaProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGtaProfileResource_Update_InvalidJSON exercises GtaProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGtaProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := GtaProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGtaProfileResource_Update_MapError exercises GtaProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGtaProfileResource_Update_MapError(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GtaProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGtaProfileResource_Delete_Happy exercises GtaProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestGtaProfileResource_Delete_Happy(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 204, "")}
	m := GtaProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtaProfileResource_Delete_NilClient exercises GtaProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGtaProfileResource_Delete_NilClient(t *testing.T) {
	r := &GtaProfileResource{}
	m := GtaProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGtaProfileResource_Delete_BuildError exercises GtaProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGtaProfileResource_Delete_BuildError(t *testing.T) {
	r := &GtaProfileResource{client: newMalformedBaseURLClient(t)}
	m := GtaProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGtaProfileResource_Delete_SendError exercises GtaProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestGtaProfileResource_Delete_SendError(t *testing.T) {
	r := &GtaProfileResource{client: newTransportErrorClient(t)}
	m := GtaProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGtaProfileResource_Delete_NotFoundSuccess exercises GtaProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestGtaProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 404, "")}
	m := GtaProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGtaProfileResource_Delete_APIError exercises GtaProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGtaProfileResource_Delete_APIError(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GtaProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_gta_profile")
}

// TestGtaProfileResource_Delete_APIErrorReadBody exercises GtaProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGtaProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &GtaProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := GtaProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
