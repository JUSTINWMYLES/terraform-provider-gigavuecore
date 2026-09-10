package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNtpServerResource_Create_Happy exercises NtpServerResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestNtpServerResource_Create_Happy(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 201, "{\"server\":\"example-id\"}")}
	m := NtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNtpServerResource_Create_NilClient exercises NtpServerResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNtpServerResource_Create_NilClient(t *testing.T) {
	r := &NtpServerResource{}
	m := NtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNtpServerResource_Create_BuildError exercises NtpServerResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNtpServerResource_Create_BuildError(t *testing.T) {
	r := &NtpServerResource{client: newMalformedBaseURLClient(t)}
	m := NtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNtpServerResource_Create_SendError exercises NtpServerResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestNtpServerResource_Create_SendError(t *testing.T) {
	r := &NtpServerResource{client: newTransportErrorClient(t)}
	m := NtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNtpServerResource_Create_APIError exercises NtpServerResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNtpServerResource_Create_APIError(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_ntp_server")
}

// TestNtpServerResource_Create_APIErrorReadBody exercises NtpServerResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNtpServerResource_Create_APIErrorReadBody(t *testing.T) {
	r := &NtpServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := NtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNtpServerResource_Create_InvalidJSON exercises NtpServerResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNtpServerResource_Create_InvalidJSON(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 201, "{{")}
	m := NtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNtpServerResource_Create_MapError exercises NtpServerResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNtpServerResource_Create_MapError(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 201, "{\"server\":12345}")}
	m := NtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNtpServerResource_Create_MissingID exercises NtpServerResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestNtpServerResource_Create_MissingID(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 201, "{}")}
	m := NtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestNtpServerResource_Create_LocationFallback exercises NtpServerResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestNtpServerResource_Create_LocationFallback(t *testing.T) {
	r := &NtpServerResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := NtpServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Server.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Server.ValueString(), "example-id")
	}
}

// TestNtpServerResource_Read_Happy exercises NtpServerResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestNtpServerResource_Read_Happy(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := NtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNtpServerResource_Read_NilClient exercises NtpServerResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNtpServerResource_Read_NilClient(t *testing.T) {
	r := &NtpServerResource{}
	m := NtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNtpServerResource_Read_BuildError exercises NtpServerResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNtpServerResource_Read_BuildError(t *testing.T) {
	r := &NtpServerResource{client: newMalformedBaseURLClient(t)}
	m := NtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNtpServerResource_Read_SendError exercises NtpServerResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestNtpServerResource_Read_SendError(t *testing.T) {
	r := &NtpServerResource{client: newTransportErrorClient(t)}
	m := NtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNtpServerResource_Read_NotFound exercises NtpServerResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestNtpServerResource_Read_NotFound(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 404, "")}
	m := NtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNtpServerResource_Read_APIError exercises NtpServerResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNtpServerResource_Read_APIError(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_ntp_server")
}

// TestNtpServerResource_Read_APIErrorReadBody exercises NtpServerResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNtpServerResource_Read_APIErrorReadBody(t *testing.T) {
	r := &NtpServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := NtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNtpServerResource_Read_InvalidJSON exercises NtpServerResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNtpServerResource_Read_InvalidJSON(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := NtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNtpServerResource_Read_MapError exercises NtpServerResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNtpServerResource_Read_MapError(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 200, "{\"server\":12345}")}
	m := NtpServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNtpServerResource_Update_Happy exercises NtpServerResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestNtpServerResource_Update_Happy(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := NtpServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNtpServerResource_Update_NilClient exercises NtpServerResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNtpServerResource_Update_NilClient(t *testing.T) {
	r := &NtpServerResource{}
	m := NtpServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNtpServerResource_Update_BuildError exercises NtpServerResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNtpServerResource_Update_BuildError(t *testing.T) {
	r := &NtpServerResource{client: newMalformedBaseURLClient(t)}
	m := NtpServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNtpServerResource_Update_SendError exercises NtpServerResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestNtpServerResource_Update_SendError(t *testing.T) {
	r := &NtpServerResource{client: newTransportErrorClient(t)}
	m := NtpServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNtpServerResource_Update_APIError exercises NtpServerResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNtpServerResource_Update_APIError(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NtpServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_ntp_server")
}

// TestNtpServerResource_Update_APIErrorReadBody exercises NtpServerResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNtpServerResource_Update_APIErrorReadBody(t *testing.T) {
	r := &NtpServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := NtpServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNtpServerResource_Update_InvalidJSON exercises NtpServerResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNtpServerResource_Update_InvalidJSON(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := NtpServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNtpServerResource_Update_MapError exercises NtpServerResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNtpServerResource_Update_MapError(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 200, "{\"server\":12345}")}
	m := NtpServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNtpServerResource_Delete_Happy exercises NtpServerResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestNtpServerResource_Delete_Happy(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 204, "")}
	m := NtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNtpServerResource_Delete_NilClient exercises NtpServerResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNtpServerResource_Delete_NilClient(t *testing.T) {
	r := &NtpServerResource{}
	m := NtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNtpServerResource_Delete_BuildError exercises NtpServerResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNtpServerResource_Delete_BuildError(t *testing.T) {
	r := &NtpServerResource{client: newMalformedBaseURLClient(t)}
	m := NtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNtpServerResource_Delete_SendError exercises NtpServerResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestNtpServerResource_Delete_SendError(t *testing.T) {
	r := &NtpServerResource{client: newTransportErrorClient(t)}
	m := NtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNtpServerResource_Delete_NotFoundSuccess exercises NtpServerResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestNtpServerResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 404, "")}
	m := NtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNtpServerResource_Delete_APIError exercises NtpServerResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNtpServerResource_Delete_APIError(t *testing.T) {
	r := &NtpServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_ntp_server")
}

// TestNtpServerResource_Delete_APIErrorReadBody exercises NtpServerResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNtpServerResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &NtpServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := NtpServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
