package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestIcapResource_Create_Happy exercises IcapResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestIcapResource_Create_Happy(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := IcapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIcapResource_Create_NilClient exercises IcapResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIcapResource_Create_NilClient(t *testing.T) {
	r := &IcapResource{}
	m := IcapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIcapResource_Create_BuildError exercises IcapResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIcapResource_Create_BuildError(t *testing.T) {
	r := &IcapResource{client: newMalformedBaseURLClient(t)}
	m := IcapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIcapResource_Create_SendError exercises IcapResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestIcapResource_Create_SendError(t *testing.T) {
	r := &IcapResource{client: newTransportErrorClient(t)}
	m := IcapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIcapResource_Create_APIError exercises IcapResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIcapResource_Create_APIError(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := IcapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_icap")
}

// TestIcapResource_Create_APIErrorReadBody exercises IcapResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIcapResource_Create_APIErrorReadBody(t *testing.T) {
	r := &IcapResource{client: newMockClientReadErrorBody(t, 501)}
	m := IcapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestIcapResource_Create_InvalidJSON exercises IcapResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestIcapResource_Create_InvalidJSON(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 201, "{{")}
	m := IcapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestIcapResource_Create_MapError exercises IcapResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestIcapResource_Create_MapError(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := IcapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestIcapResource_Create_MissingID exercises IcapResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestIcapResource_Create_MissingID(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 201, "{}")}
	m := IcapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestIcapResource_Create_LocationFallback exercises IcapResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestIcapResource_Create_LocationFallback(t *testing.T) {
	r := &IcapResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := IcapResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestIcapResource_Read_Happy exercises IcapResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestIcapResource_Read_Happy(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 200, "{}")}
	m := IcapResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestIcapResource_Read_NilClient exercises IcapResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIcapResource_Read_NilClient(t *testing.T) {
	r := &IcapResource{}
	m := IcapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIcapResource_Read_BuildError exercises IcapResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIcapResource_Read_BuildError(t *testing.T) {
	r := &IcapResource{client: newMalformedBaseURLClient(t)}
	m := IcapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIcapResource_Read_SendError exercises IcapResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestIcapResource_Read_SendError(t *testing.T) {
	r := &IcapResource{client: newTransportErrorClient(t)}
	m := IcapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIcapResource_Read_NotFound exercises IcapResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestIcapResource_Read_NotFound(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 404, "")}
	m := IcapResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestIcapResource_Read_APIError exercises IcapResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIcapResource_Read_APIError(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := IcapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_icap")
}

// TestIcapResource_Read_APIErrorReadBody exercises IcapResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIcapResource_Read_APIErrorReadBody(t *testing.T) {
	r := &IcapResource{client: newMockClientReadErrorBody(t, 501)}
	m := IcapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestIcapResource_Read_InvalidJSON exercises IcapResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestIcapResource_Read_InvalidJSON(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 200, "{{")}
	m := IcapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestIcapResource_Read_MapError exercises IcapResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestIcapResource_Read_MapError(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := IcapResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestIcapResource_Update_Happy exercises IcapResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestIcapResource_Update_Happy(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 200, "{}")}
	m := IcapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIcapResource_Update_NilClient exercises IcapResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIcapResource_Update_NilClient(t *testing.T) {
	r := &IcapResource{}
	m := IcapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIcapResource_Update_BuildError exercises IcapResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIcapResource_Update_BuildError(t *testing.T) {
	r := &IcapResource{client: newMalformedBaseURLClient(t)}
	m := IcapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIcapResource_Update_SendError exercises IcapResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestIcapResource_Update_SendError(t *testing.T) {
	r := &IcapResource{client: newTransportErrorClient(t)}
	m := IcapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIcapResource_Update_APIError exercises IcapResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIcapResource_Update_APIError(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := IcapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_icap")
}

// TestIcapResource_Update_APIErrorReadBody exercises IcapResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIcapResource_Update_APIErrorReadBody(t *testing.T) {
	r := &IcapResource{client: newMockClientReadErrorBody(t, 501)}
	m := IcapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestIcapResource_Update_InvalidJSON exercises IcapResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestIcapResource_Update_InvalidJSON(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 200, "{{")}
	m := IcapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestIcapResource_Update_MapError exercises IcapResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestIcapResource_Update_MapError(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := IcapResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestIcapResource_Delete_Happy exercises IcapResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestIcapResource_Delete_Happy(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 204, "")}
	m := IcapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIcapResource_Delete_NilClient exercises IcapResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIcapResource_Delete_NilClient(t *testing.T) {
	r := &IcapResource{}
	m := IcapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIcapResource_Delete_BuildError exercises IcapResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIcapResource_Delete_BuildError(t *testing.T) {
	r := &IcapResource{client: newMalformedBaseURLClient(t)}
	m := IcapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIcapResource_Delete_SendError exercises IcapResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestIcapResource_Delete_SendError(t *testing.T) {
	r := &IcapResource{client: newTransportErrorClient(t)}
	m := IcapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIcapResource_Delete_NotFoundSuccess exercises IcapResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestIcapResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 404, "")}
	m := IcapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIcapResource_Delete_APIError exercises IcapResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIcapResource_Delete_APIError(t *testing.T) {
	r := &IcapResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := IcapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_icap")
}

// TestIcapResource_Delete_APIErrorReadBody exercises IcapResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIcapResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &IcapResource{client: newMockClientReadErrorBody(t, 501)}
	m := IcapResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
