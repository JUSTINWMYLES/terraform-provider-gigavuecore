package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFilterTemplateResource_Create_Happy exercises FilterTemplateResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestFilterTemplateResource_Create_Happy(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFilterTemplateResource_Create_NilClient exercises FilterTemplateResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFilterTemplateResource_Create_NilClient(t *testing.T) {
	r := &FilterTemplateResource{}
	m := FilterTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFilterTemplateResource_Create_BuildError exercises FilterTemplateResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFilterTemplateResource_Create_BuildError(t *testing.T) {
	r := &FilterTemplateResource{client: newMalformedBaseURLClient(t)}
	m := FilterTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFilterTemplateResource_Create_SendError exercises FilterTemplateResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestFilterTemplateResource_Create_SendError(t *testing.T) {
	r := &FilterTemplateResource{client: newTransportErrorClient(t)}
	m := FilterTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFilterTemplateResource_Create_APIError exercises FilterTemplateResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFilterTemplateResource_Create_APIError(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_filter_template")
}

// TestFilterTemplateResource_Create_APIErrorReadBody exercises FilterTemplateResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFilterTemplateResource_Create_APIErrorReadBody(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := FilterTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFilterTemplateResource_Create_InvalidJSON exercises FilterTemplateResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFilterTemplateResource_Create_InvalidJSON(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 201, "{{")}
	m := FilterTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFilterTemplateResource_Create_MapError exercises FilterTemplateResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFilterTemplateResource_Create_MapError(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFilterTemplateResource_Create_MissingID exercises FilterTemplateResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestFilterTemplateResource_Create_MissingID(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 201, "{}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestFilterTemplateResource_Create_LocationFallback exercises FilterTemplateResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestFilterTemplateResource_Create_LocationFallback(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestFilterTemplateResource_Read_Happy exercises FilterTemplateResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestFilterTemplateResource_Read_Happy(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 200, "{}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFilterTemplateResource_Read_NilClient exercises FilterTemplateResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFilterTemplateResource_Read_NilClient(t *testing.T) {
	r := &FilterTemplateResource{}
	m := FilterTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFilterTemplateResource_Read_BuildError exercises FilterTemplateResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFilterTemplateResource_Read_BuildError(t *testing.T) {
	r := &FilterTemplateResource{client: newMalformedBaseURLClient(t)}
	m := FilterTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFilterTemplateResource_Read_SendError exercises FilterTemplateResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestFilterTemplateResource_Read_SendError(t *testing.T) {
	r := &FilterTemplateResource{client: newTransportErrorClient(t)}
	m := FilterTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFilterTemplateResource_Read_NotFound exercises FilterTemplateResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestFilterTemplateResource_Read_NotFound(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 404, "")}
	m := FilterTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFilterTemplateResource_Read_APIError exercises FilterTemplateResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFilterTemplateResource_Read_APIError(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_filter_template")
}

// TestFilterTemplateResource_Read_APIErrorReadBody exercises FilterTemplateResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFilterTemplateResource_Read_APIErrorReadBody(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := FilterTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFilterTemplateResource_Read_InvalidJSON exercises FilterTemplateResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFilterTemplateResource_Read_InvalidJSON(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 200, "{{")}
	m := FilterTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFilterTemplateResource_Read_MapError exercises FilterTemplateResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFilterTemplateResource_Read_MapError(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFilterTemplateResource_Update_Happy exercises FilterTemplateResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestFilterTemplateResource_Update_Happy(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 200, "{}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFilterTemplateResource_Update_NilClient exercises FilterTemplateResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFilterTemplateResource_Update_NilClient(t *testing.T) {
	r := &FilterTemplateResource{}
	m := FilterTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFilterTemplateResource_Update_BuildError exercises FilterTemplateResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFilterTemplateResource_Update_BuildError(t *testing.T) {
	r := &FilterTemplateResource{client: newMalformedBaseURLClient(t)}
	m := FilterTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFilterTemplateResource_Update_SendError exercises FilterTemplateResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestFilterTemplateResource_Update_SendError(t *testing.T) {
	r := &FilterTemplateResource{client: newTransportErrorClient(t)}
	m := FilterTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFilterTemplateResource_Update_APIError exercises FilterTemplateResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFilterTemplateResource_Update_APIError(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_filter_template")
}

// TestFilterTemplateResource_Update_APIErrorReadBody exercises FilterTemplateResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFilterTemplateResource_Update_APIErrorReadBody(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := FilterTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFilterTemplateResource_Update_InvalidJSON exercises FilterTemplateResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFilterTemplateResource_Update_InvalidJSON(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 200, "{{")}
	m := FilterTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFilterTemplateResource_Update_MapError exercises FilterTemplateResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFilterTemplateResource_Update_MapError(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFilterTemplateResource_Delete_Happy exercises FilterTemplateResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestFilterTemplateResource_Delete_Happy(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 204, "")}
	m := FilterTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFilterTemplateResource_Delete_NilClient exercises FilterTemplateResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFilterTemplateResource_Delete_NilClient(t *testing.T) {
	r := &FilterTemplateResource{}
	m := FilterTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFilterTemplateResource_Delete_BuildError exercises FilterTemplateResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFilterTemplateResource_Delete_BuildError(t *testing.T) {
	r := &FilterTemplateResource{client: newMalformedBaseURLClient(t)}
	m := FilterTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFilterTemplateResource_Delete_SendError exercises FilterTemplateResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestFilterTemplateResource_Delete_SendError(t *testing.T) {
	r := &FilterTemplateResource{client: newTransportErrorClient(t)}
	m := FilterTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFilterTemplateResource_Delete_NotFoundSuccess exercises FilterTemplateResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestFilterTemplateResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 404, "")}
	m := FilterTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFilterTemplateResource_Delete_APIError exercises FilterTemplateResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFilterTemplateResource_Delete_APIError(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FilterTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_filter_template")
}

// TestFilterTemplateResource_Delete_APIErrorReadBody exercises FilterTemplateResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFilterTemplateResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &FilterTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := FilterTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
