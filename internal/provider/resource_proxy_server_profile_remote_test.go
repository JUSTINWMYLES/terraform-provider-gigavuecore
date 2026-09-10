package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestProxyServerProfileResource_Create_Happy exercises ProxyServerProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestProxyServerProfileResource_Create_Happy(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestProxyServerProfileResource_Create_NilClient exercises ProxyServerProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestProxyServerProfileResource_Create_NilClient(t *testing.T) {
	r := &ProxyServerProfileResource{}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestProxyServerProfileResource_Create_BuildError exercises ProxyServerProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestProxyServerProfileResource_Create_BuildError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMalformedBaseURLClient(t)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestProxyServerProfileResource_Create_SendError exercises ProxyServerProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestProxyServerProfileResource_Create_SendError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newTransportErrorClient(t)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestProxyServerProfileResource_Create_APIError exercises ProxyServerProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestProxyServerProfileResource_Create_APIError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_proxy_server_profile")
}

// TestProxyServerProfileResource_Create_APIErrorReadBody exercises ProxyServerProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestProxyServerProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestProxyServerProfileResource_Create_InvalidJSON exercises ProxyServerProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestProxyServerProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestProxyServerProfileResource_Create_MapError exercises ProxyServerProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestProxyServerProfileResource_Create_MapError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestProxyServerProfileResource_Create_MissingID exercises ProxyServerProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestProxyServerProfileResource_Create_MissingID(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestProxyServerProfileResource_Create_LocationFallback exercises ProxyServerProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestProxyServerProfileResource_Create_LocationFallback(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestProxyServerProfileResource_Read_Happy exercises ProxyServerProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestProxyServerProfileResource_Read_Happy(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestProxyServerProfileResource_Read_NilClient exercises ProxyServerProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestProxyServerProfileResource_Read_NilClient(t *testing.T) {
	r := &ProxyServerProfileResource{}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestProxyServerProfileResource_Read_BuildError exercises ProxyServerProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestProxyServerProfileResource_Read_BuildError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMalformedBaseURLClient(t)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestProxyServerProfileResource_Read_SendError exercises ProxyServerProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestProxyServerProfileResource_Read_SendError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newTransportErrorClient(t)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestProxyServerProfileResource_Read_NotFound exercises ProxyServerProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestProxyServerProfileResource_Read_NotFound(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 404, "")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestProxyServerProfileResource_Read_APIError exercises ProxyServerProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestProxyServerProfileResource_Read_APIError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_proxy_server_profile")
}

// TestProxyServerProfileResource_Read_APIErrorReadBody exercises ProxyServerProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestProxyServerProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestProxyServerProfileResource_Read_InvalidJSON exercises ProxyServerProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestProxyServerProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestProxyServerProfileResource_Read_MapError exercises ProxyServerProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestProxyServerProfileResource_Read_MapError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestProxyServerProfileResource_Update_Happy exercises ProxyServerProfileResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestProxyServerProfileResource_Update_Happy(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestProxyServerProfileResource_Update_NilClient exercises ProxyServerProfileResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestProxyServerProfileResource_Update_NilClient(t *testing.T) {
	r := &ProxyServerProfileResource{}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestProxyServerProfileResource_Update_BuildError exercises ProxyServerProfileResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestProxyServerProfileResource_Update_BuildError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMalformedBaseURLClient(t)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestProxyServerProfileResource_Update_SendError exercises ProxyServerProfileResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestProxyServerProfileResource_Update_SendError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newTransportErrorClient(t)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestProxyServerProfileResource_Update_APIError exercises ProxyServerProfileResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestProxyServerProfileResource_Update_APIError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_proxy_server_profile")
}

// TestProxyServerProfileResource_Update_APIErrorReadBody exercises ProxyServerProfileResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestProxyServerProfileResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestProxyServerProfileResource_Update_InvalidJSON exercises ProxyServerProfileResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestProxyServerProfileResource_Update_InvalidJSON(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestProxyServerProfileResource_Update_MapError exercises ProxyServerProfileResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestProxyServerProfileResource_Update_MapError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestProxyServerProfileResource_Delete_Happy exercises ProxyServerProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestProxyServerProfileResource_Delete_Happy(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 204, "")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestProxyServerProfileResource_Delete_NilClient exercises ProxyServerProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestProxyServerProfileResource_Delete_NilClient(t *testing.T) {
	r := &ProxyServerProfileResource{}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestProxyServerProfileResource_Delete_BuildError exercises ProxyServerProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestProxyServerProfileResource_Delete_BuildError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMalformedBaseURLClient(t)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestProxyServerProfileResource_Delete_SendError exercises ProxyServerProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestProxyServerProfileResource_Delete_SendError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newTransportErrorClient(t)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestProxyServerProfileResource_Delete_NotFoundSuccess exercises ProxyServerProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestProxyServerProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 404, "")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestProxyServerProfileResource_Delete_APIError exercises ProxyServerProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestProxyServerProfileResource_Delete_APIError(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_proxy_server_profile")
}

// TestProxyServerProfileResource_Delete_APIErrorReadBody exercises ProxyServerProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestProxyServerProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ProxyServerProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := ProxyServerProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
