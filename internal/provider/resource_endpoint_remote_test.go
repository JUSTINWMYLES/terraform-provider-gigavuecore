package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestEndpointResource_Create_Happy exercises EndpointResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestEndpointResource_Create_Happy(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := EndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEndpointResource_Create_NilClient exercises EndpointResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEndpointResource_Create_NilClient(t *testing.T) {
	r := &EndpointResource{}
	m := EndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEndpointResource_Create_BuildError exercises EndpointResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEndpointResource_Create_BuildError(t *testing.T) {
	r := &EndpointResource{client: newMalformedBaseURLClient(t)}
	m := EndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEndpointResource_Create_SendError exercises EndpointResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestEndpointResource_Create_SendError(t *testing.T) {
	r := &EndpointResource{client: newTransportErrorClient(t)}
	m := EndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEndpointResource_Create_APIError exercises EndpointResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEndpointResource_Create_APIError(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_endpoint")
}

// TestEndpointResource_Create_APIErrorReadBody exercises EndpointResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEndpointResource_Create_APIErrorReadBody(t *testing.T) {
	r := &EndpointResource{client: newMockClientReadErrorBody(t, 501)}
	m := EndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestEndpointResource_Create_InvalidJSON exercises EndpointResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestEndpointResource_Create_InvalidJSON(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 201, "{{")}
	m := EndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestEndpointResource_Create_MapError exercises EndpointResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestEndpointResource_Create_MapError(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := EndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestEndpointResource_Create_MissingID exercises EndpointResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestEndpointResource_Create_MissingID(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 201, "{}")}
	m := EndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestEndpointResource_Create_LocationFallback exercises EndpointResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestEndpointResource_Create_LocationFallback(t *testing.T) {
	r := &EndpointResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := EndpointResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestEndpointResource_Read_Happy exercises EndpointResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestEndpointResource_Read_Happy(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 200, "{}")}
	m := EndpointResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestEndpointResource_Read_NilClient exercises EndpointResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEndpointResource_Read_NilClient(t *testing.T) {
	r := &EndpointResource{}
	m := EndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEndpointResource_Read_BuildError exercises EndpointResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEndpointResource_Read_BuildError(t *testing.T) {
	r := &EndpointResource{client: newMalformedBaseURLClient(t)}
	m := EndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEndpointResource_Read_SendError exercises EndpointResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestEndpointResource_Read_SendError(t *testing.T) {
	r := &EndpointResource{client: newTransportErrorClient(t)}
	m := EndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEndpointResource_Read_NotFound exercises EndpointResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestEndpointResource_Read_NotFound(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 404, "")}
	m := EndpointResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestEndpointResource_Read_APIError exercises EndpointResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEndpointResource_Read_APIError(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_endpoint")
}

// TestEndpointResource_Read_APIErrorReadBody exercises EndpointResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEndpointResource_Read_APIErrorReadBody(t *testing.T) {
	r := &EndpointResource{client: newMockClientReadErrorBody(t, 501)}
	m := EndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestEndpointResource_Read_InvalidJSON exercises EndpointResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestEndpointResource_Read_InvalidJSON(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 200, "{{")}
	m := EndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestEndpointResource_Read_MapError exercises EndpointResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestEndpointResource_Read_MapError(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := EndpointResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestEndpointResource_Delete_Happy exercises EndpointResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestEndpointResource_Delete_Happy(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 204, "")}
	m := EndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEndpointResource_Delete_NilClient exercises EndpointResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEndpointResource_Delete_NilClient(t *testing.T) {
	r := &EndpointResource{}
	m := EndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEndpointResource_Delete_BuildError exercises EndpointResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEndpointResource_Delete_BuildError(t *testing.T) {
	r := &EndpointResource{client: newMalformedBaseURLClient(t)}
	m := EndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEndpointResource_Delete_SendError exercises EndpointResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestEndpointResource_Delete_SendError(t *testing.T) {
	r := &EndpointResource{client: newTransportErrorClient(t)}
	m := EndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEndpointResource_Delete_NotFoundSuccess exercises EndpointResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestEndpointResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 404, "")}
	m := EndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEndpointResource_Delete_APIError exercises EndpointResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEndpointResource_Delete_APIError(t *testing.T) {
	r := &EndpointResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_endpoint")
}

// TestEndpointResource_Delete_APIErrorReadBody exercises EndpointResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEndpointResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &EndpointResource{client: newMockClientReadErrorBody(t, 501)}
	m := EndpointResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
