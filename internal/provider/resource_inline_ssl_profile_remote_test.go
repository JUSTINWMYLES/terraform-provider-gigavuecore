package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestInlineSslProfileResource_Create_Happy exercises InlineSslProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestInlineSslProfileResource_Create_Happy(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslProfileResource_Create_NilClient exercises InlineSslProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineSslProfileResource_Create_NilClient(t *testing.T) {
	r := &InlineSslProfileResource{}
	m := InlineSslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineSslProfileResource_Create_BuildError exercises InlineSslProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineSslProfileResource_Create_BuildError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMalformedBaseURLClient(t)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineSslProfileResource_Create_SendError exercises InlineSslProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineSslProfileResource_Create_SendError(t *testing.T) {
	r := &InlineSslProfileResource{client: newTransportErrorClient(t)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineSslProfileResource_Create_APIError exercises InlineSslProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineSslProfileResource_Create_APIError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_inline_ssl_profile")
}

// TestInlineSslProfileResource_Create_APIErrorReadBody exercises InlineSslProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineSslProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInlineSslProfileResource_Create_InvalidJSON exercises InlineSslProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInlineSslProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInlineSslProfileResource_Create_MapError exercises InlineSslProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInlineSslProfileResource_Create_MapError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInlineSslProfileResource_Create_MissingID exercises InlineSslProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestInlineSslProfileResource_Create_MissingID(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestInlineSslProfileResource_Create_LocationFallback exercises InlineSslProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestInlineSslProfileResource_Create_LocationFallback(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestInlineSslProfileResource_Read_Happy exercises InlineSslProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestInlineSslProfileResource_Read_Happy(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslProfileResource_Read_NilClient exercises InlineSslProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineSslProfileResource_Read_NilClient(t *testing.T) {
	r := &InlineSslProfileResource{}
	m := InlineSslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineSslProfileResource_Read_BuildError exercises InlineSslProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineSslProfileResource_Read_BuildError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMalformedBaseURLClient(t)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineSslProfileResource_Read_SendError exercises InlineSslProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineSslProfileResource_Read_SendError(t *testing.T) {
	r := &InlineSslProfileResource{client: newTransportErrorClient(t)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineSslProfileResource_Read_NotFound exercises InlineSslProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestInlineSslProfileResource_Read_NotFound(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 404, "")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslProfileResource_Read_APIError exercises InlineSslProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineSslProfileResource_Read_APIError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_inline_ssl_profile")
}

// TestInlineSslProfileResource_Read_APIErrorReadBody exercises InlineSslProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineSslProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInlineSslProfileResource_Read_InvalidJSON exercises InlineSslProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInlineSslProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInlineSslProfileResource_Read_MapError exercises InlineSslProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInlineSslProfileResource_Read_MapError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInlineSslProfileResource_Update_Happy exercises InlineSslProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestInlineSslProfileResource_Update_Happy(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslProfileResource_Update_NilClient exercises InlineSslProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineSslProfileResource_Update_NilClient(t *testing.T) {
	r := &InlineSslProfileResource{}
	m := InlineSslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineSslProfileResource_Update_BuildError exercises InlineSslProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineSslProfileResource_Update_BuildError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMalformedBaseURLClient(t)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineSslProfileResource_Update_SendError exercises InlineSslProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineSslProfileResource_Update_SendError(t *testing.T) {
	r := &InlineSslProfileResource{client: newTransportErrorClient(t)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineSslProfileResource_Update_APIError exercises InlineSslProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineSslProfileResource_Update_APIError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_inline_ssl_profile")
}

// TestInlineSslProfileResource_Update_APIErrorReadBody exercises InlineSslProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineSslProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInlineSslProfileResource_Update_InvalidJSON exercises InlineSslProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInlineSslProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInlineSslProfileResource_Update_MapError exercises InlineSslProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInlineSslProfileResource_Update_MapError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInlineSslProfileResource_Delete_Happy exercises InlineSslProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestInlineSslProfileResource_Delete_Happy(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 204, "")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslProfileResource_Delete_NilClient exercises InlineSslProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineSslProfileResource_Delete_NilClient(t *testing.T) {
	r := &InlineSslProfileResource{}
	m := InlineSslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineSslProfileResource_Delete_BuildError exercises InlineSslProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineSslProfileResource_Delete_BuildError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMalformedBaseURLClient(t)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineSslProfileResource_Delete_SendError exercises InlineSslProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineSslProfileResource_Delete_SendError(t *testing.T) {
	r := &InlineSslProfileResource{client: newTransportErrorClient(t)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineSslProfileResource_Delete_NotFoundSuccess exercises InlineSslProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestInlineSslProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 404, "")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslProfileResource_Delete_APIError exercises InlineSslProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineSslProfileResource_Delete_APIError(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineSslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_inline_ssl_profile")
}

// TestInlineSslProfileResource_Delete_APIErrorReadBody exercises InlineSslProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineSslProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &InlineSslProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineSslProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
