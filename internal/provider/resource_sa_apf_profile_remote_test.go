package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSaApfProfileResource_Create_Happy exercises SaApfProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestSaApfProfileResource_Create_Happy(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSaApfProfileResource_Create_NilClient exercises SaApfProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSaApfProfileResource_Create_NilClient(t *testing.T) {
	r := &SaApfProfileResource{}
	m := SaApfProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSaApfProfileResource_Create_BuildError exercises SaApfProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSaApfProfileResource_Create_BuildError(t *testing.T) {
	r := &SaApfProfileResource{client: newMalformedBaseURLClient(t)}
	m := SaApfProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSaApfProfileResource_Create_SendError exercises SaApfProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestSaApfProfileResource_Create_SendError(t *testing.T) {
	r := &SaApfProfileResource{client: newTransportErrorClient(t)}
	m := SaApfProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSaApfProfileResource_Create_APIError exercises SaApfProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSaApfProfileResource_Create_APIError(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_sa_apf_profile")
}

// TestSaApfProfileResource_Create_APIErrorReadBody exercises SaApfProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSaApfProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SaApfProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSaApfProfileResource_Create_InvalidJSON exercises SaApfProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSaApfProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := SaApfProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSaApfProfileResource_Create_MapError exercises SaApfProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSaApfProfileResource_Create_MapError(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSaApfProfileResource_Create_MissingID exercises SaApfProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestSaApfProfileResource_Create_MissingID(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestSaApfProfileResource_Create_LocationFallback exercises SaApfProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestSaApfProfileResource_Create_LocationFallback(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestSaApfProfileResource_Read_Happy exercises SaApfProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestSaApfProfileResource_Read_Happy(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSaApfProfileResource_Read_NilClient exercises SaApfProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSaApfProfileResource_Read_NilClient(t *testing.T) {
	r := &SaApfProfileResource{}
	m := SaApfProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSaApfProfileResource_Read_BuildError exercises SaApfProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSaApfProfileResource_Read_BuildError(t *testing.T) {
	r := &SaApfProfileResource{client: newMalformedBaseURLClient(t)}
	m := SaApfProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSaApfProfileResource_Read_SendError exercises SaApfProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestSaApfProfileResource_Read_SendError(t *testing.T) {
	r := &SaApfProfileResource{client: newTransportErrorClient(t)}
	m := SaApfProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSaApfProfileResource_Read_NotFound exercises SaApfProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestSaApfProfileResource_Read_NotFound(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 404, "")}
	m := SaApfProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSaApfProfileResource_Read_APIError exercises SaApfProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSaApfProfileResource_Read_APIError(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_sa_apf_profile")
}

// TestSaApfProfileResource_Read_APIErrorReadBody exercises SaApfProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSaApfProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SaApfProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSaApfProfileResource_Read_InvalidJSON exercises SaApfProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSaApfProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := SaApfProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSaApfProfileResource_Read_MapError exercises SaApfProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSaApfProfileResource_Read_MapError(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSaApfProfileResource_Update_Happy exercises SaApfProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestSaApfProfileResource_Update_Happy(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSaApfProfileResource_Update_NilClient exercises SaApfProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSaApfProfileResource_Update_NilClient(t *testing.T) {
	r := &SaApfProfileResource{}
	m := SaApfProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSaApfProfileResource_Update_BuildError exercises SaApfProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSaApfProfileResource_Update_BuildError(t *testing.T) {
	r := &SaApfProfileResource{client: newMalformedBaseURLClient(t)}
	m := SaApfProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSaApfProfileResource_Update_SendError exercises SaApfProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestSaApfProfileResource_Update_SendError(t *testing.T) {
	r := &SaApfProfileResource{client: newTransportErrorClient(t)}
	m := SaApfProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSaApfProfileResource_Update_APIError exercises SaApfProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSaApfProfileResource_Update_APIError(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_sa_apf_profile")
}

// TestSaApfProfileResource_Update_APIErrorReadBody exercises SaApfProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSaApfProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SaApfProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSaApfProfileResource_Update_InvalidJSON exercises SaApfProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSaApfProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := SaApfProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSaApfProfileResource_Update_MapError exercises SaApfProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSaApfProfileResource_Update_MapError(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSaApfProfileResource_Delete_Happy exercises SaApfProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestSaApfProfileResource_Delete_Happy(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 204, "")}
	m := SaApfProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSaApfProfileResource_Delete_NilClient exercises SaApfProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSaApfProfileResource_Delete_NilClient(t *testing.T) {
	r := &SaApfProfileResource{}
	m := SaApfProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSaApfProfileResource_Delete_BuildError exercises SaApfProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSaApfProfileResource_Delete_BuildError(t *testing.T) {
	r := &SaApfProfileResource{client: newMalformedBaseURLClient(t)}
	m := SaApfProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSaApfProfileResource_Delete_SendError exercises SaApfProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestSaApfProfileResource_Delete_SendError(t *testing.T) {
	r := &SaApfProfileResource{client: newTransportErrorClient(t)}
	m := SaApfProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSaApfProfileResource_Delete_NotFoundSuccess exercises SaApfProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestSaApfProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 404, "")}
	m := SaApfProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSaApfProfileResource_Delete_APIError exercises SaApfProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSaApfProfileResource_Delete_APIError(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SaApfProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_sa_apf_profile")
}

// TestSaApfProfileResource_Delete_APIErrorReadBody exercises SaApfProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSaApfProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &SaApfProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SaApfProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
