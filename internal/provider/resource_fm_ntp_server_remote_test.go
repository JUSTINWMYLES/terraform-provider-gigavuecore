package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFmNtpServerResource_Create_Happy exercises FmNtpServerResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestFmNtpServerResource_Create_Happy(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 201, "{\"server_host\":\"example-id\"}")}
	m := FmNtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmNtpServerResource_Create_NilClient exercises FmNtpServerResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmNtpServerResource_Create_NilClient(t *testing.T) {
	r := &FmNtpServerResource{}
	m := FmNtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmNtpServerResource_Create_BuildError exercises FmNtpServerResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmNtpServerResource_Create_BuildError(t *testing.T) {
	r := &FmNtpServerResource{client: newMalformedBaseURLClient(t)}
	m := FmNtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmNtpServerResource_Create_SendError exercises FmNtpServerResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmNtpServerResource_Create_SendError(t *testing.T) {
	r := &FmNtpServerResource{client: newTransportErrorClient(t)}
	m := FmNtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmNtpServerResource_Create_APIError exercises FmNtpServerResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmNtpServerResource_Create_APIError(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmNtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_fm_ntp_server")
}

// TestFmNtpServerResource_Create_APIErrorReadBody exercises FmNtpServerResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmNtpServerResource_Create_APIErrorReadBody(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := FmNtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFmNtpServerResource_Create_InvalidJSON exercises FmNtpServerResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFmNtpServerResource_Create_InvalidJSON(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 201, "{{")}
	m := FmNtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFmNtpServerResource_Create_MapError exercises FmNtpServerResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFmNtpServerResource_Create_MapError(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 201, "{\"server_host\":12345}")}
	m := FmNtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFmNtpServerResource_Create_MissingID exercises FmNtpServerResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestFmNtpServerResource_Create_MissingID(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 201, "{}")}
	m := FmNtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestFmNtpServerResource_Create_LocationFallback exercises FmNtpServerResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestFmNtpServerResource_Create_LocationFallback(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := FmNtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.ServerHost.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.ServerHost.ValueString(), "example-id")
	}
}

// TestFmNtpServerResource_Read_Happy exercises FmNtpServerResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestFmNtpServerResource_Read_Happy(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := FmNtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmNtpServerResource_Read_NilClient exercises FmNtpServerResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmNtpServerResource_Read_NilClient(t *testing.T) {
	r := &FmNtpServerResource{}
	m := FmNtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmNtpServerResource_Read_BuildError exercises FmNtpServerResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmNtpServerResource_Read_BuildError(t *testing.T) {
	r := &FmNtpServerResource{client: newMalformedBaseURLClient(t)}
	m := FmNtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmNtpServerResource_Read_SendError exercises FmNtpServerResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmNtpServerResource_Read_SendError(t *testing.T) {
	r := &FmNtpServerResource{client: newTransportErrorClient(t)}
	m := FmNtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmNtpServerResource_Read_NotFound exercises FmNtpServerResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestFmNtpServerResource_Read_NotFound(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 404, "")}
	m := FmNtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmNtpServerResource_Read_APIError exercises FmNtpServerResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmNtpServerResource_Read_APIError(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmNtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_fm_ntp_server")
}

// TestFmNtpServerResource_Read_APIErrorReadBody exercises FmNtpServerResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmNtpServerResource_Read_APIErrorReadBody(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := FmNtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFmNtpServerResource_Read_InvalidJSON exercises FmNtpServerResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFmNtpServerResource_Read_InvalidJSON(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := FmNtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFmNtpServerResource_Read_MapError exercises FmNtpServerResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFmNtpServerResource_Read_MapError(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 200, "{\"server_host\":12345}")}
	m := FmNtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFmNtpServerResource_Delete_Happy exercises FmNtpServerResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestFmNtpServerResource_Delete_Happy(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 204, "")}
	m := FmNtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmNtpServerResource_Delete_NilClient exercises FmNtpServerResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmNtpServerResource_Delete_NilClient(t *testing.T) {
	r := &FmNtpServerResource{}
	m := FmNtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmNtpServerResource_Delete_BuildError exercises FmNtpServerResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmNtpServerResource_Delete_BuildError(t *testing.T) {
	r := &FmNtpServerResource{client: newMalformedBaseURLClient(t)}
	m := FmNtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmNtpServerResource_Delete_SendError exercises FmNtpServerResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmNtpServerResource_Delete_SendError(t *testing.T) {
	r := &FmNtpServerResource{client: newTransportErrorClient(t)}
	m := FmNtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmNtpServerResource_Delete_NotFoundSuccess exercises FmNtpServerResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestFmNtpServerResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 404, "")}
	m := FmNtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmNtpServerResource_Delete_APIError exercises FmNtpServerResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmNtpServerResource_Delete_APIError(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmNtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_fm_ntp_server")
}

// TestFmNtpServerResource_Delete_APIErrorReadBody exercises FmNtpServerResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmNtpServerResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &FmNtpServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := FmNtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
