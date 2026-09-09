package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPortFilterResource_Create_Happy exercises PortFilterResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestPortFilterResource_Create_Happy(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 201, "{\"port\":\"example-id\"}")}
	m := PortFilterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterResource_Create_NilClient exercises PortFilterResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortFilterResource_Create_NilClient(t *testing.T) {
	r := &PortFilterResource{}
	m := PortFilterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortFilterResource_Create_BuildError exercises PortFilterResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortFilterResource_Create_BuildError(t *testing.T) {
	r := &PortFilterResource{client: newMalformedBaseURLClient(t)}
	m := PortFilterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortFilterResource_Create_SendError exercises PortFilterResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortFilterResource_Create_SendError(t *testing.T) {
	r := &PortFilterResource{client: newTransportErrorClient(t)}
	m := PortFilterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortFilterResource_Create_APIError exercises PortFilterResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortFilterResource_Create_APIError(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortFilterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_port_filter")
}

// TestPortFilterResource_Create_APIErrorReadBody exercises PortFilterResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortFilterResource_Create_APIErrorReadBody(t *testing.T) {
	r := &PortFilterResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortFilterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortFilterResource_Create_InvalidJSON exercises PortFilterResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortFilterResource_Create_InvalidJSON(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 201, "{{")}
	m := PortFilterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortFilterResource_Create_MapError exercises PortFilterResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortFilterResource_Create_MapError(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 201, "{\"port\":12345}")}
	m := PortFilterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortFilterResource_Create_MissingID exercises PortFilterResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestPortFilterResource_Create_MissingID(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 201, "{}")}
	m := PortFilterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestPortFilterResource_Create_LocationFallback exercises PortFilterResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestPortFilterResource_Create_LocationFallback(t *testing.T) {
	r := &PortFilterResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := PortFilterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Port.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Port.ValueString(), "example-id")
	}
}

// TestPortFilterResource_Read_Happy exercises PortFilterResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestPortFilterResource_Read_Happy(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 200, "{}")}
	m := PortFilterResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterResource_Read_NilClient exercises PortFilterResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortFilterResource_Read_NilClient(t *testing.T) {
	r := &PortFilterResource{}
	m := PortFilterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortFilterResource_Read_BuildError exercises PortFilterResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortFilterResource_Read_BuildError(t *testing.T) {
	r := &PortFilterResource{client: newMalformedBaseURLClient(t)}
	m := PortFilterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortFilterResource_Read_SendError exercises PortFilterResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortFilterResource_Read_SendError(t *testing.T) {
	r := &PortFilterResource{client: newTransportErrorClient(t)}
	m := PortFilterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortFilterResource_Read_NotFound exercises PortFilterResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestPortFilterResource_Read_NotFound(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 404, "")}
	m := PortFilterResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterResource_Read_APIError exercises PortFilterResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortFilterResource_Read_APIError(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := PortFilterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_port_filter")
}

// TestPortFilterResource_Read_APIErrorReadBody exercises PortFilterResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortFilterResource_Read_APIErrorReadBody(t *testing.T) {
	r := &PortFilterResource{client: newMockClientReadErrorBody(t, 500)}
	m := PortFilterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortFilterResource_Read_InvalidJSON exercises PortFilterResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortFilterResource_Read_InvalidJSON(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortFilterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortFilterResource_Read_MapError exercises PortFilterResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortFilterResource_Read_MapError(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 200, "{\"port\":12345}")}
	m := PortFilterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortFilterResource_Delete_Happy exercises PortFilterResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestPortFilterResource_Delete_Happy(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 204, "")}
	m := PortFilterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterResource_Delete_NilClient exercises PortFilterResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortFilterResource_Delete_NilClient(t *testing.T) {
	r := &PortFilterResource{}
	m := PortFilterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortFilterResource_Delete_BuildError exercises PortFilterResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortFilterResource_Delete_BuildError(t *testing.T) {
	r := &PortFilterResource{client: newMalformedBaseURLClient(t)}
	m := PortFilterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortFilterResource_Delete_SendError exercises PortFilterResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortFilterResource_Delete_SendError(t *testing.T) {
	r := &PortFilterResource{client: newTransportErrorClient(t)}
	m := PortFilterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortFilterResource_Delete_NotFoundSuccess exercises PortFilterResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestPortFilterResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 404, "")}
	m := PortFilterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterResource_Delete_APIError exercises PortFilterResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortFilterResource_Delete_APIError(t *testing.T) {
	r := &PortFilterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortFilterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_port_filter")
}

// TestPortFilterResource_Delete_APIErrorReadBody exercises PortFilterResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortFilterResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &PortFilterResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortFilterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
