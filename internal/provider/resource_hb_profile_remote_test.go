package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestHbProfileResource_Create_Happy exercises HbProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestHbProfileResource_Create_Happy(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := HbProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbProfileResource_Create_NilClient exercises HbProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHbProfileResource_Create_NilClient(t *testing.T) {
	r := &HbProfileResource{}
	m := HbProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHbProfileResource_Create_BuildError exercises HbProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHbProfileResource_Create_BuildError(t *testing.T) {
	r := &HbProfileResource{client: newMalformedBaseURLClient(t)}
	m := HbProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHbProfileResource_Create_SendError exercises HbProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestHbProfileResource_Create_SendError(t *testing.T) {
	r := &HbProfileResource{client: newTransportErrorClient(t)}
	m := HbProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHbProfileResource_Create_APIError exercises HbProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHbProfileResource_Create_APIError(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HbProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_hb_profile")
}

// TestHbProfileResource_Create_APIErrorReadBody exercises HbProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHbProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &HbProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := HbProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHbProfileResource_Create_InvalidJSON exercises HbProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHbProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := HbProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHbProfileResource_Create_MapError exercises HbProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHbProfileResource_Create_MapError(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := HbProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHbProfileResource_Create_MissingID exercises HbProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestHbProfileResource_Create_MissingID(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := HbProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestHbProfileResource_Create_LocationFallback exercises HbProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestHbProfileResource_Create_LocationFallback(t *testing.T) {
	r := &HbProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := HbProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestHbProfileResource_Read_Happy exercises HbProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestHbProfileResource_Read_Happy(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := HbProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbProfileResource_Read_NilClient exercises HbProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHbProfileResource_Read_NilClient(t *testing.T) {
	r := &HbProfileResource{}
	m := HbProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHbProfileResource_Read_BuildError exercises HbProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHbProfileResource_Read_BuildError(t *testing.T) {
	r := &HbProfileResource{client: newMalformedBaseURLClient(t)}
	m := HbProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHbProfileResource_Read_SendError exercises HbProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestHbProfileResource_Read_SendError(t *testing.T) {
	r := &HbProfileResource{client: newTransportErrorClient(t)}
	m := HbProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHbProfileResource_Read_NotFound exercises HbProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestHbProfileResource_Read_NotFound(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 404, "")}
	m := HbProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbProfileResource_Read_APIError exercises HbProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHbProfileResource_Read_APIError(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HbProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_hb_profile")
}

// TestHbProfileResource_Read_APIErrorReadBody exercises HbProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHbProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &HbProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := HbProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHbProfileResource_Read_InvalidJSON exercises HbProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHbProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := HbProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHbProfileResource_Read_MapError exercises HbProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHbProfileResource_Read_MapError(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := HbProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHbProfileResource_Update_Happy exercises HbProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestHbProfileResource_Update_Happy(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := HbProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbProfileResource_Update_NilClient exercises HbProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHbProfileResource_Update_NilClient(t *testing.T) {
	r := &HbProfileResource{}
	m := HbProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHbProfileResource_Update_BuildError exercises HbProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHbProfileResource_Update_BuildError(t *testing.T) {
	r := &HbProfileResource{client: newMalformedBaseURLClient(t)}
	m := HbProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHbProfileResource_Update_SendError exercises HbProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestHbProfileResource_Update_SendError(t *testing.T) {
	r := &HbProfileResource{client: newTransportErrorClient(t)}
	m := HbProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHbProfileResource_Update_APIError exercises HbProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHbProfileResource_Update_APIError(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HbProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_hb_profile")
}

// TestHbProfileResource_Update_APIErrorReadBody exercises HbProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHbProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &HbProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := HbProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHbProfileResource_Update_InvalidJSON exercises HbProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHbProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := HbProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHbProfileResource_Update_MapError exercises HbProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHbProfileResource_Update_MapError(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := HbProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHbProfileResource_Delete_Happy exercises HbProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestHbProfileResource_Delete_Happy(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 204, "")}
	m := HbProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbProfileResource_Delete_NilClient exercises HbProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHbProfileResource_Delete_NilClient(t *testing.T) {
	r := &HbProfileResource{}
	m := HbProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHbProfileResource_Delete_BuildError exercises HbProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHbProfileResource_Delete_BuildError(t *testing.T) {
	r := &HbProfileResource{client: newMalformedBaseURLClient(t)}
	m := HbProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHbProfileResource_Delete_SendError exercises HbProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestHbProfileResource_Delete_SendError(t *testing.T) {
	r := &HbProfileResource{client: newTransportErrorClient(t)}
	m := HbProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHbProfileResource_Delete_NotFoundSuccess exercises HbProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestHbProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 404, "")}
	m := HbProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHbProfileResource_Delete_APIError exercises HbProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHbProfileResource_Delete_APIError(t *testing.T) {
	r := &HbProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HbProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_hb_profile")
}

// TestHbProfileResource_Delete_APIErrorReadBody exercises HbProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHbProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &HbProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := HbProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
