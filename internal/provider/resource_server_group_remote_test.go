package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestServerGroupResource_Create_Happy exercises ServerGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestServerGroupResource_Create_Happy(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ServerGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerGroupResource_Create_NilClient exercises ServerGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServerGroupResource_Create_NilClient(t *testing.T) {
	r := &ServerGroupResource{}
	m := ServerGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestServerGroupResource_Create_BuildError exercises ServerGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestServerGroupResource_Create_BuildError(t *testing.T) {
	r := &ServerGroupResource{client: newMalformedBaseURLClient(t)}
	m := ServerGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestServerGroupResource_Create_SendError exercises ServerGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestServerGroupResource_Create_SendError(t *testing.T) {
	r := &ServerGroupResource{client: newTransportErrorClient(t)}
	m := ServerGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestServerGroupResource_Create_APIError exercises ServerGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestServerGroupResource_Create_APIError(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ServerGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_server_group")
}

// TestServerGroupResource_Create_APIErrorReadBody exercises ServerGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestServerGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ServerGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestServerGroupResource_Create_InvalidJSON exercises ServerGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestServerGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := ServerGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestServerGroupResource_Create_MapError exercises ServerGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestServerGroupResource_Create_MapError(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ServerGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestServerGroupResource_Create_MissingID exercises ServerGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestServerGroupResource_Create_MissingID(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := ServerGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestServerGroupResource_Create_LocationFallback exercises ServerGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestServerGroupResource_Create_LocationFallback(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ServerGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestServerGroupResource_Read_Happy exercises ServerGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestServerGroupResource_Read_Happy(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := ServerGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerGroupResource_Read_NilClient exercises ServerGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServerGroupResource_Read_NilClient(t *testing.T) {
	r := &ServerGroupResource{}
	m := ServerGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestServerGroupResource_Read_BuildError exercises ServerGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestServerGroupResource_Read_BuildError(t *testing.T) {
	r := &ServerGroupResource{client: newMalformedBaseURLClient(t)}
	m := ServerGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestServerGroupResource_Read_SendError exercises ServerGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestServerGroupResource_Read_SendError(t *testing.T) {
	r := &ServerGroupResource{client: newTransportErrorClient(t)}
	m := ServerGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestServerGroupResource_Read_NotFound exercises ServerGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestServerGroupResource_Read_NotFound(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 404, "")}
	m := ServerGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerGroupResource_Read_APIError exercises ServerGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestServerGroupResource_Read_APIError(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ServerGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_server_group")
}

// TestServerGroupResource_Read_APIErrorReadBody exercises ServerGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestServerGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ServerGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestServerGroupResource_Read_InvalidJSON exercises ServerGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestServerGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := ServerGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestServerGroupResource_Read_MapError exercises ServerGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestServerGroupResource_Read_MapError(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ServerGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestServerGroupResource_Update_Happy exercises ServerGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestServerGroupResource_Update_Happy(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := ServerGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerGroupResource_Update_NilClient exercises ServerGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServerGroupResource_Update_NilClient(t *testing.T) {
	r := &ServerGroupResource{}
	m := ServerGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestServerGroupResource_Update_BuildError exercises ServerGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestServerGroupResource_Update_BuildError(t *testing.T) {
	r := &ServerGroupResource{client: newMalformedBaseURLClient(t)}
	m := ServerGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestServerGroupResource_Update_SendError exercises ServerGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestServerGroupResource_Update_SendError(t *testing.T) {
	r := &ServerGroupResource{client: newTransportErrorClient(t)}
	m := ServerGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestServerGroupResource_Update_APIError exercises ServerGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestServerGroupResource_Update_APIError(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ServerGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_server_group")
}

// TestServerGroupResource_Update_APIErrorReadBody exercises ServerGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestServerGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ServerGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestServerGroupResource_Update_InvalidJSON exercises ServerGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestServerGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := ServerGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestServerGroupResource_Update_MapError exercises ServerGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestServerGroupResource_Update_MapError(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ServerGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestServerGroupResource_Delete_Happy exercises ServerGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestServerGroupResource_Delete_Happy(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 204, "")}
	m := ServerGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerGroupResource_Delete_NilClient exercises ServerGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServerGroupResource_Delete_NilClient(t *testing.T) {
	r := &ServerGroupResource{}
	m := ServerGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestServerGroupResource_Delete_BuildError exercises ServerGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestServerGroupResource_Delete_BuildError(t *testing.T) {
	r := &ServerGroupResource{client: newMalformedBaseURLClient(t)}
	m := ServerGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestServerGroupResource_Delete_SendError exercises ServerGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestServerGroupResource_Delete_SendError(t *testing.T) {
	r := &ServerGroupResource{client: newTransportErrorClient(t)}
	m := ServerGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestServerGroupResource_Delete_NotFoundSuccess exercises ServerGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestServerGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 404, "")}
	m := ServerGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestServerGroupResource_Delete_APIError exercises ServerGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestServerGroupResource_Delete_APIError(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ServerGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_server_group")
}

// TestServerGroupResource_Delete_APIErrorReadBody exercises ServerGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestServerGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ServerGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ServerGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
