package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestServerResource_Create_Happy exercises ServerResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestServerResource_Create_Happy(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerResource_Create_NilClient exercises ServerResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServerResource_Create_NilClient(t *testing.T) {
	r := &ServerResource{}
	m := ServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestServerResource_Create_BuildError exercises ServerResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestServerResource_Create_BuildError(t *testing.T) {
	r := &ServerResource{client: newMalformedBaseURLClient(t)}
	m := ServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestServerResource_Create_SendError exercises ServerResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestServerResource_Create_SendError(t *testing.T) {
	r := &ServerResource{client: newTransportErrorClient(t)}
	m := ServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestServerResource_Create_APIError exercises ServerResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestServerResource_Create_APIError(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_server")
}

// TestServerResource_Create_APIErrorReadBody exercises ServerResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestServerResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestServerResource_Create_InvalidJSON exercises ServerResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestServerResource_Create_InvalidJSON(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 201, "{{")}
	m := ServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestServerResource_Create_MapError exercises ServerResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestServerResource_Create_MapError(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestServerResource_Create_MissingID exercises ServerResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestServerResource_Create_MissingID(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 201, "{}")}
	m := ServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestServerResource_Create_LocationFallback exercises ServerResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestServerResource_Create_LocationFallback(t *testing.T) {
	r := &ServerResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ServerResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestServerResource_Read_Happy exercises ServerResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestServerResource_Read_Happy(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 200, "{}")}
	m := ServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerResource_Read_NilClient exercises ServerResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServerResource_Read_NilClient(t *testing.T) {
	r := &ServerResource{}
	m := ServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestServerResource_Read_BuildError exercises ServerResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestServerResource_Read_BuildError(t *testing.T) {
	r := &ServerResource{client: newMalformedBaseURLClient(t)}
	m := ServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestServerResource_Read_SendError exercises ServerResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestServerResource_Read_SendError(t *testing.T) {
	r := &ServerResource{client: newTransportErrorClient(t)}
	m := ServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestServerResource_Read_NotFound exercises ServerResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestServerResource_Read_NotFound(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 404, "")}
	m := ServerResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerResource_Read_APIError exercises ServerResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestServerResource_Read_APIError(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_server")
}

// TestServerResource_Read_APIErrorReadBody exercises ServerResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestServerResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestServerResource_Read_InvalidJSON exercises ServerResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestServerResource_Read_InvalidJSON(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 200, "{{")}
	m := ServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestServerResource_Read_MapError exercises ServerResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestServerResource_Read_MapError(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ServerResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestServerResource_Update_Happy exercises ServerResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestServerResource_Update_Happy(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 201, "{}")}
	m := ServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerResource_Update_NilClient exercises ServerResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServerResource_Update_NilClient(t *testing.T) {
	r := &ServerResource{}
	m := ServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestServerResource_Update_BuildError exercises ServerResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestServerResource_Update_BuildError(t *testing.T) {
	r := &ServerResource{client: newMalformedBaseURLClient(t)}
	m := ServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestServerResource_Update_SendError exercises ServerResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestServerResource_Update_SendError(t *testing.T) {
	r := &ServerResource{client: newTransportErrorClient(t)}
	m := ServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestServerResource_Update_APIError exercises ServerResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestServerResource_Update_APIError(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_server")
}

// TestServerResource_Update_APIErrorReadBody exercises ServerResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestServerResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestServerResource_Update_InvalidJSON exercises ServerResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestServerResource_Update_InvalidJSON(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 201, "{{")}
	m := ServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestServerResource_Update_MapError exercises ServerResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestServerResource_Update_MapError(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ServerResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestServerResource_Delete_Happy exercises ServerResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestServerResource_Delete_Happy(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 204, "")}
	m := ServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerResource_Delete_NilClient exercises ServerResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServerResource_Delete_NilClient(t *testing.T) {
	r := &ServerResource{}
	m := ServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestServerResource_Delete_BuildError exercises ServerResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestServerResource_Delete_BuildError(t *testing.T) {
	r := &ServerResource{client: newMalformedBaseURLClient(t)}
	m := ServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestServerResource_Delete_SendError exercises ServerResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestServerResource_Delete_SendError(t *testing.T) {
	r := &ServerResource{client: newTransportErrorClient(t)}
	m := ServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestServerResource_Delete_NotFoundSuccess exercises ServerResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestServerResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 404, "")}
	m := ServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerResource_Delete_APIError exercises ServerResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestServerResource_Delete_APIError(t *testing.T) {
	r := &ServerResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_server")
}

// TestServerResource_Delete_APIErrorReadBody exercises ServerResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestServerResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ServerResource{client: newMockClientReadErrorBody(t, 501)}
	m := ServerResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
