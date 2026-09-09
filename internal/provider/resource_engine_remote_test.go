package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestEngineResource_Create_Happy exercises EngineResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestEngineResource_Create_Happy(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 201, "{\"eport\":\"example-id\"}")}
	m := EngineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEngineResource_Create_NilClient exercises EngineResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEngineResource_Create_NilClient(t *testing.T) {
	r := &EngineResource{}
	m := EngineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEngineResource_Create_BuildError exercises EngineResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEngineResource_Create_BuildError(t *testing.T) {
	r := &EngineResource{client: newMalformedBaseURLClient(t)}
	m := EngineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEngineResource_Create_SendError exercises EngineResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestEngineResource_Create_SendError(t *testing.T) {
	r := &EngineResource{client: newTransportErrorClient(t)}
	m := EngineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEngineResource_Create_APIError exercises EngineResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEngineResource_Create_APIError(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EngineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_engine")
}

// TestEngineResource_Create_APIErrorReadBody exercises EngineResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEngineResource_Create_APIErrorReadBody(t *testing.T) {
	r := &EngineResource{client: newMockClientReadErrorBody(t, 501)}
	m := EngineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestEngineResource_Create_InvalidJSON exercises EngineResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestEngineResource_Create_InvalidJSON(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 201, "{{")}
	m := EngineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestEngineResource_Create_MapError exercises EngineResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestEngineResource_Create_MapError(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 201, "{\"eport\":12345}")}
	m := EngineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestEngineResource_Create_MissingID exercises EngineResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestEngineResource_Create_MissingID(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 201, "{}")}
	m := EngineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestEngineResource_Create_LocationFallback exercises EngineResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestEngineResource_Create_LocationFallback(t *testing.T) {
	r := &EngineResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := EngineResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Eport.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Eport.ValueString(), "example-id")
	}
}

// TestEngineResource_Read_Happy exercises EngineResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestEngineResource_Read_Happy(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 200, "{}")}
	m := EngineResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestEngineResource_Read_NilClient exercises EngineResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEngineResource_Read_NilClient(t *testing.T) {
	r := &EngineResource{}
	m := EngineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEngineResource_Read_BuildError exercises EngineResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEngineResource_Read_BuildError(t *testing.T) {
	r := &EngineResource{client: newMalformedBaseURLClient(t)}
	m := EngineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEngineResource_Read_SendError exercises EngineResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestEngineResource_Read_SendError(t *testing.T) {
	r := &EngineResource{client: newTransportErrorClient(t)}
	m := EngineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEngineResource_Read_NotFound exercises EngineResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestEngineResource_Read_NotFound(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 404, "")}
	m := EngineResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestEngineResource_Read_APIError exercises EngineResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEngineResource_Read_APIError(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EngineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_engine")
}

// TestEngineResource_Read_APIErrorReadBody exercises EngineResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEngineResource_Read_APIErrorReadBody(t *testing.T) {
	r := &EngineResource{client: newMockClientReadErrorBody(t, 501)}
	m := EngineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestEngineResource_Read_InvalidJSON exercises EngineResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestEngineResource_Read_InvalidJSON(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 200, "{{")}
	m := EngineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestEngineResource_Read_MapError exercises EngineResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestEngineResource_Read_MapError(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 200, "{\"eport\":12345}")}
	m := EngineResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestEngineResource_Delete_Happy exercises EngineResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestEngineResource_Delete_Happy(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 204, "")}
	m := EngineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEngineResource_Delete_NilClient exercises EngineResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEngineResource_Delete_NilClient(t *testing.T) {
	r := &EngineResource{}
	m := EngineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEngineResource_Delete_BuildError exercises EngineResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEngineResource_Delete_BuildError(t *testing.T) {
	r := &EngineResource{client: newMalformedBaseURLClient(t)}
	m := EngineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEngineResource_Delete_SendError exercises EngineResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestEngineResource_Delete_SendError(t *testing.T) {
	r := &EngineResource{client: newTransportErrorClient(t)}
	m := EngineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEngineResource_Delete_NotFoundSuccess exercises EngineResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestEngineResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 404, "")}
	m := EngineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEngineResource_Delete_APIError exercises EngineResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEngineResource_Delete_APIError(t *testing.T) {
	r := &EngineResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EngineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_engine")
}

// TestEngineResource_Delete_APIErrorReadBody exercises EngineResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEngineResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &EngineResource{client: newMockClientReadErrorBody(t, 501)}
	m := EngineResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
