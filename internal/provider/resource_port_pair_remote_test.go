package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPortPairResource_Create_Happy exercises PortPairResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestPortPairResource_Create_Happy(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := PortPairResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortPairResource_Create_NilClient exercises PortPairResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortPairResource_Create_NilClient(t *testing.T) {
	r := &PortPairResource{}
	m := PortPairResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortPairResource_Create_BuildError exercises PortPairResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortPairResource_Create_BuildError(t *testing.T) {
	r := &PortPairResource{client: newMalformedBaseURLClient(t)}
	m := PortPairResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortPairResource_Create_SendError exercises PortPairResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortPairResource_Create_SendError(t *testing.T) {
	r := &PortPairResource{client: newTransportErrorClient(t)}
	m := PortPairResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortPairResource_Create_APIError exercises PortPairResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortPairResource_Create_APIError(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortPairResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_port_pair")
}

// TestPortPairResource_Create_APIErrorReadBody exercises PortPairResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortPairResource_Create_APIErrorReadBody(t *testing.T) {
	r := &PortPairResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortPairResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortPairResource_Create_InvalidJSON exercises PortPairResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortPairResource_Create_InvalidJSON(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 201, "{{")}
	m := PortPairResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortPairResource_Create_MapError exercises PortPairResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortPairResource_Create_MapError(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := PortPairResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortPairResource_Create_MissingID exercises PortPairResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestPortPairResource_Create_MissingID(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 201, "{}")}
	m := PortPairResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestPortPairResource_Create_LocationFallback exercises PortPairResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestPortPairResource_Create_LocationFallback(t *testing.T) {
	r := &PortPairResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := PortPairResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestPortPairResource_Read_Happy exercises PortPairResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestPortPairResource_Read_Happy(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 200, "{}")}
	m := PortPairResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortPairResource_Read_NilClient exercises PortPairResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortPairResource_Read_NilClient(t *testing.T) {
	r := &PortPairResource{}
	m := PortPairResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortPairResource_Read_BuildError exercises PortPairResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortPairResource_Read_BuildError(t *testing.T) {
	r := &PortPairResource{client: newMalformedBaseURLClient(t)}
	m := PortPairResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortPairResource_Read_SendError exercises PortPairResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortPairResource_Read_SendError(t *testing.T) {
	r := &PortPairResource{client: newTransportErrorClient(t)}
	m := PortPairResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortPairResource_Read_NotFound exercises PortPairResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestPortPairResource_Read_NotFound(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 404, "")}
	m := PortPairResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortPairResource_Read_APIError exercises PortPairResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortPairResource_Read_APIError(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortPairResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_port_pair")
}

// TestPortPairResource_Read_APIErrorReadBody exercises PortPairResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortPairResource_Read_APIErrorReadBody(t *testing.T) {
	r := &PortPairResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortPairResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortPairResource_Read_InvalidJSON exercises PortPairResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortPairResource_Read_InvalidJSON(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortPairResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortPairResource_Read_MapError exercises PortPairResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortPairResource_Read_MapError(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := PortPairResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortPairResource_Update_Happy exercises PortPairResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestPortPairResource_Update_Happy(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 200, "{}")}
	m := PortPairResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortPairResource_Update_NilClient exercises PortPairResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortPairResource_Update_NilClient(t *testing.T) {
	r := &PortPairResource{}
	m := PortPairResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortPairResource_Update_BuildError exercises PortPairResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortPairResource_Update_BuildError(t *testing.T) {
	r := &PortPairResource{client: newMalformedBaseURLClient(t)}
	m := PortPairResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortPairResource_Update_SendError exercises PortPairResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortPairResource_Update_SendError(t *testing.T) {
	r := &PortPairResource{client: newTransportErrorClient(t)}
	m := PortPairResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortPairResource_Update_APIError exercises PortPairResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortPairResource_Update_APIError(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortPairResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_port_pair")
}

// TestPortPairResource_Update_APIErrorReadBody exercises PortPairResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortPairResource_Update_APIErrorReadBody(t *testing.T) {
	r := &PortPairResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortPairResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortPairResource_Update_InvalidJSON exercises PortPairResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortPairResource_Update_InvalidJSON(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortPairResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortPairResource_Update_MapError exercises PortPairResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortPairResource_Update_MapError(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := PortPairResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortPairResource_Delete_Happy exercises PortPairResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestPortPairResource_Delete_Happy(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 204, "")}
	m := PortPairResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortPairResource_Delete_NilClient exercises PortPairResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortPairResource_Delete_NilClient(t *testing.T) {
	r := &PortPairResource{}
	m := PortPairResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortPairResource_Delete_BuildError exercises PortPairResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortPairResource_Delete_BuildError(t *testing.T) {
	r := &PortPairResource{client: newMalformedBaseURLClient(t)}
	m := PortPairResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortPairResource_Delete_SendError exercises PortPairResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortPairResource_Delete_SendError(t *testing.T) {
	r := &PortPairResource{client: newTransportErrorClient(t)}
	m := PortPairResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortPairResource_Delete_NotFoundSuccess exercises PortPairResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestPortPairResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 404, "")}
	m := PortPairResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortPairResource_Delete_APIError exercises PortPairResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortPairResource_Delete_APIError(t *testing.T) {
	r := &PortPairResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortPairResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_port_pair")
}

// TestPortPairResource_Delete_APIErrorReadBody exercises PortPairResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortPairResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &PortPairResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortPairResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
