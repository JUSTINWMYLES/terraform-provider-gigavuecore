package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMapTemplateResource_Create_Happy exercises MapTemplateResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestMapTemplateResource_Create_Happy(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := MapTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapTemplateResource_Create_NilClient exercises MapTemplateResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapTemplateResource_Create_NilClient(t *testing.T) {
	r := &MapTemplateResource{}
	m := MapTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapTemplateResource_Create_BuildError exercises MapTemplateResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapTemplateResource_Create_BuildError(t *testing.T) {
	r := &MapTemplateResource{client: newMalformedBaseURLClient(t)}
	m := MapTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapTemplateResource_Create_SendError exercises MapTemplateResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapTemplateResource_Create_SendError(t *testing.T) {
	r := &MapTemplateResource{client: newTransportErrorClient(t)}
	m := MapTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapTemplateResource_Create_APIError exercises MapTemplateResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapTemplateResource_Create_APIError(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_map_template")
}

// TestMapTemplateResource_Create_APIErrorReadBody exercises MapTemplateResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapTemplateResource_Create_APIErrorReadBody(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMapTemplateResource_Create_InvalidJSON exercises MapTemplateResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMapTemplateResource_Create_InvalidJSON(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 201, "{{")}
	m := MapTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMapTemplateResource_Create_MapError exercises MapTemplateResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMapTemplateResource_Create_MapError(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := MapTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMapTemplateResource_Create_MissingID exercises MapTemplateResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestMapTemplateResource_Create_MissingID(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 201, "{}")}
	m := MapTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestMapTemplateResource_Create_LocationFallback exercises MapTemplateResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestMapTemplateResource_Create_LocationFallback(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := MapTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestMapTemplateResource_Read_Happy exercises MapTemplateResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestMapTemplateResource_Read_Happy(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 200, "{}")}
	m := MapTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapTemplateResource_Read_NilClient exercises MapTemplateResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapTemplateResource_Read_NilClient(t *testing.T) {
	r := &MapTemplateResource{}
	m := MapTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapTemplateResource_Read_BuildError exercises MapTemplateResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapTemplateResource_Read_BuildError(t *testing.T) {
	r := &MapTemplateResource{client: newMalformedBaseURLClient(t)}
	m := MapTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapTemplateResource_Read_SendError exercises MapTemplateResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapTemplateResource_Read_SendError(t *testing.T) {
	r := &MapTemplateResource{client: newTransportErrorClient(t)}
	m := MapTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapTemplateResource_Read_NotFound exercises MapTemplateResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestMapTemplateResource_Read_NotFound(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 404, "")}
	m := MapTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapTemplateResource_Read_APIError exercises MapTemplateResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapTemplateResource_Read_APIError(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_map_template")
}

// TestMapTemplateResource_Read_APIErrorReadBody exercises MapTemplateResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapTemplateResource_Read_APIErrorReadBody(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMapTemplateResource_Read_InvalidJSON exercises MapTemplateResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMapTemplateResource_Read_InvalidJSON(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 200, "{{")}
	m := MapTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMapTemplateResource_Read_MapError exercises MapTemplateResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMapTemplateResource_Read_MapError(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := MapTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMapTemplateResource_Update_Happy exercises MapTemplateResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestMapTemplateResource_Update_Happy(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 200, "{}")}
	m := MapTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapTemplateResource_Update_NilClient exercises MapTemplateResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapTemplateResource_Update_NilClient(t *testing.T) {
	r := &MapTemplateResource{}
	m := MapTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapTemplateResource_Update_BuildError exercises MapTemplateResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapTemplateResource_Update_BuildError(t *testing.T) {
	r := &MapTemplateResource{client: newMalformedBaseURLClient(t)}
	m := MapTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapTemplateResource_Update_SendError exercises MapTemplateResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapTemplateResource_Update_SendError(t *testing.T) {
	r := &MapTemplateResource{client: newTransportErrorClient(t)}
	m := MapTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapTemplateResource_Update_APIError exercises MapTemplateResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapTemplateResource_Update_APIError(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_map_template")
}

// TestMapTemplateResource_Update_APIErrorReadBody exercises MapTemplateResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapTemplateResource_Update_APIErrorReadBody(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMapTemplateResource_Update_InvalidJSON exercises MapTemplateResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMapTemplateResource_Update_InvalidJSON(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 200, "{{")}
	m := MapTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMapTemplateResource_Update_MapError exercises MapTemplateResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMapTemplateResource_Update_MapError(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := MapTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMapTemplateResource_Delete_Happy exercises MapTemplateResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestMapTemplateResource_Delete_Happy(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 204, "")}
	m := MapTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapTemplateResource_Delete_NilClient exercises MapTemplateResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapTemplateResource_Delete_NilClient(t *testing.T) {
	r := &MapTemplateResource{}
	m := MapTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMapTemplateResource_Delete_BuildError exercises MapTemplateResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMapTemplateResource_Delete_BuildError(t *testing.T) {
	r := &MapTemplateResource{client: newMalformedBaseURLClient(t)}
	m := MapTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMapTemplateResource_Delete_SendError exercises MapTemplateResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestMapTemplateResource_Delete_SendError(t *testing.T) {
	r := &MapTemplateResource{client: newTransportErrorClient(t)}
	m := MapTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMapTemplateResource_Delete_NotFoundSuccess exercises MapTemplateResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestMapTemplateResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 404, "")}
	m := MapTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMapTemplateResource_Delete_APIError exercises MapTemplateResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMapTemplateResource_Delete_APIError(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MapTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_map_template")
}

// TestMapTemplateResource_Delete_APIErrorReadBody exercises MapTemplateResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMapTemplateResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &MapTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := MapTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
