package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSslProfileResource_Create_Happy exercises SslProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestSslProfileResource_Create_Happy(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := SslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSslProfileResource_Create_NilClient exercises SslProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSslProfileResource_Create_NilClient(t *testing.T) {
	r := &SslProfileResource{}
	m := SslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSslProfileResource_Create_BuildError exercises SslProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSslProfileResource_Create_BuildError(t *testing.T) {
	r := &SslProfileResource{client: newMalformedBaseURLClient(t)}
	m := SslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSslProfileResource_Create_SendError exercises SslProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestSslProfileResource_Create_SendError(t *testing.T) {
	r := &SslProfileResource{client: newTransportErrorClient(t)}
	m := SslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSslProfileResource_Create_APIError exercises SslProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSslProfileResource_Create_APIError(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_ssl_profile")
}

// TestSslProfileResource_Create_APIErrorReadBody exercises SslProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSslProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &SslProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSslProfileResource_Create_InvalidJSON exercises SslProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSslProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := SslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSslProfileResource_Create_MapError exercises SslProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSslProfileResource_Create_MapError(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := SslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSslProfileResource_Create_MissingID exercises SslProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestSslProfileResource_Create_MissingID(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := SslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestSslProfileResource_Create_LocationFallback exercises SslProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestSslProfileResource_Create_LocationFallback(t *testing.T) {
	r := &SslProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := SslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestSslProfileResource_Read_Happy exercises SslProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestSslProfileResource_Read_Happy(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := SslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSslProfileResource_Read_NilClient exercises SslProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSslProfileResource_Read_NilClient(t *testing.T) {
	r := &SslProfileResource{}
	m := SslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSslProfileResource_Read_BuildError exercises SslProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSslProfileResource_Read_BuildError(t *testing.T) {
	r := &SslProfileResource{client: newMalformedBaseURLClient(t)}
	m := SslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSslProfileResource_Read_SendError exercises SslProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestSslProfileResource_Read_SendError(t *testing.T) {
	r := &SslProfileResource{client: newTransportErrorClient(t)}
	m := SslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSslProfileResource_Read_NotFound exercises SslProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestSslProfileResource_Read_NotFound(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 404, "")}
	m := SslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSslProfileResource_Read_APIError exercises SslProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSslProfileResource_Read_APIError(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_ssl_profile")
}

// TestSslProfileResource_Read_APIErrorReadBody exercises SslProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSslProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &SslProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSslProfileResource_Read_InvalidJSON exercises SslProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSslProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := SslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSslProfileResource_Read_MapError exercises SslProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSslProfileResource_Read_MapError(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSslProfileResource_Update_Happy exercises SslProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestSslProfileResource_Update_Happy(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := SslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSslProfileResource_Update_NilClient exercises SslProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSslProfileResource_Update_NilClient(t *testing.T) {
	r := &SslProfileResource{}
	m := SslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSslProfileResource_Update_BuildError exercises SslProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSslProfileResource_Update_BuildError(t *testing.T) {
	r := &SslProfileResource{client: newMalformedBaseURLClient(t)}
	m := SslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSslProfileResource_Update_SendError exercises SslProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestSslProfileResource_Update_SendError(t *testing.T) {
	r := &SslProfileResource{client: newTransportErrorClient(t)}
	m := SslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSslProfileResource_Update_APIError exercises SslProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSslProfileResource_Update_APIError(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_ssl_profile")
}

// TestSslProfileResource_Update_APIErrorReadBody exercises SslProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSslProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &SslProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSslProfileResource_Update_InvalidJSON exercises SslProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSslProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := SslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSslProfileResource_Update_MapError exercises SslProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSslProfileResource_Update_MapError(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSslProfileResource_Delete_Happy exercises SslProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestSslProfileResource_Delete_Happy(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 204, "")}
	m := SslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSslProfileResource_Delete_NilClient exercises SslProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSslProfileResource_Delete_NilClient(t *testing.T) {
	r := &SslProfileResource{}
	m := SslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSslProfileResource_Delete_BuildError exercises SslProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSslProfileResource_Delete_BuildError(t *testing.T) {
	r := &SslProfileResource{client: newMalformedBaseURLClient(t)}
	m := SslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSslProfileResource_Delete_SendError exercises SslProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestSslProfileResource_Delete_SendError(t *testing.T) {
	r := &SslProfileResource{client: newTransportErrorClient(t)}
	m := SslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSslProfileResource_Delete_NotFoundSuccess exercises SslProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestSslProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 404, "")}
	m := SslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSslProfileResource_Delete_APIError exercises SslProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSslProfileResource_Delete_APIError(t *testing.T) {
	r := &SslProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_ssl_profile")
}

// TestSslProfileResource_Delete_APIErrorReadBody exercises SslProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSslProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &SslProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := SslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
