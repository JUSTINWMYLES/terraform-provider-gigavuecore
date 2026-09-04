package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestHsmResource_Create_Happy exercises HsmResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestHsmResource_Create_Happy(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := HsmResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmResource_Create_NilClient exercises HsmResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHsmResource_Create_NilClient(t *testing.T) {
	r := &HsmResource{}
	m := HsmResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHsmResource_Create_BuildError exercises HsmResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHsmResource_Create_BuildError(t *testing.T) {
	r := &HsmResource{client: newMalformedBaseURLClient(t)}
	m := HsmResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHsmResource_Create_SendError exercises HsmResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestHsmResource_Create_SendError(t *testing.T) {
	r := &HsmResource{client: newTransportErrorClient(t)}
	m := HsmResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHsmResource_Create_APIError exercises HsmResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHsmResource_Create_APIError(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HsmResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_hsm")
}

// TestHsmResource_Create_APIErrorReadBody exercises HsmResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHsmResource_Create_APIErrorReadBody(t *testing.T) {
	r := &HsmResource{client: newMockClientReadErrorBody(t, 501)}
	m := HsmResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHsmResource_Create_InvalidJSON exercises HsmResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHsmResource_Create_InvalidJSON(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 201, "{{")}
	m := HsmResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHsmResource_Create_MapError exercises HsmResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHsmResource_Create_MapError(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := HsmResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHsmResource_Create_MissingID exercises HsmResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestHsmResource_Create_MissingID(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 201, "{}")}
	m := HsmResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestHsmResource_Create_LocationFallback exercises HsmResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestHsmResource_Create_LocationFallback(t *testing.T) {
	r := &HsmResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := HsmResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestHsmResource_Read_Happy exercises HsmResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestHsmResource_Read_Happy(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 200, "{}")}
	m := HsmResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmResource_Read_NilClient exercises HsmResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHsmResource_Read_NilClient(t *testing.T) {
	r := &HsmResource{}
	m := HsmResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHsmResource_Read_BuildError exercises HsmResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHsmResource_Read_BuildError(t *testing.T) {
	r := &HsmResource{client: newMalformedBaseURLClient(t)}
	m := HsmResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHsmResource_Read_SendError exercises HsmResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestHsmResource_Read_SendError(t *testing.T) {
	r := &HsmResource{client: newTransportErrorClient(t)}
	m := HsmResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHsmResource_Read_NotFound exercises HsmResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestHsmResource_Read_NotFound(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 404, "")}
	m := HsmResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmResource_Read_APIError exercises HsmResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHsmResource_Read_APIError(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HsmResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_hsm")
}

// TestHsmResource_Read_APIErrorReadBody exercises HsmResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHsmResource_Read_APIErrorReadBody(t *testing.T) {
	r := &HsmResource{client: newMockClientReadErrorBody(t, 501)}
	m := HsmResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHsmResource_Read_InvalidJSON exercises HsmResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHsmResource_Read_InvalidJSON(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 200, "{{")}
	m := HsmResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHsmResource_Read_MapError exercises HsmResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHsmResource_Read_MapError(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := HsmResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHsmResource_Update_Happy exercises HsmResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestHsmResource_Update_Happy(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 200, "{}")}
	m := HsmResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmResource_Update_NilClient exercises HsmResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHsmResource_Update_NilClient(t *testing.T) {
	r := &HsmResource{}
	m := HsmResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHsmResource_Update_BuildError exercises HsmResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHsmResource_Update_BuildError(t *testing.T) {
	r := &HsmResource{client: newMalformedBaseURLClient(t)}
	m := HsmResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHsmResource_Update_SendError exercises HsmResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestHsmResource_Update_SendError(t *testing.T) {
	r := &HsmResource{client: newTransportErrorClient(t)}
	m := HsmResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHsmResource_Update_APIError exercises HsmResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHsmResource_Update_APIError(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HsmResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_hsm")
}

// TestHsmResource_Update_APIErrorReadBody exercises HsmResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHsmResource_Update_APIErrorReadBody(t *testing.T) {
	r := &HsmResource{client: newMockClientReadErrorBody(t, 501)}
	m := HsmResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHsmResource_Update_InvalidJSON exercises HsmResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHsmResource_Update_InvalidJSON(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 200, "{{")}
	m := HsmResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHsmResource_Update_MapError exercises HsmResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHsmResource_Update_MapError(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := HsmResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHsmResource_Delete_Happy exercises HsmResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestHsmResource_Delete_Happy(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 204, "")}
	m := HsmResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmResource_Delete_NilClient exercises HsmResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHsmResource_Delete_NilClient(t *testing.T) {
	r := &HsmResource{}
	m := HsmResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHsmResource_Delete_BuildError exercises HsmResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHsmResource_Delete_BuildError(t *testing.T) {
	r := &HsmResource{client: newMalformedBaseURLClient(t)}
	m := HsmResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHsmResource_Delete_SendError exercises HsmResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestHsmResource_Delete_SendError(t *testing.T) {
	r := &HsmResource{client: newTransportErrorClient(t)}
	m := HsmResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHsmResource_Delete_NotFoundSuccess exercises HsmResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestHsmResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 404, "")}
	m := HsmResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHsmResource_Delete_APIError exercises HsmResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHsmResource_Delete_APIError(t *testing.T) {
	r := &HsmResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HsmResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_hsm")
}

// TestHsmResource_Delete_APIErrorReadBody exercises HsmResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHsmResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &HsmResource{client: newMockClientReadErrorBody(t, 501)}
	m := HsmResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
