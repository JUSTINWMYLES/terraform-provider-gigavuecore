package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestApplicationProfileResource_Create_Happy exercises ApplicationProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestApplicationProfileResource_Create_Happy(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestApplicationProfileResource_Create_NilClient exercises ApplicationProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestApplicationProfileResource_Create_NilClient(t *testing.T) {
	r := &ApplicationProfileResource{}
	m := ApplicationProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestApplicationProfileResource_Create_BuildError exercises ApplicationProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestApplicationProfileResource_Create_BuildError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMalformedBaseURLClient(t)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestApplicationProfileResource_Create_SendError exercises ApplicationProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestApplicationProfileResource_Create_SendError(t *testing.T) {
	r := &ApplicationProfileResource{client: newTransportErrorClient(t)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestApplicationProfileResource_Create_APIError exercises ApplicationProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestApplicationProfileResource_Create_APIError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_application_profile")
}

// TestApplicationProfileResource_Create_APIErrorReadBody exercises ApplicationProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestApplicationProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestApplicationProfileResource_Create_InvalidJSON exercises ApplicationProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestApplicationProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestApplicationProfileResource_Create_MapError exercises ApplicationProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestApplicationProfileResource_Create_MapError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestApplicationProfileResource_Create_MissingID exercises ApplicationProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestApplicationProfileResource_Create_MissingID(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestApplicationProfileResource_Create_LocationFallback exercises ApplicationProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestApplicationProfileResource_Create_LocationFallback(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestApplicationProfileResource_Read_Happy exercises ApplicationProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestApplicationProfileResource_Read_Happy(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestApplicationProfileResource_Read_NilClient exercises ApplicationProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestApplicationProfileResource_Read_NilClient(t *testing.T) {
	r := &ApplicationProfileResource{}
	m := ApplicationProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestApplicationProfileResource_Read_BuildError exercises ApplicationProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestApplicationProfileResource_Read_BuildError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMalformedBaseURLClient(t)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestApplicationProfileResource_Read_SendError exercises ApplicationProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestApplicationProfileResource_Read_SendError(t *testing.T) {
	r := &ApplicationProfileResource{client: newTransportErrorClient(t)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestApplicationProfileResource_Read_NotFound exercises ApplicationProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestApplicationProfileResource_Read_NotFound(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 404, "")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestApplicationProfileResource_Read_APIError exercises ApplicationProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestApplicationProfileResource_Read_APIError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_application_profile")
}

// TestApplicationProfileResource_Read_APIErrorReadBody exercises ApplicationProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestApplicationProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestApplicationProfileResource_Read_InvalidJSON exercises ApplicationProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestApplicationProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestApplicationProfileResource_Read_MapError exercises ApplicationProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestApplicationProfileResource_Read_MapError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestApplicationProfileResource_Update_Happy exercises ApplicationProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestApplicationProfileResource_Update_Happy(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestApplicationProfileResource_Update_NilClient exercises ApplicationProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestApplicationProfileResource_Update_NilClient(t *testing.T) {
	r := &ApplicationProfileResource{}
	m := ApplicationProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestApplicationProfileResource_Update_BuildError exercises ApplicationProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestApplicationProfileResource_Update_BuildError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMalformedBaseURLClient(t)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestApplicationProfileResource_Update_SendError exercises ApplicationProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestApplicationProfileResource_Update_SendError(t *testing.T) {
	r := &ApplicationProfileResource{client: newTransportErrorClient(t)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestApplicationProfileResource_Update_APIError exercises ApplicationProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestApplicationProfileResource_Update_APIError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_application_profile")
}

// TestApplicationProfileResource_Update_APIErrorReadBody exercises ApplicationProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestApplicationProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestApplicationProfileResource_Update_InvalidJSON exercises ApplicationProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestApplicationProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestApplicationProfileResource_Update_MapError exercises ApplicationProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestApplicationProfileResource_Update_MapError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestApplicationProfileResource_Delete_Happy exercises ApplicationProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestApplicationProfileResource_Delete_Happy(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 204, "")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestApplicationProfileResource_Delete_NilClient exercises ApplicationProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestApplicationProfileResource_Delete_NilClient(t *testing.T) {
	r := &ApplicationProfileResource{}
	m := ApplicationProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestApplicationProfileResource_Delete_BuildError exercises ApplicationProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestApplicationProfileResource_Delete_BuildError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMalformedBaseURLClient(t)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestApplicationProfileResource_Delete_SendError exercises ApplicationProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestApplicationProfileResource_Delete_SendError(t *testing.T) {
	r := &ApplicationProfileResource{client: newTransportErrorClient(t)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestApplicationProfileResource_Delete_NotFoundSuccess exercises ApplicationProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestApplicationProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 404, "")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestApplicationProfileResource_Delete_APIError exercises ApplicationProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestApplicationProfileResource_Delete_APIError(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ApplicationProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_application_profile")
}

// TestApplicationProfileResource_Delete_APIErrorReadBody exercises ApplicationProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestApplicationProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ApplicationProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ApplicationProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
