package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestInlineSslAppResource_Create_Happy exercises InlineSslAppResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestInlineSslAppResource_Create_Happy(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslAppResource_Create_NilClient exercises InlineSslAppResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineSslAppResource_Create_NilClient(t *testing.T) {
	r := &InlineSslAppResource{}
	m := InlineSslAppResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineSslAppResource_Create_BuildError exercises InlineSslAppResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineSslAppResource_Create_BuildError(t *testing.T) {
	r := &InlineSslAppResource{client: newMalformedBaseURLClient(t)}
	m := InlineSslAppResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineSslAppResource_Create_SendError exercises InlineSslAppResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineSslAppResource_Create_SendError(t *testing.T) {
	r := &InlineSslAppResource{client: newTransportErrorClient(t)}
	m := InlineSslAppResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineSslAppResource_Create_APIError exercises InlineSslAppResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineSslAppResource_Create_APIError(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_inline_ssl_app")
}

// TestInlineSslAppResource_Create_APIErrorReadBody exercises InlineSslAppResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineSslAppResource_Create_APIErrorReadBody(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineSslAppResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInlineSslAppResource_Create_InvalidJSON exercises InlineSslAppResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInlineSslAppResource_Create_InvalidJSON(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 201, "{{")}
	m := InlineSslAppResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInlineSslAppResource_Create_MapError exercises InlineSslAppResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInlineSslAppResource_Create_MapError(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInlineSslAppResource_Create_MissingID exercises InlineSslAppResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestInlineSslAppResource_Create_MissingID(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 201, "{}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestInlineSslAppResource_Create_LocationFallback exercises InlineSslAppResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestInlineSslAppResource_Create_LocationFallback(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestInlineSslAppResource_Read_Happy exercises InlineSslAppResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestInlineSslAppResource_Read_Happy(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 200, "{}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslAppResource_Read_NilClient exercises InlineSslAppResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineSslAppResource_Read_NilClient(t *testing.T) {
	r := &InlineSslAppResource{}
	m := InlineSslAppResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineSslAppResource_Read_BuildError exercises InlineSslAppResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineSslAppResource_Read_BuildError(t *testing.T) {
	r := &InlineSslAppResource{client: newMalformedBaseURLClient(t)}
	m := InlineSslAppResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineSslAppResource_Read_SendError exercises InlineSslAppResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineSslAppResource_Read_SendError(t *testing.T) {
	r := &InlineSslAppResource{client: newTransportErrorClient(t)}
	m := InlineSslAppResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineSslAppResource_Read_NotFound exercises InlineSslAppResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestInlineSslAppResource_Read_NotFound(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 404, "")}
	m := InlineSslAppResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslAppResource_Read_APIError exercises InlineSslAppResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineSslAppResource_Read_APIError(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_inline_ssl_app")
}

// TestInlineSslAppResource_Read_APIErrorReadBody exercises InlineSslAppResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineSslAppResource_Read_APIErrorReadBody(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineSslAppResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInlineSslAppResource_Read_InvalidJSON exercises InlineSslAppResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInlineSslAppResource_Read_InvalidJSON(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 200, "{{")}
	m := InlineSslAppResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInlineSslAppResource_Read_MapError exercises InlineSslAppResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInlineSslAppResource_Read_MapError(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInlineSslAppResource_Update_Happy exercises InlineSslAppResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestInlineSslAppResource_Update_Happy(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 200, "{}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslAppResource_Update_NilClient exercises InlineSslAppResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineSslAppResource_Update_NilClient(t *testing.T) {
	r := &InlineSslAppResource{}
	m := InlineSslAppResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineSslAppResource_Update_BuildError exercises InlineSslAppResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineSslAppResource_Update_BuildError(t *testing.T) {
	r := &InlineSslAppResource{client: newMalformedBaseURLClient(t)}
	m := InlineSslAppResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineSslAppResource_Update_SendError exercises InlineSslAppResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineSslAppResource_Update_SendError(t *testing.T) {
	r := &InlineSslAppResource{client: newTransportErrorClient(t)}
	m := InlineSslAppResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineSslAppResource_Update_APIError exercises InlineSslAppResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineSslAppResource_Update_APIError(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_inline_ssl_app")
}

// TestInlineSslAppResource_Update_APIErrorReadBody exercises InlineSslAppResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineSslAppResource_Update_APIErrorReadBody(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineSslAppResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInlineSslAppResource_Update_InvalidJSON exercises InlineSslAppResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInlineSslAppResource_Update_InvalidJSON(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 200, "{{")}
	m := InlineSslAppResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInlineSslAppResource_Update_MapError exercises InlineSslAppResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInlineSslAppResource_Update_MapError(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInlineSslAppResource_Delete_Happy exercises InlineSslAppResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestInlineSslAppResource_Delete_Happy(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 204, "")}
	m := InlineSslAppResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslAppResource_Delete_NilClient exercises InlineSslAppResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineSslAppResource_Delete_NilClient(t *testing.T) {
	r := &InlineSslAppResource{}
	m := InlineSslAppResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInlineSslAppResource_Delete_BuildError exercises InlineSslAppResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInlineSslAppResource_Delete_BuildError(t *testing.T) {
	r := &InlineSslAppResource{client: newMalformedBaseURLClient(t)}
	m := InlineSslAppResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInlineSslAppResource_Delete_SendError exercises InlineSslAppResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestInlineSslAppResource_Delete_SendError(t *testing.T) {
	r := &InlineSslAppResource{client: newTransportErrorClient(t)}
	m := InlineSslAppResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInlineSslAppResource_Delete_NotFoundSuccess exercises InlineSslAppResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestInlineSslAppResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 404, "")}
	m := InlineSslAppResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInlineSslAppResource_Delete_APIError exercises InlineSslAppResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInlineSslAppResource_Delete_APIError(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InlineSslAppResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_inline_ssl_app")
}

// TestInlineSslAppResource_Delete_APIErrorReadBody exercises InlineSslAppResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInlineSslAppResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &InlineSslAppResource{client: newMockClientReadErrorBody(t, 501)}
	m := InlineSslAppResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
