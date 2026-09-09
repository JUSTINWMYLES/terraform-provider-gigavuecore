package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestTcpProfileResource_Create_Happy exercises TcpProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestTcpProfileResource_Create_Happy(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := TcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTcpProfileResource_Create_NilClient exercises TcpProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTcpProfileResource_Create_NilClient(t *testing.T) {
	r := &TcpProfileResource{}
	m := TcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTcpProfileResource_Create_BuildError exercises TcpProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTcpProfileResource_Create_BuildError(t *testing.T) {
	r := &TcpProfileResource{client: newMalformedBaseURLClient(t)}
	m := TcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTcpProfileResource_Create_SendError exercises TcpProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestTcpProfileResource_Create_SendError(t *testing.T) {
	r := &TcpProfileResource{client: newTransportErrorClient(t)}
	m := TcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTcpProfileResource_Create_APIError exercises TcpProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTcpProfileResource_Create_APIError(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_tcp_profile")
}

// TestTcpProfileResource_Create_APIErrorReadBody exercises TcpProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTcpProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := TcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTcpProfileResource_Create_InvalidJSON exercises TcpProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTcpProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := TcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTcpProfileResource_Create_MapError exercises TcpProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTcpProfileResource_Create_MapError(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := TcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTcpProfileResource_Create_MissingID exercises TcpProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestTcpProfileResource_Create_MissingID(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := TcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestTcpProfileResource_Create_LocationFallback exercises TcpProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestTcpProfileResource_Create_LocationFallback(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := TcpProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestTcpProfileResource_Read_Happy exercises TcpProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestTcpProfileResource_Read_Happy(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := TcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTcpProfileResource_Read_NilClient exercises TcpProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTcpProfileResource_Read_NilClient(t *testing.T) {
	r := &TcpProfileResource{}
	m := TcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTcpProfileResource_Read_BuildError exercises TcpProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTcpProfileResource_Read_BuildError(t *testing.T) {
	r := &TcpProfileResource{client: newMalformedBaseURLClient(t)}
	m := TcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTcpProfileResource_Read_SendError exercises TcpProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestTcpProfileResource_Read_SendError(t *testing.T) {
	r := &TcpProfileResource{client: newTransportErrorClient(t)}
	m := TcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTcpProfileResource_Read_NotFound exercises TcpProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestTcpProfileResource_Read_NotFound(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 404, "")}
	m := TcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTcpProfileResource_Read_APIError exercises TcpProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTcpProfileResource_Read_APIError(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_tcp_profile")
}

// TestTcpProfileResource_Read_APIErrorReadBody exercises TcpProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTcpProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := TcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTcpProfileResource_Read_InvalidJSON exercises TcpProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTcpProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := TcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTcpProfileResource_Read_MapError exercises TcpProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTcpProfileResource_Read_MapError(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := TcpProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTcpProfileResource_Update_Happy exercises TcpProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestTcpProfileResource_Update_Happy(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := TcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTcpProfileResource_Update_NilClient exercises TcpProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTcpProfileResource_Update_NilClient(t *testing.T) {
	r := &TcpProfileResource{}
	m := TcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTcpProfileResource_Update_BuildError exercises TcpProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTcpProfileResource_Update_BuildError(t *testing.T) {
	r := &TcpProfileResource{client: newMalformedBaseURLClient(t)}
	m := TcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTcpProfileResource_Update_SendError exercises TcpProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestTcpProfileResource_Update_SendError(t *testing.T) {
	r := &TcpProfileResource{client: newTransportErrorClient(t)}
	m := TcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTcpProfileResource_Update_APIError exercises TcpProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTcpProfileResource_Update_APIError(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_tcp_profile")
}

// TestTcpProfileResource_Update_APIErrorReadBody exercises TcpProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTcpProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := TcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTcpProfileResource_Update_InvalidJSON exercises TcpProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTcpProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := TcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTcpProfileResource_Update_MapError exercises TcpProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTcpProfileResource_Update_MapError(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := TcpProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTcpProfileResource_Delete_Happy exercises TcpProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestTcpProfileResource_Delete_Happy(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 204, "")}
	m := TcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTcpProfileResource_Delete_NilClient exercises TcpProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTcpProfileResource_Delete_NilClient(t *testing.T) {
	r := &TcpProfileResource{}
	m := TcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTcpProfileResource_Delete_BuildError exercises TcpProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTcpProfileResource_Delete_BuildError(t *testing.T) {
	r := &TcpProfileResource{client: newMalformedBaseURLClient(t)}
	m := TcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTcpProfileResource_Delete_SendError exercises TcpProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestTcpProfileResource_Delete_SendError(t *testing.T) {
	r := &TcpProfileResource{client: newTransportErrorClient(t)}
	m := TcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTcpProfileResource_Delete_NotFoundSuccess exercises TcpProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestTcpProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 404, "")}
	m := TcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTcpProfileResource_Delete_APIError exercises TcpProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTcpProfileResource_Delete_APIError(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_tcp_profile")
}

// TestTcpProfileResource_Delete_APIErrorReadBody exercises TcpProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTcpProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &TcpProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := TcpProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
