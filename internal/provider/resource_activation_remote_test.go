package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestActivationResource_Create_Happy exercises ActivationResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestActivationResource_Create_Happy(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 200, "{\"eli_id\":\"example-id\"}")}
	m := ActivationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestActivationResource_Create_NilClient exercises ActivationResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestActivationResource_Create_NilClient(t *testing.T) {
	r := &ActivationResource{}
	m := ActivationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestActivationResource_Create_BuildError exercises ActivationResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestActivationResource_Create_BuildError(t *testing.T) {
	r := &ActivationResource{client: newMalformedBaseURLClient(t)}
	m := ActivationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestActivationResource_Create_SendError exercises ActivationResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestActivationResource_Create_SendError(t *testing.T) {
	r := &ActivationResource{client: newTransportErrorClient(t)}
	m := ActivationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestActivationResource_Create_APIError exercises ActivationResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestActivationResource_Create_APIError(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ActivationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_activation")
}

// TestActivationResource_Create_APIErrorReadBody exercises ActivationResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestActivationResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ActivationResource{client: newMockClientReadErrorBody(t, 501)}
	m := ActivationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestActivationResource_Create_InvalidJSON exercises ActivationResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestActivationResource_Create_InvalidJSON(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 200, "{{")}
	m := ActivationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestActivationResource_Create_MapError exercises ActivationResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestActivationResource_Create_MapError(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 200, "{\"eli_id\":12345}")}
	m := ActivationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestActivationResource_Create_MissingID exercises ActivationResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestActivationResource_Create_MissingID(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 200, "{}")}
	m := ActivationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestActivationResource_Create_LocationFallback exercises ActivationResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestActivationResource_Create_LocationFallback(t *testing.T) {
	r := &ActivationResource{client: newMockClientWithLocation(t, 200, "http://example.test/folders/example-id", "{}")}
	m := ActivationResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.EliId.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.EliId.ValueString(), "example-id")
	}
}

// TestActivationResource_Read_Happy exercises ActivationResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestActivationResource_Read_Happy(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 200, "{}")}
	m := ActivationResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestActivationResource_Read_NilClient exercises ActivationResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestActivationResource_Read_NilClient(t *testing.T) {
	r := &ActivationResource{}
	m := ActivationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestActivationResource_Read_BuildError exercises ActivationResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestActivationResource_Read_BuildError(t *testing.T) {
	r := &ActivationResource{client: newMalformedBaseURLClient(t)}
	m := ActivationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestActivationResource_Read_SendError exercises ActivationResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestActivationResource_Read_SendError(t *testing.T) {
	r := &ActivationResource{client: newTransportErrorClient(t)}
	m := ActivationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestActivationResource_Read_NotFound exercises ActivationResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestActivationResource_Read_NotFound(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 404, "")}
	m := ActivationResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestActivationResource_Read_APIError exercises ActivationResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestActivationResource_Read_APIError(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ActivationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_activation")
}

// TestActivationResource_Read_APIErrorReadBody exercises ActivationResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestActivationResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ActivationResource{client: newMockClientReadErrorBody(t, 501)}
	m := ActivationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestActivationResource_Read_InvalidJSON exercises ActivationResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestActivationResource_Read_InvalidJSON(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 200, "{{")}
	m := ActivationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestActivationResource_Read_MapError exercises ActivationResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestActivationResource_Read_MapError(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 200, "{\"eli_id\":12345}")}
	m := ActivationResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestActivationResource_Delete_Happy exercises ActivationResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestActivationResource_Delete_Happy(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 200, "")}
	m := ActivationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestActivationResource_Delete_NilClient exercises ActivationResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestActivationResource_Delete_NilClient(t *testing.T) {
	r := &ActivationResource{}
	m := ActivationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestActivationResource_Delete_BuildError exercises ActivationResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestActivationResource_Delete_BuildError(t *testing.T) {
	r := &ActivationResource{client: newMalformedBaseURLClient(t)}
	m := ActivationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestActivationResource_Delete_SendError exercises ActivationResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestActivationResource_Delete_SendError(t *testing.T) {
	r := &ActivationResource{client: newTransportErrorClient(t)}
	m := ActivationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestActivationResource_Delete_NotFoundSuccess exercises ActivationResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestActivationResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 404, "")}
	m := ActivationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestActivationResource_Delete_APIError exercises ActivationResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestActivationResource_Delete_APIError(t *testing.T) {
	r := &ActivationResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ActivationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_activation")
}

// TestActivationResource_Delete_APIErrorReadBody exercises ActivationResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestActivationResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ActivationResource{client: newMockClientReadErrorBody(t, 501)}
	m := ActivationResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
