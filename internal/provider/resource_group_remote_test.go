package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestGroupResource_Create_Happy exercises GroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestGroupResource_Create_Happy(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 201, "{\"name\":\"example-id\"}")}
	m := GroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGroupResource_Create_NilClient exercises GroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGroupResource_Create_NilClient(t *testing.T) {
	r := &GroupResource{}
	m := GroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGroupResource_Create_BuildError exercises GroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGroupResource_Create_BuildError(t *testing.T) {
	r := &GroupResource{client: newMalformedBaseURLClient(t)}
	m := GroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGroupResource_Create_SendError exercises GroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestGroupResource_Create_SendError(t *testing.T) {
	r := &GroupResource{client: newTransportErrorClient(t)}
	m := GroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGroupResource_Create_APIError exercises GroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGroupResource_Create_APIError(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_group")
}

// TestGroupResource_Create_APIErrorReadBody exercises GroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &GroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := GroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGroupResource_Create_InvalidJSON exercises GroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := GroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGroupResource_Create_MapError exercises GroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGroupResource_Create_MapError(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 201, "{\"name\":12345}")}
	m := GroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGroupResource_Create_MissingID exercises GroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestGroupResource_Create_MissingID(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := GroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestGroupResource_Create_LocationFallback exercises GroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestGroupResource_Create_LocationFallback(t *testing.T) {
	r := &GroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := GroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Name.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Name.ValueString(), "example-id")
	}
}

// TestGroupResource_Read_Happy exercises GroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestGroupResource_Read_Happy(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := GroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGroupResource_Read_NilClient exercises GroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGroupResource_Read_NilClient(t *testing.T) {
	r := &GroupResource{}
	m := GroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGroupResource_Read_BuildError exercises GroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGroupResource_Read_BuildError(t *testing.T) {
	r := &GroupResource{client: newMalformedBaseURLClient(t)}
	m := GroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGroupResource_Read_SendError exercises GroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGroupResource_Read_SendError(t *testing.T) {
	r := &GroupResource{client: newTransportErrorClient(t)}
	m := GroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGroupResource_Read_NotFound exercises GroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestGroupResource_Read_NotFound(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 404, "")}
	m := GroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGroupResource_Read_APIError exercises GroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGroupResource_Read_APIError(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_group")
}

// TestGroupResource_Read_APIErrorReadBody exercises GroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &GroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := GroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGroupResource_Read_InvalidJSON exercises GroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := GroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGroupResource_Read_MapError exercises GroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGroupResource_Read_MapError(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 200, "{\"name\":12345}")}
	m := GroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGroupResource_Update_Happy exercises GroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestGroupResource_Update_Happy(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := GroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGroupResource_Update_NilClient exercises GroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGroupResource_Update_NilClient(t *testing.T) {
	r := &GroupResource{}
	m := GroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGroupResource_Update_BuildError exercises GroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGroupResource_Update_BuildError(t *testing.T) {
	r := &GroupResource{client: newMalformedBaseURLClient(t)}
	m := GroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGroupResource_Update_SendError exercises GroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestGroupResource_Update_SendError(t *testing.T) {
	r := &GroupResource{client: newTransportErrorClient(t)}
	m := GroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGroupResource_Update_APIError exercises GroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGroupResource_Update_APIError(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_group")
}

// TestGroupResource_Update_APIErrorReadBody exercises GroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &GroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := GroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGroupResource_Update_InvalidJSON exercises GroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := GroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGroupResource_Update_MapError exercises GroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGroupResource_Update_MapError(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 200, "{\"name\":12345}")}
	m := GroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGroupResource_Delete_Happy exercises GroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestGroupResource_Delete_Happy(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 204, "")}
	m := GroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGroupResource_Delete_NilClient exercises GroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGroupResource_Delete_NilClient(t *testing.T) {
	r := &GroupResource{}
	m := GroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGroupResource_Delete_BuildError exercises GroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGroupResource_Delete_BuildError(t *testing.T) {
	r := &GroupResource{client: newMalformedBaseURLClient(t)}
	m := GroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGroupResource_Delete_SendError exercises GroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestGroupResource_Delete_SendError(t *testing.T) {
	r := &GroupResource{client: newTransportErrorClient(t)}
	m := GroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGroupResource_Delete_NotFoundSuccess exercises GroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 404, "")}
	m := GroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGroupResource_Delete_APIError exercises GroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGroupResource_Delete_APIError(t *testing.T) {
	r := &GroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_group")
}

// TestGroupResource_Delete_APIErrorReadBody exercises GroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &GroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := GroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
