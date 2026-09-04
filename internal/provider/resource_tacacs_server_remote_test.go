package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestTacacsServerResource_Create_Happy exercises TacacsServerResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestTacacsServerResource_Create_Happy(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 201, "{\"server_address\":\"example-id\"}")}
	m := TacacsServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTacacsServerResource_Create_NilClient exercises TacacsServerResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTacacsServerResource_Create_NilClient(t *testing.T) {
	r := &TacacsServerResource{}
	m := TacacsServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTacacsServerResource_Create_BuildError exercises TacacsServerResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTacacsServerResource_Create_BuildError(t *testing.T) {
	r := &TacacsServerResource{client: newMalformedBaseURLClient(t)}
	m := TacacsServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTacacsServerResource_Create_SendError exercises TacacsServerResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestTacacsServerResource_Create_SendError(t *testing.T) {
	r := &TacacsServerResource{client: newTransportErrorClient(t)}
	m := TacacsServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTacacsServerResource_Create_APIError exercises TacacsServerResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTacacsServerResource_Create_APIError(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TacacsServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_tacacs_server")
}

// TestTacacsServerResource_Create_APIErrorReadBody exercises TacacsServerResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTacacsServerResource_Create_APIErrorReadBody(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := TacacsServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTacacsServerResource_Create_InvalidJSON exercises TacacsServerResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTacacsServerResource_Create_InvalidJSON(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 201, "{{")}
	m := TacacsServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTacacsServerResource_Create_MapError exercises TacacsServerResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTacacsServerResource_Create_MapError(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 201, "{\"server_address\":12345}")}
	m := TacacsServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTacacsServerResource_Create_MissingID exercises TacacsServerResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestTacacsServerResource_Create_MissingID(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 201, "{}")}
	m := TacacsServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestTacacsServerResource_Create_LocationFallback exercises TacacsServerResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestTacacsServerResource_Create_LocationFallback(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := TacacsServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.ServerAddress.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.ServerAddress.ValueString(), "example-id")
	}
}

// TestTacacsServerResource_Read_Happy exercises TacacsServerResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestTacacsServerResource_Read_Happy(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := TacacsServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTacacsServerResource_Read_NilClient exercises TacacsServerResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTacacsServerResource_Read_NilClient(t *testing.T) {
	r := &TacacsServerResource{}
	m := TacacsServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTacacsServerResource_Read_BuildError exercises TacacsServerResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTacacsServerResource_Read_BuildError(t *testing.T) {
	r := &TacacsServerResource{client: newMalformedBaseURLClient(t)}
	m := TacacsServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTacacsServerResource_Read_SendError exercises TacacsServerResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestTacacsServerResource_Read_SendError(t *testing.T) {
	r := &TacacsServerResource{client: newTransportErrorClient(t)}
	m := TacacsServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTacacsServerResource_Read_NotFound exercises TacacsServerResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestTacacsServerResource_Read_NotFound(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 404, "")}
	m := TacacsServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTacacsServerResource_Read_APIError exercises TacacsServerResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTacacsServerResource_Read_APIError(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TacacsServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_tacacs_server")
}

// TestTacacsServerResource_Read_APIErrorReadBody exercises TacacsServerResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTacacsServerResource_Read_APIErrorReadBody(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := TacacsServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTacacsServerResource_Read_InvalidJSON exercises TacacsServerResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTacacsServerResource_Read_InvalidJSON(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := TacacsServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTacacsServerResource_Read_MapError exercises TacacsServerResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTacacsServerResource_Read_MapError(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 200, "{\"server_address\":12345}")}
	m := TacacsServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTacacsServerResource_Update_Happy exercises TacacsServerResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestTacacsServerResource_Update_Happy(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := TacacsServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTacacsServerResource_Update_NilClient exercises TacacsServerResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTacacsServerResource_Update_NilClient(t *testing.T) {
	r := &TacacsServerResource{}
	m := TacacsServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTacacsServerResource_Update_BuildError exercises TacacsServerResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTacacsServerResource_Update_BuildError(t *testing.T) {
	r := &TacacsServerResource{client: newMalformedBaseURLClient(t)}
	m := TacacsServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTacacsServerResource_Update_SendError exercises TacacsServerResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestTacacsServerResource_Update_SendError(t *testing.T) {
	r := &TacacsServerResource{client: newTransportErrorClient(t)}
	m := TacacsServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTacacsServerResource_Update_APIError exercises TacacsServerResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTacacsServerResource_Update_APIError(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TacacsServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_tacacs_server")
}

// TestTacacsServerResource_Update_APIErrorReadBody exercises TacacsServerResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTacacsServerResource_Update_APIErrorReadBody(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := TacacsServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTacacsServerResource_Update_InvalidJSON exercises TacacsServerResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTacacsServerResource_Update_InvalidJSON(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := TacacsServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTacacsServerResource_Update_MapError exercises TacacsServerResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTacacsServerResource_Update_MapError(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 200, "{\"server_address\":12345}")}
	m := TacacsServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTacacsServerResource_Delete_Happy exercises TacacsServerResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestTacacsServerResource_Delete_Happy(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 204, "")}
	m := TacacsServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTacacsServerResource_Delete_NilClient exercises TacacsServerResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTacacsServerResource_Delete_NilClient(t *testing.T) {
	r := &TacacsServerResource{}
	m := TacacsServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTacacsServerResource_Delete_BuildError exercises TacacsServerResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTacacsServerResource_Delete_BuildError(t *testing.T) {
	r := &TacacsServerResource{client: newMalformedBaseURLClient(t)}
	m := TacacsServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTacacsServerResource_Delete_SendError exercises TacacsServerResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestTacacsServerResource_Delete_SendError(t *testing.T) {
	r := &TacacsServerResource{client: newTransportErrorClient(t)}
	m := TacacsServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTacacsServerResource_Delete_NotFoundSuccess exercises TacacsServerResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestTacacsServerResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 404, "")}
	m := TacacsServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTacacsServerResource_Delete_APIError exercises TacacsServerResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTacacsServerResource_Delete_APIError(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TacacsServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_tacacs_server")
}

// TestTacacsServerResource_Delete_APIErrorReadBody exercises TacacsServerResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTacacsServerResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &TacacsServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := TacacsServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
