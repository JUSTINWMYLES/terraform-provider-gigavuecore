package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestHeaderStripResource_Create_Happy exercises HeaderStripResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestHeaderStripResource_Create_Happy(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 201, "{\"box_id\":\"example-id\"}")}
	m := HeaderStripResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHeaderStripResource_Create_NilClient exercises HeaderStripResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHeaderStripResource_Create_NilClient(t *testing.T) {
	r := &HeaderStripResource{}
	m := HeaderStripResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHeaderStripResource_Create_BuildError exercises HeaderStripResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHeaderStripResource_Create_BuildError(t *testing.T) {
	r := &HeaderStripResource{client: newMalformedBaseURLClient(t)}
	m := HeaderStripResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHeaderStripResource_Create_SendError exercises HeaderStripResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestHeaderStripResource_Create_SendError(t *testing.T) {
	r := &HeaderStripResource{client: newTransportErrorClient(t)}
	m := HeaderStripResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHeaderStripResource_Create_APIError exercises HeaderStripResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHeaderStripResource_Create_APIError(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HeaderStripResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_header_strip")
}

// TestHeaderStripResource_Create_APIErrorReadBody exercises HeaderStripResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHeaderStripResource_Create_APIErrorReadBody(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientReadErrorBody(t, 501)}
	m := HeaderStripResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHeaderStripResource_Create_InvalidJSON exercises HeaderStripResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHeaderStripResource_Create_InvalidJSON(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 201, "{{")}
	m := HeaderStripResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHeaderStripResource_Create_MapError exercises HeaderStripResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHeaderStripResource_Create_MapError(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 201, "{\"box_id\":12345}")}
	m := HeaderStripResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHeaderStripResource_Create_MissingID exercises HeaderStripResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestHeaderStripResource_Create_MissingID(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 201, "{}")}
	m := HeaderStripResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestHeaderStripResource_Create_LocationFallback exercises HeaderStripResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestHeaderStripResource_Create_LocationFallback(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := HeaderStripResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.BoxId.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.BoxId.ValueString(), "example-id")
	}
}

// TestHeaderStripResource_Read_Happy exercises HeaderStripResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestHeaderStripResource_Read_Happy(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 200, "{}")}
	m := HeaderStripResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestHeaderStripResource_Read_NilClient exercises HeaderStripResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHeaderStripResource_Read_NilClient(t *testing.T) {
	r := &HeaderStripResource{}
	m := HeaderStripResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHeaderStripResource_Read_BuildError exercises HeaderStripResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHeaderStripResource_Read_BuildError(t *testing.T) {
	r := &HeaderStripResource{client: newMalformedBaseURLClient(t)}
	m := HeaderStripResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHeaderStripResource_Read_SendError exercises HeaderStripResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestHeaderStripResource_Read_SendError(t *testing.T) {
	r := &HeaderStripResource{client: newTransportErrorClient(t)}
	m := HeaderStripResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHeaderStripResource_Read_NotFound exercises HeaderStripResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestHeaderStripResource_Read_NotFound(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 404, "")}
	m := HeaderStripResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestHeaderStripResource_Read_APIError exercises HeaderStripResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHeaderStripResource_Read_APIError(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HeaderStripResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_header_strip")
}

// TestHeaderStripResource_Read_APIErrorReadBody exercises HeaderStripResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHeaderStripResource_Read_APIErrorReadBody(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientReadErrorBody(t, 501)}
	m := HeaderStripResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestHeaderStripResource_Read_InvalidJSON exercises HeaderStripResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestHeaderStripResource_Read_InvalidJSON(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 200, "{{")}
	m := HeaderStripResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestHeaderStripResource_Read_MapError exercises HeaderStripResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestHeaderStripResource_Read_MapError(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 200, "{\"box_id\":12345}")}
	m := HeaderStripResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestHeaderStripResource_Delete_Happy exercises HeaderStripResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestHeaderStripResource_Delete_Happy(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 204, "")}
	m := HeaderStripResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHeaderStripResource_Delete_NilClient exercises HeaderStripResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHeaderStripResource_Delete_NilClient(t *testing.T) {
	r := &HeaderStripResource{}
	m := HeaderStripResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHeaderStripResource_Delete_BuildError exercises HeaderStripResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHeaderStripResource_Delete_BuildError(t *testing.T) {
	r := &HeaderStripResource{client: newMalformedBaseURLClient(t)}
	m := HeaderStripResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHeaderStripResource_Delete_SendError exercises HeaderStripResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestHeaderStripResource_Delete_SendError(t *testing.T) {
	r := &HeaderStripResource{client: newTransportErrorClient(t)}
	m := HeaderStripResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHeaderStripResource_Delete_NotFoundSuccess exercises HeaderStripResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestHeaderStripResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 404, "")}
	m := HeaderStripResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHeaderStripResource_Delete_APIError exercises HeaderStripResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHeaderStripResource_Delete_APIError(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HeaderStripResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_header_strip")
}

// TestHeaderStripResource_Delete_APIErrorReadBody exercises HeaderStripResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHeaderStripResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &HeaderStripResource{client: newMockClientReadErrorBody(t, 501)}
	m := HeaderStripResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
