package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPortGroupResource_Create_Happy exercises PortGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestPortGroupResource_Create_Happy(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := PortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortGroupResource_Create_NilClient exercises PortGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortGroupResource_Create_NilClient(t *testing.T) {
	r := &PortGroupResource{}
	m := PortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortGroupResource_Create_BuildError exercises PortGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortGroupResource_Create_BuildError(t *testing.T) {
	r := &PortGroupResource{client: newMalformedBaseURLClient(t)}
	m := PortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortGroupResource_Create_SendError exercises PortGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortGroupResource_Create_SendError(t *testing.T) {
	r := &PortGroupResource{client: newTransportErrorClient(t)}
	m := PortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortGroupResource_Create_APIError exercises PortGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortGroupResource_Create_APIError(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_port_group")
}

// TestPortGroupResource_Create_APIErrorReadBody exercises PortGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &PortGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortGroupResource_Create_InvalidJSON exercises PortGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := PortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortGroupResource_Create_MapError exercises PortGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortGroupResource_Create_MapError(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := PortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortGroupResource_Create_MissingID exercises PortGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestPortGroupResource_Create_MissingID(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := PortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestPortGroupResource_Create_LocationFallback exercises PortGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestPortGroupResource_Create_LocationFallback(t *testing.T) {
	r := &PortGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := PortGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestPortGroupResource_Read_Happy exercises PortGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestPortGroupResource_Read_Happy(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := PortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortGroupResource_Read_NilClient exercises PortGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortGroupResource_Read_NilClient(t *testing.T) {
	r := &PortGroupResource{}
	m := PortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortGroupResource_Read_BuildError exercises PortGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortGroupResource_Read_BuildError(t *testing.T) {
	r := &PortGroupResource{client: newMalformedBaseURLClient(t)}
	m := PortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortGroupResource_Read_SendError exercises PortGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortGroupResource_Read_SendError(t *testing.T) {
	r := &PortGroupResource{client: newTransportErrorClient(t)}
	m := PortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortGroupResource_Read_NotFound exercises PortGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestPortGroupResource_Read_NotFound(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 404, "")}
	m := PortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortGroupResource_Read_APIError exercises PortGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortGroupResource_Read_APIError(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_port_group")
}

// TestPortGroupResource_Read_APIErrorReadBody exercises PortGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &PortGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortGroupResource_Read_InvalidJSON exercises PortGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortGroupResource_Read_MapError exercises PortGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortGroupResource_Read_MapError(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := PortGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortGroupResource_Update_Happy exercises PortGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestPortGroupResource_Update_Happy(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := PortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortGroupResource_Update_NilClient exercises PortGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortGroupResource_Update_NilClient(t *testing.T) {
	r := &PortGroupResource{}
	m := PortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortGroupResource_Update_BuildError exercises PortGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortGroupResource_Update_BuildError(t *testing.T) {
	r := &PortGroupResource{client: newMalformedBaseURLClient(t)}
	m := PortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortGroupResource_Update_SendError exercises PortGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortGroupResource_Update_SendError(t *testing.T) {
	r := &PortGroupResource{client: newTransportErrorClient(t)}
	m := PortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortGroupResource_Update_APIError exercises PortGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortGroupResource_Update_APIError(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_port_group")
}

// TestPortGroupResource_Update_APIErrorReadBody exercises PortGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &PortGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortGroupResource_Update_InvalidJSON exercises PortGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortGroupResource_Update_MapError exercises PortGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortGroupResource_Update_MapError(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := PortGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortGroupResource_Delete_Happy exercises PortGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestPortGroupResource_Delete_Happy(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 204, "")}
	m := PortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortGroupResource_Delete_NilClient exercises PortGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortGroupResource_Delete_NilClient(t *testing.T) {
	r := &PortGroupResource{}
	m := PortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortGroupResource_Delete_BuildError exercises PortGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortGroupResource_Delete_BuildError(t *testing.T) {
	r := &PortGroupResource{client: newMalformedBaseURLClient(t)}
	m := PortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortGroupResource_Delete_SendError exercises PortGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortGroupResource_Delete_SendError(t *testing.T) {
	r := &PortGroupResource{client: newTransportErrorClient(t)}
	m := PortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortGroupResource_Delete_NotFoundSuccess exercises PortGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestPortGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 404, "")}
	m := PortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortGroupResource_Delete_APIError exercises PortGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortGroupResource_Delete_APIError(t *testing.T) {
	r := &PortGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_port_group")
}

// TestPortGroupResource_Delete_APIErrorReadBody exercises PortGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &PortGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
