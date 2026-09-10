package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestIbPathwayResource_Create_Happy exercises IbPathwayResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestIbPathwayResource_Create_Happy(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := IbPathwayResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIbPathwayResource_Create_NilClient exercises IbPathwayResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIbPathwayResource_Create_NilClient(t *testing.T) {
	r := &IbPathwayResource{}
	m := IbPathwayResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIbPathwayResource_Create_BuildError exercises IbPathwayResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIbPathwayResource_Create_BuildError(t *testing.T) {
	r := &IbPathwayResource{client: newMalformedBaseURLClient(t)}
	m := IbPathwayResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIbPathwayResource_Create_SendError exercises IbPathwayResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestIbPathwayResource_Create_SendError(t *testing.T) {
	r := &IbPathwayResource{client: newTransportErrorClient(t)}
	m := IbPathwayResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIbPathwayResource_Create_APIError exercises IbPathwayResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIbPathwayResource_Create_APIError(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := IbPathwayResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_ib_pathway")
}

// TestIbPathwayResource_Create_APIErrorReadBody exercises IbPathwayResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIbPathwayResource_Create_APIErrorReadBody(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientReadErrorBody(t, 501)}
	m := IbPathwayResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestIbPathwayResource_Create_InvalidJSON exercises IbPathwayResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestIbPathwayResource_Create_InvalidJSON(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 201, "{{")}
	m := IbPathwayResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestIbPathwayResource_Create_MapError exercises IbPathwayResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestIbPathwayResource_Create_MapError(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := IbPathwayResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestIbPathwayResource_Create_MissingID exercises IbPathwayResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestIbPathwayResource_Create_MissingID(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 201, "{}")}
	m := IbPathwayResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestIbPathwayResource_Create_LocationFallback exercises IbPathwayResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestIbPathwayResource_Create_LocationFallback(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := IbPathwayResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestIbPathwayResource_Read_Happy exercises IbPathwayResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestIbPathwayResource_Read_Happy(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 200, "{}")}
	m := IbPathwayResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestIbPathwayResource_Read_NilClient exercises IbPathwayResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIbPathwayResource_Read_NilClient(t *testing.T) {
	r := &IbPathwayResource{}
	m := IbPathwayResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIbPathwayResource_Read_BuildError exercises IbPathwayResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIbPathwayResource_Read_BuildError(t *testing.T) {
	r := &IbPathwayResource{client: newMalformedBaseURLClient(t)}
	m := IbPathwayResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIbPathwayResource_Read_SendError exercises IbPathwayResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestIbPathwayResource_Read_SendError(t *testing.T) {
	r := &IbPathwayResource{client: newTransportErrorClient(t)}
	m := IbPathwayResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIbPathwayResource_Read_NotFound exercises IbPathwayResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestIbPathwayResource_Read_NotFound(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 404, "")}
	m := IbPathwayResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestIbPathwayResource_Read_APIError exercises IbPathwayResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIbPathwayResource_Read_APIError(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := IbPathwayResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_ib_pathway")
}

// TestIbPathwayResource_Read_APIErrorReadBody exercises IbPathwayResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIbPathwayResource_Read_APIErrorReadBody(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientReadErrorBody(t, 501)}
	m := IbPathwayResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestIbPathwayResource_Read_InvalidJSON exercises IbPathwayResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestIbPathwayResource_Read_InvalidJSON(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 200, "{{")}
	m := IbPathwayResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestIbPathwayResource_Read_MapError exercises IbPathwayResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestIbPathwayResource_Read_MapError(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := IbPathwayResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestIbPathwayResource_Update_Happy exercises IbPathwayResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestIbPathwayResource_Update_Happy(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 200, "{}")}
	m := IbPathwayResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIbPathwayResource_Update_NilClient exercises IbPathwayResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIbPathwayResource_Update_NilClient(t *testing.T) {
	r := &IbPathwayResource{}
	m := IbPathwayResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIbPathwayResource_Update_BuildError exercises IbPathwayResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIbPathwayResource_Update_BuildError(t *testing.T) {
	r := &IbPathwayResource{client: newMalformedBaseURLClient(t)}
	m := IbPathwayResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIbPathwayResource_Update_SendError exercises IbPathwayResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestIbPathwayResource_Update_SendError(t *testing.T) {
	r := &IbPathwayResource{client: newTransportErrorClient(t)}
	m := IbPathwayResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIbPathwayResource_Update_APIError exercises IbPathwayResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIbPathwayResource_Update_APIError(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := IbPathwayResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_ib_pathway")
}

// TestIbPathwayResource_Update_APIErrorReadBody exercises IbPathwayResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIbPathwayResource_Update_APIErrorReadBody(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientReadErrorBody(t, 501)}
	m := IbPathwayResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestIbPathwayResource_Update_InvalidJSON exercises IbPathwayResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestIbPathwayResource_Update_InvalidJSON(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 200, "{{")}
	m := IbPathwayResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestIbPathwayResource_Update_MapError exercises IbPathwayResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestIbPathwayResource_Update_MapError(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := IbPathwayResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestIbPathwayResource_Delete_Happy exercises IbPathwayResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestIbPathwayResource_Delete_Happy(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 204, "")}
	m := IbPathwayResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIbPathwayResource_Delete_NilClient exercises IbPathwayResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIbPathwayResource_Delete_NilClient(t *testing.T) {
	r := &IbPathwayResource{}
	m := IbPathwayResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIbPathwayResource_Delete_BuildError exercises IbPathwayResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIbPathwayResource_Delete_BuildError(t *testing.T) {
	r := &IbPathwayResource{client: newMalformedBaseURLClient(t)}
	m := IbPathwayResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIbPathwayResource_Delete_SendError exercises IbPathwayResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestIbPathwayResource_Delete_SendError(t *testing.T) {
	r := &IbPathwayResource{client: newTransportErrorClient(t)}
	m := IbPathwayResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIbPathwayResource_Delete_NotFoundSuccess exercises IbPathwayResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestIbPathwayResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 404, "")}
	m := IbPathwayResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIbPathwayResource_Delete_APIError exercises IbPathwayResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIbPathwayResource_Delete_APIError(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := IbPathwayResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_ib_pathway")
}

// TestIbPathwayResource_Delete_APIErrorReadBody exercises IbPathwayResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIbPathwayResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &IbPathwayResource{client: newMockClientReadErrorBody(t, 501)}
	m := IbPathwayResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
