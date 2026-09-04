package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestHsmGroupResource_Create_Happy exercises HsmGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestHsmGroupResource_Create_Happy(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := HsmGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmGroupResource_Create_NilClient exercises HsmGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHsmGroupResource_Create_NilClient(t *testing.T) {
	r := &HsmGroupResource{}
	m := HsmGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHsmGroupResource_Create_BuildError exercises HsmGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHsmGroupResource_Create_BuildError(t *testing.T) {
	r := &HsmGroupResource{client: newMalformedBaseURLClient(t)}
	m := HsmGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHsmGroupResource_Create_SendError exercises HsmGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestHsmGroupResource_Create_SendError(t *testing.T) {
	r := &HsmGroupResource{client: newTransportErrorClient(t)}
	m := HsmGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHsmGroupResource_Create_APIError exercises HsmGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHsmGroupResource_Create_APIError(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HsmGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_hsm_group")
}

// TestHsmGroupResource_Create_APIErrorReadBody exercises HsmGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHsmGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := HsmGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHsmGroupResource_Create_InvalidJSON exercises HsmGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHsmGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := HsmGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHsmGroupResource_Create_MapError exercises HsmGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHsmGroupResource_Create_MapError(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := HsmGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHsmGroupResource_Create_MissingID exercises HsmGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestHsmGroupResource_Create_MissingID(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := HsmGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestHsmGroupResource_Create_LocationFallback exercises HsmGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestHsmGroupResource_Create_LocationFallback(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := HsmGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestHsmGroupResource_Read_Happy exercises HsmGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestHsmGroupResource_Read_Happy(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := HsmGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmGroupResource_Read_NilClient exercises HsmGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHsmGroupResource_Read_NilClient(t *testing.T) {
	r := &HsmGroupResource{}
	m := HsmGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHsmGroupResource_Read_BuildError exercises HsmGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHsmGroupResource_Read_BuildError(t *testing.T) {
	r := &HsmGroupResource{client: newMalformedBaseURLClient(t)}
	m := HsmGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHsmGroupResource_Read_SendError exercises HsmGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestHsmGroupResource_Read_SendError(t *testing.T) {
	r := &HsmGroupResource{client: newTransportErrorClient(t)}
	m := HsmGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHsmGroupResource_Read_NotFound exercises HsmGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestHsmGroupResource_Read_NotFound(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 404, "")}
	m := HsmGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmGroupResource_Read_APIError exercises HsmGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHsmGroupResource_Read_APIError(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HsmGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_hsm_group")
}

// TestHsmGroupResource_Read_APIErrorReadBody exercises HsmGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHsmGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := HsmGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHsmGroupResource_Read_InvalidJSON exercises HsmGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHsmGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := HsmGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHsmGroupResource_Read_MapError exercises HsmGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHsmGroupResource_Read_MapError(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := HsmGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHsmGroupResource_Update_Happy exercises HsmGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestHsmGroupResource_Update_Happy(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := HsmGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmGroupResource_Update_NilClient exercises HsmGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHsmGroupResource_Update_NilClient(t *testing.T) {
	r := &HsmGroupResource{}
	m := HsmGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHsmGroupResource_Update_BuildError exercises HsmGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHsmGroupResource_Update_BuildError(t *testing.T) {
	r := &HsmGroupResource{client: newMalformedBaseURLClient(t)}
	m := HsmGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHsmGroupResource_Update_SendError exercises HsmGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestHsmGroupResource_Update_SendError(t *testing.T) {
	r := &HsmGroupResource{client: newTransportErrorClient(t)}
	m := HsmGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHsmGroupResource_Update_APIError exercises HsmGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHsmGroupResource_Update_APIError(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HsmGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_hsm_group")
}

// TestHsmGroupResource_Update_APIErrorReadBody exercises HsmGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHsmGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := HsmGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHsmGroupResource_Update_InvalidJSON exercises HsmGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHsmGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := HsmGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHsmGroupResource_Update_MapError exercises HsmGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHsmGroupResource_Update_MapError(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := HsmGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHsmGroupResource_Delete_Happy exercises HsmGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestHsmGroupResource_Delete_Happy(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 204, "")}
	m := HsmGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmGroupResource_Delete_NilClient exercises HsmGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHsmGroupResource_Delete_NilClient(t *testing.T) {
	r := &HsmGroupResource{}
	m := HsmGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHsmGroupResource_Delete_BuildError exercises HsmGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHsmGroupResource_Delete_BuildError(t *testing.T) {
	r := &HsmGroupResource{client: newMalformedBaseURLClient(t)}
	m := HsmGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHsmGroupResource_Delete_SendError exercises HsmGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestHsmGroupResource_Delete_SendError(t *testing.T) {
	r := &HsmGroupResource{client: newTransportErrorClient(t)}
	m := HsmGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHsmGroupResource_Delete_NotFoundSuccess exercises HsmGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestHsmGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 404, "")}
	m := HsmGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmGroupResource_Delete_APIError exercises HsmGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHsmGroupResource_Delete_APIError(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HsmGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_hsm_group")
}

// TestHsmGroupResource_Delete_APIErrorReadBody exercises HsmGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHsmGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &HsmGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := HsmGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
